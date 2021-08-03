package main

import "math/big"

type Done struct {
	ONE          *big.Int
	THREE        *big.Int
	tmp          *big.Int
	tmpAdd       *big.Int
	provenUpTo   *big.Int
	sparseProven []*big.Int
}

func NewDone(size int) *Done {
	done := Done{
		ONE:          big.NewInt(1),
		THREE:        big.NewInt(3),
		tmp:          big.NewInt(0),
		tmpAdd:       big.NewInt(0),
		provenUpTo:   big.NewInt(1),
		sparseProven: make([]*big.Int, size),
	}
	for i := 0; i < size; i++ {
		done.sparseProven[i] = big.NewInt(1)
	}
	return &done
}

func (d *Done) KnownToConverge(i *big.Int) bool {
	return d.provenUpTo.Cmp(i) >= 0
}

func (d *Done) AddKnownToConverge(i *big.Int) {
	d.tmpAdd.Set(i)
	d.tmpAdd.Sub(d.tmp, d.provenUpTo)
	if d.tmpAdd.IsInt64() && d.tmpAdd.Int64() == 1 {
		d.provenUpTo.Add(d.provenUpTo, d.ONE)
	} else {
		d.insertIntoSparseProven(i)
	}
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

}

func (d *Done) VerifyConverges(i *big.Int, adding bool) bool {
	if adding {
		defer d.AddKnownToConverge(i)
	}
	var loops int64 = 0
	d.tmp.Set(i)
	for {
		//print(d.tmp.Text(10), " ")
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
		if adding {
			d.AddKnownToConverge(d.tmp)
		}
		loops++
		if loops > 2^32 { // FIXME
			return false
		}
	}
}
