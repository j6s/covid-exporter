package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Country struct {
	Country     string
	CountryCode string
	Slug        string
	Date        string

	NewConfirmed   int
	TotalConfirmed int
	NewDeaths      int
	TotalDeaths    int
	NewRecovered   int
	TotalRecovered int
}

type SummaryResponse struct {
	Message   string
	Countries []Country
	Global    struct {
		NewConfirmed   int
		TotalConfirmed int
		NewDeaths      int
		TotalDeaths    int
		NewRecovered   int
		TotalRecovered int
	}
}

var (
	listen = flag.String("listen", ":9084", "Host and port to listen on")
)

const (
	endpoint = "https://api.covid19api.com/summary"
)

func getData() (error, SummaryResponse) {
	data := SummaryResponse{}

	response, err := http.Get(endpoint)
	if err != nil {
		return err, data
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err, data
	}

	err = json.Unmarshal(body, &data)

	return nil, data
}

func gatherMetrics() *prometheus.Registry {
	cases := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: "covid_cases"},
		[]string{"status", "country"},
	)

	registry := prometheus.NewRegistry()
	registry.Register(cases)

	err, data := getData()
	if err != nil {
		// TODO handle error
		log.Fatal(err)
	}

	for _, country := range data.Countries {
		cases.WithLabelValues("confirmed", country.CountryCode).Set(float64(country.TotalConfirmed))
		cases.WithLabelValues("dead", country.CountryCode).Set(float64(country.TotalDeaths))
		cases.WithLabelValues("recovered", country.CountryCode).Set(float64(country.TotalRecovered))
	}

	return registry
}

func main() {
	flag.Parse()

	http.HandleFunc("/metrics", func(response http.ResponseWriter, request *http.Request) {
		registry := gatherMetrics()

		promhttp.HandlerFor(
			registry,
			promhttp.HandlerOpts{},
		).ServeHTTP(response, request)
	})

	log.Printf("Starting to listen on %s", *listen)
	log.Fatal(http.ListenAndServe(*listen, nil))
}
