package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/metrics/", metricsPage)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
<head><title>Sfcc Exporter</title></head>
			 <body>
			 <h1>Sfcc Exporter</h1>
			 <p><a href='/metrics/'>sfcc metrics</a></p>
			 </body>
			 </html>`))
	})
	log.Fatal(http.ListenAndServe(":9240", nil))
}
