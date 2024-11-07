package sharedServices

import (
	"net"
	"regexp"

	vals "github.com/sty-holdings/sharedServices/v2024/validators"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// GetIPAddresses - returns a list of all IP addresses
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetIPAddresses() (ipList []string) {

	var (
		err            error
		tNetAddr       []net.Addr
		tNetInterfaces []net.Interface
	)

	if tNetInterfaces, err = net.Interfaces(); err == nil {
		// handle err
		for _, i := range tNetInterfaces {
			tNetAddr, err = i.Addrs()
			// handle err
			for _, addr := range tNetAddr {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}
				ipList = append(ipList, ip.String())
			}
		}
	}

	return
}

// GetIPv4Addresses - returns a list of all IPv4 addresses
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetIPv4Addresses() (ipV4List []string) {

	for _, ipAddress := range GetIPAddresses() {
		if vals.IsIPv4Valid(ipAddress) {
			ipV4List = append(ipV4List, ipAddress)
		}
	}

	return
}

// GetIPv6Addresses - returns a list of all IPv6 addresses
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetIPv6Addresses() (ipV6List []string) {

	for _, ipAddress := range GetIPAddresses() {
		if vals.IsIPv6Valid(ipAddress) {
			ipV6List = append(ipV6List, ipAddress)
		}
	}

	return
}
