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
	var resp string

	for _, metric := range metrics {
		resp += fmt.Sprintf("%s\n%s\n%s{site=\"%s\"} %v\n", helpPromoTotal, typePromoTotal, namePromoTotal, metric.Site, metric.PromotionEnabled+metric.PromotionDisabled)
		resp += fmt.Sprintf("%s\n%s\n%s{site=\"%s\"} %v\n", helpPromoEnable, typePromoEnable, namePromoEnable, metric.Site, metric.PromotionEnabled)
		resp += fmt.Sprintf("%s\n%s\n%s{site=\"%s\"} %v\n", helpPromoDisable, typePromoDisable, namePromoDisable, metric.Site, metric.PromotionDisabled)
	}
	return []byte(resp)
}
