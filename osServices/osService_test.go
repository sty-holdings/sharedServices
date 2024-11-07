package sharedServices

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGetIPAddresses(tPtr *testing.T) {
	var (
		ipAddresses        []string
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(
		tFunctionName, func(tPtr *testing.T) {
			// Success
			if ipAddresses = GetIPAddresses(); len(ipAddresses) == 0 {
				tPtr.Errorf("%v FAILED - THIS IS NOT WORKING. THERE SHOULD ALWAYS BE AN IP ADDRESS.", tFunctionName)
			} else {
				for _, address := range ipAddresses {
					fmt.Println("IP: ", address)
				}
			}
		},
	)
}

func TestGetIPv4Addresses(tPtr *testing.T) {
	var (
		ipAddresses        []string
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(
		tFunctionName, func(tPtr *testing.T) {
			// Success
			if ipAddresses = GetIPv4Addresses(); len(ipAddresses) == 0 {
				tPtr.Errorf("%v FAILED - THIS IS NOT WORKING. THERE SHOULD ALWAYS BE AN IPv4 ADDRESS.", tFunctionName)
			} else {
				for _, address := range ipAddresses {
					fmt.Println("IP: ", address)
				}
			}
		},
	)
}

func TestGetIPv6Addresses(tPtr *testing.T) {
	var (
		ipAddresses        []string
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(
		tFunctionName, func(tPtr *testing.T) {
			// Success
			if ipAddresses = GetIPv6Addresses(); len(ipAddresses) == 0 {
				fmt.Printf("%v WARNING - Your system may not have an IPv6 assigned.", tFunctionName)
			} else {
				for _, address := range ipAddresses {
					fmt.Println("IP: ", address)
				}
			}
		},
	)
}
