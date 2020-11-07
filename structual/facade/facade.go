package structual

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// CurrentWeatherDataRetriever 接口
type CurrentWeatherDataRetriever interface {
	GetByCityAndCountryCode(city, coutryCode string) (Weather, error)
	GetByGeoCoordinates(lat, lon float32) (Weather, error)
}

// CurrentWeatherData 结构
type CurrentWeatherData struct {
	APIKey string
}

// Weather 结构
type Weather struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Cod   int    `json:"cod"`
	Coord struct {
		Lon float32 `json:"lon"`
		Lat float32 `json:"lat"`
	} `json:"corrd"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float32 `json:"temp"`
		Pressure float32 `json:"pressure"`
		Humidity float32 `json:"humidity"`
		TempMin  float32 `json:"temp_min"`
		TempMax  float32 `json:"temp_max"`
	} `json:"main"`
	Wind struct {
		Speed float32 `json:"spped"`
		Deg   float32 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Rain struct {
		ThreeHours float32 `json:"3h"`
	} `json:"rain"`
	Dt  uint32 `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float32 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
}

// GetByGeoCoordinates 实现接口
func (c *CurrentWeatherData) GetByGeoCoordinates(lat, lon float32) (weather *Weather, err error) {
	return c.doRequest(
		fmt.Sprintf("http:://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&APPID=%s", lat, lon, c.APIKey))
}

// GetByCityAndCountryCode 实现接口
func (c *CurrentWeatherData) GetByCityAndCountryCode(city, coutryCode string) (weather *Weather, err error) {
	return c.doRequest(
		fmt.Sprintf("http://api.openweathermap.org/data/2.5/weatherq=%s,%s&APPID=%s", city, coutryCode, c.APIKey))
}

func (c *CurrentWeatherData) doRequest(url string) (weather *Weather, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		byt, errMsg := ioutil.ReadAll(resp.Body)
		if errMsg != nil {
			errMsg = fmt.Errorf("%s", string(byt))
		}
		err = fmt.Errorf("status code was %d, aborting. Error message was:\n%s", resp.StatusCode, errMsg)
		return
	}

	weather, err = c.responseParser(resp.Body)
	resp.Body.Close()
	return
}

func (c *CurrentWeatherData) responseParser(body io.Reader) (*Weather, error) {
	w := new(Weather)
	err := json.NewDecoder(body).Decode(w)
	if err != nil {
		return nil, err
	}
	return w, nil
}
