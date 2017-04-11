package weather

//import "fmt"
import "net/http"
import "io/ioutil"
import "encoding/json"
//import "reflect"
import "strconv"

type Baidu_Weather struct {
	url string
	resp_body map[string] interface{}
	access_key string
}

func request_baidu_weather(url string) (map[string] interface{}, error){

	var json_map map[string] interface{} = nil
	resp, err := http.Get(url)
	if err == nil {
		//var body = nil
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			json.Unmarshal(body, &json_map)
			//fmt.Println(json_map, "==>", reflect.TypeOf(json_map))
		}

		resp.Body.Close()
	}

	return json_map, nil
}

func (baidu *Baidu_Weather)Init(ak string) (error) {
	baidu.access_key = ak
	return nil
}

func (baidu *Baidu_Weather)Do_Request(location string) (err error) {
	baidu.url = "http://api.map.baidu.com/telematics/v3/weather?location='" + location + "'&output=json&ak=" + baidu.access_key
	baidu.resp_body, err = request_baidu_weather(baidu.url)
	return
}

func (baidu *Baidu_Weather)Get_PM25() (pm25 int, err error) {
	var results_resp = baidu.resp_body["results"]
	pm25, _ = strconv.Atoi( ((results_resp.([] interface{}))[0].(map[string] interface{}))["pm25"].(string) )
	return
}

/*
func main() {
	var location string
	var pm25 int

	location = "沈阳"
	fmt.Println(location, "天气情况：")
	var bw baidu_weather
	err := bw.Do_Request(location)
	if err == nil {
		pm25, _ = bw.Get_PM25()
	}
	fmt.Println("PM2.5: ", pm25)
}
*/

