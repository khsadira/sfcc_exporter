package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getCampaignMetrics(metric Metrics, target string, token string) (Metrics, error) {

	bufTrue, err := getCampaignJSON(target, token, "true")
	if err != nil {
		return metric, err
	}
	bufFalse, err := getCampaignJSON(target, token, "false")
	if err != nil {
		return metric, err
	}
	var scan Scan

	json.Unmarshal(bufTrue, &scan)
	metric.CampaignEnable = scan.Total

	json.Unmarshal(bufFalse, &scan)
	metric.CampaignDisable = scan.Total
	return metric, nil
}

func getCampaignJSON(target string, token string, search string) ([]byte, error) {
	client := &http.Client{}
	jsBody := fmt.Sprintf(`{"query":{"text_query":{"fields":["enabled"],"search_phrase":"%s"}},"select" : "(**)"}`, search)
	jsonBody := []byte(jsBody)
	query := fmt.Sprintf("https://store-dev.ubi.com/s/-/dw/data/v19_8/sites/%s/campaign_search", target)
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
