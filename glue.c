#include "glue.h"


AvahiServiceBrowser*
go_avahi_service_browser_new(
    AvahiClient* client,
    AvahiIfIndex index,
    AvahiProtocol proto,
    const char* serviceType,
    const char* domain,
    AvahiLookupFlags flags,
    AvahiServiceBrowserCallback callback,
    uintptr_t ptr)
{
    return avahi_service_browser_new(
        client,
        index,
		proto,
        serviceType,
        domain,
        flags,
		callback, // callback
		(void*) ptr);                       // userdata
}

