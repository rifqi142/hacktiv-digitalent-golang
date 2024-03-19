package handler

import (
	"encoding/json"
	"hacktiv-digitalent-golang/assignment/assignment-3/service"
	"hacktiv-digitalent-golang/assignment/assignment-3/util"
	"net/http"
	"time"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	status := service.GenerateRandomStatusValue()

	err := util.WriteStatusFile(status, "status.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assignment/assignment-3/static/index.html")
	
}

// reload status every 15 seconds
func StatusAutoReload() {
	ticker := time.NewTicker(15 * time.Second)
	go func() {
		for range ticker.C {
			status := service.GenerateRandomStatusValue()
			err := util.WriteStatusFile(status, "status.json")
			if err != nil {
				panic(err)
			}
		}
	}()
}