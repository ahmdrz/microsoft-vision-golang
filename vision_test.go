package vision

import (
	"fmt"
	"testing"
)

var key string = ""

func TestAnalyze(t *testing.T) {
	vision, _ := New(key)
	vision.Analyze("https://portalstoragewuprod2.azureedge.net/vision/Analysis/1.jpg", VisualFeatures{
		Categories: true,
	})
}

func TestMakeString(t *testing.T) {
	vis := VisualFeatures{}
	vis.Adult = true
	vis.Categories = true

	fmt.Println(vis.String())
}
