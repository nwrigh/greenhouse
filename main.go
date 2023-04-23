package main

import (
	"context"
	"fmt"
	"greenhouse/repo"
	"greenhouse/sensor"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	dataPin      = 4
	outputFile   = "/data/temperature_data.txt"
	readInterval = 30 * time.Second
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	// Start your main processing function as a goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		processData(ctx)
	}()

	// Set up signal capturing
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Wait for a signal
	<-sigChan

	// Cancel the context to stop the processData function
	cancel()

	// Wait for the processData function to complete
	wg.Wait()

	fmt.Println("Graceful shutdown complete")
}

func processData(ctx context.Context) {
	sensor := &sensor.DHT22Sensor{
		Pin:          dataPin,
		Retries:      10,
		IsFahrenheit: false,
	}

	writer := &repo.DataFileWriter{
		Filename: outputFile,
	}

	for {
		select {
		case <-ctx.Done():
			// Cleanup resources here, if necessary
			fmt.Println("Shutting down gracefully...")
			return
		default:
			// Read temperature and humidity from the DHT22 sensor
			temperature, humidity, err := sensor.Read()
			if err != nil {
				fmt.Printf("Error reading sensor data: %v\n", err)
			} else {
				// Log the sensor data
				fmt.Printf("Temperature: %.2f °C, Humidity: %.2f %%\n", temperature, humidity)

				// Write the sensor data to the output file
				data := fmt.Sprintf("%s Temperature: %.2f °C, Humidity: %.2f %%\n", time.Now().Format(time.RFC3339), temperature, humidity)
				err := writer.Write(data)
				if err != nil {
					fmt.Printf("Error writing sensor data to file: %v\n", err)
				}
			}

			// Sleep for the specified interval before the next reading
			time.Sleep(readInterval)
		}
	}
}
