package avahi

// #include "glue.h"
import "C"

import ()

type Watch C.AvahiWatch
type WatchEvent C.AvahiWatchEvent

type Timeout C.AvahiTimeout

type TimeoutCallback func()

type Poll interface {
	GetPoll() *C.AvahiPoll
}

/*
type Poll interface {
    NewWatch(fd uintptr, ev WatchEvent) *Watch
    WatchUpdate(watch *Watch, event WatchEvent)
    WatchGetEvents(watch *Watch) WatchEvent
    WatchFree(watch *Watch)

    NewTimeout(tv time.Duration, cb TimeoutCallback) *Timeout
    TimeoutUpdate(timeout *Timeout, tv time.Duration)
    TimeoutFree(timeout *Timeout)
}
*/
