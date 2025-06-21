package internal

import (
	"encoding/json"
	"net/http"
)

type CalcRequest struct {
	FirstNum  float64 `json:"first_num"`
	SecondNum float64 `json:"second_num"`
	Operation string  `json:"operation"`
}

type CalcResp struct {
	Result float64 `json:"result"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only post method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CalcRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	var result float64
	switch req.Operation {
	case "Addition":
		result = req.FirstNum + req.SecondNum
	default:
		http.Error(w, "unsupported operation", http.StatusBadRequest)
	}
	resp := CalcResp{Result: result}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
