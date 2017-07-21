package main

import (
	"strconv"
)

func main() {

	p := palindrom()
	println("Answer is ", p)

}

/* A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
Find the largest palindrome made from the product of two 3-digit numbers. */
func palindrom() int {

	var pali int
	var good int

	for i := 1; i <= 999; i++ {

		for j := 1; j <= 999; j++ {
			x := i * j

			//Let's find out how many digits x have!
			s := strconv.Itoa(x)
			length := len(s)

			//Is number of digits even? Otherwise we don't care!
			if length%2 == 0 {
				count := length / 2

				//Loop through numbers and for each match, count + 1 to good
				for z := 0; z < count; z++ {
					if s[z:z+1] == s[length-z-1:length-z] {
						good += 1
					}
				}

				//We need to check just half of digits, good equal count means that it is palindrom!
				if good == count && x > pali {
					pali = x
					println(pali, " - ", i, " - ", j)
				}

				//Reset count - honestly don't know how to reset it in more noble way
				good = 0

			}

		}
	}

	return pali

}
