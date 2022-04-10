package api_model

type CurrentJsonResponse struct {
	Location *Location `json:"location,omitempty" form:"location,omitempty"` //TODO: Write general description for this field
	Current  *Current  `json:"current,omitempty" form:"current,omitempty"`   //TODO: Write general description for this field
}

/*
 * Structure for the custom type Location
 */
type Location struct {
	Name           *string  `json:"name,omitempty" form:"name,omitempty"`                       //Local area name.
	Region         *string  `json:"region,omitempty" form:"region,omitempty"`                   //Local area region.
	Country        *string  `json:"country,omitempty" form:"country,omitempty"`                 //Country
	Lat            *float64 `json:"lat,omitempty" form:"lat,omitempty"`                         //Area latitude
	Lon            *float64 `json:"lon,omitempty" form:"lon,omitempty"`                         //Area longitude
	TzId           *string  `json:"tz_id,omitempty" form:"tz_id,omitempty"`                     //Time zone
	LocaltimeEpoch *int64   `json:"localtime_epoch,omitempty" form:"localtime_epoch,omitempty"` //Local date and time in unix time
	Localtime      *string  `json:"localtime,omitempty" form:"localtime,omitempty"`             //Local date and time
}

/*
 * Structure for the custom type Current
 */
type Current struct {
	LastUpdatedEpoch *int64     `json:"last_updated_epoch,omitempty" form:"last_updated_epoch,omitempty"` //Local time when the real time data was updated in unix time.
	LastUpdated      *string    `json:"last_updated,omitempty" form:"last_updated,omitempty"`             //Local time when the real time data was updated.
	TempC            *float64   `json:"temp_c,omitempty" form:"temp_c,omitempty"`                         //Temperature in celsius
	TempF            *float64   `json:"temp_f,omitempty" form:"temp_f,omitempty"`                         //Temperature in fahrenheit
	IsDay            *int64     `json:"is_day,omitempty" form:"is_day,omitempty"`                         //1 = Yes 0 = No <br />Whether to show day condition icon or night icon
	Condition        *Condition `json:"condition,omitempty" form:"condition,omitempty"`                   //TODO: Write general description for this field
	WindMph          *float64   `json:"wind_mph,omitempty" form:"wind_mph,omitempty"`                     //Wind speed in miles per hour
	WindKph          *float64   `json:"wind_kph,omitempty" form:"wind_kph,omitempty"`                     //Wind speed in kilometer per hour
	WindDegree       *int64     `json:"wind_degree,omitempty" form:"wind_degree,omitempty"`               //Wind direction in degrees
	WindDir          *string    `json:"wind_dir,omitempty" form:"wind_dir,omitempty"`                     //Wind direction as 16 point compass. e.g.: NSW
	PressureMb       *float64   `json:"pressure_mb,omitempty" form:"pressure_mb,omitempty"`               //Pressure in millibars
	PressureIn       *float64   `json:"pressure_in,omitempty" form:"pressure_in,omitempty"`               //Pressure in inches
	PrecipMm         *float64   `json:"precip_mm,omitempty" form:"precip_mm,omitempty"`                   //Precipitation amount in millimeters
	PrecipIn         *float64   `json:"precip_in,omitempty" form:"precip_in,omitempty"`                   //Precipitation amount in inches
	Humidity         *int64     `json:"humidity,omitempty" form:"humidity,omitempty"`                     //Humidity as percentage
	Cloud            *int64     `json:"cloud,omitempty" form:"cloud,omitempty"`                           //Cloud cover as percentage
	FeelslikeC       *float64   `json:"feelslike_c,omitempty" form:"feelslike_c,omitempty"`               //Feels like temperature as celcius
	FeelslikeF       *float64   `json:"feelslike_f,omitempty" form:"feelslike_f,omitempty"`               //Feels like temperature as fahrenheit
	VisKm            *float64   `json:"vis_km,omitempty" form:"vis_km,omitempty"`                         //TODO: Write general description for this field
	VisMiles         *float64   `json:"vis_miles,omitempty" form:"vis_miles,omitempty"`                   //TODO: Write general description for this field
	Uv               *float64   `json:"uv,omitempty" form:"uv,omitempty"`                                 //UV Index
	GustMph          *float64   `json:"gust_mph,omitempty" form:"gust_mph,omitempty"`                     //Wind gust in miles per hour
	GustKph          *float64   `json:"gust_kph,omitempty" form:"gust_kph,omitempty"`                     //Wind gust in kilometer per hour
}
type Condition struct {
	Text *string `json:"text,omitempty" form:"text,omitempty"` //Weather condition text
	Icon *string `json:"icon,omitempty" form:"icon,omitempty"` //Weather icon url
	Code *int64  `json:"code,omitempty" form:"code,omitempty"` //Weather condition unique code.
}
