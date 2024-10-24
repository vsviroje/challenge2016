package temp

func init() {
	distributorData = make(map[string]distributor)
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
}
