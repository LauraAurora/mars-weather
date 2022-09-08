package nasa_api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"time"
)

type Soles struct {
	Id                        string `json:"id"`
	Terrestrial_date          string `json:"terrestrialDate"`
	Sol                       string `json:"sol"`
	Season                    string `json:"season"`
	Min_temp                  string `json:"min_temp"`
	Max_temp                  string `json:"max_temp"`
	Pressure                  string `json:"pressure"`
	Pressure_string           string `json:"pressure_string"`
	Atmo_opacity              string `json:"atmo_opacity"`
	Sunrise                   string `json:"sunrise"`
	Sunset                    string `json:"sunset"`
	Local_uv_irradiance_index string `json:"local_uv_irradiance_index"`
	Min_gts_temp              string `json:"min_gts_temp"`
	Max_gts_temp              string `json:"max_gts_temp"`
}

func GetData() Soles {
	var result Soles
	currentTime := time.Now().Add(-120*time.Hour)
	url := "https://cab.inta-csic.es/rems/wp-content/plugins/marsweather-widget/api.php?"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return result
	}
	req.Header.Add("Cookie", "qtrans_cookie_test=qTranslate+Cookie+Test; PHPSESSID=q14ial9jve90djjl8dlclmr8l3")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return result
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return result
	}
	
	parse := string(body)
	substring := parse[12827:13187]

	json.Unmarshal([]byte(substring), &result)
	result.Terrestrial_date = currentTime.Format("01-02-2006")
	fmt.Println(result)
	return result
}
