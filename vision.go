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

func (vision *Vision) Analyze(url string, order VisualFeatures) (VisionResult, error) {
	query, err := order.String()
	if err != nil {
		panic(err)
	}
	apiURL := URL + "/analyze?visualFeatures=" + query

	req, err := http.NewRequest("POST", apiURL, strings.NewReader("{\"url\":\""+url+"\"}"))
	if err != nil {
		return VisionResult{}, err
	}
	req.Header.Set("Ocp-Apim-Subscription-Key", vision.BingKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return VisionResult{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)

		result := VisionResult{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return VisionResult{}, err
		}

		vision.LastRequestID = result.RequestID
		return result, nil
	}

	if resp.StatusCode == 400 || resp.StatusCode == 415 || resp.StatusCode == 500 {
		body, _ := ioutil.ReadAll(resp.Body)

		result := Error{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			return VisionResult{}, err
		}

		return VisionResult{}, fmt.Errorf(result.Code)
	}

	return VisionResult{}, fmt.Errorf("Unknown Error Occured , Check the key")
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
