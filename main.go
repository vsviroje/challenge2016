package main

import (
	"challenge2016/temp"
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Println("Welcome to cinema distributor validator")

		fmt.Println("Select from below operation:")
		fmt.Println("1.Add distributor")
		fmt.Println("2.Validate location of cinema for distributor")
		fmt.Println("enter 1 or 2 of choice:")

		in := temp.GetInputFromConsole()
		if in == "1" {
			fmt.Println("is sub distributor:(y)")
			in = temp.GetInputFromConsole()
			parentDis := ""
			if in == "y" {
				fmt.Println("enter name of parent distributor:")
				parentDis = temp.GetInputFromConsole()
			}

			fmt.Println("enter name of distributor:")
			disName := temp.GetInputFromConsole()

			fmt.Println("enter cinema location:")
			fmt.Println("Format:countryCode-stateCode-cityCode | eg: IN-MH-YEOLA")
			location := temp.GetInputFromConsole()

			fmt.Println("is above location is excluded:(y)")
			isExcluded := temp.GetInputFromConsole()

			err := temp.AddCinemaLocToDistribution(disName, location, isExcluded == "y", parentDis != "", parentDis)
			if err != nil {
				fmt.Println("Failed to add:", err.Error())
				continue
			}

			fmt.Println("Distributor successfully added")
		} else if in == "2" {

			fmt.Println("enter name of distributor:")
			disName := temp.GetInputFromConsole()

			fmt.Println("enter cinema location:")
			fmt.Println("Format:countryCode-stateCode-cityCode | eg: IN-MH-YEOLA")
			location := temp.GetInputFromConsole()

			err := temp.IsDistributionAllowed(disName, location)
			if err != nil {
				fmt.Println("Failed to validate:", err.Error())
				continue
			}
			fmt.Println("Distribution is allowed")
		} else {
			fmt.Println("wrong input:", in, len(in))
			os.Exit(-1)
		}
		fmt.Println("-----------------------------------------------------")
	}
}
