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

var WIZ_BULB_PORT = "38899"
var MS_LAG time.Duration = 1
var RGB_MAX float64 = 255
var RGB_MIN float64 = 0
var isRunning bool = true
var BRIGHTNESS = 100

var r float64 = 167.00
var g float64 = 151.00
var b float64 = 116.00
var light_maps_lol = map[string]int{
	"192.168.1.127": 1, // living room window bottom
	"192.168.1.100": 2, // living room window top
	"192.168.1.126": 3, // ac living room bottom
	"192.168.1.101": 4, // ac living room top
	"192.168.1.124": 5, // kitchen to fridge
	"192.168.1.125": 6, // kitchen 1 away
	"192.168.1.128": 7, // bathroom 1
	"192.168.1.129": 8, // bathroom 2
	"192.168.1.121": 9, // bedroom

}
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

	fmt.Printf("\n-- STARTING %d WIZ LIGHTBULBS", len(light_maps_lol))

	// Connect to all lights
	for bulbIp, _ := range light_maps_lol {
		selected_light, err := net.Dial("udp", fmt.Sprintf("%s:%s", bulbIp, WIZ_BULB_PORT))
		if err != nil {
			log.Panic()
		}
		selected_light.Write([]byte(fmt.Sprintf(`{"method": "setPilot", "params":{"state": true, "dimming": %d }}`, 10)))
		selected_light.Write([]byte(`{"method": "setPilot", "params":{"state": true, "dimming": "` + "50" + `"}}`))
		selected_light.Write([]byte(`{"method": "setPilot", "params":{"state": true, "dimming": "` + "75" + `"}}`))
		selected_light.Write([]byte(`{"method": "setPilot", "params":{"state": true, "dimming": "` + "100" + `"}}`))
		selected_light.Write([]byte(
			fmt.Sprintf(
				`{"method": "setPilot", "params":{"state": true, "r": %f, "g": %f, "b": %f}}`,
				r+rand.Float64()*(10 - -10),
				g+rand.Float64()*(10 - -10),
				b+rand.Float64()*(10 - -10))))
	}
	time.Sleep(3 * time.Second)

	DisneyFlick()
}

func DisneyFlick() {

	for isRunning {
		for bulbIp, _ := range light_maps_lol {

			selected_light, err := net.Dial("udp", fmt.Sprintf("%s:%s", bulbIp, WIZ_BULB_PORT))
			if err != nil {
				log.Panic()
			}

			selected_light.Write([]byte(fmt.Sprintf(`{"method": "setPilot", "params":{"state": true, "dimming": %d }}`, 100)))
			selected_light.Write([]byte(
				fmt.Sprintf(
					`{"method": "setPilot", "params":{"state": true, "r": %f, "g": %f, "b": %f}}`,
					r+rand.Float64()*(10 - -10),
					g+rand.Float64()*(10 - -10),
					b+rand.Float64()*(10 - -10))))

			time.Sleep(250 * time.Millisecond)
			_, err = selected_light.Write([]byte(fmt.Sprintf(`{"method": "setPilot", "params":{"state": true, "dimming": %d }}`, 10)))
			if err != nil {
				fmt.Errorf("couldn't get response")
				continue
			}
			time.Sleep(250 * time.Millisecond)
			// fmt.Printf("--> LIGHT No. %d  (%s) --> \"r\": %f, \"g\": %f, \"b\": %f", bulbId, bulbIp, r, g, b)
		}
		if !isRunning {
			break
		}
	}
}

func RainbowMadness() {

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
