package main

import (
	"fmt"

	"github.com/wdfky/interception"
)

const (
	SCANCODE_X   = 0x2D
	SCANCODE_Y   = 0x15
	SCANCODE_ESC = 0x01
)

func main() {
	inteception := interception.New()
	defer inteception.Destroy()
	inteception.SetFilter(inteception.IsKeyBoard(), interception.FILTER_KEY_DOWN|interception.FILTER_KEY_UP)
	cnt := 0
	emptySrtoke := &interception.KeyBoardStroke{}
	for cnt < 1000 {
		cnt++
		device := inteception.Wait()
		fmt.Println(device)
		if inteception.Receive(inteception.Wait(), &emptySrtoke) == 0 {
			break
		}
		if emptySrtoke.Code == SCANCODE_X {
			emptySrtoke.Code = SCANCODE_Y
		}
		inteception.Send(device, &emptySrtoke)
	}
}
