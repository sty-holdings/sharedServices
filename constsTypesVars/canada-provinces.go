// Package constant_type_vars
/*
This file contains USA states and postal codes

RESTRICTIONS:
	- Do not edit this comment section.

NOTES:
    To improve code readability, the constant names do not follow camelCase.
	Do not remove IDE inspection directives

COPYRIGHT and WARRANTY:
	Copyright 2022
	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.

*/
package constant_type_vars

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
