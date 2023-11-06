package observer

import (
	"testing"
	"time"
)

func TestObserver(t *testing.T) {
	server := &WeatherForecastServer{}

	client1 := &Client{name: "client1"}
	client2 := &Client{name: "client2"}

	server.RegisterObserver(client1)
	server.RegisterObserver(client2)

	// Simulating weather updates
	event1 := WeatherChangeEvent{Temperature: 25.5, Humidity: 70.2, Pressure: 1013.5}
	server.SetMeasurements(event1)
	time.Sleep(2 * time.Second)

	event2 := WeatherChangeEvent{Temperature: 26.8, Humidity: 68.9, Pressure: 1012.8}
	server.SetMeasurements(event2)
	time.Sleep(2 * time.Second)

	event3 := WeatherChangeEvent{Temperature: 24.3, Humidity: 72.1, Pressure: 1014.2}
	server.SetMeasurements(event3)

	// Unregistering an observer
	server.RemoveObserver(client2)

	// Simulating
	time.Sleep(2 * time.Second)
	event4 := WeatherChangeEvent{Temperature: 22.2, Humidity: 66.6, Pressure: 1013.1}
	server.SetMeasurements(event4)

	time.Sleep(time.Second)
}
