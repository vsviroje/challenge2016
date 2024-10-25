package temp

import (
	"errors"
	"fmt"
)

func AddCinemaLocToDistribution(distributorName string, cinemaLocation string, isExcludeData bool,
	isSubDistri bool, parentDistribName string) error {
	locationData, err := getLocationData(cinemaLocation)
	if err != nil {
		return err
	}

	ele := distributorData[distributorName]
	if isSubDistri {
		_, isOk := distributorData[parentDistribName]
		if !isOk {
			return errors.New("parentDistributorName does not exist")
		}
		ele.parentDistributorName = parentDistribName
	}

	if isExcludeData {
		if ele.excluded == nil {
			ele.excluded = make(map[string]map[string]map[string]interface{})
		}
		_, isOk := ele.excluded[locationData.countryName]
		if !isOk {
			ele.excluded[locationData.countryName] = make(map[string]map[string]interface{})
		}
		if locationData.stateName != "" {
			_, isOk := ele.excluded[locationData.countryName][locationData.stateName]
			if !isOk {
				ele.excluded[locationData.countryName][locationData.stateName] = make(map[string]interface{})
			}
		}
		if locationData.cityName != "" {
			_, isOk := ele.excluded[locationData.countryName][locationData.stateName][locationData.cityName]
			if !isOk {
				ele.excluded[locationData.countryName][locationData.stateName][locationData.cityName] = true
			}
		}
	} else {
		if ele.included == nil {
			ele.included = make(map[string]map[string]map[string]interface{})
		}
		_, isOk := ele.included[locationData.countryName]
		if !isOk {
			ele.included[locationData.countryName] = make(map[string]map[string]interface{})
		}
		if locationData.stateName != "" {
			_, isOk := ele.included[locationData.countryName][locationData.stateName]
			if !isOk {
				ele.included[locationData.countryName][locationData.stateName] = make(map[string]interface{})
			}
		}
		if locationData.cityName != "" {
			_, isOk := ele.included[locationData.countryName][locationData.stateName][locationData.cityName]
			if !isOk {
				ele.included[locationData.countryName][locationData.stateName][locationData.cityName] = true
			}
		}
	}

	distributorData[distributorName] = ele
	return nil
}

func IsDistributionAllowed(distributorName string, cinemaLocation string) error {
	ele, isOk := distributorData[distributorName]
	if !isOk {
		return errors.New("distributorName does not exist")
	}

	locationData, err := getLocationData(cinemaLocation)
	if err != nil {
		return err
	}

	locaValidDetails := validateDistrubutionLocation(ele, locationData)
	fmt.Println(locaValidDetails)
	if (locationData.len == 3 && !(locaValidDetails.countryLevel && locaValidDetails.stateLevel && locaValidDetails.cityLevel)) ||
		(locationData.len == 2 && !(locaValidDetails.countryLevel && locaValidDetails.stateLevel)) ||
		(locationData.len == 1 && !(locaValidDetails.countryLevel)) {
		return errors.New("distrubutor not allowed")
	}

	return nil
}

func validateDistrubutionLocation(ele distributor, locationData *locationData) *locationAllowed {
	var locAllowed *locationAllowed

	if ele.parentDistributorName != "" {
		locAllowed = validateDistrubutionLocation(distributorData[ele.parentDistributorName], locationData)
		fmt.Println(locAllowed)
	}

	if locAllowed == nil {
		locAllowed = &locationAllowed{}
	}

	if (locationData.len == 3 && (locAllowed.countryLevel && locAllowed.stateLevel && locAllowed.cityLevel)) ||
		(locationData.len == 2 && (locAllowed.countryLevel && locAllowed.stateLevel)) ||
		(locationData.len == 1 && (locAllowed.countryLevel)) ||
		ele.parentDistributorName == "" {

		locAllowed.countryLevel = true
		stateList, isCountryOk := ele.excluded[locationData.countryName]
		if isCountryOk && len(stateList) > 0 {
			locAllowed.stateLevel = true

			cityList, isStateOk := ele.excluded[locationData.countryName][locationData.stateName]
			if isStateOk && len(cityList) > 0 {
				locAllowed.cityLevel = true

				city, isCityOk := ele.excluded[locationData.countryName][locationData.stateName][locationData.cityName]
				if isCityOk && city != nil {
					locAllowed.cityLevel = false
					return locAllowed
				}
				locAllowed.cityLevel = false
				if ele.parentDistributorName != "" {
					return locAllowed
				}

			} else if isStateOk && len(cityList) == 0 {
				locAllowed.cityLevel = false
				return locAllowed
			}
			locAllowed.stateLevel = false
			if ele.parentDistributorName != "" {
				return locAllowed
			}
		} else if isCountryOk && len(stateList) == 0 {
			locAllowed.stateLevel = false
			return locAllowed
		}
		locAllowed.countryLevel = false
		if ele.parentDistributorName != "" {
			return locAllowed
		}
	}

	locAllowed.countryLevel = true
	stateList, isCountryOk := ele.included[locationData.countryName]
	if isCountryOk && len(stateList) > 0 {
		locAllowed.stateLevel = true

		cityList, isStateOk := ele.included[locationData.countryName][locationData.stateName]
		if isStateOk && len(cityList) > 0 {
			locAllowed.cityLevel = true

			city, isCityOk := ele.included[locationData.countryName][locationData.stateName][locationData.cityName]
			if isCityOk && city != nil {
				return locAllowed
			}
			locAllowed.cityLevel = false
			return locAllowed
		} else if isStateOk && len(cityList) == 0 {
			locAllowed.cityLevel = true
			return locAllowed
		}
		locAllowed.stateLevel = false
		return locAllowed
	} else if isCountryOk && len(stateList) == 0 {
		locAllowed.stateLevel = true
		return locAllowed
	}
	locAllowed.countryLevel = false
	return locAllowed

}
