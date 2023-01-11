package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

var urlStatic string

func init() {
	fmt.Println("FOO:", os.Getenv("urlApi"))
	urlStatic = os.Getenv("urlApi")
}

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
	getListing(urlStatic)
}

func getListing(url string) ([]ArrivalRes, error) {
	client := &http.Client{}
	req, err := initReq(url)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("error status code: %d", res.StatusCode))
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return deserialize(body)
}

func initReq(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Basic TXVmTXVmOjNxZlA0NmV6QHZLZjR5IQ==")
	return req, nil
}

func deserialize(res []byte) ([]ArrivalRes, error) {
	var data []ArrivalRes
	if err := json.Unmarshal(res, &data); err != nil {
		fmt.Println("failed to unmarshal:", err)
		return nil, err
	}
	return data, nil
}
