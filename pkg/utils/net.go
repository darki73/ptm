package utils

import "net"

// IsValidIPWithSubnet checks if the given string is a valid IP with subnet.
func IsValidIPWithSubnet(address string) bool {
	_, _, err := net.ParseCIDR(address)
	return err == nil
}

// IsValidIP checks if the given string is a valid IP.
func IsValidIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil
}

// IsInSameNetwork checks if the given IP address is in the same network.
func IsInSameNetwork(ip string, network string) bool {
	parsedIP := net.ParseIP(ip)
	_, parsedNetwork, err := net.ParseCIDR(network)
	if err != nil {
		return false
	}
	return parsedNetwork.Contains(parsedIP)
}
