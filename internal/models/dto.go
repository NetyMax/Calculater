package models

type CalcResp struct {
	Result float64 `json:"result"`
}

type CalcRequest struct {
	FirstNum  float64 `json:"first_num"`
	SecondNum float64 `json:"second_num"`
	Operation string  `json:"operation"`
}
