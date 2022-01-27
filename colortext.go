package main

import "fmt"

//:::::::::::::::::::::::::: AESTETICS:::::::::::::::::::::

//Color type for color coded text messages for aestetics
type Color string

const (
	//ColorBlack ..
	ColorBlack Color = "\u001b[30m"
	//ColorRed ...for errors
	ColorRed = "\u001b[31m"
	//ColorGreen ...for successes
	ColorGreen = "\u001b[32m"
	//ColorYellow ...for warnings
	ColorYellow = "\u001b[33m"
	//ColorBlue ...for instructions
	ColorBlue = "\u001b[34m"
	//ColorReset ...
	ColorReset = "\u001b[0m"
)

//colorize ...for aestetics
func colorize(color Color, message string) {
	fmt.Print(string(color), message, string(ColorReset))
	//return string(color)
}

//trivColorize applies Colors to string
func trivColorize(color []Color, message []string) error {
	for i, v := range color {
		colorize(v, message[i])
		//msg = append(msg, colorize(v, message[i]))
	}
	return nil
}
