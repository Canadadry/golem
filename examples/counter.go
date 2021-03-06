package main

import (
	"demo/common"
	"fmt"
	"github.com/canadadry/golem"
)

const css = `
		* { margin: 0; padding: 0; box-sizing: border-box; user-select: none; }
		body { height: 100vh; display: flex; align-items: center; justify-content: center; background-color: #f1c40f; font-family: 'Helvetika Neue', Arial, sans-serif; font-size: 28px; }
		.counter-container { display: flex; flex-direction: column; align-items: center; }
		.counter { text-transform: uppercase; color: #fff; font-weight: bold; font-size: 3rem; }
		.btn-row { display: flex; align-items: center; margin: 1rem; }
		.btn { cursor: pointer; min-width: 4em; padding: 1em; border-radius: 5px; text-align: center; margin: 0 1rem; box-shadow: 0 6px #8b5e00; color: white; background-color: #E4B702; position: relative; font-weight: bold; }
		.btn:hover { box-shadow: 0 4px #8b5e00; top: 2px; }
		.btn:active{ box-shadow: 0 1px #8b5e00; top: 5px; }
	`

func Init() interface{} {
	return int(0)
}

func View(model interface{}) golem.Element {
	return golem.Div(golem.Props{Class: "counter-container"},
		golem.Div(golem.Props{Class: "counter", Id: "counter", InnerText: fmt.Sprintf("%v", model)}),
		golem.Div(golem.Props{Class: "btn-row"},
			golem.Div(golem.Props{
				Class:     "btn btn-incr",
				InnerText: "+1",
				EventListener: map[string]string{
					"click": "Plus",
				},
			}),
			golem.Div(golem.Props{
				Class:     "btn btn-decr",
				InnerText: "-1",
				EventListener: map[string]string{
					"click": "Minus",
				},
			}),
		),
	)
}

func Update(m interface{}, event string) (interface{}, string) {
	c, ok := m.(int)
	if !ok {
		c, ok = Init().(int)
		if !ok {
			panic("init func should return an int")
		}
	}

	switch event {
	case "Plus":
		c += 1
	case "Minus":
		c -= 1
	}
	return c, ""
}

func main() {
	if err := common.Run(480, 320, css, Init, View, Update); err != nil {
		fmt.Println(err)
	}
}
