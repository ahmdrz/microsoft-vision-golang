// vision project vision.go
package vision

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

func New(key string) (*Vision, error) {
	if len(key) < 10 {
		return nil, fmt.Errorf("Invalid Key")
	}
	return &Vision{
		BingKey: key,
	}, nil
}

func (vision *Vision) Analyze(url string, order VisualFeatures) VisionResult {
	query, err := order.String()
	if err != nil {
		panic(err)
	}
	apiURL := URL + "/analyze?visualFeatures=" + query

	req, err := http.NewRequest("POST", apiURL, strings.NewReader("{\"url\":\""+url+"\"}"))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Ocp-Apim-Subscription-Key", vision.BingKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	result := VisionResult{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}

	vision.LastRequestID = result.RequestID
	return result
}

func (order VisualFeatures) String() (string, error) {
	var (
		v = make([]string, 0)
		s = reflect.ValueOf(order)
		t = reflect.TypeOf(order)
	)

	for i := 0; i < s.NumField(); i++ {
		if s.Field(i).Interface().(bool) == true {
			v = append(v, t.Field(i).Name)
		}
	}

	if len(v) == 0 {
		return "", fmt.Errorf("Empty v")
	}

	result := ""
	for _, item := range v {
		result += item + ","
	}

	return result[:len(result)-1], nil
}
