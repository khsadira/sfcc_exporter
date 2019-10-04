package main

type Metrics struct {
	Site               string
	PromotionEnable    int
	PromotionDisable   int
	CouponEnable       int
	CouponDisable      int
	CampaignEnable     int
	CampaignDisable    int
	OrderComplete      int
	OrderCompleteToday int
}

type Scan struct {
	Total int `json:"total"`
	Count int `json:"count"`
}

type JsOrders struct {
	Hits []struct {
		Data struct {
			LastModified string `json:"last_modified"`
		} `json:"data"`
	} `json:"hits"`
}

type Sites struct {
	Data []struct {
		Type string `json:"_type"`
		ID   string `json:"id"`
	} `json:"data"`
	Total int `json:"total"`
	Count int `json:"count"`
}
