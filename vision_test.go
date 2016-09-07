package vision

import (
	"fmt"
	"testing"
)

var key string = ""

func TestAnalyze(t *testing.T) {
	vision, err := New(key)
	if err != nil {
		t.Fatal(err)
	}
	result, err := vision.Analyze("https://portalstoragewuprod2.azureedge.net/vision/Analysis/1.jpg", VisualFeatures{
		Tags: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestTag(t *testing.T) {
	vision, err := New(key)
	if err != nil {
		t.Fatal(err)
	}
	result, err := vision.Tag("https://portalstoragewuprod2.azureedge.net/vision/Analysis/1.jpg")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func TestMakeString(t *testing.T) {
	vis := VisualFeatures{}
	vis.Adult = true
	vis.Categories = true

	fmt.Println(vis.String())
}
