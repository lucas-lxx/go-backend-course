package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

type contextKey string

var UserIDKey contextKey = "userID"

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
)

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

type NormalTruck struct {
	id    string
	cargo int
}

func (t *NormalTruck) LoadCargo() error {
	t.cargo += 1
	return nil
}

func (t *NormalTruck) UnloadCargo() error {
	t.cargo = 0
	return nil
}

type EletricTruck struct {
	id      string
	cargo   int
	battery float64
}

func (e *EletricTruck) LoadCargo() error {
	e.cargo += 1
	e.battery += (-1)
	return nil
}

func (e *EletricTruck) UnloadCargo() error {
	e.cargo = 0
	e.battery += (-1)
	return nil
}

func processTruck(ctx context.Context, truck Truck) error {
	log.Printf("processing... truck %+v\n", truck)

	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	// Simulate some processing time
	time.Sleep(time.Second)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("error unloading cargo: %w", err)
	}

	log.Printf("processed truck %+v\n", truck)

	return nil
}

func processFleet(ctx context.Context, trucks []Truck) error {
	var wg sync.WaitGroup

	for _, t := range trucks {
		wg.Add(1)
		go func(t Truck) {
			if err := processTruck(ctx, t); err != nil {
				log.Println(err)
			}
			wg.Done()
		}(t)
	}
	wg.Wait()

	return nil
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, UserIDKey, 42)
	// var normalTruck *NormalTruck = &NormalTruck{id: "1"}
	fleet := []Truck{
		&NormalTruck{id: "NT1", cargo: 0},
		&EletricTruck{id: "ET2", cargo: 0, battery: 100},
		&NormalTruck{id: "NT3", cargo: 0},
		&EletricTruck{id: "ET4", cargo: 0, battery: 100},
	}

	processFleet(ctx, fleet)
	log.Println("All trucks were processed")
}
