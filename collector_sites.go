package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getSiteMetrics(token string) []string {
	var sites []string

	sites = findSites(sites, token, 0, 200)
	return sites
}

func findSites(sites []string, token string, start int, count int) []string {
	buf, _ := getSiteJSON(token, start, count)

	var data Sites
	json.Unmarshal(buf, &data)

	for i := 0; i < data.Count; i++ {
		sites = append(sites, data.Data[i].ID)
	}
	if data.Total >= start + data.Count {
		sites = findSites(sites, token, start + count, count)
	}
	return sites
}

func getSiteJSON(token string, start int, count int) ([]byte, error) {
	client := &http.Client{}
	query := fmt.Sprintf("%s/s/-/dw/data/v19_8/sites?start=%d&count=%d", hostname, start, count)
	req, err := http.NewRequest("GET", query, nil)
	req.Header.Add("Authorization", "Bearer "+token)
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