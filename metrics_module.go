package main

import (
	"fmt"
	"log"
	"net/http"
)

func metricsPage(w http.ResponseWriter, r *http.Request) {
	var resp []byte

	query, ok := r.URL.Query()["site"]

	if !ok || len(query[0]) < 0 {
		log.Println("Url param 'site' is missing (by example: /metrics/?site={id}")
		resp = []byte("No report query found")
	} else {
		resp = getMetricsSFCC(query)
	}
	w.Write(resp)
}

func getMetricsSFCC(query []string) []byte {
	var metrics []Metrics

	token, err := getToken("CLIENT_ID_SFCC", "CLIENT_PW_SFCC")
	if err != nil {
		log.Println(err)
		return []byte("Access token wasn't generated")
	}
	for _, target := range query {
		metrics = append(metrics, fillMetrics(target, token))
	}
	resp := metricsToByte(metrics)
	return (resp)
}

func metricsToByte(metrics []Metrics) []byte {
	var resp, respTotal, respEnable, respDisable string

	for _, metric := range metrics {
		respTotal += fmt.Sprintf("%s{site=\"%s\"} %v\n", namePromoTotal, metric.Site, metric.PromotionEnabled+metric.PromotionDisabled)
		respEnable += fmt.Sprintf("%s{site=\"%s\"} %v\n", namePromoEnable, metric.Site, metric.PromotionEnabled)
		respDisable += fmt.Sprintf("%s{site=\"%s\"} %v\n", namePromoDisable, metric.Site, metric.PromotionDisabled)
	}
	resp = fmt.Sprintf("%s\n%s\n%s%s\n%s\n%s%s\n%s\n%s", helpPromoTotal, typePromoTotal, respTotal, helpPromoEnable, typePromoEnable, respEnable, helpPromoDisable, typePromoDisable, respDisable)
	return []byte(resp)
}
