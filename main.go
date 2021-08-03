package main

import "math/big"

func converges(n *big.Int) bool {

}

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

			converges(n)



		}
	}

}
