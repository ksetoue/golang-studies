package main

import "fmt"

var lightSwitchIsOn bool = false

func main() {
	printLightSwitchState()
	toggleLightSwitch()
	printLightSwitchState()
	toggleLightSwitch()
	printLightSwitchState()
}

func printLightSwitchState() {
	fmt.Println("The lightswitch is off:", lightSwitchIsOn)
}

func toggleLightSwitch() {
	lightSwitchIsOn = !lightSwitchIsOn
}
