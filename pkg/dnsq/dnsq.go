package dnsq

import "net"

// Formatting Helpers
func trimTrailingDot(hostname string) string {
	cleaned := hostname
	last := len(hostname) - 1
	if last >= 0 && hostname[last] == '.' {
		cleaned = hostname[:last]
	}

	return cleaned
}

func GetMXRecords(query string) []string {
	hosts := []string{}
	mxs, error := net.LookupMX(query)
	if error == nil {
		for _, mx := range mxs {
			hosts = append(hosts, trimTrailingDot(mx.Host))
		}
	}
	return hosts
}

func GetCNAMERecord(query string) []string {
	cname, _ := net.LookupCNAME(query)
	return []string{trimTrailingDot(cname)}
}

func GetIPRecord(query string) []string {
	ips := []string{}
	results, error := net.LookupIP(query)
	if error == nil {
		for _, result := range results {
			ips = append(ips, result.String())
		}
	}

	return ips
}

func GetReverseIPRecord(query string) []string {
	hostnames := []string{}
	results, error := net.LookupAddr(query)
	if error == nil {
		for _, result := range results {
			hostnames = append(hostnames, trimTrailingDot(result))
		}
	}

	return hostnames
}
