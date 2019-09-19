package main

const (
	namePromo      = "sfcc_promotion"
	namePromoTrue  = namePromo + "_true"
	namePromoFalse = namePromo + "_false"
	namePromoTotal = namePromo + "_total"
	helpPromoTrue  = "# HELP " + namePromoTrue + ""
	typePromoTrue  = "# TYPE " + namePromoTrue + " gauge"
	helpPromoFalse = "# HELP " + namePromoFalse + ""
	typePromoFalse = "# TYPE " + namePromoFalse + " gauge"
	helpPromoTotal = "# HELP " + namePromoTotal + ""
	typePromoTotal = "# TYPE " + namePromoTotal + " gauge"
)
