package main

import (
	"math/big"
	"testing"
)

func TestDoneSimplest(t *testing.T) {
	d := NewDone(10)
	if !d.KnownToConverge(d.ONE) {
		t.Fail()
	}
	d.tmp.Set(d.ONE)
	for i := 0; i < 1000; i++ {
		d.tmp = d.tmp.Add(d.tmp, d.ONE)
		if d.KnownToConverge(d.tmp) {
			t.Fail()
		}
		d.AddKnownToConverge(d.tmp)
		if !d.KnownToConverge(d.tmp) {
			t.Fail()
		}
	}
}

func TestDoneStateFor10(t *testing.T) {
	d := NewDone(10)
	// 10, 5, 16, 8, 4, 2, 1
	d.VerifyConverges(big.NewInt(10), true)
	// provenUpTo: 2, sparseProven: 4, 5, 8, 10, 16
	//println(d.String())
	if d.provenUpTo.Int64() != 2 {
		t.Fail()
	}
	if d.sparseProven[0].Int64() != 4 ||
		d.sparseProven[1].Int64() != 5 ||
		d.sparseProven[2].Int64() != 8 ||
		d.sparseProven[3].Int64() != 10 ||
		d.sparseProven[4].Int64() != 16 {
		t.Fail()
	}
}

func TestDoneStateFor3(t *testing.T) {
	d := NewDone(10)
	d.VerifyConverges(big.NewInt(3), true)
	// 3, 10, 5, 16, 8, 4, 2, 1
	// provenUpTo: 5, sparseProven: 8, 10, 16
	//println(d.String())
	if d.provenUpTo.Int64() != 5 {
		t.Fail()
	}
	if d.sparseProven[0].Int64() != 8 ||
		d.sparseProven[1].Int64() != 10 ||
		d.sparseProven[2].Int64() != 16 {
		t.Fail()
	}

	if !d.KnownToConverge(big.NewInt(4)) {
		t.Fail()
	}
	if d.KnownToConverge(big.NewInt(9)) {
		t.Fail()
	}
	if !d.KnownToConverge(big.NewInt(10)) {
		t.Fail()
	}
	if d.KnownToConverge(big.NewInt(200)) {
		t.Fail()
	}
}

func TestMiniPerf(t *testing.T) {
	d := NewDone(1)
	for i := 0; i < 1_000; i++ {
		if d.VerifyConverges(big.NewInt(int64(i)), false) {
			d.VerifyConverges(big.NewInt(int64(i)), true)
		}
	}
}
