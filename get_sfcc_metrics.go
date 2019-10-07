package main

import (
	"fmt"
	"log"
)

func getMetricsSFCC() []byte {
	//	var metrics []Metrics

	token, err := getToken("CLIENT_ID_SFCC", "CLIENT_PW_SFCC")
	if err != nil {
		log.Println(err)
		return []byte("Access token wasn't generated")
	}
	sites, len := getSiteMetrics(token)
	var met [VAL]Metrics
	b := make(chan int, len)
	for i, target := range sites {
		//	metrics = append(metrics, fillMetrics(target, token))
		go func(metrics *Metrics, target string, token string, b chan int) {
			*(metrics) = fillMetrics(target, token)
			b <- 1
		}(&met[i], target, token, b)
	}
	for i := 0; i < len; i++ {
		<-b
	}
	resp := metricsToByte(met, len)
	return resp
}

func metricsToByte(metrics [VAL]Metrics, len int) []byte {
	var resp string
	var promoTotal, promoEnable, promoDisable string
	var couponTotal, couponEnable, couponDisable string
	var campaignTotal, campaignEnable, campaignDisable string
	var orderComplete, orderCompleteToday string

	for i, metric := range metrics {
		if i == len {
			break
		}
		//PROMO_VAR
		promoTotal += fmt.Sprintf("%s{site=\"%s\"} %v\n", namePromoTotal, metric.Site, metric.PromotionEnable+metric.PromotionDisable)
		promoEnable += fmt.Sprintf("%s{site=\"%s\"} %v\n", namePromoEnable, metric.Site, metric.PromotionEnable)
		promoDisable += fmt.Sprintf("%s{site=\"%s\"} %v\n", namePromoDisable, metric.Site, metric.PromotionDisable)

		//COUPON_VAR
		couponTotal += fmt.Sprintf("%s{site=\"%s\"} %v\n", nameCouponTotal, metric.Site, metric.CouponEnable+metric.CouponDisable)
		couponEnable += fmt.Sprintf("%s{site=\"%s\"} %v\n", nameCouponEnable, metric.Site, metric.CouponEnable)
		couponDisable += fmt.Sprintf("%s{site=\"%s\"} %v\n", nameCouponDisable, metric.Site, metric.CouponDisable)

		//CAMPAIGN_VAR
		campaignTotal += fmt.Sprintf("%s{site=\"%s\"} %v\n", nameCampaignTotal, metric.Site, metric.CampaignEnable+metric.CampaignDisable)
		campaignEnable += fmt.Sprintf("%s{site=\"%s\"} %v\n", nameCampaignEnable, metric.Site, metric.CampaignEnable)
		campaignDisable += fmt.Sprintf("%s{site=\"%s\"} %v\n", nameCampaignDisable, metric.Site, metric.CampaignDisable)

		//ORDER_VAR
		orderComplete += fmt.Sprintf("%s{site=\"%s\"} %v\n", nameOrderComplete, metric.Site, metric.OrderComplete)
		orderCompleteToday += fmt.Sprintf("%s{site=\"%s\"} %v\n", nameOrderCompleteToday, metric.Site, metric.OrderCompleteToday)

	}
	resp = fmt.Sprintf("%s\n%s\n%s%s\n%s\n%s%s\n%s\n%s", helpPromoTotal, typePromoTotal, promoTotal, helpPromoEnable, typePromoEnable, promoEnable, helpPromoDisable, typePromoDisable, promoDisable)
	resp += fmt.Sprintf("%s\n%s\n%s%s\n%s\n%s%s\n%s\n%s", helpCouponTotal, typeCouponTotal, couponTotal, helpCouponEnable, typeCouponEnable, couponEnable, helpCouponDisable, typeCouponDisable, couponDisable)
	resp += fmt.Sprintf("%s\n%s\n%s%s\n%s\n%s%s\n%s\n%s", helpCampaignTotal, typeCampaignTotal, campaignTotal, helpCampaignEnable, typeCampaignEnable, campaignEnable, helpCampaignDisable, typeCampaignDisable, campaignDisable)
	resp += fmt.Sprintf("%s\n%s\n%s", helpOrderComplete, typeOrderComplete, orderComplete)
	resp += fmt.Sprintf("%s\n%s\n%s", helpOrderCompleteToday, typeOrderCompleteToday, orderCompleteToday)
	return []byte(resp)
}
