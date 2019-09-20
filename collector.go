package main

import (
	"log"
)

func fillMetrics(target string, token string) Metrics {
	var metric Metrics

	metric, err := getPromoMetrics(metric, target, token)
	if err != nil {
		log.Println(err)
	}
	metric, err = getCouponMetrics(metric, target, token)
	if err != nil {
		log.Println(err)
	}
	metric, err = getCampaignMetrics(metric, target, token)
	if err != nil {
		log.Println(err)
	}
	metric.Site = target
	return metric
}
