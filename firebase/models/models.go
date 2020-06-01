package models

//ShortLinksRequestModel
type ShortLinksRequestModel struct {
	LongDynamicLink string                       `json:"longDynamicLink"`
	Suffix          ShortLinksOptionRequestModel `json:"suffix"`
}

//ShortLinksOptionRequestModel
type ShortLinksOptionRequestModel struct {
	Option string `json:"option"`
}

//ShortLinksResponseModel
type ShortLinksResponseModel struct {
	ShortLink string `json:"shortLink"`
	Warning   []struct {
		WarningCode    string `json:"warningCode"`
		WarningMessage string `json:"warningMessage"`
	} `json:"warning"`
	PreviewLink string `json:"previewLink"`
}

//ShortLinksErrorResponseModel
type ShortLinksErrorResponseModel struct {
	Error errorModel `json:"error"`
}

//errorModel
type errorModel struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}
