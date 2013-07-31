// Package randomdata implements a bunch of simple ways to generate (pseudo) random data
package randomdata

import (
	"math/rand"
	"strings"
	"time"
)

const (
	Male         int = 0
	Female       int = 1
	RandomGender int = 2
)

const (
	FullCountry      = 0
	TwoCharCountry   = 1
	ThreeCharCountry = 2
)

var firstNamesMale = []string{
	"Jacob", "Mason", "Ethan", "Noah", "William", "Liam", "Jayden",
	"Michael", "Alexander", "Aiden", "Daniel", "Matthew", "Elijah",
	"James", "Anthony", "Benjamin", "Joshua", "Andrew", "David", "Joseph"}

var firstNamesFemale = []string{
	"Sophia", "Emma", "Isabella", "Olivia", "Ava", "Emily", "Abigail",
	"Mia", "Madison", "Elizabeth", "Chloe", "Ella", "Avery", "Addison",
	"Aubrey", "Lily", "Natalie", "Sofia", "Charlotte", "Zoey"}

var lastNames = []string{
	"Smith", "Johnson", "Williams", "Jones", "Brown", "Davis", "Miller",
	"Wilson", "Moore", "Taylor", "Anderson", "Thomas", "Jackson", "White",
	"Harris", "Martin", "Thompson", "Garcia", "Martinez", "Robinson"}

var domains = []string{"test.com"}

var paragraphs = []string{}

// Fetched from the world bank at http://siteresources.worldbank.org/DATASTATISTICS/Resources/CLASS.XLS
var countries = []string{"Afghanistan", "Albania", "Algeria", "American Samoa",
	"Andorra", "Angola", "Antigua and Barbuda", "Argentina", "Armenia", "Aruba", "Australia",
	"Austria", "Azerbaijan", "Bahamas, The", "Bahrain", "Bangladesh", "Barbados", "Belarus",
	"Belgium", "Belize", "Benin", "Bermuda", "Bhutan", "Bolivia", "Bosnia and Herzegovina",
	"Botswana", "Brazil", "Brunei Darussalam", "Bulgaria", "Burkina Faso", "Burundi", "Cambodia",
	"Cameroon", "Canada", "Cape Verde", "Cayman Islands", "Central African Republic", "Chad",
	"Channel Islands", "Chile", "China", "Colombia", "Comoros", "Congo, Dem. Rep.", "Congo, Rep.",
	"Costa Rica", "Côte d'Ivoire", "Croatia", "Cuba", "Cyprus", "Czech Republic", "Denmark", "Djibouti",
	"Dominica", "Dominican Republic", "Ecuador", "Egypt, Arab Rep.", "El Salvador", "Equatorial Guinea",
	"Eritrea", "Estonia", "Ethiopia", "Faeroe Islands", "Fiji", "Finland", "France", "French Polynesia",
	"Gabon", "Gambia, The", "Georgia", "Germany", "Ghana", "Greece", "Greenland", "Grenada", "Guam",
	"Guatemala", "Guinea", "Guinea-Bissau", "Guyana", "Haiti", "Honduras", "Hong Kong, China", "Hungary",
	"Iceland", "India", "Indonesia", "Iran, Islamic Rep.", "Iraq", "Ireland", "Isle of Man", "Israel",
	"Italy", "Jamaica", "Japan", "Jordan", "Kazakhstan", "Kenya", "Kiribati", "Korea, Dem. Rep.", "Korea, Rep.",
	"Kuwait", "Kyrgyz Republic", "Lao PDR", "Latvia", "Lebanon", "Lesotho", "Liberia", "Libya", "Liechtenstein",
	"Lithuania", "Luxembourg", "Macao, China", "Macedonia, FYR", "Madagascar", "Malawi", "Malaysia", "Maldives",
	"Mali", "Malta", "Marshall Islands", "Mauritania", "Mauritius", "Mayotte", "Mexico", "Micronesia, Fed. Sts.",
	"Moldova", "Monaco", "Mongolia", "Montenegro", "Morocco", "Mozambique", "Myanmar", "Namibia", "Nepal",
	"Netherlands", "Netherlands Antilles", "New Caledonia", "New Zealand", "Nicaragua", "Niger", "Nigeria",
	"Northern Mariana Islands", "Norway", "Oman", "Pakistan", "Palau", "Panama", "Papua New Guinea", "Paraguay",
	"Peru", "Philippines", "Poland", "Portugal", "Puerto Rico", "Qatar", "Romania", "Russian Federation", "Rwanda",
	"Samoa", "San Marino", "São Tomé and Principe", "Saudi Arabia", "Senegal", "Serbia", "Seychelles", "Sierra Leone",
	"Singapore", "Slovak Republic", "Slovenia", "Solomon Islands", "Somalia", "South Africa", "Spain", "Sri Lanka",
	"St. Kitts and Nevis", "St. Lucia", "St. Vincent and the Grenadines", "Sudan", "Suriname", "Swaziland", "Sweden",
	"Switzerland", "Syrian Arab Republic", "Tajikistan", "Tanzania", "Thailand", "Timor-Leste", "Togo", "Tonga",
	"Trinidad and Tobago", "Tunisia", "Turkey", "Turkmenistan", "Uganda", "Ukraine", "United Arab Emirates",
	"United Kingdom", "United States", "Uruguay", "Uzbekistan", "Vanuatu", "Venezuela, RB", "Vietnam",
	"Virgin Islands (U.S.)", "West Bank and Gaza", "Yemen, Rep.", "Zambia", "Zimbabwe"}

var countriesThreeChars = []string{
	"AFG", "ALB", "DZA", "ASM", "ADO", "AGO", "ATG", "ARG", "ARM", "ABW", "AUS", "AUT", "AZE", "BHS", "BHR", "BGD",
	"BRB", "BLR", "BEL", "BLZ", "BEN", "BMU", "BTN", "BOL", "BIH", "BWA", "BRA", "BRN", "BGR", "BFA", "BDI", "KHM",
	"CMR", "CAN", "CPV", "CYM", "CAF", "TCD", "CHI", "CHL", "CHN", "COL", "COM", "ZAR", "COG", "CRI", "CIV", "HRV",
	"CUB", "CUW", "CYP", "CZE", "DNK", "DJI", "DMA", "DOM", "ECU", "EGY", "SLV", "GNQ", "ERI", "EST", "ETH", "FRO",
	"FJI", "FIN", "FRA", "PYF", "GAB", "GMB", "GEO", "DEU", "GHA", "GRC", "GRL", "GRD", "GUM", "GTM", "GIN", "GNB",
	"GUY", "HTI", "HND", "HKG", "HUN", "ISL", "IND", "IDN", "IRN", "IRQ", "IRL", "IMY", "ISR", "ITA", "JAM", "JPN",
	"JOR", "KAZ", "KEN", "KIR", "PRK", "KOR", "KSV", "KWT", "KGZ", "LAO", "LVA", "LBN", "LSO", "LBR", "LBY", "LIE",
	"LTU", "LUX", "MAC", "MKD", "MDG", "MWI", "MYS", "MDV", "MLI", "MLT", "MHL", "MRT", "MUS", "MEX", "FSM", "MDA",
	"MCO", "MNG", "MNE", "MAR", "MOZ", "MMR", "NAM", "NPL", "NLD", "NCL", "NZL", "NIC", "NER", "NGA", "MNP", "NOR",
	"OMN", "PAK", "PLW", "PAN", "PNG", "PRY", "PER", "PHL", "POL", "PRT", "PRI", "QAT", "ROM", "RUS", "RWA", "WSM",
	"SMR", "STP", "SAU", "SEN", "SRB", "SYC", "SLE", "SGP", "SXM", "SVK", "SVN", "SLB", "SOM", "ZAF", "SSD", "ESP",
	"LKA", "KNA", "LCA", "MAF", "VCT", "SDN", "SUR", "SWZ", "SWE", "CHE", "SYR", "TJK", "TZA", "THA", "TMP", "TGO",
	"TON", "TTO", "TUN", "TUR", "TKM", "TCA", "TUV", "UGA", "UKR", "ARE", "GBR", "USA", "URY", "UZB", "VUT", "VEN",
	"VNM", "VIR", "WBG", "YEM", "ZMB", "ZWE"}

var countriesTwoChars = []string{
	"AF", "AX", "AL", "DZ", "AS", "AD", "AO", "AI", "AQ", "AG", "AR", "AM", "AW", "AU", "AT", "AZ", "BS", "BH", "BD",
	"BB", "BY", "BE", "BZ", "BJ", "BM", "BT", "BO", "BQ", "BA", "BW", "BV", "BR", "IO", "BN", "BG", "BF", "BI", "KH",
	"CM", "CA", "CV", "KY", "CF", "TD", "CL", "CN", "CX", "CC", "CO", "KM", "CG", "CD", "CK", "CR", "CI", "HR", "CU",
	"CW", "CY", "CZ", "DK", "DJ", "DM", "DO", "EC", "EG", "SV", "GQ", "ER", "EE", "ET", "FK", "FO", "FJ", "FI", "FR",
	"GF", "PF", "TF", "GA", "GM", "GE", "DE", "GH", "GI", "GR", "GL", "GD", "GP", "GU", "GT", "GG", "GN", "GW", "GY",
	"HT", "HM", "VA", "HN", "HK", "HU", "IS", "IN", "ID", "IR", "IQ", "IE", "IM", "IL", "IT", "JM", "JP", "JE", "JO",
	"KZ", "KE", "KI", "KP", "KR", "KW", "KG", "LA", "LV", "LB", "LS", "LR", "LY", "LI", "LT", "LU", "MO", "MK", "MG",
	"MW", "MY", "MV", "ML", "MT", "MH", "MQ", "MR", "MU", "YT", "MX", "FM", "MD", "MC", "MN", "ME", "MS", "MA", "MZ",
	"MM", "NA", "NR", "NP", "NL", "NC", "NZ", "NI", "NE", "NG", "NU", "NF", "MP", "NO", "OM", "PK", "PW", "PS", "PA",
	"PG", "PY", "PE", "PH", "PN", "PL", "PT", "PR", "QA", "RE", "RO", "RU", "RW", "BL", "SH", "KN", "LC", "MF", "PM",
	"VC", "WS", "SM", "ST", "SA", "SN", "RS", "SC", "SL", "SG", "SX", "SK", "SI", "SB", "SO", "ZA", "GS", "SS", "ES",
	"LK", "SD", "SR", "SJ", "SZ", "SE", "CH", "SY", "TW", "TJ", "TZ", "TH", "TL", "TG", "TK", "TO", "TT", "TN", "TR",
	"TM", "TC", "TV", "UG", "UA", "AE", "GB", "US", "UM", "UY", "UZ", "VU", "VE", "VN", "VG", "VI", "WF", "EH", "YE",
	"ZM", "ZW"}

var cities = []string{
	"Derby Center", "New Deal", "Cienega Springs", "Ransom Canyon", "Burrton", "Hoonah", "Lucien", "San Martin",
	"Buffalo City", "Skidaway Island", "Kingsbridge", "Berkhamsted", "Bury", "Brandwell", "Campden", "Plympton",
	"Baldock", "Northleach", "Newstead"}

var states = []string{"Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware",
	"Florida", "Georgia", "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine",
	"Maryland", "Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri", "Montana", "Nebraska", "Nevada",
	"New Hampshire", "New Jersey", "New Mexico", "New York", "North Carolina", "North Dakota", "Ohio", "Oklahoma", "Oregon",
	"Pennsylvania", "Rhode Island", "South Carolina", "South Dakota", "Tennessee", "Texas", "Utah", "Vermont", "Virginia",
	"Washington", "West Virginia", "Wisconsin", "Wyoming"}

// Returns a random part of a slice
func randomFrom(source []string) string {
	rand.Seed(time.Now().UnixNano())
	return source[rand.Intn(len(source))]
}

// Returns a random first name
func FirstName(gender int) string {
	var name = ""
	switch gender {
	case Male:
		name = randomFrom(firstNamesMale)
		break
	case Female:
		name = randomFrom(firstNamesFemale)
		break
	default:
		rand.Seed(time.Now().UnixNano())
		name = FirstName(rand.Intn(2))
		break
	}
	return name
}

// returns a random last name
func LastName() string {
	return randomFrom(lastNames)
}

// returns a combinaton of FirstName LastName randomized
func FullName(gender int) string {
	return FirstName(gender) + " " + LastName()
}

// returns a random email
func Email() string {
	return strings.ToLower(FirstName(RandomGender)+LastName()) + "@" + randomFrom(domains)
}

// returns a random country
func Country(countryStyle int64) string {
	country := ""
	switch countryStyle {

	default:

	case FullCountry:
		country = randomFrom(countries)
		break
	case TwoCharCountry:
		country = randomFrom(countriesTwoChars)
		break

	case ThreeCharCountry:
		country = randomFrom(countriesThreeChars)
		break
	}
	return country
}

// returns a random city
func City() string {
	return randomFrom(cities)
}

func State() string {
	return randomFrom(states)
}

// returns a random number, if only one integer is supplied it is treated as the max value to return
// if a second argument is supplied it returns a number between (and including) the two numbers
func Number(numberRange ...int) int {
	nr := 0
	rand.Seed(time.Now().UnixNano())
	if len(numberRange) > 1 {
		nr = 1
		nr = rand.Intn(numberRange[1]-numberRange[0]) + numberRange[0]
	} else {
		nr = rand.Intn(numberRange[0])
	}
	return nr
}