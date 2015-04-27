package stories

import (
	"encoding/json"
	"net/http"

	"github.com/mickelsonm/storymaker/models/stories"
)

func GetAllTimelines(w http.ResponseWriter, r *http.Request) {
	timelines, err := stories.GetAllTimelines()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(timelines)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
