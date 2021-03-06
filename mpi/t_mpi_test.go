// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mpi

import (
	"testing"

	"github.com/cpmech/gosl/chk"
)

func TestMpi01(tst *testing.T) {

	//verbose()
	chk.PrintTitle("Mpi01. Start, IsOn, World Communicator")

	switchMPI()
	if !IsOn() {
		tst.Errorf("MPI should be on\n")
		return
	}

	wcomm := NewCommunicator(nil)
	zcomm := NewCommunicator([]int32{0})
	chk.Int32(tst, "World rank", wcomm.Rank(), 0)
	chk.Int32(tst, "Local rank", zcomm.Rank(), 0)

	wcomm.Barrier()
	zcomm.Barrier()
}
