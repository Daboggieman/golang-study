package piscine

// AdvancedSortWordArr sorts the slice a in-place using the comparator f.
// The comparator f must return 1 when a > b, 0 when a == b, and -1 when a < b.
func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if f(a[j], a[j+1]) == 1 {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
