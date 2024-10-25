package temp

import (
	"bufio"
	"encoding/csv"
	"errors"
	"log"
	"os"
	"strings"
)

func init() {
	prepareCityData("./cities.csv")
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func prepareCityData(filePath string) {
	lines := readCsvFile(filePath)
	cityDataFromCsv = make(map[string]map[string]map[string]interface{})

	for i := 1; i < len(lines); i++ {

		if _, ok := cityDataFromCsv[lines[i][2]]; !ok {
			cityDataFromCsv[lines[i][2]] = make(map[string]map[string]interface{})
		}

		if _, ok := cityDataFromCsv[lines[i][2]][lines[i][1]]; !ok {
			cityDataFromCsv[lines[i][2]][lines[i][1]] = make(map[string]interface{})
		}

		cityDataFromCsv[lines[i][2]][lines[i][1]][lines[i][0]] = true
	}
}

func GetInputFromConsole() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func getLocationData(cinemaLocation string) (*locationData, error) {
	var ele locationData

	data := strings.Split(cinemaLocation, "-")
	ele.len = len(data)

	if len(data) == 3 {
		ele.countryName = data[0]
		ele.stateName = data[1]
		ele.cityName = data[2]

	} else if len(data) == 2 {
		ele.countryName = data[0]
		ele.stateName = data[1]

	} else if len(data) == 1 {
		ele.countryName = data[0]
	} else {
		return nil, errors.New("empty cinemaLocation")
	}

	if ele.countryName != "" {
		if _, isOk := cityDataFromCsv[ele.countryName]; !isOk {
			return nil, errors.New("invalid country")
		}
	}

	if ele.stateName != "" {
		if _, isOk := cityDataFromCsv[ele.countryName][ele.stateName]; !isOk {
			return nil, errors.New("invalid state")
		}
	}

	if ele.cityName != "" {
		if _, isOk := cityDataFromCsv[ele.countryName][ele.stateName][ele.cityName]; !isOk {
			return nil, errors.New("invalid city")
		}
	}

	return &ele, nil
}
