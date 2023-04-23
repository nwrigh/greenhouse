package main

import (
	"context"
	"fmt"
	"greenhouse/repo"
	"greenhouse/sensor"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
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
	for {
		select {
		case <-ctx.Done():
			// Cleanup resources here, if necessary
			fmt.Println("Shutting down gracefully...")
			return
		default:
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
			fmt.Println("Performing task...")
			time.Sleep(time.Second * 5)
		}
	}
}
