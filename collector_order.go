package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getOrderMetrics(metric *Metrics, target string, token string, c chan bool) {
	var scan Scan

	bufExported, _ := getOrderJSON(target, token, "exported")
	json.Unmarshal(bufExported, &scan)
	metric.OrderComplete = scan.Total

	c <- true
}

func getOrderJSON(target string, token string, search string) ([]byte, error) {
	client := &http.Client{}
	jsBody := fmt.Sprintf(`{"query":{"text_query":{"fields":["export_status"],"search_phrase":"%s"}},"select":"(**)","count":1}`, search)
	jsonBody := []byte(jsBody)
	query := fmt.Sprintf("%s/s/%s/dw/shop/v19_8/order_search", hostname, target)
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