package main

import (
	"github.com/h2non/gock"
	"github.com/nbio/st"
	"reflect"
	"testing"
)

func TestAbs(t *testing.T) {
	a := 1
	if a != 1 {
		t.Errorf("Abs(-1) = %d; want 1", a)
	}
}

func TestClient(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	var apiReturn []ArrivalRes
	apiReturn = append(apiReturn, ArrivalRes{
		Icao24:                           "1",
		FirstSeen:                        0,
		EstDepartureAirport:              nil,
		LastSeen:                         0,
		EstArrivalAirport:                "",
		Callsign:                         "",
		EstDepartureAirportHorizDistance: nil,
		EstDepartureAirportVertDistance:  nil,
		EstArrivalAirportHorizDistance:   0,
		EstArrivalAirportVertDistance:    0,
		DepartureAirportCandidatesCount:  0,
		ArrivalAirportCandidatesCount:    0,
	})
	apiReturn = append(apiReturn, ArrivalRes{
		Icao24:                           "2",
		FirstSeen:                        0,
		EstDepartureAirport:              nil,
		LastSeen:                         0,
		EstArrivalAirport:                "",
		Callsign:                         "",
		EstDepartureAirportHorizDistance: nil,
		EstDepartureAirportVertDistance:  nil,
		EstArrivalAirportHorizDistance:   0,
		EstArrivalAirportVertDistance:    0,
		DepartureAirportCandidatesCount:  0,
		ArrivalAirportCandidatesCount:    0,
	})
	gock.New("http://localhost:8080/api").
		MatchParam("airport", "EDDF").
		MatchParam("begin", "1517227200").
		MatchParam("end", "1517230800").
		Get("/flights/arrival").
		Reply(200).
		JSON(apiReturn)

	res, err := getListing("http://localhost:8080/api/flights/arrival?airport=EDDF&begin=1517227200&end=1517230800")
	if err != nil {
		t.Errorf("getListing error %v", err)
	}
	if !reflect.DeepEqual(res, apiReturn) {
		t.Errorf("the res and apiReturn is not egals")
	}
	st.Expect(t, gock.IsDone(), true)
}
