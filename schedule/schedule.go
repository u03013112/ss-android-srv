package schedule

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/u03013112/ss-android/android"
	"github.com/u03013112/ss-android/sql"
)

// Init :
func Init() {
	sql.GetInstance().AutoMigrate(&Check{})
	go func() {
		for {
			if e := check(); e != nil {
				fmt.Printf("ckecj err:%s\n", e)
				continue
			}
			time.Sleep(time.Second * 300)
		}
	}()
}

type used struct {
	Used  string `json:"used,omitempty"`
	Error string `json:"error,omitempty"`
}

func getJMSUsed() (float64, error) {
	resp, err := http.Get("http://frp.u03013112.win:28000/used")
	if err != nil {
		fmt.Println(err)
		return 0, errors.New("getUsed() get err")
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var u used
			if e := json.Unmarshal(body, &u); e != nil {
				return 0, errors.New("getUsed() Unmarshal err")
			}
			uf, _ := strconv.ParseFloat(u.Used, 64)
			return uf, nil
		}
	}
	return 0, errors.New("getUsed() err")
}

type passwd struct {
	Passwd string
	Error  string
}

func changePasswd() (string, error) {
	resp, err := http.Get("http://frp.u03013112.win:28000/changepasswd")
	if err != nil {
		fmt.Println(err)
		return "", errors.New("changePasswd() get err")
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var p passwd
			if e := json.Unmarshal(body, &p); e != nil {
				return "", errors.New("changePasswd() Unmarshal err")
			}
			if p.Error != "" {
				return p.Passwd, errors.New(p.Error)
			}
			return p.Passwd, nil
		}
	}
	return "", errors.New("changePasswd() err")
}

// Check : 监测日志
type Check struct {
	sql.BaseModel
	TotalUsedInJMS    float64
	TotalUsedInSystem float64
	UsedInJMS         float64
	UsedInSystem      float64
	ChangePasswd      string
}

func getUsedInSystem() float64 {
	var rx float64
	var tx float64
	row := sql.GetInstance().Raw("SELECT SUM(used_rx_traffic)/1024/1024/1024 as rx , SUM(used_tx_traffic)/1024/1024/1024 as tx FROM android_user").Row()
	row.Scan(&rx, &tx)
	return rx + tx
}

// 监测是否有流量外流
func check() error {
	ju, e := getJMSUsed()
	if e != nil {
		return e
	}
	if ju == 0 {
		// 目前有几率取到0流量
		return errors.New("ju = 0")
	}
	su := getUsedInSystem()

	var c Check
	db := sql.GetInstance().Last(&c)
	if db.RowsAffected == 0 {
		fmt.Printf("ju=%f su=%f\n", ju, su)
		record := &Check{
			TotalUsedInJMS:    ju,
			TotalUsedInSystem: su,
		}
		db = sql.GetInstance().Create(record)
		if db.Error != nil {
			fmt.Printf("create err:%s\n", db.Error)
		}
		return nil
	}

	record := &Check{}
	record.UsedInJMS = ju - c.TotalUsedInJMS
	record.UsedInSystem = su - c.TotalUsedInSystem
	record.TotalUsedInJMS = ju
	record.TotalUsedInSystem = su

	if (record.UsedInSystem > 0 && (record.UsedInJMS/record.UsedInSystem) > 5) ||
		(record.UsedInSystem == 0 && record.UsedInJMS > 1) { //jms中使用超过系统统计的5倍，就改密码
		p, e := changePasswd()
		if e != nil {
			return e
		}
		record.ChangePasswd = p
		android.GrpcSetPasswd(p)
	}
	sql.GetInstance().Create(record)

	return nil
}
