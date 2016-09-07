package vision

type Vision struct {
	BingKey       string
	LastRequestID string
}

const (
	URL string = "https://api.projectoxford.ai/vision/v1.0"
)

/*
Categories - categorizes image content according to a taxonomy defined in documentation.
Tags - tags the image with a detailed list of words related to the image content.
Description - describes the image content with a complete English sentence.
Faces - detects if faces are present. If present, generate coordinates, gender and age.
ImageType - detects if image is clipart or a line drawing.
Color - determines the accent color, dominant color, and whether an image is black&white.
Adult - detects if the image is pornographic in nature (depicts nudity or a sex act). Sexually suggestive content is also detected.
*/
type VisualFeatures struct {
	Categories  bool
	Tags        bool
	Description bool
	Faces       bool
	ImageType   bool
	Color       bool
	Adult       bool
}

type VisionResult struct {
	RequestID   string      `json:"requestId"`
	Categories  []Category  `json:"categories"`
	Adult       Adult       `json:"adult"`
	Faces       []Face      `json:"faces"`
	ImageType   ImageType   `json:"imageType`
	Color       Color       `json:"color"`
	MetaData    MetaData    `json:"metadata"`
	Description Description `json:"description"`
	Tags        []Tag       `json:"tags"`
}

type Tag struct {
	Name       string  `json:"name"`
	Confidence float64 `json:"confidence"`
}

type Description struct {
	Tags     []string  `json:"tags"`
	Captions []Caption `json:"captions`
}

type Caption struct {
	Text       string  `json:"text"`
	Confidence float64 `json:"confidence`
}

type MetaData struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Format string `json:"format"`
}

type Color struct {
	DominantColorForeground string   `json:"dominantColorForeground"`
	DominantColorBackground string   `json:"dominantColorBackground"`
	AccentColor             string   `json:"accentColor"`
	IsBWImg                 bool     `json:"isBWImg`
	DominantColors          []string `json:"dominantColors"`
}

type ImageType struct {
	ClipArtType     int `json:"clipArtType"`
	LineDrawingType int `json:"lineDrawingType"`
}

type Face struct {
	Age           int       `json:"age"`
	Gender        string    `json:"gender"`
	FaceRectangle Rectangle `json:"faceRectangle"`
}

type Rectangle struct {
	Top    int `json:"top"`
	Left   int `json:"left"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Adult struct {
	IsAdultContent bool    `json:"isAdultContent"`
	IsRacyContent  bool    `json:"isRacyContent"`
	AdultScore     float64 `json:"adultScore"`
	RacyScore      float64 `json:"racyScore"`
}

type Category struct {
	Name  string  `json:"name"`
	Score float32 `json:"score"`
}
