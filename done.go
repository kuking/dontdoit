package main

import "math/big"

type Done struct {
	ONE          *big.Int
	THREE        *big.Int
	tmp          *big.Int
	provenUpTo   *big.Int
	sparseProven []*big.Int
}

func NewDone(size int) *Done {
	return &Done{
		ONE:          big.NewInt(1),
		THREE:        big.NewInt(3),
		tmp:          big.NewInt(0),
		provenUpTo:   big.NewInt(1),
		sparseProven: make([]*big.Int, size),
	}
}

func (d *Done) KnownToConverge(i *big.Int) bool {
	return d.provenUpTo.Cmp(i) >= 0
}

func (d *Done) AddKnownToConverge(i *big.Int) {
	d.tmp.Sub(i, d.provenUpTo)
	if d.tmp.IsUint64() && d.tmp.Uint64() == 1 {
		d.provenUpTo.Add(d.provenUpTo, d.ONE)
	} else {
		panic("implement sparseProven")
	}
}

func (d *Done) VerifyConverges(i *big.Int, adding bool) bool {
	if adding {
		d.AddKnownToConverge(i)
	}
	var loops int64 = 0
	d.tmp.Set(i)
	defer println()
	for {
		print(d.tmp.Text(10), " ")
		if d.KnownToConverge(d.tmp) {
			return true
		}

		if d.tmp.Bit(0) == 0 {
			d.tmp.Rsh(d.tmp, 1) // divide by 2
		} else {
			d.tmp = d.tmp.Mul(d.tmp, d.THREE)
			d.tmp = d.tmp.Add(d.tmp, d.ONE)
		}
		if adding {
			d.AddKnownToConverge(d.tmp)
		}
		loops++
		if loops > 2^32 { // FIXME
			return false
		}
	}
}
