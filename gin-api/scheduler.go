package scheduler

import (
	"log"
	"time"
)

func StartScheduler() {
	ticker := time.NewTicker(5 * time.Minute) // Adjust interval as needed
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				fetchCatabusData()
				fetchWeatherData()
				fetchTrafficData()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func fetchCatabusData() {
	// Call CATABUS API and store data
	log.Println("Fetching CATABUS data...")
	// Your API logic here
}

func fetchWeatherData() {
	// Call OpenWeatherMap API and store data
	log.Println("Fetching weather data...")
	// Your API logic here
}

func fetchTrafficData() {
	// Call Google Maps API and store data
	log.Println("Fetching traffic data...")
	// Your API logic here
}
