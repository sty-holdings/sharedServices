package sharedServices

// List of values that Currency can take.
//
//goland:noinspection ALL
const (
	CURRENCY_AED = "aed" // United Arab Emirates Dirham
	CURRENCY_AFN = "afn" // Afghan Afghani
	CURRENCY_ALL = "all" // Albanian Lek
	CURRENCY_AMD = "amd" // Armenian Dram
	CURRENCY_ANG = "ang" // Netherlands Antillean Gulden
	CURRENCY_AOA = "aoa" // Angolan Kwanza
	CURRENCY_ARS = "ars" // Argentine Peso
	CURRENCY_AUD = "aud" // Australian Dollar
	CURRENCY_AWG = "awg" // Aruban Florin
	CURRENCY_AZN = "azn" // Azerbaijani Manat
	CURRENCY_BAM = "bam" // Bosnia & Herzegovina Convertible Mark
	CURRENCY_BBD = "bbd" // Barbadian Dollar
	CURRENCY_BDT = "bdt" // Bangladeshi Taka
	CURRENCY_BGN = "bgn" // Bulgarian Lev
	CURRENCY_BIF = "bif" // Burundian Franc
	CURRENCY_BMD = "bmd" // Bermudian Dollar
	CURRENCY_BND = "bnd" // Brunei Dollar
	CURRENCY_BOB = "bob" // Bolivian Boliviano
	CURRENCY_BRL = "brl" // Brazilian Real
	CURRENCY_BSD = "bsd" // Bahamian Dollar
	CURRENCY_BWP = "bwp" // Botswana Pula
	CURRENCY_BZD = "bzd" // Belize Dollar
	CURRENCY_CAD = "cad" // Canadian Dollar
	CURRENCY_CDF = "cdf" // Congolese Franc
	CURRENCY_CHF = "chf" // Swiss Franc
	CURRENCY_CLP = "clp" // Chilean Peso
	CURRENCY_CNY = "cny" // Chinese Renminbi Yuan
	CURRENCY_COP = "cop" // Colombian Peso
	CURRENCY_CRC = "crc" // Costa Rican Colón
	CURRENCY_CVE = "cve" // Cape Verdean Escudo
	CURRENCY_CZK = "czk" // Czech Koruna
	CURRENCY_DJF = "djf" // Djiboutian Franc
	CURRENCY_DKK = "dkk" // Danish Krone
	CURRENCY_DOP = "dop" // Dominican Peso
	CURRENCY_DZD = "dzd" // Algerian Dinar
	CURRENCY_EEK = "eek" // Estonian Kroon
	CURRENCY_EGP = "egp" // Egyptian Pound
	CURRENCY_ETB = "etb" // Ethiopian Birr
	CURRENCY_EUR = "eur" // Euro
	CURRENCY_FJD = "fjd" // Fijian Dollar
	CURRENCY_FKP = "fkp" // Falkland Islands Pound
	CURRENCY_GBP = "gbp" // British Pound
	CURRENCY_GEL = "gel" // Georgian Lari
	CURRENCY_GIP = "gip" // Gibraltar Pound
	CURRENCY_GMD = "gmd" // Gambian Dalasi
	CURRENCY_GNF = "gnf" // Guinean Franc
	CURRENCY_GTQ = "gtq" // Guatemalan Quetzal
	CURRENCY_GYD = "gyd" // Guyanese Dollar
	CURRENCY_HKD = "hkd" // Hong Kong Dollar
	CURRENCY_HNL = "hnl" // Honduran Lempira
	CURRENCY_HRK = "hrk" // Croatian Kuna
	CURRENCY_HTG = "htg" // Haitian Gourde
	CURRENCY_HUF = "huf" // Hungarian Forint
	CURRENCY_IDR = "idr" // Indonesian Rupiah
	CURRENCY_ILS = "ils" // Israeli New Sheqel
	CURRENCY_INR = "inr" // Indian Rupee
	CURRENCY_ISK = "isk" // Icelandic Króna
	CURRENCY_JMD = "jmd" // Jamaican Dollar
	CURRENCY_JPY = "jpy" // Japanese Yen
	CURRENCY_KES = "kes" // Kenyan Shilling
	CURRENCY_KGS = "kgs" // Kyrgyzstani Som
	CURRENCY_KHR = "khr" // Cambodian Riel
	CURRENCY_KMF = "kmf" // Comorian Franc
	CURRENCY_KRW = "krw" // South Korean Won
	CURRENCY_KYD = "kyd" // Cayman Islands Dollar
	CURRENCY_KZT = "kzt" // Kazakhstani Tenge
	CURRENCY_LAK = "lak" // Lao Kip
	CURRENCY_LBP = "lbp" // Lebanese Pound
	CURRENCY_LKR = "lkr" // Sri Lankan Rupee
	CURRENCY_LRD = "lrd" // Liberian Dollar
	CURRENCY_LSL = "lsl" // Lesotho Loti
	CURRENCY_LTL = "ltl" // Lithuanian Litas
	CURRENCY_LVL = "lvl" // Latvian Lats
	CURRENCY_MAD = "mad" // Moroccan Dirham
	CURRENCY_MDL = "mdl" // Moldovan Leu
	CURRENCY_MGA = "mga" // Malagasy Ariary
	CURRENCY_MKD = "mkd" // Macedonian Denar
	CURRENCY_MNT = "mnt" // Mongolian Tögrög
	CURRENCY_MOP = "mop" // Macanese Pataca
	CURRENCY_MRO = "mro" // Mauritanian Ouguiya
	CURRENCY_MUR = "mur" // Mauritian Rupee
	CURRENCY_MVR = "mvr" // Maldivian Rufiyaa
	CURRENCY_MWK = "mwk" // Malawian Kwacha
	CURRENCY_MXN = "mxn" // Mexican Peso
	CURRENCY_MYR = "myr" // Malaysian Ringgit
	CURRENCY_MZN = "mzn" // Mozambican Metical
	CURRENCY_NAD = "nad" // Namibian Dollar
	CURRENCY_NGN = "ngn" // Nigerian Naira
	CURRENCY_NIO = "nio" // Nicaraguan Córdoba
	CURRENCY_NOK = "nok" // Norwegian Krone
	CURRENCY_NPR = "npr" // Nepalese Rupee
	CURRENCY_NZD = "nzd" // New Zealand Dollar
	CURRENCY_PAB = "pab" // Panamanian Balboa
	CURRENCY_PEN = "pen" // Peruvian Nuevo Sol
	CURRENCY_PGK = "pgk" // Papua New Guinean Kina
	CURRENCY_PHP = "php" // Philippine Peso
	CURRENCY_PKR = "pkr" // Pakistani Rupee
	CURRENCY_PLN = "pln" // Polish Złoty
	CURRENCY_PYG = "pyg" // Paraguayan Guaraní
	CURRENCY_QAR = "qar" // Qatari Riyal
	CURRENCY_RON = "ron" // Romanian Leu
	CURRENCY_RSD = "rsd" // Serbian Dinar
	CURRENCY_RUB = "rub" // Russian Ruble
	CURRENCY_RWF = "rwf" // Rwandan Franc
	CURRENCY_SAR = "sar" // Saudi Riyal
	CURRENCY_SBD = "sbd" // Solomon Islands Dollar
	CURRENCY_SCR = "scr" // Seychellois Rupee
	CURRENCY_SEK = "sek" // Swedish Krona
	CURRENCY_SGD = "sgd" // Singapore Dollar
	CURRENCY_SHP = "shp" // Saint Helenian Pound
	CURRENCY_SLL = "sll" // Sierra Leonean Leone
	CURRENCY_SOS = "sos" // Somali Shilling
	CURRENCY_SRD = "srd" // Surinamese Dollar
	CURRENCY_STD = "std" // São Tomé and Príncipe Dobra
	CURRENCY_SVC = "svc" // Salvadoran Colón
	CURRENCY_SZL = "szl" // Swazi Lilangeni
	CURRENCY_THB = "thb" // Thai Baht
	CURRENCY_TJS = "tjs" // Tajikistani Somoni
	CURRENCY_TOP = "top" // Tongan Paʻanga
	CURRENCY_TRY = "try" // Turkish Lira
	CURRENCY_TTD = "ttd" // Trinidad and Tobago Dollar
	CURRENCY_TWD = "twd" // New Taiwan Dollar
	CURRENCY_TZS = "tzs" // Tanzanian Shilling
	CURRENCY_UAH = "uah" // Ukrainian Hryvnia
	CURRENCY_UGX = "ugx" // Ugandan Shilling
	CURRENCY_USD = "usd" // United States Dollar
	CURRENCY_UYU = "uyu" // Uruguayan Peso
	CURRENCY_UZS = "uzs" // Uzbekistani Som
	CURRENCY_VEF = "vef" // Venezuelan Bolívar
	CURRENCY_VND = "vnd" // Vietnamese Đồng
	CURRENCY_VUV = "vuv" // Vanuatu Vatu
	CURRENCY_WST = "wst" // Samoan Tala
	CURRENCY_XAF = "xaf" // Central African Cfa Franc
	CURRENCY_XCD = "xcd" // East Caribbean Dollar
	CURRENCY_XOF = "xof" // West African Cfa Franc
	CURRENCY_XPF = "xpf" // Cfp Franc
	CURRENCY_YER = "yer" // Yemeni Rial
	CURRENCY_ZAR = "zar" // South African Rand
	CURRENCY_ZMW = "zmw" // Zambian Kwacha
)

var (
	CurrencyValidValues = []string{
		CURRENCY_AED,
		CURRENCY_AFN,
		CURRENCY_ALL,
		CURRENCY_AMD,
		CURRENCY_ANG,
		CURRENCY_AOA,
		CURRENCY_ARS,
		CURRENCY_AUD,
		CURRENCY_AWG,
		CURRENCY_AZN,
		CURRENCY_BAM,
		CURRENCY_BBD,
		CURRENCY_BDT,
		CURRENCY_BGN,
		CURRENCY_BIF,
		CURRENCY_BMD,
		CURRENCY_BND,
		CURRENCY_BOB,
		CURRENCY_BRL,
		CURRENCY_BSD,
		CURRENCY_BWP,
		CURRENCY_BZD,
		CURRENCY_CAD,
		CURRENCY_CDF,
		CURRENCY_CHF,
		CURRENCY_CLP,
		CURRENCY_CNY,
		CURRENCY_COP,
		CURRENCY_CRC,
		CURRENCY_CVE,
		CURRENCY_CZK,
		CURRENCY_DJF,
		CURRENCY_DKK,
		CURRENCY_DOP,
		CURRENCY_DZD,
		CURRENCY_EEK,
		CURRENCY_EGP,
		CURRENCY_ETB,
		CURRENCY_EUR,
		CURRENCY_FJD,
		CURRENCY_FKP,
		CURRENCY_GBP,
		CURRENCY_GEL,
		CURRENCY_GIP,
		CURRENCY_GMD,
		CURRENCY_GNF,
		CURRENCY_GTQ,
		CURRENCY_GYD,
		CURRENCY_HKD,
		CURRENCY_HNL,
		CURRENCY_HRK,
		CURRENCY_HTG,
		CURRENCY_HUF,
		CURRENCY_IDR,
		CURRENCY_ILS,
		CURRENCY_INR,
		CURRENCY_ISK,
		CURRENCY_JMD,
		CURRENCY_JPY,
		CURRENCY_KES,
		CURRENCY_KGS,
		CURRENCY_KHR,
		CURRENCY_KMF,
		CURRENCY_KRW,
		CURRENCY_KYD,
		CURRENCY_KZT,
		CURRENCY_LAK,
		CURRENCY_LBP,
		CURRENCY_LKR,
		CURRENCY_LRD,
		CURRENCY_LSL,
		CURRENCY_LTL,
		CURRENCY_LVL,
		CURRENCY_MAD,
		CURRENCY_MDL,
		CURRENCY_MGA,
		CURRENCY_MKD,
		CURRENCY_MNT,
		CURRENCY_MOP,
		CURRENCY_MRO,
		CURRENCY_MUR,
		CURRENCY_MVR,
		CURRENCY_MWK,
		CURRENCY_MXN,
		CURRENCY_MYR,
		CURRENCY_MZN,
		CURRENCY_NAD,
		CURRENCY_NGN,
		CURRENCY_NIO,
		CURRENCY_NOK,
		CURRENCY_NPR,
		CURRENCY_NZD,
		CURRENCY_PAB,
		CURRENCY_PEN,
		CURRENCY_PGK,
		CURRENCY_PHP,
		CURRENCY_PKR,
		CURRENCY_PLN,
		CURRENCY_PYG,
		CURRENCY_QAR,
		CURRENCY_RON,
		CURRENCY_RSD,
		CURRENCY_RUB,
		CURRENCY_RWF,
		CURRENCY_SAR,
		CURRENCY_SBD,
		CURRENCY_SEK,
		CURRENCY_SGD,
		CURRENCY_SHP,
		CURRENCY_SLL,
		CURRENCY_SOS,
		CURRENCY_SRD,
		CURRENCY_STD,
		CURRENCY_SVC,
		CURRENCY_SZL,
		CURRENCY_THB,
		CURRENCY_TJS,
		CURRENCY_TOP,
		CURRENCY_TRY,
		CURRENCY_TTD,
		CURRENCY_TWD,
		CURRENCY_TZS,
		CURRENCY_UAH,
		CURRENCY_UGX,
		CURRENCY_USD,
		CURRENCY_UYU,
		CURRENCY_UZS,
		CURRENCY_VEF,
		CURRENCY_VND,
		CURRENCY_VUV,
		CURRENCY_WST,
		CURRENCY_XAF,
		CURRENCY_XCD,
		CURRENCY_XOF,
		CURRENCY_XPF,
		CURRENCY_YER,
		CURRENCY_ZAR,
		CURRENCY_ZMW,
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
