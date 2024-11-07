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

// List of values that Currency can take.
//
//goland:noinspection ALL
const (
	CurrencyAED = "aed" // United Arab Emirates Dirham
	CurrencyAFN = "afn" // Afghan Afghani
	CurrencyALL = "all" // Albanian Lek
	CurrencyAMD = "amd" // Armenian Dram
	CurrencyANG = "ang" // Netherlands Antillean Gulden
	CurrencyAOA = "aoa" // Angolan Kwanza
	CurrencyARS = "ars" // Argentine Peso
	CurrencyAUD = "aud" // Australian Dollar
	CurrencyAWG = "awg" // Aruban Florin
	CurrencyAZN = "azn" // Azerbaijani Manat
	CurrencyBAM = "bam" // Bosnia & Herzegovina Convertible Mark
	CurrencyBBD = "bbd" // Barbadian Dollar
	CurrencyBDT = "bdt" // Bangladeshi Taka
	CurrencyBGN = "bgn" // Bulgarian Lev
	CurrencyBIF = "bif" // Burundian Franc
	CurrencyBMD = "bmd" // Bermudian Dollar
	CurrencyBND = "bnd" // Brunei Dollar
	CurrencyBOB = "bob" // Bolivian Boliviano
	CurrencyBRL = "brl" // Brazilian Real
	CurrencyBSD = "bsd" // Bahamian Dollar
	CurrencyBWP = "bwp" // Botswana Pula
	CurrencyBZD = "bzd" // Belize Dollar
	CurrencyCAD = "cad" // Canadian Dollar
	CurrencyCDF = "cdf" // Congolese Franc
	CurrencyCHF = "chf" // Swiss Franc
	CurrencyCLP = "clp" // Chilean Peso
	CurrencyCNY = "cny" // Chinese Renminbi Yuan
	CurrencyCOP = "cop" // Colombian Peso
	CurrencyCRC = "crc" // Costa Rican Colón
	CurrencyCVE = "cve" // Cape Verdean Escudo
	CurrencyCZK = "czk" // Czech Koruna
	CurrencyDJF = "djf" // Djiboutian Franc
	CurrencyDKK = "dkk" // Danish Krone
	CurrencyDOP = "dop" // Dominican Peso
	CurrencyDZD = "dzd" // Algerian Dinar
	CurrencyEEK = "eek" // Estonian Kroon
	CurrencyEGP = "egp" // Egyptian Pound
	CurrencyETB = "etb" // Ethiopian Birr
	CurrencyEUR = "eur" // Euro
	CurrencyFJD = "fjd" // Fijian Dollar
	CurrencyFKP = "fkp" // Falkland Islands Pound
	CurrencyGBP = "gbp" // British Pound
	CurrencyGEL = "gel" // Georgian Lari
	CurrencyGIP = "gip" // Gibraltar Pound
	CurrencyGMD = "gmd" // Gambian Dalasi
	CurrencyGNF = "gnf" // Guinean Franc
	CurrencyGTQ = "gtq" // Guatemalan Quetzal
	CurrencyGYD = "gyd" // Guyanese Dollar
	CurrencyHKD = "hkd" // Hong Kong Dollar
	CurrencyHNL = "hnl" // Honduran Lempira
	CurrencyHRK = "hrk" // Croatian Kuna
	CurrencyHTG = "htg" // Haitian Gourde
	CurrencyHUF = "huf" // Hungarian Forint
	CurrencyIDR = "idr" // Indonesian Rupiah
	CurrencyILS = "ils" // Israeli New Sheqel
	CurrencyINR = "inr" // Indian Rupee
	CurrencyISK = "isk" // Icelandic Króna
	CurrencyJMD = "jmd" // Jamaican Dollar
	CurrencyJPY = "jpy" // Japanese Yen
	CurrencyKES = "kes" // Kenyan Shilling
	CurrencyKGS = "kgs" // Kyrgyzstani Som
	CurrencyKHR = "khr" // Cambodian Riel
	CurrencyKMF = "kmf" // Comorian Franc
	CurrencyKRW = "krw" // South Korean Won
	CurrencyKYD = "kyd" // Cayman Islands Dollar
	CurrencyKZT = "kzt" // Kazakhstani Tenge
	CurrencyLAK = "lak" // Lao Kip
	CurrencyLBP = "lbp" // Lebanese Pound
	CurrencyLKR = "lkr" // Sri Lankan Rupee
	CurrencyLRD = "lrd" // Liberian Dollar
	CurrencyLSL = "lsl" // Lesotho Loti
	CurrencyLTL = "ltl" // Lithuanian Litas
	CurrencyLVL = "lvl" // Latvian Lats
	CurrencyMAD = "mad" // Moroccan Dirham
	CurrencyMDL = "mdl" // Moldovan Leu
	CurrencyMGA = "mga" // Malagasy Ariary
	CurrencyMKD = "mkd" // Macedonian Denar
	CurrencyMNT = "mnt" // Mongolian Tögrög
	CurrencyMOP = "mop" // Macanese Pataca
	CurrencyMRO = "mro" // Mauritanian Ouguiya
	CurrencyMUR = "mur" // Mauritian Rupee
	CurrencyMVR = "mvr" // Maldivian Rufiyaa
	CurrencyMWK = "mwk" // Malawian Kwacha
	CurrencyMXN = "mxn" // Mexican Peso
	CurrencyMYR = "myr" // Malaysian Ringgit
	CurrencyMZN = "mzn" // Mozambican Metical
	CurrencyNAD = "nad" // Namibian Dollar
	CurrencyNGN = "ngn" // Nigerian Naira
	CurrencyNIO = "nio" // Nicaraguan Córdoba
	CurrencyNOK = "nok" // Norwegian Krone
	CurrencyNPR = "npr" // Nepalese Rupee
	CurrencyNZD = "nzd" // New Zealand Dollar
	CurrencyPAB = "pab" // Panamanian Balboa
	CurrencyPEN = "pen" // Peruvian Nuevo Sol
	CurrencyPGK = "pgk" // Papua New Guinean Kina
	CurrencyPHP = "php" // Philippine Peso
	CurrencyPKR = "pkr" // Pakistani Rupee
	CurrencyPLN = "pln" // Polish Złoty
	CurrencyPYG = "pyg" // Paraguayan Guaraní
	CurrencyQAR = "qar" // Qatari Riyal
	CurrencyRON = "ron" // Romanian Leu
	CurrencyRSD = "rsd" // Serbian Dinar
	CurrencyRUB = "rub" // Russian Ruble
	CurrencyRWF = "rwf" // Rwandan Franc
	CurrencySAR = "sar" // Saudi Riyal
	CurrencySBD = "sbd" // Solomon Islands Dollar
	CurrencySCR = "scr" // Seychellois Rupee
	CurrencySEK = "sek" // Swedish Krona
	CurrencySGD = "sgd" // Singapore Dollar
	CurrencySHP = "shp" // Saint Helenian Pound
	CurrencySLL = "sll" // Sierra Leonean Leone
	CurrencySOS = "sos" // Somali Shilling
	CurrencySRD = "srd" // Surinamese Dollar
	CurrencySTD = "std" // São Tomé and Príncipe Dobra
	CurrencySVC = "svc" // Salvadoran Colón
	CurrencySZL = "szl" // Swazi Lilangeni
	CurrencyTHB = "thb" // Thai Baht
	CurrencyTJS = "tjs" // Tajikistani Somoni
	CurrencyTOP = "top" // Tongan Paʻanga
	CurrencyTRY = "try" // Turkish Lira
	CurrencyTTD = "ttd" // Trinidad and Tobago Dollar
	CurrencyTWD = "twd" // New Taiwan Dollar
	CurrencyTZS = "tzs" // Tanzanian Shilling
	CurrencyUAH = "uah" // Ukrainian Hryvnia
	CurrencyUGX = "ugx" // Ugandan Shilling
	CurrencyUSD = "usd" // United States Dollar
	CurrencyUYU = "uyu" // Uruguayan Peso
	CurrencyUZS = "uzs" // Uzbekistani Som
	CurrencyVEF = "vef" // Venezuelan Bolívar
	CurrencyVND = "vnd" // Vietnamese Đồng
	CurrencyVUV = "vuv" // Vanuatu Vatu
	CurrencyWST = "wst" // Samoan Tala
	CurrencyXAF = "xaf" // Central African Cfa Franc
	CurrencyXCD = "xcd" // East Caribbean Dollar
	CurrencyXOF = "xof" // West African Cfa Franc
	CurrencyXPF = "xpf" // Cfp Franc
	CurrencyYER = "yer" // Yemeni Rial
	CurrencyZAR = "zar" // South African Rand
	CurrencyZMW = "zmw" // Zambian Kwacha
)

var (
	CurrencyValidValues = []string{
		CurrencyAED,
		CurrencyAFN,
		CurrencyALL,
		CurrencyAMD,
		CurrencyANG,
		CurrencyAOA,
		CurrencyARS,
		CurrencyAUD,
		CurrencyAWG,
		CurrencyAZN,
		CurrencyBAM,
		CurrencyBBD,
		CurrencyBDT,
		CurrencyBGN,
		CurrencyBIF,
		CurrencyBMD,
		CurrencyBND,
		CurrencyBOB,
		CurrencyBRL,
		CurrencyBSD,
		CurrencyBWP,
		CurrencyBZD,
		CurrencyCAD,
		CurrencyCDF,
		CurrencyCHF,
		CurrencyCLP,
		CurrencyCNY,
		CurrencyCOP,
		CurrencyCRC,
		CurrencyCVE,
		CurrencyCZK,
		CurrencyDJF,
		CurrencyDKK,
		CurrencyDOP,
		CurrencyDZD,
		CurrencyEEK,
		CurrencyEGP,
		CurrencyETB,
		CurrencyEUR,
		CurrencyFJD,
		CurrencyFKP,
		CurrencyGBP,
		CurrencyGEL,
		CurrencyGIP,
		CurrencyGMD,
		CurrencyGNF,
		CurrencyGTQ,
		CurrencyGYD,
		CurrencyHKD,
		CurrencyHNL,
		CurrencyHRK,
		CurrencyHTG,
		CurrencyHUF,
		CurrencyIDR,
		CurrencyILS,
		CurrencyINR,
		CurrencyISK,
		CurrencyJMD,
		CurrencyJPY,
		CurrencyKES,
		CurrencyKGS,
		CurrencyKHR,
		CurrencyKMF,
		CurrencyKRW,
		CurrencyKYD,
		CurrencyKZT,
		CurrencyLAK,
		CurrencyLBP,
		CurrencyLKR,
		CurrencyLRD,
		CurrencyLSL,
		CurrencyLTL,
		CurrencyLVL,
		CurrencyMAD,
		CurrencyMDL,
		CurrencyMGA,
		CurrencyMKD,
		CurrencyMNT,
		CurrencyMOP,
		CurrencyMRO,
		CurrencyMUR,
		CurrencyMVR,
		CurrencyMWK,
		CurrencyMXN,
		CurrencyMYR,
		CurrencyMZN,
		CurrencyNAD,
		CurrencyNGN,
		CurrencyNIO,
		CurrencyNOK,
		CurrencyNPR,
		CurrencyNZD,
		CurrencyPAB,
		CurrencyPEN,
		CurrencyPGK,
		CurrencyPHP,
		CurrencyPKR,
		CurrencyPLN,
		CurrencyPYG,
		CurrencyQAR,
		CurrencyRON,
		CurrencyRSD,
		CurrencyRUB,
		CurrencyRWF,
		CurrencySAR,
		CurrencySBD,
		CurrencySEK,
		CurrencySGD,
		CurrencySHP,
		CurrencySLL,
		CurrencySOS,
		CurrencySRD,
		CurrencySTD,
		CurrencySVC,
		CurrencySZL,
		CurrencyTHB,
		CurrencyTJS,
		CurrencyTOP,
		CurrencyTRY,
		CurrencyTTD,
		CurrencyTWD,
		CurrencyTZS,
		CurrencyUAH,
		CurrencyUGX,
		CurrencyUSD,
		CurrencyUYU,
		CurrencyUZS,
		CurrencyVEF,
		CurrencyVND,
		CurrencyVUV,
		CurrencyWST,
		CurrencyXAF,
		CurrencyXCD,
		CurrencyXOF,
		CurrencyXPF,
		CurrencyYER,
		CurrencyZAR,
		CurrencyZMW,
	}
)

func IsValidCurrency(currency string) bool {

	for _, c := range CurrencyValidValues {
		if c == currency {
			return true
		}
	}

	return false
}
