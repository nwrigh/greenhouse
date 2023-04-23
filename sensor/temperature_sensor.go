package sensor

import (
	"fmt"
	"github.com/d2r2/go-dht"
	"log"
)

type DHT22Sensor struct {
	Pin          int
	Retries      int
	IsFahrenheit bool
}

func (s *DHT22Sensor) Read() (float32, float32, error) {
	temp, hum, _, err := dht.ReadDHTxxWithRetry(dht.DHT22, s.Pin, s.IsFahrenheit, s.Retries)
	if err != nil {
		log.Printf("Error reading DHT22 sensor: %v", err)
		return 0, 0, err
	}

	return temp, hum, nil
}

func (s *DHT22Sensor) String(temp float32, hum float32) string {
	return fmt.Sprintf("Temperature: %.2f°C, Humidity: %.2f%%\n", temp, hum)
}
