package server
import (
	"encoding/json"
	"fmt"
	"net/http"

	"dev_ops_th/scraper"
)

func StartServer() {
	http.HandleFunc("/daily-quotes", func(w http.ResponseWriter, r *http.Request) {
		quotes, err := scraper.ScrapeQuotes()
		if err != nil {
			http.Error(w, "Failed to scrape quotes", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err = json.NewEncoder(w).Encode(quotes)
		if err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8080", nil)
	fmt.Println("Server started at http://localhost:8000/daily-quotes")
}