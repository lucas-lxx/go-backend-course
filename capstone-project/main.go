package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	ErrTruckNotFound      = errors.New("Truck not found")
	ErrTruckAlreadyExists = errors.New("Truck already exists")
)

type FleetManager interface {
	AddTruck(id string, cargo int) error
	GetTruck(id string) (*Truck, error)
	RemoveTruck(id string) error
	UpdateTruck(id string, cargo int) error
}

type Truck struct {
	ID    string
	Cargo int
}

type truckManager struct {
	trucks map[string]*Truck
}

func NewTruckManager() *truckManager {
	return &truckManager{
		trucks: make(map[string]*Truck),
	}
}

func (t *truckManager) AddTruck(id string, cargo int) error {
	new_truck := &Truck{id, cargo}
	if _, ok := t.trucks[id]; ok {
		return ErrTruckAlreadyExists
	}
	t.trucks[id] = new_truck
	return nil
}

func (t *truckManager) GetTruck(id string) (*Truck, error) {
	truck, ok := t.trucks[id]
	if !ok {
		return nil, ErrTruckNotFound
	}
	return truck, nil
}

func (t *truckManager) RemoveTruck(id string) error {
	if _, ok := t.trucks[id]; ok {
		delete(t.trucks, id)
		return nil
	}
	return ErrTruckNotFound
}

func (t *truckManager) UpdateTruck(id string, cargo int) error {
	truck, err := t.GetTruck(id)
	if err != nil {
		return err
	}
	truck.Cargo = cargo
	return nil
}

func main() {
	var f FleetManager
	f = NewTruckManager()
	fmt.Println(f)
	fmt.Println("===The truck manager===")
	for {
		fmt.Println("options:")
		fmt.Println("create new truck:   c <id> <cargo>")
		fmt.Println("read truck:         r <id>")
		fmt.Println("update truck:       u <id> <cargo>")
		fmt.Println("delete truck:       d <id>\n")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		input = strings.TrimSpace(input)
		args := strings.Fields(input)
		switch args[0] {
		case "c":

		case "r":
		case "u":
		case "d":
		}
	}
}
