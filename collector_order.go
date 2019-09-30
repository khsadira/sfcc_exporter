package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func getOrderMetrics(metric *Metrics, target string, token string, c chan bool) {
	var scan Scan
	var jsOrders JsOrders

	bufExported, _ := getOrderJSON(target, token, "exported", 0)
	json.Unmarshal(bufExported, &scan)
	json.Unmarshal(bufExported, &jsOrders)

	total := findNbOrderLastFiveMin(jsOrders, 0, 200, scan.Total, target, token)
	metric.OrderComplete = scan.Total
	metric.OrderCompleteToday = total

	c <- true
}

func getOrderJSON(target string, token string, search string, start int) ([]byte, error) {
	client := &http.Client{}
	jsBody := fmt.Sprintf(`{"query":{"text_query":{"fields":["export_status"],"search_phrase":"%s"}},"select":"(**)","count":200,"start":%d,"sorts":[{"field":"last_modified","sort_order":"desc"}]}`, search, start)
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

func findNbOrderLastFiveMin(scan JsOrders, start int, count int, total int, target string, token string) int {

	var ret int
	var last  bool

	date := time.Now().Unix() - (5 * 60)
	for _, a := range scan.Hits {
		s, _ := time.Parse(time.RFC3339, a.Data.LastModified)
		if s.Unix() >= date {
			ret += 1
			last = true
		} else {
			last = false
		}
	}
	if last == true && total > start+count {
		//find next 200 by start
		var jsOrders JsOrders

		start += count
		bufExported, _ := getOrderJSON(target, token, "exported", start)
		json.Unmarshal(bufExported, &jsOrders)
		ret += findNbOrderLastFiveMin(jsOrders, start, count, total, target, token)
	}
	return ret
}
