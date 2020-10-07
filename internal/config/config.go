package config

//DEBUG struct to help debugging
type DEBUG struct {
	ASSET string
}
type ENV struct {
	URL string
}

//DebugData returns data to help debugging
func DebugData(isDebug bool) string {
	debug := new(DEBUG)
	debug.ASSET = "HGLG11"

	return debug.ASSET
}

func GetURL() string {
	env := new(ENV)
	env.URL = "https://statusinvest.com.br/fundos-imobiliarios/"

	return env.URL
}
