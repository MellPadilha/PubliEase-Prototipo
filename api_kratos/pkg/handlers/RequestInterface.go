package handlers

import (
	"net/http"
)

func RequestInterface(w http.ResponseWriter, r *http.Request) {
	//var todo models.Todo

	//err := json.NewDecoder(r.Body).Decode(&todo)

	//if err != nil {
	//	log.Printf("Erro, fudeu")
	//	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	//	return
	//}

	//var resp map[string]any

	//if err != nil {
	//	resp = map[string]any{
	//		"Error":   true,
	//		"Message": fmt.Sprintf("IIhhhh deu erro", err),
	//	}
	//
	//} else {
	//	resp = map[string]any{
	//		"Error":   false,
	//		"Message": fmt.Sprintf("Ihh"),
	//	}
	//}

	//w.Header().Add("Content-Type", "application/json")

	//json.NewEncoder(w).Encode(resp)
}
