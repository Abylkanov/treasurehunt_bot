package handlers

func getPhoto(year string, series string, value int) interface{} {
	var imageMap map[string]string
	var keys []string

	// Инициализируем карту и ключи в зависимости от года и серии
	if year == "2023" && series == "th" {
		imageMap = imageMapTH2023
		keys = keysTH2023
	} else if year == "2023" && series == "supers" {
		imageMap = imageMapSTH2023
		keys = keysSTH2023
	} else if year == "2024" && series == "th" {
		imageMap = imageMapTH2024
		keys = keysTH2024
	} else if year == "2024" && series == "supers" {
		imageMap = imageMapSTH2024
		keys = keysSTH2024
	} else {
		return nil
	}

	if value == 0 {
		return imageMap // Возвращаем всю мапу
	} else if value > 0 && value <= len(keys) {
		key := keys[value-1]                         // Индексируем с нуля
		return map[string]string{key: imageMap[key]} // Возвращаем только выбранный элемент
	}

	return nil
}

var keysTH2023 = []string{
	"95 Jeep Cherokee",
	"2020 Ram 1500 Rebel",
	"BMW R nineT Racer",
	"Ducati 1199 Panigale",
	"Madfast",
	"Mad Propz",
	"Mod Rod",
	"Raijin Express",
	"Rise N Climb",
	"Surf Crate",
	"Time Shifter",
	"Tooned Volkswagen Golf Mk1",
	"Toyota Land Cruiser",
	"Volkswagen Baja Bug",
	"Donut Drifter",
}

var keysSTH2023 = []string{
	"'65 Mercury Comet Cyclone",
	"'68 COPO Camaro",
	"'68 Corvette–Gas Monkey Garage",
	"'69 Shelby GT-500",
	"'82 Toyota Supra",
	"1968 Mazda Cosmo Sport",
	"Classic TV Series Batmobile",
	"Datsun 510 Wagon",
	"Glory Chaser",
	"Lotus Evija",
	"Mercedes-Benz 300 SL",
	"Mighty K",
	"Porsche 935",
	"Renault Sport R.S. 01",
	"Volvo 240 Drift Wagon",
}

var keysTH2024 = []string{
	"'47 Chevy Fleetline",
	"'59 Chevy Impala",
	"Bone Shaker",
	"Draggin' Wagon",
	"Ain't Fare",
	"Batman Forever Batmobile",
	"Car-de-Asada",
	"Custom '53 Chevy",
	"DMC DeLorean",
	"Ford Mustang Mach-E 1400",
	"Honda Super Cub Custom",
	"Hot Wheels Ford Transit Connect",
	"Porsche 928S Safari",
	"Purple Passion",
	"Tooligan",
}

var keysSTH2024 = []string{
	"'18 Camaro SS",
	"'60s Fiat 500D Modificado",
	"'71 El Camino",
	"'77 Pontiac Firebird TA",
	"'83 Chevy Silverado",
	"'89 Mercedes-Benz 560 SEC AMG",
	"'96 Nissan 180SX Type X",
	"24 Lamborghini Huracan LP 620-2 Super Trofeo",
	"Volvo P1800 Gasser",
	"Audi 90 quattro",
	"BMW 507",
	"Celero GT",
	"Ford Escort RS2000",
	"Mazda 787B",
	"Mitsubishi Pajero Evolution",
}

// Аналогично для других мап

var imageMapTH2023 = map[string]string{
	"95 Jeep Cherokee":           "./images/TH-2023/95-Jeep-Cherokee-1.jpg",
	"2020 Ram 1500 Rebel":        "./images/TH-2023/2020-Ram-1500-Rebel-1.jpg",
	"BMW R nineT Racer":          "./images/TH-2023/BMW-R-nineT-Racer-TH-1.jpg",
	"Ducati 1199 Panigale":       "./images/TH-2023/Ducati-1199-Panigale-1.jpg",
	"Madfast":                    "./images/TH-2023/Madfast-1.jpg",
	"Mad Propz":                  "./images/TH-2023/Mad-Propz-1.jpg",
	"Mod Rod":                    "./images/TH-2023/Mod-Rod-1.jpg",
	"Raijin Express":             "./images/TH-2023/Raijin-Express-23-TH-1.jpg",
	"Rise N Climb":               "./images/TH-2023/Rise-N-Climb-C-0.jpg",
	"Surf Crate":                 "./images/TH-2023/Surf-Crate-1.jpg",
	"Time Shifter":               "./images/TH-2023/Time-Shifter-0.jpg",
	"Tooned Volkswagen Golf Mk1": "./images/TH-2023/Tooned-Volkswagen-Golf-Mk1-1.jpg",
	"Toyota Land Cruiser":        "./images/TH-2023/Toyota-Land-Cruiser-1-1.jpg",
	"Volkswagen Baja Bug":        "./images/TH-2023/Volkswagen-Baja-Bug-1.jpg",
	"Donut Drifter":              "./images/TH-2023/Donut-Drifter-TH-0.jpg",
}

var imageMapSTH2023 = map[string]string{
	"'65 Mercury Comet Cyclone":      "./images/STH-2023/'65-Mercury-Comet-Cyclone-Int-Card.jpg",
	"'68 COPO Camaro":                "./images/STH-2023/'68-COPO-Camaro-US-Card.jpg",
	"'68 Corvette–Gas Monkey Garage": "./images/STH-2023/'68-Corvette-–-Gas-Monkey-Garage-US-1.jpg",
	"'69 Shelby GT-500":              "./images/STH-2023/'69-Shelby-GT-500-Int.jpg",
	"'82 Toyota Supra":               "./images/STH-2023/'82-Toyota-Supra-US.jpg",
	"1968 Mazda Cosmo Sport":         "./images/STH-2023/1968-Mazda-Cosmo-Sport-US-Card.jpg",
	"Classic TV Series Batmobile":    "./images/STH-2023/Classic-TV-Series-Batmobile-Tooned-STH-Card.jpg",
	"Datsun 510 Wagon":               "./images/STH-2023/Datsun-510-Wagon.jpg",
	"Glory Chaser":                   "./images/STH-2023/Glory-Chaser-US-Card-1.jpg",
	"Lotus Evija":                    "./images/STH-2023/Lotus-Evija-US-Card.jpg",
	"Mercedes-Benz 300 SL":           "./images/STH-2023/Mercedes-Benz-300-SL-STH-US.jpg",
	"Mighty K":                       "./images/STH-2023/Mighty-K-US-Card.jpg",
	"Porsche 935":                    "./images/STH-2023/Porsche-935-Int-Card.jpg",
	"Renault Sport R.S. 01":          "./images/STH-2023/Renault-Sport-RS01-US-Card.jpg",
	"Volvo 240 Drift Wagon":          "./images/STH-2023/Volvo-240-Drift-Wagon-US-Card.jpg",
}

var imageMapTH2024 = map[string]string{
	"'47 Chevy Fleetline":             "./images/TH-2024/'47-Chevy-Fleetline-1.jpg",
	"'59 Chevy Impala":                "./images/TH-2024/2024-59-Chevy-Impala-1.jpg",
	"Bone Shaker":                     "./images/TH-2024/2024-Bone-Shaker-C-1.jpg",
	"Draggin' Wagon":                  "./images/TH-2024/2024-Draggin-Wagon-1.jpg",
	"Ain't Fare":                      "./images/TH-2024/Aint-Fare-1.jpg",
	"Batman Forever Batmobile":        "./images/TH-2024/Batman-Forever-Batmobile-1.jpg",
	"Car-de-Asada":                    "./images/TH-2024/Car-de-Asada-1.jpg",
	"Custom '53 Chevy":                "./images/TH-2024/Custom-53-Chevy-1.jpg",
	"DMC DeLorean":                    "./images/TH-2024/DMC-DeLorean-1.jpg",
	"Ford Mustang Mach-E 1400":        "./images/TH-2024/Ford-Mustang-Mach-E-1400-1.jpg",
	"Honda Super Cub Custom":          "./images/TH-2024/Honda-Super-Cub-Custom-1-1.jpg",
	"Hot Wheels Ford Transit Connect": "./images/TH-2024/Hot-Wheels-Ford-Transit-Connect-1.jpg",
	"Porsche 928S Safari":             "./images/TH-2024/Porsche-928S-Safari-1a.jpg",
	"Purple Passion":                  "./images/TH-2024/Purple-Passion-TH-1.jpg",
	"Tooligan":                        "./images/TH-2024/Tooligan-1.jpg",
}

var imageMapSTH2024 = map[string]string{
	"'18 Camaro SS":                                "./images/STH-2024/'18-Camaro-SS-US-Card.jpg",
	"'60s Fiat 500D Modificado":                    "./images/STH-2024/'60s-Fiat-500D-Modificado-US-Card.jpg",
	"'71 El Camino":                                "./images/STH-2024/'71-El-Camino-US-Card.jpg",
	"'77 Pontiac Firebird TA":                      "./images/STH-2024/'77-Pontiac-Firebird-TA-US-Card.jpg",
	"'83 Chevy Silverado":                          "./images/STH-2024/'83-Chevy-Silverado-US-Card.jpg",
	"'89 Mercedes-Benz 560 SEC AMG":                "./images/STH-2024/'89-Mercedes-Benz-560-SEC-AMG-US-Card.jpg",
	"'96 Nissan 180SX Type X":                      "./images/STH-2024/'96-Nissan-180SX-Type-X-US-Card.jpg",
	"24 Lamborghini Huracan LP 620-2 Super Trofeo": "./images/STH-2024/24-Lamborghini-Huracan-LP-620-2-Super-Trofeo-INT-Card.jpg",
	"Volvo P1800 Gasser":                           "./images/STH-2024/2024-Volvo-P1800-Gasser-US-Card.jpg",
	"Audi 90 quattro":                              "./images/STH-2024/Audi-90-quattro-US-Card.jpg",
	"BMW 507":                                      "./images/STH-2024/BMW-507-Int-Card.jpg",
	"Celero GT":                                    "./images/STH-2024/Celero-GT-US-Card.jpg",
	"Ford Escort RS2000":                           "./images/STH-2024/Ford-Escort-RS2000-Int-Card.jpg",
	"Mazda 787B":                                   "./images/STH-2024/Mazda-787B-US-Card.jpg",
	"Mitsubishi Pajero Evolution":                  "./images/STH-2024/Mitsubishi-Pajero-Evolution-US-Card.jpg",
}
