package main

import (
	"log"
)

func fillMetrics(target string, token string) Metrics {
	var metric Metrics

	metric, err := getPromoMetrics(target, token)
	metric.Site = target
	if err != nil {
		log.Println(err)
	}
	return metric
}
