package main

import (
	"errors"
	"fmt"
	"sync"
)

var (
	ErrTruckNotFound      = errors.New("Truck not found")
	ErrTruckAlreadyExists = errors.New("Truck already exists")
	ErrInvalidInput       = errors.New("One or more inputs were invalid")
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
	sync.RWMutex
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
	t.Lock()
	defer t.Unlock()
	t.trucks[id] = new_truck
	return nil
}

func (t *truckManager) GetTruck(id string) (*Truck, error) {
	t.RLock()
	defer t.RUnlock()
	truck, ok := t.trucks[id]
	if !ok {
		return nil, ErrTruckNotFound
	}
	return truck, nil
}

func (t *truckManager) RemoveTruck(id string) error {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.trucks[id]; ok {
		delete(t.trucks, id)
		return nil
	}
	return ErrTruckNotFound
}

func (t *truckManager) UpdateTruck(id string, cargo int) error {
	t.Lock()
	defer t.Unlock()
	truck, err := t.GetTruck(id)
	if err != nil {
		return err
	}
	truck.Cargo = cargo
	return nil
}

func (t *truckManager) ShowAllTrucks() {
	for _, val := range t.trucks {
		fmt.Printf("%v ", val)
	}
	fmt.Println()
}

// func main() {
// 	var f FleetManager
// 	f = NewTruckManager()
// 	fmt.Println(f)
// 	fmt.Println("===The truck manager===")
// 	for {
// 		fmt.Println("options:")
// 		fmt.Println("create new truck:   c <id> <cargo>")
// 		fmt.Println("read truck:         r <id>")
// 		fmt.Println("update truck:       u <id> <cargo>")
// 		fmt.Println("delete truck:       d <id>")
// 		reader := bufio.NewReader(os.Stdin)
// 		input, err := reader.ReadString('\n')
// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 		input = strings.TrimSpace(input)
// 		args := strings.Fields(input)
// 		cmd, id := args[0], args[1]
// 		var cargo int
// 		if cmd == "c" || cmd == "u" {
// 			cargo, err = strconv.Atoi(args[2])
// 			if err != nil {
// 				fmt.Errorf("%v\n", ErrInvalidInput.Error())
// 				continue
// 			}
// 		}
// 		switch args[0] {
// 		case "c":
// 			f.AddTruck(id, cargo)
// 		case "r":
// 			t, err := f.GetTruck(id)
// 			if err != nil {
// 				fmt.Println(ErrTruckNotFound.Error())
// 				continue
// 			}
// 			fmt.Printf("%v\n\n", *t)
// 		case "u":
// 			f.UpdateTruck(id, cargo)
// 		case "d":
// 			f.RemoveTruck(id)
// 		default:
// 			fmt.Errorf("%v\n", ErrInvalidInput.Error())
// 		}
// 	}
// }
