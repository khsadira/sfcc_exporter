package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Scan struct {
	Total int `json:"total"`
}

func getPromoMetrics(target string, token string) (Metrics, error) {
	var metric Metrics

	bufTrue, err := getPromoJSON(target, token, "true")
	if err != nil {
		return metric, err
	}
	bufFalse, err := getPromoJSON(target, token, "false")
	if err != nil {
		return metric, err
	}
	var scan Scan

	json.Unmarshal(bufTrue, &scan)
	metric.PromotionEnabled = scan.Total

	json.Unmarshal(bufFalse, &scan)
	metric.PromotionDisabled = scan.Total
	return metric, nil
}

func getPromoJSON(target string, token string, search string) ([]byte, error) {
	client := &http.Client{}
	jsBody := fmt.Sprintf(`{"query":{"text_query":{"fields":["enabled"],"search_phrase":"%s"}},"select" : "(**)"}`, search)
	jsonBody := []byte(jsBody)
	query := fmt.Sprintf("https://store-dev.ubi.com/s/-/dw/data/v19_8/sites/%s/promotion_search", target)
	req, err := http.NewRequest("POST", query, bytes.NewBuffer(jsonBody))
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return []byte(""), err
	}
	if err != nil {
		log.Println(err)
		return []byte(""), err
	}
	defer resp.Body.Close()
	buf, _ := ioutil.ReadAll(resp.Body)

	return buf, nil
}
