package main

import "math/big"

type Done struct {
	ONE          *big.Int
	tmp          *big.Int
	provenUpTo   *big.Int
	sparseProven []*big.Int
}

func NewDone(size int) *Done {
	return &Done{
		ONE:          big.NewInt(1),
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
