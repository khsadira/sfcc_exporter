package main

type Metrics struct {
	Site              string
	PromotionEnabled  int
	PromotionDisabled int
	PromotionTotal    int
	CouponEnabled     int
	CouponDisabled    int
	CouponTotal       int
}

type Scan struct {
	Total int `json:"total"`
}
