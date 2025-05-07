package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("processTruck", func(t *testing.T) {
		t.Run("Should load and unload a truck cargo", func(t *testing.T) {
			nt := &NormalTruck{id: "1", cargo: 42}
			et := &EletricTruck{id: "2"}

			if err := processTruck(nt); err != nil {
				t.Fatalf("error processing truck: %s", err)
			}

			if err := processTruck(et); err != nil {
				t.Fatalf("error processing truck: %s", err)
			}

			// asserting
			if nt.cargo != 0 {
				t.Fatal("normal truck cargo should be 0")
			}

			if et.battery != -2 {
				t.Fatal("eletric truck battery should be -2")
			}
		})
	})
}
