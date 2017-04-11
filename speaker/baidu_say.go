package speaker

import (
	"os"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
//	"fmt"
//	"time"
)

type Baidu_Say struct {
	api_key string
	secret string
	token string
	cuid string
}

func (baidu *Baidu_Say)Init(ak string, sec string, cuid string) error {
	baidu.api_key = ak
	baidu.secret = sec
	baidu.cuid = cuid

	var tok []byte
	var err error = nil
	tok_file := ak + ".tok"
	if _, err = os.Stat(tok_file); err == nil {
		tok, err = ioutil.ReadFile(tok_file)
		if err == nil {
			baidu.token = string(tok)
		}
	}

	return err
}

func (baidu *Baidu_Say)Create_Token() (string, error) {
	var json_map map[string] string = nil
	var token string = ""
	url_create_token := "https://openapi.baidu.com/oauth/2.0/token?grant_type=client_credentials&client_id=" + baidu.api_key + "&client_secret=" + baidu.secret
	resp, err := http.Post(url_create_token, "", nil)
	if err == nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			json.Unmarshal(body, &json_map)
		}
		content_type := resp.Header.Get("Content-Type")
		if resp.StatusCode == 200 && content_type == "application/json" {
			token = baidu.parse_token(json_map)
			if token != "" {
				var tok_data = []byte(token)
				ioutil.WriteFile(baidu.api_key + ".tok", tok_data, 0666)
				baidu.token = token
			}
		}
		resp.Body.Close()
	}
	return token, err
}

func (baidu *Baidu_Say)parse_token(body_json map[string] string) (string) {
	return body_json["access_token"]
}


func (baidu *Baidu_Say)Do_Text2audio(text string) (error, int) {
	var err_no int = 0
	var err error
	var resp *http.Response

	if baidu.token == "" {
		return nil, 502
	}

	url_t2a := "http://tsn.baidu.com/text2audio?tex=" + text + "&lan=zh&cuid="+ baidu.cuid + "&ctp=1&tok=" + baidu.token
	resp, err = http.Get(url_t2a)
	if err == nil {
		content_type := resp.Header.Get("Content-Type")
		if content_type == "application/json" {
			// json
			type Err_resp struct {
				Err_msg string `json:"err_msg"`
				Err_no int `json:"err_no"`
			}
			var resp_content Err_resp
			var body []byte

			body, err = ioutil.ReadAll(resp.Body)
			if err == nil && resp.StatusCode == 200{
				json.Unmarshal(body, &resp_content)
				err_no = resp_content.Err_no
				err = errors.New(resp_content.Err_msg)
			}
			resp.Body.Close()
		} else if content_type == "audio/mp3" {
			// audio mp3
			var mp3_data []byte
			mp3_data, err = ioutil.ReadAll(resp.Body)
			if err == nil {
				// write to file
				err = ioutil.WriteFile("./trans.mp3", mp3_data, 0666)
			}
			resp.Body.Close()
		} else {
			// unknow
			err = errors.New("Unknow Content Type ! => " + content_type + " | url: " + url_t2a)
		}
	}
	return err, err_no
}

/*
func main() {
	fmt.Println("Speaker Start .")
	var bs Baidu_Say

	//baidu Init
	bs.Init()

	say_content := "你好，我是lucy。"
	for {
		// baidu Do_Text2audio
		err, rcode := bs.Do_Text2audio(say_content)

		// baidu Create_Token
		if err == nil && rcode == 502 {
			bs.Create_Token()
		} else if rcode != 0 {
			fmt.Print(err)
			fmt.Println("\t code: %d", rcode)
		} else {
			fmt.Printf("Baidu Say: %s\n", say_content)
		}

		time.Sleep(5 * time.Second)
	}
}
*/

