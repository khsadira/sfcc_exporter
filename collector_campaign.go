package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getCampaignMetrics(metric *Metrics, target string, token string, c chan bool) {

	bufTrue, _ := getCampaignJSON(target, token, "true")
	bufFalse, _ := getCampaignJSON(target, token, "false")
	var scan Scan

	json.Unmarshal(bufTrue, &scan)
	(*metric).CampaignEnable = scan.Total

	json.Unmarshal(bufFalse, &scan)
	(*metric).CampaignDisable = scan.Total
	c <- true
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
