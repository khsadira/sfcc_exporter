package main

import (
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
	var orderComplete, orderCompleteLastFive string

	for i, metric := range metrics {
		if i == len {
			break
		}
		//PROMO_VAR
		promoTotal += promMetrics(namePromoTotal, metric.Site, metric.PromoEnable+metric.PromoDisable)
		promoEnable += promMetrics(namePromoEnable, metric.Site, metric.PromoEnable)
		promoDisable += promMetrics(namePromoDisable, metric.Site, metric.PromoDisable)

		//COUPON_VAR
		couponTotal += promMetrics(nameCouponTotal, metric.Site, metric.CouponEnable+metric.CouponDisable)
		couponEnable += promMetrics(nameCouponEnable, metric.Site, metric.CouponEnable)
		couponDisable += promMetrics(nameCouponDisable, metric.Site, metric.CouponDisable)

		//CAMPAIGN_VAR
		campaignTotal += promMetrics(nameCampaignTotal, metric.Site, metric.CampaignEnable+metric.CampaignDisable)
		campaignEnable += promMetrics(nameCampaignEnable, metric.Site, metric.CampaignEnable)
		campaignDisable += promMetrics(nameCampaignDisable, metric.Site, metric.CampaignDisable)

		//ORDER_VAR
		orderComplete += promMetrics(nameOrderComplete, metric.Site, metric.OrderComplete)
		orderCompleteLastFive += promMetrics(nameOrderCompleteLastFive, metric.Site, metric.OrderCompleteLastFive)
	}
	//PROMO_RESP
	resp += promDesc(namePromoTotal, "Total numbers of promotions", "gauge", promoTotal)
	resp += promDesc(namePromoEnable, "Numbers of promotions enabled", "gauge", promoEnable)
	resp += promDesc(namePromoDisable, "Numbers of promotions disabled", "gauge", promoDisable)

	//COUPON_RESP
	resp += promDesc(nameCouponTotal, "Total numbers of coupons", "gauge", couponTotal)
	resp += promDesc(nameCouponEnable, "Numbers of coupons enabled", "gauge", couponEnable)
	resp += promDesc(nameCouponDisable, "Numbers of coupons disabled", "gauge", couponDisable)

	//CAMPAIGN_RESP
	resp += promDesc(nameCampaignTotal, "Total numbers of campaigns", "gauge", campaignTotal)
	resp += promDesc(nameCampaignEnable, "Numbers of campaigns enabled", "gauge", campaignEnable)
	resp += promDesc(nameCampaignDisable, "Numbers of campaigns disabled", "gauge", campaignDisable)

	//ORDER_RESP
	resp += promDesc(nameOrderComplete, "Numbers of orders completed", "gauge", orderComplete)
	resp += promDesc(nameOrderCompleteLastFive, "Numbers of orders completed since five minutes", "gauge", orderCompleteLastFive)

	return []byte(resp)
}
