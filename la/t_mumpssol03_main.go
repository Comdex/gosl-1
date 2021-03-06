// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
	"github.com/cpmech/gosl/la"
	"github.com/cpmech/gosl/mpi"
)

func main() {

	mpi.Start(false)
	defer mpi.Stop(false)

	myrank := mpi.Rank()
	if myrank == 0 {
		io.Pf("\nTest MUMPS Sol 03\n")
	}

	var t la.TripletC
	switch mpi.Size() {
	case 1:
		t.Init(5, 5, 13, true)
		t.Put(0, 0, 1.0, 0)
		t.Put(0, 0, 1.0, 0)
		t.Put(1, 0, 3.0, 0)
		t.Put(0, 1, 3.0, 0)
		t.Put(2, 1, -1.0, 0)
		t.Put(4, 1, 4.0, 0)
		t.Put(1, 2, 4.0, 0)
		t.Put(2, 2, -3.0, 0)
		t.Put(3, 2, 1.0, 0)
		t.Put(4, 2, 2.0, 0)
		t.Put(2, 3, 2.0, 0)
		t.Put(1, 4, 6.0, 0)
		t.Put(4, 4, 1.0, 0)
	case 2:
		if myrank == 0 {
			t.Init(5, 5, 6, true)
			t.Put(0, 0, 1.0, 0)
			t.Put(0, 0, 1.0, 0)
			t.Put(1, 0, 3.0, 0)
			t.Put(0, 1, 3.0, 0)
			t.Put(2, 1, -1.0, 0)
			t.Put(4, 1, 4.0, 0)
		} else {
			t.Init(5, 5, 7, true)
			t.Put(1, 2, 4.0, 0)
			t.Put(2, 2, -3.0, 0)
			t.Put(3, 2, 1.0, 0)
			t.Put(4, 2, 2.0, 0)
			t.Put(2, 3, 2.0, 0)
			t.Put(1, 4, 6.0, 0)
			t.Put(4, 4, 1.0, 0)
		}
	default:
		chk.Panic("this test needs 1 or 2 procs")
	}

	b := []complex128{8.0, 45.0, -3.0, 3.0, 19.0}
	x_correct := []complex128{1, 2, 3, 4, 5}
	sum_b_to_root := false
	la.RunMumpsTestC(&t, 1e-14, b, x_correct, sum_b_to_root)
}
