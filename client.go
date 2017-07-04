package avahi

// #include "glue.h"
// void go_client_callback(AvahiClient* client, AvahiClientState state, void* userdata);
import "C"

import (
	"log"
	"runtime"
	"unsafe"
)

const (
	CLIENT_CONNECTING         = C.AVAHI_CLIENT_CONNECTING
	CLIENT_FAILURE            = C.AVAHI_CLIENT_FAILURE
	CLIENT_IGNORE_USER_CONFIG = C.AVAHI_CLIENT_IGNORE_USER_CONFIG
	CLIENT_NO_FAIL            = C.AVAHI_CLIENT_NO_FAIL
	CLIENT_S_COLLISION        = C.AVAHI_CLIENT_S_COLLISION
	CLIENT_S_REGISTERING      = C.AVAHI_CLIENT_S_REGISTERING
	CLIENT_S_RUNNING          = C.AVAHI_CLIENT_S_RUNNING
)

type Protocol C.AvahiProtocol

const (
	ProtoInet   = Protocol(C.AVAHI_PROTO_INET)
	ProtoInet6  = Protocol(C.AVAHI_PROTO_INET6)
	ProtoUnspec = Protocol(C.AVAHI_PROTO_UNSPEC)
)

func (p Protocol) String() string {
	return C.GoString(C.avahi_proto_to_string(C.AvahiProtocol(p)))
}

const (
	IF_UNSPEC = C.AVAHI_IF_UNSPEC
)

type ClientFlag C.AvahiClientFlags

const (
	ClientNoFail = ClientFlag(C.AVAHI_CLIENT_NO_FAIL)
)

type Client struct {
	Logger *log.Logger

	ptr *C.AvahiClient
}

//export go_client_callback
func go_client_callback(s *C.AvahiClient, state C.AvahiClientState, userdata unsafe.Pointer) {
	log.Printf("client callback")
}

func NewClient(poll Poll, flags ClientFlag) (*Client, error) {
	var errCode C.int

	cl, _ := C.avahi_client_new(poll.GetPoll(), C.AvahiClientFlags(flags), nil, C.go_client_callback, &errCode)
	if cl == nil {
		return nil, Error(errCode)
	}
	ret := &Client{ptr: cl}
	runtime.SetFinalizer(ret, func(c *Client) {
		c.free()
	})
	return ret, nil
}

func (c *Client) GetPtr() *C.AvahiClient {
	return c.ptr
}

func (c *Client) free() {
	C.avahi_client_free(c.ptr)
}
