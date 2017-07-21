package hello

import (
	"fmt"
	app "github.com/kooksee/gin-demo/app"
)

type Hello struct {
	name string
}

func (h *Hello) Say(name string) {
	h.name = name
	fmt.Println("Hello")
}

func (h *Hello) GetName() string {
	print(h.name)
	return h.name
}

func init() {
	s := app.GetInstance()
	s.Set("hh", &Hello{})

	hh, ok := s.Get("hh1").(*Hello)
	if ok {
		hh.Say("Hello")
	} else {
		println(ok)
	}

}

func InitHello() {
	s := app.GetInstance()
	s.Set("h", &Hello{})
}
