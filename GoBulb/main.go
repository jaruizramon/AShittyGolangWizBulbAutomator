package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// DEFAULT VARIABLES (SOFT CODABLE)
var WIZ_BULB_PORT = "38899"
var MS_LAG time.Duration = 10
var RGB_MAX float64 = 255
var RGB_MIN float64 = 0
var isRunning bool = true
var BRIGHTNESS = 100
var uniDim float64 = 100

var r float64 = 255
var g float64 = 0
var b float64 = 0

var light_maps_lol = map[string]int{
	"192.168.1.142": 1, // living room window bottom
	"192.168.1.105": 2, // living room window top
	"192.168.1.144": 3, // kitchen 1 away
	"192.168.1.103": 4, // ac living room bottom
	"192.168.1.141": 5, // ac living room top
	"192.168.1.143": 6, // kitchen to fridge
	"192.168.1.145": 7, // bathroom 2
	"192.168.1.106": 8, // bathroom 1
	"192.168.1.107": 9, // bedroom

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
		time.Sleep(25 * time.Millisecond)
		selected_light.Write([]byte(fmt.Sprintf(`{"method": "setPilot", "params":{"state": true, "dimming": %d }}`, 25)))
		time.Sleep(25 * time.Millisecond)
		selected_light.Write([]byte(fmt.Sprintf(`{"method": "setPilot", "params":{"state": true, "dimming": %d }}`, 65)))
		time.Sleep(25 * time.Millisecond)
		selected_light.Write([]byte(fmt.Sprintf(`{"method": "setPilot", "params":{"state": true, "dimming": %d }}`, 70)))
		time.Sleep(25 * time.Millisecond)
		selected_light.Write([]byte(fmt.Sprintf(`{"method": "setPilot", "params":{"state": true, "dimming": %f }}`, uniDim)))
		time.Sleep(25 * time.Millisecond)
		selected_light.Write([]byte(
			fmt.Sprintf(
				`{"method": "setPilot", "params":{"state": true, "r": %f, "g": %f, "b": %f}}`,
				r+rand.Float64()*(10 - -10),
				g+rand.Float64()*(10 - -10),
				b+rand.Float64()*(10 - -10))))
	}
	time.Sleep(3 * time.Second)
	fmt.Println("\n\nDefine final RGB\n\n")
	var in string = ""
	fmt.Scanln(&in)

	ss := strings.Split(in, ",")
	fmt.Println(len(ss))

	r, err := strconv.ParseFloat(ss[0], 64)
	if err != nil {
		r = 255.00
	}
	g, err := strconv.ParseFloat(ss[1], 64)
	if err != nil {
		g = 255.00
	}
	b, err := strconv.ParseFloat(ss[2], 64)
	if err != nil {
		b = 255.00
	}

	fmt.Println(r, g, b)
	fmt.Println("\n\n")

	fmt.Println("\n\n1. Disney  2. Rainbow Madness 3. Cascaron")
	in = ""
	fmt.Scanln(&in)

	if in == "1" {
		DisneyFlick()
	} else if in == "2" {
		RainbowMadness()
	} else if in == "3" {
		Cascaron()
	}

}

func FlickerLights(selected_light net.Conn, flickers int, lag time.Duration) {
	for i := 0; i <= flickers; i++ {
		selected_light.Write([]byte(fmt.Sprintf(`{"method": "setPilot", "params":{"state": true, "dimming": %f }}`, 100.00)))
		time.Sleep(lag * time.Millisecond)
		selected_light.Write([]byte(fmt.Sprintf(`{"method": "setPilot", "params":{"state": true, "dimming": %f }}`, 11.00)))
		time.Sleep(lag * time.Millisecond)
	}

}

func FlickerLightsRGBDimDefaultColour(selected_light net.Conn, flickers int, lag time.Duration, r1 float64, g1 float64, b1 float64) {
	for i := 0; i <= flickers; i++ {
		selected_light.Write([]byte(fmt.Sprintf(`{"method": "setPilot", "params":{"state": true,"speed":%f , "r": %f, "g": %f, "b": %f}}`, 100.00, r1, g1, b1)))
		selected_light.Write([]byte(fmt.Sprintf(`{"method": "setPilot", "params":{"state": true,"speed":%f , "dimming": %f }}`, 90.00, 100.00)))
		time.Sleep(lag * time.Millisecond)
		selected_light.Write([]byte(fmt.Sprintf(`{"method": "setPilot", "params":{"state": true,"speed":%f , "r": %f, "g": %f, "b": %f}}`, 100.00, r, g, b)))
		selected_light.Write([]byte(fmt.Sprintf(`{"method": "setPilot", "params":{"state": true,"speed":%f , "dimming": %f }}`, 90.00, 10.00)))
		time.Sleep(lag * time.Millisecond)
	}

}

func FlickerLightsRGBBright(selected_light net.Conn, flickers int, lag time.Duration, r1 float64, g1 float64, b1 float64, r2 float64, g2 float64, b2 float64) {
	for i := 0; i <= flickers; i++ {
		selected_light.Write([]byte(fmt.Sprintf(`{"method": "setPilot", "params":{"state": true,"speed":%f , "r": %f, "g": %f, "b": %f}}`, 100.00, r1, g1, b1)))
		selected_light.Write([]byte(fmt.Sprintf(`{"method": "setPilot", "params":{"state": true,"speed":%f , "dimming": %f }}`, 10.00, 10.00)))
		time.Sleep(lag * time.Millisecond)
		selected_light.Write([]byte(fmt.Sprintf(`{"method": "setPilot", "params":{"state": true,"speed":%f , "r": %f, "g": %f, "b": %f}}`, 100.00, r2, g2, b2)))
		selected_light.Write([]byte(fmt.Sprintf(`{"method": "setPilot", "params":{"state": true,"speed":%f , "dimming": %f }}`, 10.00, 100.00)))
		time.Sleep(lag * time.Millisecond)
	}

}

func FlickerLightsRGBRand(selected_light net.Conn, flickers int, lag time.Duration) {
	for i := 0; i <= flickers; i++ {
		r = GetRandomRangedFloat(55, 255)
		g = GetRandomRangedFloat(60, 155)
		b = GetRandomRangedFloat(40, 200)
		selected_light.Write([]byte(fmt.Sprintf(`{"method": "setPilot", "params":{"state": true,"speed":%f , "r": %f, "g": %f, "b": %f}}`, 100.00, r, g, b)))
		selected_light.Write([]byte(fmt.Sprintf(`{"method": "setPilot", "params":{"state": true,"speed":%f , "dimming": %f }}`, 10.00, 100.00)))
		time.Sleep(lag * time.Millisecond)

		r = GetRandomRangedFloat(0, 255)
		g = GetRandomRangedFloat(0, 255)
		b = GetRandomRangedFloat(0, 255)

		selected_light.Write([]byte(fmt.Sprintf(`{"method": "setPilot", "params":{"state": true,"speed":%f , "r": %f, "g": %f, "b": %f}}`, 100.00, r, g, b)))
		selected_light.Write([]byte(fmt.Sprintf(`{"method": "setPilot", "params":{"state": true,"speed":%f , "dimming": %f }}`, 10.00, 11.00)))
		time.Sleep(lag * time.Millisecond)
	}

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

			time.Sleep(100 * time.Millisecond)
			_, err = selected_light.Write([]byte(fmt.Sprintf(`{"method": "setPilot", "params":{"state": true, "dimming": %d }}`, 10)))
			if err != nil {
				fmt.Errorf("couldn't get response")
				continue
			}
			time.Sleep(100 * time.Millisecond)
			// fmt.Printf("--> LIGHT No. %d  (%s) --> \"r\": %f, \"g\": %f, \"b\": %f", bulbId, bulbIp, r, g, b)
		}
		if !isRunning {
			break
		}
	}
}

func Cascaron() {

	for isRunning {
		for bulbIp, _ := range light_maps_lol {

			selected_light, err := net.Dial("udp", fmt.Sprintf("%s:%s", bulbIp, WIZ_BULB_PORT))
			uniDim = 10
			selected_light.Write([]byte(fmt.Sprintf(`{"method": "setPilot", "params":{"state": true, "dimming": %f }}`, uniDim)))
			if err != nil {
				log.Panic()
			}

			selected_light.Write([]byte(
				fmt.Sprintf(
					`{"method": "setPilot", "params":{"state": true, "r": %f, "g": %f, "b": %f}}`,
					GetRandomRangedFloat(0, 255),
					GetRandomRangedFloat(0, 255),
					GetRandomRangedFloat(0, 255))))

			time.Sleep(250 * time.Millisecond)

			// fmt.Printf("--> LIGHT No. %d  (%s) --> \"r\": %f, \"g\": %f, \"b\": %f", bulbId, bulbIp, r, g, b)
		}
		for bulbIp, _ := range light_maps_lol {
			selected_light, err := net.Dial("udp", fmt.Sprintf("%s:%s", bulbIp, WIZ_BULB_PORT))

			//FlickerLights(selected_light, 40, 50)
			//FlickerLightsRGB(selected_light, 10, 250, 0, 204, 204, 255, 255, 0)
			//FlickerLightsRGBRand(selected_light, 10, 200)
			FlickerLightsRGBDimDefaultColour(selected_light, 10, 166, 255, 100, 50)

			selected_light.Write([]byte(fmt.Sprintf(`{"method": "setPilot", "params":{"state": true, "dimming": %f }}`, 10.00)))
			if err != nil {
				log.Panic()
			}

			selected_light.Write([]byte(
				fmt.Sprintf(
					`{"method": "setPilot", "params":{"state": true, "r": %f, "g": %f, "b": %f}}`,
					GetRandomRangedFloat(0, 255),
					GetRandomRangedFloat(0, 255),
					GetRandomRangedFloat(0, 255))))
			time.Sleep(10 * time.Millisecond)

			// fmt.Printf("--> LIGHT No. %d  (%s) --> \"r\": %f, \"g\": %f, \"b\": %f", bulbId, bulbIp, r, g, b)
		}
		if !isRunning {
			break
		}

		time.Sleep(3 * time.Second)
	}

}

func GetRandomRangedFloat(min float64, max float64) float64 {
	return rand.Float64() * ((max - min + 1) + min)
}

func RainbowMadness() {

	for isRunning {
		for bulbIp, _ := range light_maps_lol {

			r = rand.Float64() * (RGB_MAX - RGB_MIN)
			g = rand.Float64() * (RGB_MAX - RGB_MIN)
			b = rand.Float64() * (RGB_MAX - RGB_MIN)

			time.Sleep(MS_LAG * time.Millisecond)

			selected_light, err := net.Dial("udp", fmt.Sprintf("%s:%s", bulbIp, WIZ_BULB_PORT))
			uniDim = GetRandomRangedFloat(10, 100)
			selected_light.Write([]byte(fmt.Sprintf(`{"method": "setPilot", "params":{"state": true, "dimming": %f }}`, uniDim)))
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
