package model

type Input struct {
	Cards string `json:"input" binding:"required"`
}

type ConfigJudge struct {
	SortNums         []int
	CountMarks       int
	Variation        int
	NumCountsOverlap []int
	UniqNumsSize     int
}

type WeatherApi struct {
	Lat            int      `json:"lat"`
	Lon            int      `json:"lon"`
	Timezone       string   `json:"timezone"`
	TimezoneOffset int      `json:"timezoneoffset"`
	Hourly         []Hourly `json:"hourly"`
}

type Hourly struct {
	Dt        int     `json:"dt"`
	Sunrise   int     `json:"sunrise"`
	Sunset    int     `json:"sunset"`
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	DewPoint  float64 `json:"dew_point"`
	Uvi       float64 `json:"uvi"`
	Clouds    int     `json:"clouds"`
	Visavity  int     `json:"visavity"`
	WindSpeed float64 `json:"wind_speed"`
	WindDeg   int     `json:"wind_deg"`
	WindGust  float64 `json:"wind_gust"`
	Weaher    []Weather
	Pop       float64 `json:"pop"`
}

type Weather struct {
	Id          string `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}
