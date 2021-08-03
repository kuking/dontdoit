package main

import "math/big"

func main() {
	println("dont do it")

	n := big.NewInt(1)

	done := NewDone(1000)
	for {
		if done.KnownToConverge(n) { // known to converge
			println(n.Text(10), "known to converge")
			n.Add(n, done.ONE)
		} else if n.Bit(0) == 0 { // odd number, skip
			println(n.Text(10), "known to converge as it is odd")
			done.AddKnownToConverge(n)
			n.Add(n, done.ONE)
		} else { // even, needs to check
			if done.VerifyConverges(n, false) {
				println(n.Text(10), "verified to converge")
				done.VerifyConverges(n, true)
				done.AddKnownToConverge(n)
				n.Add(n, done.ONE)
			}

		}
	}

}
