package vision

import (
	"fmt"
	"testing"
)

var key string = "4f558a77503b46549ac0d784c1651d9e" // this code is sample , forget this :) my token has been changed !

func TestOCR(t *testing.T) {
	vision, err := New(key)
	if err != nil {
		t.Log(err)
		return
	}
	result, err := vision.OCR("https://portalstoragewuprod2.azureedge.net/vision/OpticalCharacterRecognition/1.jpg", OCROption{
		Language: LANG_English,
	})
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(result.String())
}

func TestAnalyze(t *testing.T) {
	vision, err := New(key)
	if err != nil {
		t.Log(err)
		return
	}
	result, err := vision.Analyze("https://portalstoragewuprod2.azureedge.net/vision/Analysis/1.jpg", VisualFeatures{
		Tags: true,
	})
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(result)
}

func TestTag(t *testing.T) {
	vision, err := New(key)
	if err != nil {
		t.Log(err)
		return
	}
	result, err := vision.Tag("https://portalstoragewuprod2.azureedge.net/vision/Analysis/1.jpg")
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(result)
}

func TestMakeString(t *testing.T) {
	vis := VisualFeatures{}
	vis.Adult = true
	vis.Categories = true

	fmt.Println(vis.String())
}
