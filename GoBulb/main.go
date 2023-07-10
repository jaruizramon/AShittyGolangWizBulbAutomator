package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

func main() {

	a_, err := net.Dial("udp", fmt.Sprintf("%s:%s", "192.168.1.101", "38899"))
	if err != nil {
		panic("could not connect to wiz light")
	} else {
		fmt.Println("Executed Bedroom")
	}
	b_, err := net.Dial("udp", fmt.Sprintf("%s:%s", "192.168.1.100", "38899"))
	if err != nil {
		panic("could not connect to wiz light")
	} else {
		fmt.Println("Executed Living Room A")
	}
	c_, err := net.Dial("udp", fmt.Sprintf("%s:%s", "192.168.1.121", "38899"))
	if err != nil {
		panic("could not connect to wiz light")
	} else {
		fmt.Println("Executed Living Room B")
	}

	// MAX BRIGHTNESS
	a_.Write([]byte(`{"method": "setPilot", "params":{"state": true, "dimming": 100}}`))
	b_.Write([]byte(`{"method": "setPilot", "params":{"state": true, "dimming": 100}}`))
	c_.Write([]byte(`{"method": "setPilot", "params":{"state": true, "dimming": 100}}`))

	var isRunning bool = true
	for isRunning {
		var (
			r float64 = 64
			g float64 = 128
			b float64 = 32
		)

		if r <= 0 || g <= 0 || b <= 0 {
			r = r + 32
			g = g + 64
			b = b + 40
		} else if r >= 255 || g >= 255 || b >= 255 {
			r = r + 15
			g = g + 45
			b = b + 70
		} else {
			r = rand.Float64() * 64
			g = rand.Float64() * 32
			b = rand.Float64() * 12
		}

		//--- GOTTA TURN THESE INTO ITERABLES TO MAKE IT MORE EFFICIENT AND MAKE THIS CODING WAR CRIME LESS CRIMINAL START ---.
		a_.Write([]byte(
			fmt.Sprintf(
				`{"method": "setPilot", "params":{"state": true, "r": %f, "g": %f, "b": %f}}`, r+64, g+32, b+16)))
		// fmt.Printf(`L1 - > {"method": "setPilot", "params":{"state": true, "r": %f, "g": %f, "b": %f}}`+"\n", r, g, b)
		b_.Write([]byte(
			fmt.Sprintf(
				`{"method": "setPilot", "params":{"state": true, "r": %f, "g": %f, "b": %f}}`, r+16, g+48, b+64)))
		// fmt.Printf(`L2 - > {"method": "setPilot", "params":{"state": true, "r": %f, "g": %f, "b": %f}}`+"\n", r, g, b)
		c_.Write([]byte(
			fmt.Sprintf(
				`{"method": "setPilot", "params":{"state": true, "r": %f, "g": %f, "b": %f}}`, r+32, g+8, b+16)))
		// fmt.Printf(`L3 - > {"method": "setPilot", "params":{"state": true, "r": %f, "g": %f, "b": %f}}`+"\n", r, g, b)

		//--- GOTTA TURN THESE INTO ITERABLES TO MAKE IT MORE EFFICIENT AND MAKE THIS CODING WAR CRIME LESS CRIMINAL END ---.
		time.Sleep(456 * time.Millisecond)
	}

}
