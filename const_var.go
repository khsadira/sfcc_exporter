package main

const (
	namePromo        = "sfcc_promotions"
	namePromoEnable  = namePromo + "_enabled"
	namePromoDisable = namePromo + "_disabled"
	namePromoTotal   = namePromo + "_total"
	helpPromoEnable  = "# HELP " + namePromoEnable + " Number of promotion enabled"
	typePromoEnable  = "# TYPE " + namePromoEnable + " gauge"
	helpPromoDisable = "# HELP " + namePromoDisable + " Number of promotion disabled"
	typePromoDisable = "# TYPE " + namePromoDisable + " gauge"
	helpPromoTotal   = "# HELP " + namePromoTotal + " Total number of promotion"
	typePromoTotal   = "# TYPE " + namePromoTotal + " gauge"
)
