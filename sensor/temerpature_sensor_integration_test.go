//go:build integration
// +build integration

package sensor

import (
	"errors"
	"testing"
)

func TestDHT22Sensor(t *testing.T) {
	tests := []struct {
		name          string
		pin           int
		retries       int
		isFahrenheit  bool
		expectedError error
	}{
		{
			name:          "Valid pin",
			pin:           4, // Replace with a valid GPIO pin number for your Raspberry Pi setup
			retries:       10,
			isFahrenheit:  false,
			expectedError: nil,
		},
		{
			name:          "Invalid pin",
			pin:           -1,
			retries:       10,
			isFahrenheit:  false,
			expectedError: errors.New("Error reading DHT22 sensor"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sensor := &DHT22Sensor{
				Pin:          tt.pin,
				Retries:      tt.retries,
				IsFahrenheit: tt.isFahrenheit,
			}

			_, _, err := sensor.Read()

			if tt.expectedError != nil && err == nil {
				t.Errorf("Expected error, but got none")
			}

			if tt.expectedError == nil && err != nil {
				t.Errorf("Error reading DHT22 sensor: %v", err)
			}
		})
	}
}
