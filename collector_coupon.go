package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getCouponMetrics(metric *Metrics, target string, token string, c chan bool) {

	bufTrue, _ := getCouponJSON(target, token, "true")
	bufFalse, _ := getCouponJSON(target, token, "false")
	var scan Scan

	json.Unmarshal(bufTrue, &scan)
	(*metric).CouponEnable = scan.Total

	json.Unmarshal(bufFalse, &scan)
	(*metric).CouponDisable = scan.Total
	c <- true
}

func getCouponJSON(target string, token string, search string) ([]byte, error) {
	client := &http.Client{}
	jsBody := fmt.Sprintf(`{"query":{"text_query":{"fields":["enabled"],"search_phrase":"%s"}},"select":"(**)","count":1}`, search)
	jsonBody := []byte(jsBody)
	query := fmt.Sprintf("%s/s/-/dw/data/v19_8/sites/%s/coupon_search", hostname, target)
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
