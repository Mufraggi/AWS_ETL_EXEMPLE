package clientApi

import (
	"github.com/h2non/gock"
	"github.com/nbio/st"
	"reflect"
	"testing"
)

func initArrival() []ArrivalRes {
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
	return apiReturn
}

func TestClientWork(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution
	t.Setenv("urlApi", "http://localhost:8080/api/flights/arrival?airport=EDDF&begin=1517227200&end=1517230800")
	apiReturn := initArrival()
	cApi := New()
	gock.New("http://localhost:8080/api").
		MatchParam("airport", "EDDF").
		MatchParam("begin", "1517227200").
		MatchParam("end", "1517230800").
		Get("/flights/arrival").
		Reply(200).
		JSON(apiReturn)

	res, err := cApi.GetListing()
	if err != nil {
		t.Errorf("GetListing error %v", err)
	}
	if !reflect.DeepEqual(res, apiReturn) {
		t.Errorf("the res and apiReturn is not egals")
	}
	st.Expect(t, gock.IsDone(), true)
}

func TestClientWorkEmptyList(t *testing.T) {
	t.Setenv("urlApi", "http://localhost:8080/api/flights/arrival?airport=EDDF&begin=1517227200&end=1517230800")
	defer gock.Off() // Flush pending mocks after test execution
	apiReturn := make([]ArrivalRes, 0)
	cApi := New()
	gock.New("http://localhost:8080/api").
		MatchParam("airport", "EDDF").
		MatchParam("begin", "1517227200").
		MatchParam("end", "1517230800").
		Get("/flights/arrival").
		Reply(200).
		JSON(apiReturn)

	res, err := cApi.GetListing()
	if err != nil {
		t.Errorf("GetListing error %v", err)
	}
	if !reflect.DeepEqual(res, apiReturn) {
		t.Errorf("the res and apiReturn is not egals")
	}
	st.Expect(t, gock.IsDone(), true)
}

func TestClientReturns404(t *testing.T) {
	defer gock.Off()
	t.Setenv("urlApi", "http://localhost:8080/api/flights/arrival?airport=EDDF&begin=1517227200&end=1517230800")
	cApi := New()
	gock.New("http://localhost:8080/api").
		MatchParam("airport", "EDDF").
		MatchParam("begin", "1517227200").
		MatchParam("end", "1517230800").
		Get("/flights/arrival").
		Reply(404)
	_, err := cApi.GetListing()
	if err == nil {
		t.Errorf("GetListing does not returing error")
	}
	st.Expect(t, gock.IsDone(), true)
}
