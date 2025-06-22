package internal

import (
	"encoding/json"
	"net/http"

	"github.com/Max/Calculation/Calculater/models"
)

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only post method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.CalcRequest
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
	resp := models.CalcResp{Result: result}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
