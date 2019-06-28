package learn

import (
	"encoding/json"
	"fmt"
)

// 只要是可导出成员（变量首字母大写），都可以转成json。
// 因成员变量test是不可导出的，故无法转成json
// 如果变量打上了json标签，如Name旁边的 `json:"name"` ，
// 那么转化成的json key就用该标签“name”，否则取变量名作为key
type User struct {
	Name    string `json:"name"`
	Age     int
	Address string
	test    string
}

func Learn11() {
	var user = User{"loren", 18, "laoshan", "ttttt"}
	jsonData, err := json.Marshal(&user)
	fmt.Println(err, string(jsonData))
	str := `{"name": "runoob","age": 18,"address": "laoshan"}`
	var user2 User
	err2 := json.Unmarshal([]byte(str), &user2)
	fmt.Printf("%v - %v\n", err2, user2.Age)
	//interface
	var t TestInterface = &user2
	t.interFun1(999)
	fmt.Printf("%v - %v\n", err2, user2.Age)
}

func (u *User) interFun1(s int) {
	u.Age = s
}
func (u *User) interFun2(s string) {
	u.Name = s
}

type TestInterface interface {
	interFun1(s int)
	interFun2(s string)
}
