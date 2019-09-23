package main

func fillMetrics(target string, token string) Metrics {
	var metric Metrics

	c := make(chan bool, 3)
	go getPromoMetrics(&metric, target, token, c)
	go getCouponMetrics(&metric, target, token, c)
	go getCampaignMetrics(&metric, target, token, c)

	for i := 0; i < 3; i++ {
		<-c
	}
	metric.Site = target
	return metric
}
