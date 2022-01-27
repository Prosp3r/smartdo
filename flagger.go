package main

import (
	"flag"
	"fmt"
	"os"
)

var flagSet = flag.NewFlagSet("trivaform", flag.ContinueOnError)

func init() { flagSet.Usage = usage }

//func init() { flagSet.Version = version }

var (
	versionNum = "0.1 beta - Prosp3r" //temporary placeholder for version information
)

func version() {
	fmt.Printf(` Trivaform Version: %s`, versionNum)
}

//usage displays command line usage information
func usage() {
	fmt.Fprintf(
		os.Stderr, `
Smartdo is a commandline tool for deploying and querying an ethereum ERC20 smart contract.

Usage:
		./smartdo <command> -[arguments]

sourcefile
		Source file location for private keyfile e.g ./wallet

Arguments are:
		h	Display this usage guide
		deploy  Deploys the smart contract to a blockchain network
		adduser Create a new ethereum blockchain compatible encrypted account

		v	Display version number

	smartdo

(c) 2022 smartdo all rights reserved.
		`)
}

var sourceFile string //= "..."
var defautlFileSrc = "hotel.csv"
var port int

//DEFINE FLAGS
var (
	h      = flag.Bool("h", false, "Display usage guide")
	ordern = flag.Bool("ordern", false, "Sort/Order alphabeticaly by name when writing to file")
	orderr = flag.Bool("orderr", false, "Sort/Order by rating score when writing to file")
	//src    = flag.String("src", sourceFile, "Source file location to be converted e.g. newfile.csv")
	stats = flag.Bool("stats", false, "Display stats at the end of operation")
	v     = flag.Bool("v", false, "Display version number")
)

func flagger() bool {
	//parse them
	// flag.StringVar(&sourceFile, "src", "hotel.csv", "Source file location to be converted e.g. newfile.csv")
	flag.Parse()

	colorx := []Color{ColorBlue, ColorYellow, ColorRed}
	messagex := []string{"S M A ", "R T  ", "D O"}

	if *h {
		fmt.Print("W E L C O M E   T O   ")
		trivColorize(colorx, messagex)
		usage()
		os.Exit(1)
	}

	if *v {
		trivColorize(colorx, messagex)
		version()
		os.Exit(1)
	}

	// //Determing alternative source file
	// if len(sourceFile) > 3 {
	// 	//fmt.Printf("Source provided : %v \n", sourceFile)
	// 	if _, err := os.Stat(sourceFile); err == nil {
	// 		//exists
	// 		fmt.Printf("Source provided : %v \n", sourceFile)
	// 		FileSrc = sourceFile
	// 		fmt.Printf("Attempting to read from: %v \n", FileSrc)
	// 	} else {
	// 		fmt.Printf("The file %v does not appear to exist. \n Please note that full path not supported by TrivaForm for now \n... bye now \n", sourceFile)
	// 		os.Exit(1)
	// 	}

	// } else {
	// 	fmt.Printf("No custom source provided \n")
	// 	FileSrc = defautlFileSrc
	// 	fmt.Printf("Attempting to read from default file source: %v \n", FileSrc)
	// }
	// //display valid and invalid entry count stats after operations
	// if *stats {
	// 	readStats = true
	// }

	// if *ordern {
	// 	OrderByName = true
	// }

	// if *orderr {
	// 	OrderByRatings = true
	// }
	// //Implement outfile flag

	//end outfile flag
	return true
}
