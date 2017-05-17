package driver

type Storager interface {
	ListConfigs(prefix string) []ServerConfigResult
	AddConfig(key, value string)
	RemoveConfig(key, value string)
	GetConfig(key string) []ServerConfigResult
	CountConfigKey(key string) int

	AddData(name, dtype, value, description string) error
	ListData(map[string]string) []DataResult
	RemoveData(name string) error
	AddDataToProfile(profile, data string)
	RemoveDataFromProfile(profile, data string)
	CountData() int

	AddProfile(name string) error
	ListProfile(filter map[string]string) map[string][]string
	RemoveProfile(name string) error
	CountProfile() int

	AddIP(ipaddr, netmask, gateway string) error
	ListIP(filter map[string]string) []IPResult
	RemoveIP(ipaddr string) error
	CountIP() int

	AddInterface(index int, mac, ip string) error
	ListInterface(filter map[string]string) []InterfaceResult
	RemoveInterface(mac string) error
	CountInterface() int

	AddHost(enable bool, name, fqdn, uuid, profile string, interfaces []string) error
	ListHost(filter map[string]string) []HostResult
	RemoveHost(name string) error
	EnableHost(name string) error
	DisableHost(name string) error
	CountHost() int

	GetIDFromIP(ip string) int

	GetID(srvid int) int
	GetHostname(srvid int) string
	GetFQDN(srvid int) string
	GetUserData(srvid int) string
	GetVendorData(srvid int) string
	GetPublicKeys(srvid int) []string
	GetRegion(srvid int) string
	GetInterfaces(srvid int) []string
	GetInterfacesType(srvid int, itype string) []int
	GetEnumeratedInterface(srvid int, itype string, index int) []string
	GetInterfaceMACAddress(srvid int, itype string, index int) string
	GetInterfaceType(srvid int, itype string, index int) string
	GetInterfaceIPv4Address(srvid int, itype string, index int) string
	GetInterfaceIPv4Netmask(srvid int, itype string, index int) string
	GetInterfaceIPv4Gateway(srvid int, itype string, index int) string
	GetDNSIndex(srvid int) []string
	GetDNSNameservers(srvid int) []string
	GetDNSSearchDomains(srvid int) []string
	GetDNSOptions(srvid int) []string
	GetTags(srvid int) []string
	GetKeys(srvid int) []string
	GetKey(srvid int, key string) string

	End()
}
