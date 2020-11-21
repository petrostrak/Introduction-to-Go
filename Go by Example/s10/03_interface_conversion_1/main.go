/*
Conversion using strconv package. The most common numeric
conversions are Atoi(arrayofstring to int) and Itoa(int to
arrayofstring).

func Atoi(s string) (int, error)
func Itoa(i int) string

Converting strings to values
func ParseBool(str string) (bool, error)
func ParseFloat(s string, bitSize int) (float64, error)
func ParseInt(s string, base int, bitSize int) (i int64, err error)
func ParseUint(s string, base int, bitSize int) (uint64, error)

Converting values to strings
func FormatBool(b bool) string
func FormatFloat(f float64, fmt byte, prec, bitSize int) string
func FormatInt(i int64, base int) string
func FormatUint(i uint64, base int) string
*/

package main

import "fmt"

func main() {
	var i = 10
	var f = 10.25
	fmt.Println(f + float64(i))
	fmt.Println(int(f) + i)

	var r rune = 'a'
	var i2 int32 = 'b'

	fmt.Println(r, i2)
	fmt.Println(string(r), string(i2))

	b := []byte{'t', 'e', 's', 't'}
	fmt.Println(b, string(b))

	s := "test"
	fmt.Println([]byte(s), s)
}
