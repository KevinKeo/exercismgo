/*
Package leap implements the solution for the leap exercice in excercism.io.
*/
package leap

// testVersion should match the targetTestVersion in the test file.
const testVersion = 2

//IsLeapYear return true if the year is a leap.
func IsLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}
	if year%100 == 0 {
		return false
	}
	if year%4 == 0 {
		return true
	}
	return false
}
