package main

import (
	"fmt"
	"greenhouse/repo"
	"greenhouse/sensor"
	"log"
	"time"
)

func main() {
	sensor := &sensor.DHT22Sensor{
		Pin:          7, // Replace with a valid GPIO pin number for your Raspberry Pi setup
		Retries:      10,
		IsFahrenheit: false,
	}

	writer := &repo.DataFileWriter{
		Filename: "temperature_data.txt",
	}

	for {
		temperature, humidity, err := sensor.Read()
		if err != nil {
			log.Printf("Error reading sensor data: %v", err)
		} else {
			data := fmt.Sprintf("Temperature: %.2f°C, Humidity: %.2f%%\n", temperature, humidity)
			fmt.Print(data)

			err := writer.Write(data)
			if err != nil {
				log.Printf("Error writing data to file: %v", err)
			}
		}

		time.Sleep(30 * time.Second)
	}
}
