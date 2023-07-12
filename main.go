package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func main() {

	var WIZ_BULB_PORT = "38899"
	var MS_LAG time.Duration = 5
	var RGB_MAX float64 = 255
	var RGB_MIN float64 = 0
	var isRunning bool = true
	var BRIGHTNESS = "100"

	light_maps_lol := map[string]int{
		"192.168.1.101": 1,
		"192.168.1.124": 2,
		"192.168.1.121": 3,
		"192.168.1.126": 4,
		"192.168.1.125": 5,
		"192.168.1.127": 6,
		"192.168.1.100": 7,
		"192.168.1.128": 8,
		"192.168.1.129": 9,
	}

	fmt.Printf("\n-- STARTING %d WIZ LIGHTBULBS", len(light_maps_lol))

	var r float64 = 167.00
	var g float64 = 151.00
	var b float64 = 116.00

	// Connect to all lights
	for bulbIp, _ := range light_maps_lol {
		selected_light, err := net.Dial("udp", fmt.Sprintf("%s:%s", bulbIp, WIZ_BULB_PORT))
		if err != nil {
			log.Panic()
		}
		selected_light.Write([]byte(`{"method": "setPilot", "params":{"state": true, "dimming": "` + BRIGHTNESS + "}}"))
		selected_light.Write([]byte(
			fmt.Sprintf(
				`{"method": "setPilot", "params":{"state": true, "r": %f, "g": %f, "b": %f}}`,
				r+rand.Float64()*(10 - -10),
				g-rand.Float64()*(10 - -10),
				b+rand.Float64()*(10 - -10))))

	}

	time.Sleep(3 * time.Second)

	for isRunning {
		for bulbIp, _ := range light_maps_lol {

			r = rand.Float64() * (RGB_MAX - RGB_MIN)
			g = rand.Float64() * (RGB_MAX - RGB_MIN)
			b = rand.Float64() * (RGB_MAX - RGB_MIN)

			time.Sleep(MS_LAG * time.Millisecond)

			selected_light, err := net.Dial("udp", fmt.Sprintf("%s:%s", bulbIp, WIZ_BULB_PORT))
			if err != nil {
				log.Panic()
			}
			time.Sleep(MS_LAG * time.Millisecond)

			selected_light.Write([]byte(
				fmt.Sprintf(
					`{"method": "setPilot", "params":{"state": true, "r": %f, "g": %f, "b": %f}}`, r, g, b)))

			// fmt.Printf("--> LIGHT No. %d  (%s) --> \"r\": %f, \"g\": %f, \"b\": %f", bulbId, bulbIp, r, g, b)

		}

		if !isRunning {
			break
		}

	}

}
