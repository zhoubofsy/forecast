package main

import "github.com/zhoubofsy/weather_forecast/weather"
import "github.com/zhoubofsy/weather_forecast/speaker"
import "fmt"
import "strconv"
import "encoding/json"
import "io/ioutil"

type WFConfig struct {
	Speaker_AK string `json:"speaker_ak"`
	Speaker_Sec string `json:"speaker_ack"`
	Speaker_cuid string `json:"speaker_cuid"`
	Weather_AK string `json:"weather_ak"`
	Weather_loc string `json:"weather_location"`
}

func main() {
	// load weather forecast config
	raw_config, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println("Read Config Failure => ", err)
		return
	}

	var config WFConfig
	if err = json.Unmarshal(raw_config, &config); err != nil {
		fmt.Println("Config Parse Failure => ", err)
	}

	ak := config.Speaker_AK
	sec := config.Speaker_Sec
	cuid := config.Speaker_cuid
	location := config.Weather_loc

	var bw weather.Baidu_Weather
	var pm25 int

	bw.Init(ak)
	err = bw.Do_Request(location)
	if err == nil {
		pm25, _ = bw.Get_PM25()
	}

	var bs speaker.Baidu_Say
	content := location + "当前PM2.5为" + strconv.Itoa(pm25)

	bs.Init(ak, sec, cuid)
	err, rcode := bs.Do_Text2audio(content)
	if err == nil && rcode == 502 {
		bs.Create_Token()
		err, rcode = bs.Do_Text2audio(content)
	}

	fmt.Printf("BaiduSay: %s\n", content)
}

