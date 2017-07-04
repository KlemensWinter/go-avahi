package main

import (
    "log"
	"github.com/KlemensWinter/go-avahi"
)

func foo() error {
	return avahi.ErrInvalidDnsError
}

func main() {
	poll := avahi.NewSimplePoll()
	defer poll.Quit()

    client, err := avahi.NewClient(poll, 0)
    if err != nil {
        log.Fatalf("error: %s", err)
    }


    sb, err := avahi.NewServiceBrowser(client, avahi.IF_UNSPEC, avahi.ProtoUnspec, "_http._tcp")
    if err != nil {
        log.Fatalf("error: %s", err)
    }

    go func() {
        for {
            res := <-sb.Results
            log.Printf("result: %#v", res)
        }
    }()



    if err := poll.Loop(); err != nil {
        log.Fatalf("error: %s", err)
    }


	return
}
