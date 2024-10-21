package kickcom

// ImageV1 is a representation of an image in API v1.
type ImageV1 struct {
	Srcset string `json:"srcset"`
	Src    string `json:"src"`
}
