package weather

// Location is a type that represents a city. With more time I would have
// abstracted this into an interface so the actual type of a city would be
// a (hidden) implementation detail. This would be helpful because
// as things grow a string might not be the correct data type.
// This would also facilitate searching by coordinates and what not.
type Location struct {
	Lat float32
	Lon float32
}

// Weather is weather related information for a specific city.
type WeatherItem struct {
	FeelsLike   float32  `json:"feelsLike"`
	Temp        float32  `json:"temp"`
	WindSpeed   float32  `json:"windSpeed"`
	WindDir     int      `json:"windDirection"`
	Description []string `json:"description"`
}

// weather is the type that is returned by the weather api.
type weather struct {
	Code    int    `json:"cod"`
	Message string `json:"message"`
	Hourly  []struct {
		Temp      float32 `json:"temp"`
		FeelsLike float32 `json:"feels_like"`
		Humidity  float32 `json:"humidity"`
		Pressure  float32 `json:"pressure"`
		WindSpeed float32 `json:"wind_speed"`
		WindDir   int     `json:"wind_deg"`
		Weather   []struct {
			Description string `json:"description"`
		} `json:"weather"`
	} `json:"hourly"`
}
