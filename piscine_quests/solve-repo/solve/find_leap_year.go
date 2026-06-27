package solve

import "fmt"

func FindLeapYear() {
	var InsYear int
	var CurrentYear int

	fmt.Println("Enter current year:")
	fmt.Println("Enter year to check if it is a leap year:")
	_, err := fmt.Scan(&InsYear)
	_, err = fmt.Scan(&CurrentYear)
	if InsYear < CurrentYear {
		fmt.Println("you are checking for a year in the past, that's alright")
	} else if InsYear == CurrentYear {
		fmt.Println("yup!! you are checking for the current year")
	}
	if err != nil {
		fmt.Println("Invalid year, please insert a correct year:", err)
	}
}
