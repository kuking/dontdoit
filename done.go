package main

import "math/big"

type Done struct {
	ONE          *big.Int
	THREE        *big.Int
	tmp          *big.Int
	tmpAdd       *big.Int
	tmpPump      *big.Int
	provenUpTo   *big.Int
	sparseProven []*big.Int
}

func NewDone(size int) *Done {
	done := Done{
		ONE:          big.NewInt(1),
		THREE:        big.NewInt(3),
		tmp:          big.NewInt(0),
		tmpAdd:       big.NewInt(0),
		tmpPump:      big.NewInt(0),
		provenUpTo:   big.NewInt(1),
		sparseProven: make([]*big.Int, size),
	}
	for i := 0; i < size; i++ {
		done.sparseProven[i] = big.NewInt(1)
	}
	return &done
}

func (d *Done) KnownToConverge(i *big.Int) bool {
	if d.provenUpTo.Cmp(i) >= 0 {
		return true
	}
	for idx := 0; idx < len(d.sparseProven); idx++ {
		cmp := d.sparseProven[idx].Cmp(i)
		if cmp == 0 {
			return true
		} else if cmp > 0 {
			return false
		}
	}
	return false
}

func (d *Done) AddKnownToConverge(i *big.Int) {
	if d.KnownToConverge(i) {
		return
	}
	d.insertIntoSparseProven(i)
	d.pumpProvenUpToIfPossible()
}

func (d *Done) insertIntoSparseProven(i *big.Int) {
	for idx := 0; idx < len(d.sparseProven); idx++ {
		cmp := d.sparseProven[idx].Cmp(i)
		if cmp == 0 {
			return // already there
		} else if cmp > 0 || d.sparseProven[idx].Cmp(d.ONE) == 0 {
			for idx2 := len(d.sparseProven) - 2; idx2 >= idx; idx2-- {
				d.sparseProven[idx2+1].Set(d.sparseProven[idx2])
			}
			d.sparseProven[idx].Set(i)
			return
		}
	}
}

func (d *Done) String() (out string) {
	out = "ProvenUpTo: " + d.provenUpTo.Text(10) + " SparseProven: "
	for i := 0; i < len(d.sparseProven); i++ {
		out = out + d.sparseProven[i].Text(10) + " "
		if d.sparseProven[i].Cmp(d.ONE) == 0 {
			break
		}
	}
	return
}

func (d *Done) pumpProvenUpToIfPossible() {
	for {
		d.tmpPump.Set(d.sparseProven[0])
		d.tmpPump.Sub(d.tmpPump, d.provenUpTo)
		if d.tmpPump.IsInt64() && d.tmpPump.Int64() == 1 {
			d.provenUpTo.Set(d.sparseProven[0])
			for idx := 0; idx < len(d.sparseProven)-2; idx++ {
				d.sparseProven[idx].Set(d.sparseProven[idx+1])
			}
			d.sparseProven[len(d.sparseProven)-1].Set(d.ONE)
		} else {
			return
		}
	}
}

func (d *Done) VerifyConverges(i *big.Int, adding bool) bool {
	if adding {
		defer d.AddKnownToConverge(i)
	}
	var loops int64 = 0
	d.tmp.Set(i)
	for {
		if d.KnownToConverge(d.tmp) {
			return true
		}
		if adding {
			d.AddKnownToConverge(d.tmp)
		}
		if d.tmp.Bit(0) == 0 {
			d.tmp.Rsh(d.tmp, 1) // divide by 2
		} else {
			d.tmp = d.tmp.Mul(d.tmp, d.THREE)
			d.tmp = d.tmp.Add(d.tmp, d.ONE)
		}
		loops++
		if loops > 1<<32 { // FIXME
			println("Gave up verifying", i.Text(10), " after 2^32 iterations.")
			return false
		}
	}
}
