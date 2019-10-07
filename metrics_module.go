package main

import (
	"net/http"
)

func metricsPage(w http.ResponseWriter, r *http.Request) {
	var resp []byte
	/*
		query, ok := r.URL.Query()["site"]

		if !ok || len(query[0]) < 0 {
			log.Println("Url param 'site' is missing (by example: /metrics/?site={id}")
			resp = []byte("No report query found")
		} else {
			resp = getMetricsSFCC(query)
		}*/
	resp = getMetricsSFCC()
	w.Write(resp)
}
