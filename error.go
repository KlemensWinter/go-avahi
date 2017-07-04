package avahi

// #include "glue.h"
import "C"

import (
	"fmt"
)

type Error int

// Error Codes
//
// see http://avahi.org/download/doxygen/error_8h.html
const (
	ErrOk                    = Error(C.AVAHI_OK)
	ErrAccessDendied         = Error(C.AVAHI_ERR_ACCESS_DENIED)
	ErrBadState              = Error(C.AVAHI_ERR_BAD_STATE)
	ErrCollision             = Error(C.AVAHI_ERR_COLLISION)
	ErrDBusError             = Error(C.AVAHI_ERR_DBUS_ERROR)
	ErrDisconnected          = Error(C.AVAHI_ERR_DISCONNECTED)
	ErrDnsFormerr            = Error(C.AVAHI_ERR_DNS_FORMERR)
	ErrDnsNoAuth             = Error(C.AVAHI_ERR_DNS_NOTAUTH)
	ErrDnsNotImp             = Error(C.AVAHI_ERR_DNS_NOTIMP)
	ErrDnsNotZone            = Error(C.AVAHI_ERR_DNS_NOTZONE)
	ErrDnsNxDomain           = Error(C.AVAHI_ERR_DNS_NXDOMAIN)
	ErrDnsNxrrSet            = Error(C.AVAHI_ERR_DNS_NXRRSET)
	ErrDnsRefused            = Error(C.AVAHI_ERR_DNS_REFUSED)
	ErrDnsServfail           = Error(C.AVAHI_ERR_DNS_SERVFAIL)
	ErrDnsYxDomain           = Error(C.AVAHI_ERR_DNS_YXDOMAIN)
	ErrDnsYxrrSet            = Error(C.AVAHI_ERR_DNS_YXRRSET)
	ErrFailure               = Error(C.AVAHI_ERR_FAILURE)
	ErrInvalidAddress        = Error(C.AVAHI_ERR_INVALID_ADDRESS)
	ErrInvalidArgument       = Error(C.AVAHI_ERR_INVALID_ARGUMENT)
	ErrInvalidConfig         = Error(C.AVAHI_ERR_INVALID_CONFIG)
	ErrInvalidDnsClass       = Error(C.AVAHI_ERR_INVALID_DNS_CLASS)
	ErrInvalidDnsError       = Error(C.AVAHI_ERR_INVALID_DNS_ERROR)
	ErrInvalidDnsType        = Error(C.AVAHI_ERR_INVALID_DNS_TYPE)
	ErrInvalidDomainName     = Error(C.AVAHI_ERR_INVALID_DOMAIN_NAME)
	ErrInvalidFlags          = Error(C.AVAHI_ERR_INVALID_FLAGS)
	ErrInvalidHostName       = Error(C.AVAHI_ERR_INVALID_HOST_NAME)
	ErrInvalidInterface      = Error(C.AVAHI_ERR_INVALID_INTERFACE)
	ErrInvalidKey            = Error(C.AVAHI_ERR_INVALID_KEY)
	ErrInvalidObject         = Error(C.AVAHI_ERR_INVALID_OBJECT)
	ErrInvalidOperation      = Error(C.AVAHI_ERR_INVALID_OPERATION)
	ErrInvalidPacket         = Error(C.AVAHI_ERR_INVALID_PACKET)
	ErrInvalidPort           = Error(C.AVAHI_ERR_INVALID_PORT)
	ErrInvalidProtocol       = Error(C.AVAHI_ERR_INVALID_PROTOCOL)
	ErrInvalidRdata          = Error(C.AVAHI_ERR_INVALID_RDATA)
	ErrInvalidRecord         = Error(C.AVAHI_ERR_INVALID_RECORD)
	ErrInvalidServiceName    = Error(C.AVAHI_ERR_INVALID_SERVICE_NAME)
	ErrInvalidServiceSubtype = Error(C.AVAHI_ERR_INVALID_SERVICE_SUBTYPE)
	ErrInvalidServiceType    = Error(C.AVAHI_ERR_INVALID_SERVICE_TYPE)
	ErrInvalidTTL            = Error(C.AVAHI_ERR_INVALID_TTL)
	ErrIsEmpty               = Error(C.AVAHI_ERR_IS_EMPTY)
	ErrIsPattern             = Error(C.AVAHI_ERR_IS_PATTERN)
	ErrMax                   = Error(C.AVAHI_ERR_MAX)
	ErrNoChange              = Error(C.AVAHI_ERR_NO_CHANGE)
	ErrNoDaemon              = Error(C.AVAHI_ERR_NO_DAEMON)
	ErrNoMemory              = Error(C.AVAHI_ERR_NO_MEMORY)
	ErrNoNetwork             = Error(C.AVAHI_ERR_NO_NETWORK)
	ErrNotFound              = Error(C.AVAHI_ERR_NOT_FOUND)
	ErrNotPermitted          = Error(C.AVAHI_ERR_NOT_PERMITTED)
	ErrNotSupported          = Error(C.AVAHI_ERR_NOT_SUPPORTED)
	ErrOs                    = Error(C.AVAHI_ERR_OS)
	ErrTimeout               = Error(C.AVAHI_ERR_TIMEOUT)
	ErrTooManyClients        = Error(C.AVAHI_ERR_TOO_MANY_CLIENTS)
	ErrTooManyEntries        = Error(C.AVAHI_ERR_TOO_MANY_ENTRIES)
	ErrTooManyObjects        = Error(C.AVAHI_ERR_TOO_MANY_OBJECTS)
	ErrVersionMismatch       = Error(C.AVAHI_ERR_VERSION_MISMATCH)
)

func (err Error) Error() string {
	ret, _ := C.avahi_strerror(C.int(err))
	return fmt.Sprint("avahi: ", C.GoString(ret))
}
