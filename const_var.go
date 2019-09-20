package main

const (
	//PROMOTION VAR
	namePromo        = "sfcc_promotions"
	namePromoEnable  = namePromo + "_enabled"
	namePromoDisable = namePromo + "_disabled"
	namePromoTotal   = namePromo + "_total"
	helpPromoEnable  = "# HELP " + namePromoEnable + " Number of promotions enabled"
	typePromoEnable  = "# TYPE " + namePromoEnable + " gauge"
	helpPromoDisable = "# HELP " + namePromoDisable + " Number of promotions disabled"
	typePromoDisable = "# TYPE " + namePromoDisable + " gauge"
	helpPromoTotal   = "# HELP " + namePromoTotal + " Total number of promotions"
	typePromoTotal   = "# TYPE " + namePromoTotal + " gauge"

	//COUPON VAR
	nameCoupon        = "sfcc_coupons"
	nameCouponEnable  = nameCoupon + "_enabled"
	nameCouponDisable = nameCoupon + "_disabled"
	nameCouponTotal   = nameCoupon + "_total"
	helpCouponEnable  = "# HELP " + nameCouponEnable + " Number of coupons enabled"
	typeCouponEnable  = "# TYPE " + nameCouponEnable + " gauge"
	helpCouponDisable = "# HELP " + nameCouponDisable + " Number of coupons disabled"
	typeCouponDisable = "# TYPE " + nameCouponDisable + " gauge"
	helpCouponTotal   = "# HELP " + nameCouponTotal + " Total number of coupons"
	typeCouponTotal   = "# TYPE " + nameCouponTotal + " gauge"
)
