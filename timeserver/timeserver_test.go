package timeserver

import (
	"net/http"
	"testing"

	"io/ioutil"
	"net/http/httptest"
)

func BenchmarkTimer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TimeHandler := timer{timerProv{}}

		ts := httptest.NewServer(http.HandlerFunc(TimeHandler.RealTime))

		defer ts.Close()

		resp, err := http.Get(ts.URL)
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

	}

}

type fakeTimerProv struct{}

func (ft fakeTimerProv) getCurrentTime() string {
	return "2017-04-27 11:49:33.7154352 +0300 EEST"
}

func TestGetTimeOK(t *testing.T) {
	t.Log("Start")
	var expectedResult = "2017-04-27 11:49:33.7154352 +0300 EEST"

	fakeTimeHandler := timer{fakeTimerProv{}}

	ts := httptest.NewServer(http.HandlerFunc(fakeTimeHandler.RealTime))

	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if string(body) != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, string(body))
	}
	t.Log("Finish")

}
/*
func TestGetTimeFail(t *testing.T) {
	t.Log("Start")
	var expectedResult = "2017-04-26 11:49:33.7154352 +0300 EEST"

	fakeTimeHandler := timer{fakeTimerProv{}}

	ts := httptest.NewServer(http.HandlerFunc(fakeTimeHandler.RealTime))

	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if string(body) != expectedResult {
		t.Errorf("Expected %s but got %s", expectedResult, string(body))
	}
	t.Log("Finish")

}
*/