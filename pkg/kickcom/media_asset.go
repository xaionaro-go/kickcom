package kickcom

import "time"

type MediaAsset struct {
	ID               uint64 `json:"id"`
	ModelType        string `json:"model_type"`
	ModelID          uint64 `json:"model_id"`
	CollectionName   string `json:"collection_name"`
	Name             string `json:"name"`
	FileName         string `json:"file_name"`
	MimeType         string `json:"mime_type"`
	Disk             string `json:"disk"`
	Size             uint64 `json:"size"`
	Manipulations    []any  `json:"manipulations"`
	CustomProperties struct {
		GeneratedConversions struct {
			Fullsize bool `json:"fullsize"`
		} `json:"generated_conversions"`
	} `json:"custom_properties"`
	/*ResponsiveImages struct {
		Fullsize struct {
			Urls      []string `json:"urls"`
			Base64Svg string   `json:"base64svg"`
		} `json:"fullsize"`
	} `json:"responsive_images"`*/
	OrderColumn     uint64    `json:"order_column"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	UUID            string    `json:"uuid"`
	ConversionsDisk string    `json:"conversions_disk"`
}
