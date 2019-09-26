package main

func fillMetrics(target string, token string) Metrics {
	var metric Metrics

	c := make(chan bool, 4)
	go getPromoMetrics(&metric, target, token, c)
	go getCouponMetrics(&metric, target, token, c)
	go getCampaignMetrics(&metric, target, token, c)
	go getOrderMetrics(&metric, target, token, c)

	for i := 0; i < 4; i++ {
		<-c
	}
	metric.Site = target
	return metric
}
