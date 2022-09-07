package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
)

// MX Lookups
func mxHandler(response http.ResponseWriter, request *http.Request) {
	responseMap := map[string]interface{}{}
	query := getQuery(request)
	if len(query) == 0 {
		response.WriteHeader(400)
		return
	}
	responseMap["answer"] = getMXRecords(query)
	writeJsonResponse(response, responseMap)
}

func getMXRecords(query string) []string {
	hosts := []string{}
	mxs, error := net.LookupMX(query)
	if error == nil {
		for _, mx := range mxs {
			hosts = append(hosts, trimTrailingDot(mx.Host))
		}
	}
	return hosts
}

// CNAME Lookups
func cnameHandler(response http.ResponseWriter, request *http.Request) {
	responseMap := map[string]interface{}{}
	query := getQuery(request)
	if len(query) == 0 {
		response.WriteHeader(400)
		return
	}
	responseMap["answer"] = getCNAMERecord(query)
	writeJsonResponse(response, responseMap)
}

func getCNAMERecord(query string) string {
	cname, _ := net.LookupCNAME(query)
	return trimTrailingDot(cname)
}

// IP Lookups
func ipHandler(response http.ResponseWriter, request *http.Request) {
	responseMap := map[string]interface{}{}
	query := getQuery(request)
	if len(query) == 0 {
		response.WriteHeader(400)
		return
	}
	responseMap["answer"] = getIPRecord(query)
	writeJsonResponse(response, responseMap)
}

func getIPRecord(query string) []string {
	ips := []string{}
	results, error := net.LookupIP(query)
	if error == nil {
		for _, result := range results {
			ips = append(ips, result.String())
		}
	}

	return ips
}

// Reverse IP Lookups
func reverseIPHandler(response http.ResponseWriter, request *http.Request) {
	responseMap := map[string]interface{}{}
	query := getQuery(request)
	if len(query) == 0 {
		response.WriteHeader(400)
		return
	}
	responseMap["answer"] = getReverseIPRecord(query)
	writeJsonResponse(response, responseMap)
}

func getReverseIPRecord(query string) []string {
	hostnames := []string{}
	results, error := net.LookupAddr(query)
	if error == nil {
		for _, result := range results {
			hostnames = append(hostnames, trimTrailingDot(result))
		}
	}

	return hostnames
}

// Web Helpers
func getQuery(request *http.Request) string {
	query := ""
	if request.Method == "POST" {
		request.ParseForm()
		query = request.Form.Get("q")
	} else if request.Method == "GET" {
		query = request.URL.Query().Get("q")
	}

	return query
}

func writeJsonResponse(response http.ResponseWriter, data map[string]interface{}) {
	responseJson, _ := json.Marshal(data)
	response.Write([]byte(responseJson))
}

// Formatting Helpers
func trimTrailingDot(hostname string) string {
	cleaned := hostname
	last := len(hostname) - 1
	if last >= 0 && hostname[last] == '.' {
		cleaned = hostname[:last]
	}

	return cleaned
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/mx", mxHandler)
	mux.HandleFunc("/cname", cnameHandler)
	mux.HandleFunc("/ip", ipHandler)
	mux.HandleFunc("/reverse", reverseIPHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
