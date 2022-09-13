package app

import (
	"dnsq/pkg/dnsq"
	"encoding/json"
	"log"
	"net/http"
)

type fn func(string) []string

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

func genericHandler(f fn) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		responseMap := map[string]interface{}{}
		query := getQuery(request)
		if len(query) == 0 {
			response.WriteHeader(400)
			return
		}
		responseMap["answer"] = f(query)
		writeJsonResponse(response, responseMap)
	})
}

func Run() {
	mux := http.NewServeMux()
	mux.Handle("/mx", genericHandler(dnsq.GetMXRecords))
	mux.Handle("/cname", genericHandler(dnsq.GetCNAMERecord))
	mux.Handle("/ip", genericHandler(dnsq.GetIPRecord))
	mux.Handle("/reverse", genericHandler(dnsq.GetReverseIPRecord))
	log.Fatal(http.ListenAndServe(":8080", mux))
}
