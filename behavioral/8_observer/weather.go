package observer

import (
	"fmt"
)

// WeatherChangeEvent represents the event containing weather changes
type WeatherChangeEvent struct {
	Temperature float64
	Humidity    float64
	Pressure    float64

	// Add any additional indexes or properties as needed
}

// Observer interface represents the observers that will receive weather updates
type Observer interface {
	Update(event WeatherChangeEvent)
}

// Subject interface represents the subject or publisher that will notify observers of weather changes
type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers(event WeatherChangeEvent)
}

// WeatherForecastServer represents the weather forecast server, which is the subject or publisher
type WeatherForecastServer struct {
	observers []Observer
}

func (w *WeatherForecastServer) RegisterObserver(observer Observer) {
	w.observers = append(w.observers, observer)
}

func (w *WeatherForecastServer) RemoveObserver(observer Observer) {
	for i, obs := range w.observers {
		if obs == observer {
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
			break
		}
	}
}

func (w *WeatherForecastServer) NotifyObservers(event WeatherChangeEvent) {
	for _, observer := range w.observers {
		observer.Update(event)
	}
}

func (w *WeatherForecastServer) SetMeasurements(event WeatherChangeEvent) {
	w.NotifyObservers(event)
}

// Client represents a client that acts as an observer
type Client struct {
	name string
}

func (c *Client) Update(event WeatherChangeEvent) {
	fmt.Printf("Client %s received weather update - Temperature: %.2f, Humidity: %.2f, Pressure: %.2f\n",
		c.name, event.Temperature, event.Humidity, event.Pressure)
	// Handle the additional indexes or properties as needed
}
