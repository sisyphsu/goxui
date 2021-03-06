package main

import (
	"github.com/go-eden/goxui"
	"github.com/go-eden/goxui/util"
	slog "github.com/go-eden/slf4go"
	"runtime"
	"time"
)

type Param struct {
	Name string
	Age  int64
}

type Status struct {
	Flag bool
	Root Root
	User User
}

func (s *Status) Test(a int64, b float64) []string {
	return []string{util.ToString(a), util.ToString(b)}
}

type Root struct {
	Number  int32
	Number2 int64
	Str     string
	Body    Body
}

func (r *Root) Test(s1 string, s2 string, m map[string]interface{}) []interface{} {
	return []interface{}{s1, s2, m}
}

func (r *Root) Test0(s1 string, s2 string) []interface{} {
	return []interface{}{s1, s2}
}
func (r *Root) Test1(s1 string, s2 string) []interface{} {
	return []interface{}{s1, s2}
}
func (r *Root) Test2(s1 string, s2 string) []interface{} {
	return []interface{}{s1, s2}
}

func (r *Root) Test3(s string, p Param) {
	slog.Infof("########## s[%v], param[%v]", s, p)
}

type Body struct {
	Real float64
}

func (b *Body) Test(a1, a2, a3 interface{}) User {
	return User{true, "哈哈", 99, 88.33}
}

func (b *Body) Test1(a1, a2, a3 interface{}) *User {
	return &User{false, "你妹", 999, 888.33}
}

type User struct {
	Enable bool
	Name   string
	Age    int
	Score  float32
}

func (u *User) ChangeInfo(name string, age int, score float32) {
	u.Name = name
	u.Age = age
	u.Score = score
}

func (u User) QueryInfo() (string, bool) {
	return u.Name, u.Enable
}

/**
  uilib测试程序
*/
func main() {
	runtime.LockOSThread()

	goxui.Init()

	goxui.BindObject(new(User))

	go func() {
		time.Sleep(time.Second * 5)
		goxui.TriggerEvent("event_bool", true)
		goxui.TriggerEvent("event_int", 10000)
		goxui.TriggerEvent("event_long", 10000000)
		goxui.TriggerEvent("event_double", 10000.4444)
		goxui.TriggerEvent("event_string", "fdasfadsfasdfdafdsafdsa")
		goxui.TriggerEvent("event_object", Param{"啦啦啦", 3333333})
		goxui.TriggerEvent("event_array", []Param{{"啦啦啦", 3333333}, {"啦啦啦444", 3333333}})
	}()

	goxui.StartRelative("ui", "fulltest.qml")
}
