package weatherbit

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// CurrentObservation holds the result of our query to Weatherbit API.
type CurrentObservation struct {
	Data []struct {
		Rh       int     `json:"rh"`
		Pod      string  `json:"pod"`
		Pres     float64 `json:"pres"`
		Timezone string  `json:"timezone"`
		Weather  struct {
			Icon        string `json:"icon"`
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"weather"`
		CountryCode  string      `json:"country_code"`
		Clouds       int         `json:"clouds"`
		Vis          int         `json:"vis"`
		WindSpd      float64     `json:"wind_spd"`
		WindCdirFull string      `json:"wind_cdir_full"`
		AppTemp      float64     `json:"app_temp"`
		Lon          float64     `json:"lon"`
		StateCode    string      `json:"state_code"`
		Ts           int         `json:"ts"`
		ElevAngle    int         `json:"elev_angle"`
		HAngle       int         `json:"h_angle"`
		Dewpt        float64     `json:"dewpt"`
		ObTime       string      `json:"ob_time"`
		Uv           int         `json:"uv"`
		Sunset       string      `json:"sunset"`
		Sunrise      string      `json:"sunrise"`
		CityName     string      `json:"city_name"`
		Precip       interface{} `json:"precip"`
		Station      string      `json:"station"`
		Lat          float64     `json:"lat"`
		Dhi          int         `json:"dhi"`
		Datetime     string      `json:"datetime"`
		Temp         float64     `json:"temp"`
		WindDir      int         `json:"wind_dir"`
		Slp          int         `json:"slp"`
		WindCdir     string      `json:"wind_cdir"`
	} `json:"data"`
	Count int `json:"count"`
}

// GetWeather is used to obtain weather data from the specified city.
func GetWeather(country string, city string, apiKey string) (CurrentObservation, error) {
	url := "https://api.weatherbit.io/v2.0/current"
	// make new request
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")

	// build query string
	q := req.URL.Query()
	q.Add("key", apiKey)
	q.Add("city", city)
	q.Add("country", country)
	req.URL.RawQuery = q.Encode()

	// getting response
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return CurrentObservation{}, err
	}

	// read the body to bytes for un-marshalling
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return CurrentObservation{}, err
	}

	// If statusCode is not good, return error string
	switch res.StatusCode {
	case 200:
		response := CurrentObservation{}
		json.Unmarshal(body, &response)
		return response, nil
	default:
		// Turn response into string and return it
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		responseBody := buf.String()
		err = errors.New(responseBody)

		return CurrentObservation{}, err
	}
}

