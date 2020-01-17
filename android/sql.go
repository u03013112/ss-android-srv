package android

import (
	"time"

	"github.com/u03013112/ss-android/sql"
)

// User : android 用户表，用于存储UUID和有效期
type User struct {
	sql.BaseModel
	UUID           string    `json:"uuid,omitempty"`
	Token          string    `json:"token,omitempty"`
	ExpireDate     time.Time `json:"expireDate,omitempty"`
	TotalRxTraffic int64     `json:"totalRxTraffic,omitempty"`
	UsedRxTraffic  int64     `json:"usedRxTraffic,omitempty"`
	BaseRxTraffic  int64     `json:"baseRxTraffic,omitempty"`
	UsedTxTraffic  int64     `json:"usedTxTraffic,omitempty"`
	BaseTxTraffic  int64     `json:"baseTxTraffic,omitempty"`
	Online         bool      `json:"online,omitempty"`
	IP             string    `json:"ip,omitempty"`
	ClientVersion  string    `json:"clientVersion,omitempty"`
}

// TableName :
func (User) TableName() string {
	return "android_user"
}

// InitDB : 初始化表格，建议在整个数据库初始化之后调用
func InitDB() {
	sql.GetInstance().AutoMigrate(&User{}, &Production{}, &GoogleAd{}, &Bill{}, &LoginLog{})
}

func getOrCreateUserByUUID(uuid string) User {
	var user User
	db := sql.GetInstance().First(&user, "uuid = ?", uuid)
	if db.RowsAffected == 0 {
		user.UUID = uuid
		user.ExpireDate = time.Unix(0, 0)
		user.Token = ""
		user.TotalRxTraffic = 0
		user.UsedRxTraffic = 0
		user.BaseRxTraffic = 0
		user.Online = false
		sql.GetInstance().Create(&user)
	}
	return user
}

func (u *User) updateToken(token string) {
	sql.GetInstance().Model(u).Update(User{
		Token: token,
	})
}

func (u *User) updateUserInfo(token string, ip string, version string) {
	sql.GetInstance().Model(u).Update(User{
		Token:         token,
		IP:            ip,
		ClientVersion: version,
	})
}

func getUserByToken(token string) (*User, error) {
	var user User
	db := sql.GetInstance().Model(&User{}).Where("token = ?", token).First(&user)
	if db.RowsAffected == 1 {
		return &user, nil
	}
	return nil, db.Error
}

func (u *User) update() error {
	db := sql.GetInstance().Save(u)
	return db.Error
}

// Production :
type Production struct {
	sql.BaseModel
	Description string
	Time        int64
	Total       int64
	Price       int64
	Hiden       int64
}

func (p *Production) findAll() ([]Production, error) {
	pList := []Production{}
	db := sql.GetInstance().Find(&pList)
	return pList, db.Error
}

func (p *Production) getByID() error {
	db := sql.GetInstance().Where(p).First(p)
	return db.Error
}

// Bill :
type Bill struct {
	sql.BaseModel
	UserID       int64 `json:"userID,omitempty"`
	ProductionID int64 `json:"productionID,omitempty"`
}

// TableName :
func (Bill) TableName() string {
	return "android_bills"
}

func (b *Bill) record() {
	sql.GetInstance().Create(b)
}

// GoogleAd :
type GoogleAd struct {
	sql.BaseModel
	Gid string
}

func getGoogleAd() string {
	var g GoogleAd
	sql.GetInstance().First(&g)
	return g.Gid
}

func getGoogleAdFree() string {
	var g GoogleAd
	sql.GetInstance().First(&g, 2)
	return g.Gid
}

// LoginLog :
type LoginLog struct {
	sql.BaseModel
	UserID int64  `json:"userID,omitempty"`
	IP     string `json:"ip,omitempty"`
}

// TableName :
func (LoginLog) TableName() string {
	return "android_login_logs"
}

func (b *LoginLog) record() {
	sql.GetInstance().Create(b)
}
