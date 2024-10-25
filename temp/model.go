package temp

func init() {
	distributorData = make(map[string]distributor)
	distributorData = map[string]distributor{
		"Distributor1": {
			included: map[string]map[string]map[string]interface{}{
				"IN": nil,
				"CO": make(map[string]map[string]interface{}),
			},
			excluded: map[string]map[string]map[string]interface{}{
				"IN": {
					"TN": {
						"WLGTN": true,
					},
					"KA": nil,
				},
				"ID": nil,
			},
		},
		"Distributor2": {
			parentDistributorName: "Distributor1",
			included: map[string]map[string]map[string]interface{}{
				"IN": nil,
			},
			excluded: map[string]map[string]map[string]interface{}{
				"IN": {
					"TN": {
						"UDIPT": true,
					},
				},
			},
		},
		"Distributor3": {
			parentDistributorName: "Distributor2",
			included: map[string]map[string]map[string]interface{}{
				"IN": {
					"KA": {
						"YADGR": true,
					},
				},
			},
			excluded: nil,
		},
	}
}

var cityDataFromCsv map[string]map[string]map[string]interface{}
var distributorData map[string]distributor

type distributor struct {
	parentDistributorName string
	included              map[string]map[string]map[string]interface{}
	excluded              map[string]map[string]map[string]interface{}
}

type locationData struct {
	countryName string
	stateName   string
	cityName    string
	len         int
}

type locationAllowed struct {
	countryLevel bool
	stateLevel   bool
	cityLevel    bool
}
