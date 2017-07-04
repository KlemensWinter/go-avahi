package avahi

// #include "glue.h"
import "C"

import (
	"fmt"
	"runtime"
	"unsafe"
)

type BrowserEvent C.AvahiBrowserEvent

const (
	BrowserAllForNow      = BrowserEvent(C.AVAHI_BROWSER_ALL_FOR_NOW)
	BrowserCacheExhausted = BrowserEvent(C.AVAHI_BROWSER_CACHE_EXHAUSTED)
	BrowserFailure        = BrowserEvent(C.AVAHI_BROWSER_FAILURE)
	BrowserNew            = BrowserEvent(C.AVAHI_BROWSER_NEW)
	BrowserRemove         = BrowserEvent(C.AVAHI_BROWSER_REMOVE)
)

func (be BrowserEvent) String() string {
	switch be {
	case BrowserAllForNow:
		return "BrowserAllFoNow"
	case BrowserCacheExhausted:
		return "BrowserCacheExhausted"
	case BrowserFailure:
		return "BrowserFailure"
	case BrowserNew:
		return "BrowserNew"
	case BrowserRemove:
		return "BrowserRemove"
	}
	return fmt.Sprintf("BrowserEvent(0x%x)", int(be))
}

type ServiceBrowser struct {
    Results chan *BrowserResult

	ptr *C.AvahiServiceBrowser
}

type LookupResultFlags C.AvahiLookupResultFlags

type BrowserResult struct {
    IfIndex int
    Proto   Protocol
    Event   BrowserEvent
    Name    string
    Type    string
    Domain string
    Flags   LookupResultFlags
}

var browsers = map[uintptr]*ServiceBrowser{}

//export go_service_browser_callback
func go_service_browser_callback(b *C.AvahiServiceBrowser, interf C.AvahiIfIndex, proto C.AvahiProtocol,
	event C.AvahiBrowserEvent, cname, ctype, cdomain *C.char, flags C.AvahiLookupResultFlags,
	userdata C.uintptr_t) {

    sb := uintptr(userdata)

    browser, ok := browsers[sb]
    if !ok {
        panic("internal error: browser not found")
    }

    bcd := &BrowserResult{
        IfIndex: int(interf),
        Proto: Protocol(proto),
        Event: BrowserEvent(event),
        Name: C.GoString(cname),
        Type:   C.GoString(ctype),
        Domain: C.GoString(cdomain),
        Flags: LookupResultFlags(flags),
    }
    browser.Results <-bcd
}


func NewServiceBrowser(client *Client, interf int, proto Protocol, serviceType string) (*ServiceBrowser, error) {
	cServiceType := C.CString(serviceType)

	ret := &ServiceBrowser{
        Results: make(chan *BrowserResult),
    }

    id := uintptr(unsafe.Pointer(ret))

	sb, _ := C.go_avahi_service_browser_new(
        client.GetPtr(),
        C.AvahiIfIndex(interf),
		C.AvahiProtocol(proto),
        cServiceType,
        nil,
        0,
		(*[0]byte)(C.go_service_browser_callback), // callback
		C.uintptr_t(id))                       // userdata
	if sb == nil {
		return nil, fmt.Errorf("avahi: browser_new() failed")
	}

    browsers[uintptr(unsafe.Pointer(ret))] = ret

	ret.ptr = sb
	runtime.SetFinalizer(ret, func(sb *ServiceBrowser) {
        delete(browsers, uintptr(unsafe.Pointer(sb)))
		sb.free()
	})
	return ret, nil
}

func (sb *ServiceBrowser) free() {
	C.avahi_service_browser_free(sb.ptr)
	sb.ptr = nil
}
