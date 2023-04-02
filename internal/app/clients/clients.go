package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	//URI scheme and host for the National Weather Service API
	nwsAPIEndPointBase = "https://api.weather.gov"
	//URI path for the regular forecast - 7 days with each day and night
	nwsAPIForecastURI = "/forecast"
	//URI path for the Hourly Forecast
	nwsAPIHourlyForecastURI = "/forecast/hourly"
)

type weatherClientHeaders struct {
	UserAgent string `json:"user-agent"`
}

type weatherClientHourlyPeriods struct {
	Name             string `json:"name"`
	Number           int    `json:"number"`
	Temperature      int    `json:"temperature"`
	TemperatureUnit  string `json:"temperatureUnit"`
	ShortForcast     string `json:"shortForecast"`
	DetailedForecast string `json:"detailedForecast"`
	WindSpeed        string `json:"windSpeed"`
	WindDirection    string `json:"windDirection"`
}

type weatherClientHourlyProperties struct {
	Updated string                       `json:"updated"`
	Units   string                       `json:"us"`
	Periods []weatherClientHourlyPeriods `json:"periods"`
}

type weatherClientHourlyResponse struct {
	Details weatherClientHourlyProperties `json:"properties"`
}

//TODO - add input and return signature
func Weather() {
	//Set the request header with a user-agent field.
	//This is required for authentication of the request by the National Weather Service API.
	//Documentation says in the future the user-agent field to be replaced by API key but no known date
	//Documentation at https://www.weather.gov/documentation/services-web-api
	//Excerpt from National Weather Service API documentation around authentication:
	//Authentication
	//A User Agent is required to identify your application. This string can be anything, and the more unique to your application the less likely it will be affected by a security event.
	//If you include contact information (website or email), we can contact you if your string is associated to a security event.
	//This will be replaced with an API key in the future.
	var weatherHeaders = weatherClientHeaders{
		UserAgent: "JeremiahBrooks, jeremiah.brooks.987@gmail.com@gmail.com",
	}
	//headers.userAgent = "JeremiahBrooks, jeremiah.brooks.987@gmail.com@gmail.com"

	//Geoloaction coordinates
	// - max 4 decimal point precision
	//geoLocationCoordinates := "38.6749,-94.9132"
	gridX := "32"
	gridY := "32"
	office := "EAX"

	//Set the full National Weather Service API endpoint
	//Details:
	// obtain the office and gridx and gridy
	// https://api.weather.gov/points/{lat},{lon}
	// https://api.weather.gov/points/39.7456,-97.0892
	// Get forcast
	// https://api.weather.gov/gridpoints/{office}/{gridx},{gridy}/forecast
	// Example:
	// https://api.weather.gov/gridpoints/TOP/32,81/forecast
	// You can also get the hour-by-hour forecast from the forecastHourly property.
	// For our example
	// https://api.weather.gov/gridpoints/TOP/32,81/forecast/hourly

	// Standard Forcast
	//requestURL := fmt.Sprintf("http://localhost:%d", serverPort)
	nwsAPIEndPointFullURL := nwsAPIEndPointBase + "/gridpoints/" + office + "/" + gridX + "," + gridY + nwsAPIForecastURI

	// Hourly
	//nwsAPIEndPointFullURL := nwsAPIEndPointBase + "/gridpoints/" + office + "/" + gridX + "," + gridY + nwsAPIHourlyForecastURI

	//Create request to the National Weather Service API endpoint to obtain the weather data

	req, err := http.NewRequest(http.MethodGet, nwsAPIEndPointFullURL, nil)
	if err != nil {
		//TODO - add structured logging
		fmt.Printf("client: could not create request: %s\n", err)
		//TODO - add return
		os.Exit(1)
	}
	req.Header.Set("User-Agent", weatherHeaders.UserAgent)

	hc := http.Client{Timeout: time.Duration(15) * time.Second}

	res, err := hc.Do(req)
	if err != nil {
		//TODO - add structured logging
		fmt.Printf("client: error making http request: %s\n", err)
		//TODO - add return
		os.Exit(1)
	}

	//DEBUG Logging
	//TODO - add structured logging
	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		//TODO - add structured logging
		fmt.Printf("client: could not read response body: %s\n", err)
		//TODO - add return
		os.Exit(1)
	}
	//fmt.Printf("client: response body: %s\n", resBody)

	var weatherData weatherClientHourlyResponse

	err = json.Unmarshal(resBody, &weatherData)
	if err != nil {
		//TODO - add return and structured logging
		fmt.Println(err)
	}

	b, err := json.MarshalIndent(weatherData, "", "  ")
	if err != nil {
		//TODO - add return and structured logging
		fmt.Println(err)
	}
	//TODO - add return
	fmt.Print(string(b))

}
