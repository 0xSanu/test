package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type OccupancyData struct {
	RoomID        string  `json:"room_id"`
	OccupancyRate float64 `json:"occupancy_rate"`
	AverageRate   float64 `json:"average_rate"`
	HighestRate   float64 `json:"highest_rate"`
	LowestRate    float64 `json:"lowest_rate"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// Dummy data
var dummyData = map[string]OccupancyData{
	"101": {
		RoomID:        "101",
		OccupancyRate: 85.0,
		AverageRate:   150.0,
		HighestRate:   250.0,
		LowestRate:    120.0,
	},
	"102": {
		RoomID:        "102",
		OccupancyRate: 90.0,
		AverageRate:   175.0,
		HighestRate:   275.0,
		LowestRate:    130.0,
	},
	"103": {
		RoomID:        "103",
		OccupancyRate: 78.0,
		AverageRate:   145.0,
		HighestRate:   240.0,
		LowestRate:    100.0,
	},
	"104": {
		RoomID:        "104",
		OccupancyRate: 65.0,
		AverageRate:   130.0,
		HighestRate:   200.0,
		LowestRate:    90.0,
	},
}

func getRoomOccupancy(roomID string) (OccupancyData, error) {
	if data, exists := dummyData[roomID]; exists {
		return data, nil
	}
	return OccupancyData{}, fmt.Errorf("room ID not found")
}

func occupancyHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Parsing room id from url
	roomID := r.URL.Query().Get("room_id")
	if roomID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Missing room_id parameter"})
		return
	}

	occupancyData, err := getRoomOccupancy(roomID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Room not found"})
		return
	}

	// Returning data as JSON
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(occupancyData)
}

func main() {
	// Route for the API
	http.HandleFunc("/occupancy", occupancyHandler)

	// Listening the
	fmt.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
