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
	log.Fatal(http.ListenAndServe(":8080", mux))
}
