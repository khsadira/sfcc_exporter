package main

type Metrics struct {
	Site                  string
	PromoEnable           int
	PromoDisable          int
	CouponEnable          int
	CouponDisable         int
	CampaignEnable        int
	CampaignDisable       int
	OrderComplete         int
	OrderCompleteLastFive int
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
