package main

import (
	"fmt"
	"math/big"
)

func main() {
	println("dont do it")
	c := 0
	n := big.NewInt(1)
	done := NewDone(100000)
	for {
		if done.KnownToConverge(n) { // known to converge
			//println(n.Text(10), "already known to converge")
			n.Add(n, done.ONE)
		} else if n.Bit(0) == 0 { // odd number, skip
			//println(n.Text(10), "known to converge as it is even and all previous converge")
			done.AddKnownToConverge(n)
			n.Add(n, done.ONE)
		} else { // even, needs to check
			if done.VerifyConverges(n, false) {
				//println(n.Text(10), "verified to converge")
				done.VerifyConverges(n, true)
				done.AddKnownToConverge(n)
				n.Add(n, done.ONE)
			} else {
				panic(fmt.Sprintln(n.Text(10), "could not be verified"))
			}
		}

		c++
		if c == 1000 {
			println(n.Text(10), " verified; ", done.String()[0:100])
			c = 0
		}
	}

}
