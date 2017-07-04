#pragma once

#include <avahi-common/cdecl.h>
#include <avahi-common/defs.h>
#include <avahi-common/watch.h>
#include <avahi-common/error.h>
#include <avahi-client/client.h>
#include <avahi-client/lookup.h>

#include <stddef.h>
#include <stdlib.h>

/////////////////////////////
// Browser stuff
/////////////////////////////

AvahiServiceBrowser*
go_avahi_service_browser_new(
    AvahiClient* client,
    AvahiIfIndex index,
    AvahiProtocol proto,
    const char* serviceType,
    const char* domain,
    AvahiLookupFlags flags,
    AvahiServiceBrowserCallback callback,
    uintptr_t ptr);

void
go_service_browser_callback(
    AvahiServiceBrowser *b,
    AvahiIfIndex interface,
    AvahiProtocol protocol,
    AvahiBrowserEvent event,
    char *name, // const!
    char *type, // const!
    char *domain, // const!
    AvahiLookupResultFlags flags,
    uintptr_t userdata);

