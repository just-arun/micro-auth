package model

type ResponseParams struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func (ResponseParams) TableName() string {
	return "response_params"
}