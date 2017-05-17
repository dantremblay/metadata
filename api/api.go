package api

import (
	"github.com/kassisol/metadata/api/metadata"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func API(addr string) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(ServerIP())

	e.GET("/", indexHandle)

	v1 := e.Group("/metadata/v1")

	v1.GET("", metadata.AllHandle)
	v1.GET("/", metadata.AllHandle)
	v1.GET(".json", metadata.AllJsonHandle)

	v1.GET("/id", metadata.IDHandle)
	v1.GET("/hostname", metadata.HostnameHandle)
	v1.GET("/fqdn", metadata.FQDNHandle)
	v1.GET("/user-data", metadata.UserDataHandle)
	v1.GET("/vendor-data", metadata.VendorDataHandle)
	v1.GET("/public-keys", metadata.PublicKeysHandle)
	v1.GET("/region", metadata.RegionHandle)

	v1.GET("/interfaces/", metadata.NetworkInterfacesIndexHandle)
	v1.GET("/interfaces/:type/", metadata.NetworkInterfaceTypeIndexHandle)
	v1.GET("/interfaces/:type/:num/", metadata.NetworkEnumeratedInterfaceIndexHandle)
	v1.GET("/interfaces/:type/:num/mac", metadata.NetworkInterfaceMACAddressHandle)
	v1.GET("/interfaces/:type/:num/type", metadata.NetworkInterfaceTypeHandle)
	v1.GET("/interfaces/:type/:num/ipv4/", metadata.NetworkInterfaceIPv4IndexHandle)
	v1.GET("/interfaces/:type/:num/ipv4/address", metadata.NetworkInterfaceIPv4AddressHandle)
	v1.GET("/interfaces/:type/:num/ipv4/netmask", metadata.NetworkInterfaceIPv4NetmaskHandle)
	v1.GET("/interfaces/:type/:num/ipv4/gateway", metadata.NetworkInterfaceIPv4GatewayHandle)

	v1.GET("/dns/", metadata.DnsIndexHandle)
	v1.GET("/dns/nameservers", metadata.DnsNameserversHandle)
	v1.GET("/dns/searchdomains", metadata.DnsSearchDomainsHandle)
	v1.GET("/dns/options", metadata.DnsOptionsHandle)

	v1.GET("/tags/", metadata.TagsIndexHandle)

	v1.GET("/keys/", metadata.KeysIndexHandle)
	v1.GET("/keys/:name", metadata.KeyNameHandle)

	e.Logger.Fatal(e.Start(addr))
}
