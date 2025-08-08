package server
import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/danielmunoz1solo/dev-ops-project-soLo/pkg/scraper"
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

	log.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}