package config

//DEBUG struct to help debugging
type DEBUG struct {
	Asset                string
	DesiredMonthlyIncome string
}
type ENV struct {
	URL string
}

//DebugData returns data to help debugging
func DebugData() *DEBUG {
	debug := new(DEBUG)
	debug.Asset = "HGLG11"
	debug.DesiredMonthlyIncome = "1000"

	return debug
}

func GetURL() string {
	env := new(ENV)
	env.URL = "https://statusinvest.com.br/fundos-imobiliarios/"

	return env.URL
}
