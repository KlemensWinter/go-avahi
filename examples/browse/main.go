package main

import (
	"fmt"
	"github.com/KlemensWinter/go-avahi"
)

func foo() error {
	return avahi.ERR_INVALID_DNS_ERROR
}

func main() {
	poll := avahi.NewSimplePoll()
	defer poll.Quit()

	return
}
