package main

import (
	"fmt"
	"log"

	"github.com/paypal/gatt"
	"github.com/paypal/gatt/examples/option"
	"github.com/suapapa/go_eddystone"
)

const (
	advTypeAllUUID16     = 0x03 // Complete List of 16-bit Service Class UUIDs
	advTypeServiceData16 = 0x16 // Service Data - 16-bit UUID
)

const (
	advFlagGeneralDiscoverable = 0x02
	advFlagLEOnly              = 0x04
)

func main() {
	f, err := eddystone.MakeURLFrame("http://google.com", -20)
	if err != nil {
		panic(err)
	}
	fmt.Println(f)

	// Advertise as an Eddystone beacon
	a := &gatt.AdvPacket{}
	if runtime.GOOS != "darwin" { // flag not set if darwin
		a.AppendFlags(advFlagGeneralDiscoverable | advFlagLEOnly)
	}
	a.AppendField(advTypeAllUUID16, eddystone.SvcUUIDBytes)
	a.AppendField(advTypeServiceData16, append(eddystone.SvcUUIDBytes, f...))

	fmt.Println(a.Len(), a)

	d, err := gatt.NewDevice(option.DefaultServerOptions...)
	if err != nil {
		log.Fatalf("Failed to open device, err: %s", err)
	}

	// Register optional handlers.
	d.Handle(
		gatt.CentralConnected(func(c gatt.Central) { fmt.Println("Connect: ", c.ID()) }),
		gatt.CentralDisconnected(func(c gatt.Central) { fmt.Println("Disconnect: ", c.ID()) }),
	)

	// A mandatory handler for monitoring device state.
	onStateChanged := func(d gatt.Device, s gatt.State) {
		fmt.Printf("State: %s\n", s)
		switch s {
		case gatt.StatePoweredOn:
			d.Advertise(a)
		default:
			log.Println(s)
		}
	}

	d.Init(onStateChanged)
	select {}
}
