package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const urlStatic = "https://opensky-network.org/api/flights/arrival?airport=EDDF&begin=1517227200&end=1517230800"

//func init() {
//}

type ArrivalRes struct {
	Icao24                           string  `json:"icao24"`
	FirstSeen                        int     `json:"firstSeen"`
	EstDepartureAirport              *string `json:"estDepartureAirport"`
	LastSeen                         int     `json:"lastSeen"`
	EstArrivalAirport                string  `json:"estArrivalAirport"`
	Callsign                         string  `json:"callsign"`
	EstDepartureAirportHorizDistance *int    `json:"estDepartureAirportHorizDistance"`
	EstDepartureAirportVertDistance  *int    `json:"estDepartureAirportVertDistance"`
	EstArrivalAirportHorizDistance   int     `json:"estArrivalAirportHorizDistance"`
	EstArrivalAirportVertDistance    int     `json:"estArrivalAirportVertDistance"`
	DepartureAirportCandidatesCount  int     `json:"departureAirportCandidatesCount"`
	ArrivalAirportCandidatesCount    int     `json:"arrivalAirportCandidatesCount"`
}

func main() {
	aa, _ := getListing(urlStatic)
	for _, a := range aa {
		fmt.Println(a)
	}
}

func getListing(url string) ([]ArrivalRes, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	req.Header.Add("Authorization", "Basic TXVmTXVmOjNxZlA0NmV6QHZLZjR5IQ==")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return deserialize(body)
}
func deserialize(res []byte) ([]ArrivalRes, error) {
	var data []ArrivalRes
	if err := json.Unmarshal(res, &data); err != nil {
		fmt.Println("failed to unmarshal:", err)
		return nil, err
	}
	return data, nil
}

//https://medium.com/zus-health/mocking-outbound-http-requests-in-go-youre-probably-doing-it-wrong-60373a38d2aa
