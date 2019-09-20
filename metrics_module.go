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
	var promoTotal, promoEnable, promoDisable string
	var couponTotal, couponEnable, couponDisable string

	for _, metric := range metrics {
		//PROMO_VAR
		promoTotal += fmt.Sprintf("%s{site=\"%s\"} %v\n", namePromoTotal, metric.Site, metric.PromotionEnabled+metric.PromotionDisabled)
		promoEnable += fmt.Sprintf("%s{site=\"%s\"} %v\n", namePromoEnable, metric.Site, metric.PromotionEnabled)
		promoDisable += fmt.Sprintf("%s{site=\"%s\"} %v\n", namePromoDisable, metric.Site, metric.PromotionDisabled)

		//COUPON_VAR
		couponTotal += fmt.Sprintf("%s{site=\"%s\"} %v\n", nameCouponTotal, metric.Site, metric.CouponEnabled+metric.CouponDisabled)
		couponEnable += fmt.Sprintf("%s{site=\"%s\"} %v\n", nameCouponEnable, metric.Site, metric.CouponEnabled)
		couponDisable += fmt.Sprintf("%s{site=\"%s\"} %v\n", nameCouponDisable, metric.Site, metric.CouponDisabled)
	}
	resp = fmt.Sprintf("%s\n%s\n%s%s\n%s\n%s%s\n%s\n%s", helpPromoTotal, typePromoTotal, promoTotal, helpPromoEnable, typePromoEnable, promoEnable, helpPromoDisable, typePromoDisable, promoDisable)
	resp += fmt.Sprintf("%s\n%s\n%s%s\n%s\n%s%s\n%s\n%s", helpCouponTotal, typeCouponTotal, couponTotal, helpCouponEnable, typeCouponEnable, couponEnable, helpCouponDisable, typeCouponDisable, couponDisable)
	return []byte(resp)
}
