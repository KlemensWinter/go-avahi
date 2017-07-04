package avahi

// #cgo pkg-config: avahi-client
// #include <avahi-common/defs.h>
// #include <avahi-common/simple-watch.h>
import "C"

import (
	"runtime"
)

// #include <avahi-client/client.h>

// Simple Poll
type SimplePoll struct {
	c *C.AvahiSimplePoll
}

func NewSimplePoll() (ret *SimplePoll) {
	ptr, _ := C.avahi_simple_poll_new()
	if ptr == nil {
		panic("simple_poll_new() failed")
	}

	ret = &SimplePoll{ptr}
	runtime.SetFinalizer(ret, func(sp *SimplePoll) {
		C.avahi_simple_poll_free(sp.c)
	})
	return
}

func (sp *SimplePoll) Loop() (err error) {
	ret, err := C.avahi_simple_poll_loop(sp.c)
	if err != nil {
		panic(err)
	}
	return Error(ret)
}

func (sp *SimplePoll) Quit() {
	_, err := C.avahi_simple_poll_quit(sp.c)
	if err != nil {
		panic(err)
	}
	return
}

func (sp *SimplePoll) GetPoll() *C.AvahiPoll {
	return C.avahi_simple_poll_get(sp.c)
}
