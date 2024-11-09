package sharedServices

//goland:noinspection ALL
const (
	CAN_AB                   = "AB" // Alberta
	CAN_BC                   = "BC" // British Columbia
	CAN_MB                   = "MB" // Manitoba
	CAN_NB                   = "NB" // New Brunswick
	CAN_NL                   = "NL" // Newfoundland and Labrador
	CAN_NT                   = "NT" // Northwest Territories
	CAN_NS                   = "NS" // Nova Scotia
	CAN_NU                   = "NU" // Nunavut
	CAN_ON                   = "ON" // Ontario
	CAN_PE                   = "PE" // Prince Edward Island
	CAN_QC                   = "QC" // Quebec
	CAN_SK                   = "SK" // Saskatchewan
	CAN_YT                   = "YT" // Yukon
	CAN_ALL_PROVIDENCE_CODES = "AB BC MB NB NL NT NS NU ON PE QC SK YT"
)

type CanadaProvinceInfo struct {
	englishName string
	frenchName  string
}

var (
	canadianProvinceInfo = map[string]CanadaProvinceInfo{
		"AB": {"Alberta", ""},
		"BC": {"British Columbia", ""},
		"MB": {"Manitoba", ""},
		"NB": {"New Brunswick", ""},
		"NL": {"Newfoundland and Labrador", ""},
		"NT": {"Northwest Territories", ""},
		"NS": {"Nova Scotia", ""},
		"NU": {"Nunavut", ""},
		"ON": {"Ontario", ""},
		"PE": {"Prince Edward Island", ""},
		"QC": {"Quebec", ""},
		"SK": {"Saskatchewan", ""},
		"YT": {"Yukon", ""},
	}
)

func GetEnglishProvinceName(postalCode string) string {
	return canadianProvinceInfo[postalCode].englishName
}

func GetFrenchProvinceName(postalCode string) string {
	return canadianProvinceInfo[postalCode].frenchName
}
