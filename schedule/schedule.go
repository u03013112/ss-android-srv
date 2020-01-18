package schedule

import (
	"fmt"
)

// Init :
func Init() {
	go func() {
		for {
			changePasswd()
			time.Sleep(time.Second * 60)
		}
	}()
}

func changePasswd() {
	fmt.Printf("changePasswd\n")

}
