package entity

type Response struct {
	Data             interface{}       `json:"data"`
	Error            bool              `json:"error"`
	ErrorText        string            `json:"errorText"`
	AdditionalErrors map[string]string `json:"additionalErrors"`
}
