package app

import (
	"dnsq/pkg/dnsq"
	"encoding/json"
	"log"
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
	responseMap["answer"] = dnsq.GetMXRecords(query)
	writeJsonResponse(response, responseMap)
}

// CNAME Lookups
func cnameHandler(response http.ResponseWriter, request *http.Request) {
	responseMap := map[string]interface{}{}
	query := getQuery(request)
	if len(query) == 0 {
		response.WriteHeader(400)
		return
	}
	responseMap["answer"] = dnsq.GetCNAMERecord(query)
	writeJsonResponse(response, responseMap)
}

// IP Lookups
func ipHandler(response http.ResponseWriter, request *http.Request) {
	responseMap := map[string]interface{}{}
	query := getQuery(request)
	if len(query) == 0 {
		response.WriteHeader(400)
		return
	}
	responseMap["answer"] = dnsq.GetIPRecord(query)
	writeJsonResponse(response, responseMap)
}

// Reverse IP Lookups
func reverseIPHandler(response http.ResponseWriter, request *http.Request) {
	responseMap := map[string]interface{}{}
	query := getQuery(request)
	if len(query) == 0 {
		response.WriteHeader(400)
		return
	}
	responseMap["answer"] = dnsq.GetReverseIPRecord(query)
	writeJsonResponse(response, responseMap)
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

func Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/mx", mxHandler)
	mux.HandleFunc("/cname", cnameHandler)
	mux.HandleFunc("/ip", ipHandler)
	mux.HandleFunc("/reverse", reverseIPHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
