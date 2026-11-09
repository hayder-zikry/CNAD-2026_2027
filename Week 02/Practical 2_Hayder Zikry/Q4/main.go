package main

import "fmt"

type MonthData struct {
	Month    string
	Rainfall int
	Days     int
}

func main() {
	yr2017 := []MonthData{
		{"January", 220, 18}, {"February", 10, 2}, {"March", 20, 3},
		{"April", 10, 3}, {"May", 22, 6}, {"June", 24, 4},
		{"July", 168, 13}, {"August", 198, 15}, {"September", 188, 15},
		{"October", 213, 18}, {"November", 260, 19}, {"December", 278, 19},
	}

	yr2018 := []MonthData{
		{"January", 245, 18}, {"February", 41, 6}, {"March", 201, 13},
		{"April", 197, 16}, {"May", 175, 15}, {"June", 172, 14},
		{"July", 166, 14}, {"August", 198, 18}, {"September", 185, 15},
		{"October", 216, 16}, {"November", 256, 19}, {"December", 278, 20},
	}

	yr2019 := []MonthData{
		{"January", 250, 17}, {"February", 40, 7}, {"March", 200, 12},
		{"April", 195, 15}, {"May", 175, 15}, {"June", 170, 13},
		{"July", 165, 13}, {"August", 190, 15}, {"September", 180, 14},
		{"October", 210, 16}, {"November", 250, 19}, {"December", 275, 19},
	}

	rainfall_list := map[string][]MonthData{
		"2017": yr2017,
		"2018": yr2018,
		"2019": yr2019,
	}

	// 4(a)
	rainfall_year_list := map[string]int{
		"2017": 0,
		"2018": 0,
		"2019": 0,
	}

	// 4(b)
	rainfall_lowest_list := map[string]string{
		"2017": "",
		"2018": "",
		"2019": "",
	}

	for year, months := range rainfall_list {
		totalDays := 0
		lowestRainfall := -1
		lowestMonth := ""

		for _, data := range months {
			totalDays += data.Days

			if lowestRainfall == -1 || data.Rainfall <= lowestRainfall {
				lowestRainfall = data.Rainfall
				lowestMonth = data.Month[:3]
			}
		}

		rainfall_year_list[year] = totalDays
		rainfall_lowest_list[year] = lowestMonth
	}

	// 4(a)
	fmt.Printf("Total number of days of rainfall for each year: map[2017:%d 2018:%d 2019:%d]\n",
		rainfall_year_list["2017"], rainfall_year_list["2018"], rainfall_year_list["2019"])
	// 4(b)
	fmt.Printf("Month in which there was lowest rainfall for each year: map[2017:%s 2018:%s 2019:%s]\n",
		rainfall_lowest_list["2017"], rainfall_lowest_list["2018"], rainfall_lowest_list["2019"])

	fmt.Println("\nProgram exited.")
}
