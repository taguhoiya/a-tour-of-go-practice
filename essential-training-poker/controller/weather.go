package controller

import (
	"poker/model"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func NeedUmbrella(c *gin.Context) {

	req, err := http.NewRequest(http.MethodGet, "https://api.openweathermap.org/data/2.5/onecall", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("content-type", "application/json")
	q := req.URL.Query()
	q.Add("lat", "34")
	q.Add("lon", "135")
	q.Add("appid", "8e55ee2e871bb393dfbc59545d05ddb3")
	q.Add("exclude", "minutely,current,daily,alerts")
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var weaherApi model.WeatherApi
	if err = json.Unmarshal([]byte(body), &weaherApi); err != nil {
		fmt.Println(err)
		return
	}

	currentTime := time.Now()
	tokyo, _ := time.LoadLocation("Asia/Tokyo")
	if err != nil {
			fmt.Println(err)
			return
	}
	now := time.Now().In(tokyo)
	condition := int(Bod(now).Sub(currentTime).Hours()) + 1
	todayPops := []float64{}
	counter := 1
	for _, v := range weaherApi.Hourly{
		if counter == condition + 1 {
			break
		}
		todayPops = append(todayPops, v.Pop)
		counter++
	}
	averagePop := sum(todayPops) / float64(len(todayPops))

	var result string
	if averagePop < 0.3 {
		result = "今日は傘が不要です。"
	} else {
		result = "今日は傘が必要です。"
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status": 200,
		"result": result,
	})
}

func Bod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day + 1, 0, 0, 0, 0, t.Location())
}

func sum(slice []float64) float64 {
	sum := 0.0
	for _, v := range slice {
		sum += v
	}
	return sum
}