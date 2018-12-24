package models

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAddUser(t *testing.T) {
	user := User{
		Id:"abc123",
		Username:"韦祖健",
		Password:"123456",
		Profile:Profile{
			Gender:"male",
			Age:100,
			Address:"6666",
			Email:"12313123@gg.com",
		},
	}

	convey.Convey("测试AddUser", t, func() {
		AddUser(user)
		convey.So(UserList["abc123"].Username, convey.ShouldEqual, "韦祖健")
		convey.So(UserList["abc123"].Password, convey.ShouldEqual, "123456")
		//...
	})
}