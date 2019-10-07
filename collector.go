package main

func fillMetrics(target string, token string) Metrics {
	var metrics Metrics

	c := make(chan bool, 4)
	go getOrderMetrics(&metrics, target, token, c)
	go getPromoMetrics(&metrics, target, token, c)
	go getCouponMetrics(&metrics, target, token, c)
	go getCampaignMetrics(&metrics, target, token, c)

	for i := 0; i < 4; i++ {
		<-c
	}
	metrics.Site = target
	return metrics
}
