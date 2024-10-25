package handlers

func getSelectedMessage(year string, series string) string {
	switch year {
	case "2023":
		switch series {
		case "th":
			return list2023th
		case "supers":
			return list2023supers
		default:
			return "unknown series"
		}
	case "2024":
		switch series {
		case "th":
			return list2024th
		case "supers":
			return list2024supers
		default:
			return "unknown series"
		}
	default:
		return "unknown year"
	}
}

var list2023th = ` 
1. '95 Jeep Cherokee
2. 2020 Ram 1500 Rebel
3. BMW R nineT Racer
4. Ducati 1199 Panigale
5. Madfast
6. Mad Propz
7. Mod Rod
8. Raijin Express
9. Rise N Climb
10. Surf Crate
11. Time Shifter
12. Tooned Volkswagen Golf Mk1
13. Toyota Land Cruiser
14. Volkswagen Baja Bug
15. Donut Drifter
`

var list2023supers = `
1. '65 Mercury Comet Cyclone
2. '68 COPO Camaro
3. '68 Corvette–Gas Monkey Garage
4. '69 Shelby GT-500
5. '82 Toyota Supra
6. 1968 Mazda Cosmo Sport
7. Classic TV Series Batmobile
8. Datsun 510 Wagon
9. Glory Chaser
10. Lotus Evija
11. Mercedes-Benz 300 SL
12. Mighty K
13. Porsche 935
14. Renault Sport R.S. 01
15. Volvo 240 Drift Wagon
`

var list2024th = `
1. '47Chevy Fleetline
2. '59 Chevy Impala
3. Bone Shaker
4. Draggin' Wagon
5. Ain't Fare
6. Batman Forever Batmobile
7. Car-de-Asada
8. Custom '53 Chevy
9. DMC DeLorean
10. Ford Mustang Mach-E 1400
11. Honda Super Cub Custom
12. Hot Wheels Ford Transit Connect
13. Porsche 928S Safari
14. Purple Passion
15. Tooligan
`

var list2024supers = `
1. '18 Camaro SS
2. '60s Fiat 500D Modificado
3. '71 El Camino
4. '77 Pontiac Firebird TA
5. '83 Chevy Silverado
6. '89 Mercedes-Benz 560 SEC AMG
7. '96 Nissan 180SX Type X
8. 24 Lamborghini Huracan LP 620-2 Super Trofeo
9. Volvo P1800 Gasser
10. Audi 90 quattro
11. BMW 507
12. Celero GT
13. Ford Escort RS2000
14. Mazda 787B
15. Mitsubishi Pajero Evolution
`
