package main

import "os"
import (
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/brutella/log"
)

func turnLightOn() {
	log.Println("Turn Light On")
}

func turnLightOff() {
	log.Println("Turn Light Off")
}

func main() {
	log.Verbose = true
	log.Info = true

	info := accessory.Info{
		Name:         "Personal Light Bulb",
		Manufacturer: "Matthias",
	}

	acc := accessory.NewLightbulb(info)

	acc.Lightbulb.On.OnValueRemoteUpdate(func(on bool) {
		if on == true {
			turnLightOn()
		} else {
			turnLightOff()
		}
	})

	t, err := hc.NewIPTransport(hc.Config{Pin: "32191123", Port: os.Getenv("PORT"), IP: os.Getenv("IP")}, acc.Accessory)
	if err != nil {
		log.Fatal(err)
	}

	hc.OnTermination(func() {
		t.Stop()
	})

	t.Start()
}
