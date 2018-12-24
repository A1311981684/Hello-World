package models

import (
	"github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
)

func TestAddUser(t *testing.T) {
	user := User{
		Id:       "abc123",
		Username: "韦祖健",
		Password: "123456",
		Profile: Profile{
			Gender:  "male",
			Age:     100,
			Address: "6666",
			Email:   "12313123@gg.com",
		},
	}

	convey.Convey("测试AddUser", t, func() {
		AddUser(user)
		convey.So(UserList["abc123"].Username, convey.ShouldEqual, "韦祖健")
		convey.So(UserList["abc123"].Password, convey.ShouldEqual, "123456")
		//...
	})
}
func TestGetUser(t *testing.T) {
	convey.Convey("测试GETuser", t, func() {
		u, err := GetUser("user_11111")
		if err != nil {
			log.Print("不存在")
			panic("not exist")
		}
		convey.So(err, convey.ShouldBeNil)
		convey.So(u.Id, convey.ShouldEqual, "user_11111")
		u, err = GetUser("小行星")
		log.Print("不存在")
		//panic("not exist")
		convey.So(err, convey.ShouldNotBeNil)
		//判断map中存不存在这个value
		_, ok := UserList["小行星"]
		convey.So(ok, convey.ShouldEqual, false)
	})
}
