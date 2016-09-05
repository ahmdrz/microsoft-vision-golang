package vision

type Vision struct {
	BingKey string
}

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

const (
	URL string = "https://api.projectoxford.ai/vision/v1.0"
)
