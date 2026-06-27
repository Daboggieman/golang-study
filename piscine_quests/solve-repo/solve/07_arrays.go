package solve

import (
	"fmt"
	// "strconv"
)

func PracArr(n int) {
	randten := []int{2, 4, 3, 1, 7, 8, 5, 6, 9, 0, 11}

	for i := 0; i <= len(randten)-1; i++ {
		fmt.Print(randten[i], ",", " ")
	}

	for i := 11; i <= n; {
		randten = append (randten, i)
		if i != n{
		fmt.Print(randten[i], ",", " ")	
		} else {
			fmt.Print(randten[i])	
		}
		i++ 
	}

	// fmt.Println(randten[:4])

	//the block below is still under construction, fix later
	/*teststr := "this is the test string that we want to print"
	nowconv := strconv.Atoi(teststr)
	teststrslc := []int{nowconv}

	for i := 0; i <= len(teststrslc)-1; i++ {
		backconv := strconv.Itoa(teststrslc[i])
		fmt.Print(backconv, ",", " ")
	}

	fmt.Println(teststrslc[0]) */
}