package temp

import (
	"errors"
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

	if !validateDistrubutionLocation(ele, locationData) {
		return errors.New("distrubution not allowed")
	}

	return nil
}

func validateDistrubutionLocation(ele distributor, locationData *locationData) bool {
	isAllowed := false
	if ele.parentDistributorName != "" {
		isAllowed = validateDistrubutionLocation(distributorData[ele.parentDistributorName], locationData)
	}

	if isAllowed {
		stateList, isStateOk := ele.excluded[locationData.countryName]
		if isStateOk && len(stateList) > 0 {

			cityList, isCityOk := ele.excluded[locationData.countryName][locationData.stateName]
			if isCityOk && len(cityList) > 0 {

				temp, isOk := ele.excluded[locationData.countryName][locationData.stateName][locationData.cityName]
				if isOk && temp != nil {
					return false
				}
				return true

			} else if isCityOk && len(cityList) == 0 {
				return false
			}
			return true

		} else if isStateOk && len(stateList) == 0 {
			return false
		}
		return true
	}

	stateList, isStateOk := ele.included[locationData.countryName]
	if isStateOk && len(stateList) > 0 {

		cityList, isCityOk := ele.included[locationData.countryName][locationData.stateName]
		if isCityOk && len(cityList) > 0 {

			temp, isOk := ele.included[locationData.countryName][locationData.stateName][locationData.cityName]
			if isOk && temp != nil {
				return true
			}
			return false

		} else if isCityOk && len(cityList) == 0 {
			return true
		}
		return false

	} else if isStateOk && len(stateList) == 0 {
		return true
	}
	return false

}
