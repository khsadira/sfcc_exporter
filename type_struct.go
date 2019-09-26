package main

type Metrics struct {
	Site             string
	PromotionEnable  int
	PromotionDisable int
	CouponEnable     int
	CouponDisable    int
	CampaignEnable   int
	CampaignDisable  int
	OrderComplete    int
}

type Scan struct {
	Total int `json:"total"`
}
