// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package la

import (
	"math"
	"math/cmplx"
	"testing"

	"github.com/cpmech/gosl/chk"
)

func checkResid(tst *testing.T, a *Matrix, x, b Vector, tolNorm float64) {
	r := NewVector(len(x))
	r.Apply(-1, b)           // r := -b
	MatVecMulAdd(r, 1, a, x) // r += 1*a*x
	resid := r.Norm()
	if resid > tolNorm {
		tst.Errorf("residual is too large: %g\n", resid)
		return
	}
}

func checkResidC(tst *testing.T, a *MatrixC, x, b VectorC, tolNorm float64) {
	r := NewVectorC(len(x))
	r.Apply(-1, b)            // r = -b
	MatVecMulAddC(r, 1, a, x) // r += 1*a*x
	resid := cmplx.Abs(r.Norm())
	if resid > tolNorm {
		tst.Errorf("residual is too large: %g\n", resid)
		return
	}
}

func calcLLt(L *Matrix) (LLt *Matrix) {
	LLt = NewMatrix(L.M, L.M)
	for i := 0; i < L.M; i++ {
		for j := 0; j < L.M; j++ {
			for k := 0; k < L.M; k++ {
				LLt.Add(i, j, L.Get(i, k)*L.Get(j, k))
			}
		}
	}
	return
}

func TestCholesky01(tst *testing.T) {

	//verbose()
	chk.PrintTitle("Cholesky 01")

	a := NewMatrixSlice([][]float64{
		{25.0, 15.0, -5.0},
		{15.0, 18.0, 0.0},
		{-5.0, 0.0, 11.0},
	})

	L := NewMatrix(3, 3)
	Cholesky(L, a) // L is such that: A = L * transp(L)
	LLt := calcLLt(L)
	chk.Matrix(tst, "a = LLt", 1e-17, LLt.GetSlice(), a.GetSlice())
	chk.Matrix(tst, "L", 1e-17, L.GetSlice(), [][]float64{
		{+5, 0, 0},
		{+3, 3, 0},
		{-1, 1, 3},
	})
}

func TestCholesky02(tst *testing.T) {

	//verbose()
	chk.PrintTitle("Cholesky 02")

	a := NewMatrixSlice([][]float64{
		{2, 1, 1, 3, 2},
		{1, 2, 2, 1, 1},
		{1, 2, 9, 1, 5},
		{3, 1, 1, 7, 1},
		{2, 1, 5, 1, 8},
	})
	L := NewMatrix(5, 5)
	Cholesky(L, a)
	chk.Matrix(tst, "a = LLt", 1e-15, calcLLt(L).GetSlice(), a.GetSlice())
	chk.Matrix(tst, "L", 1e-15, L.GetSlice(), [][]float64{
		{math.Sqrt2, 0, 0, 0, 0},
		{1.0 / math.Sqrt2, math.Sqrt(3.0 / 2.0), 0, 0, 0},
		{1.0 / math.Sqrt2, math.Sqrt(3.0 / 2.0), math.Sqrt(7.0), 0, 0},
		{3.0 / math.Sqrt2, -1.0 / math.Sqrt(6.0), 0, math.Sqrt(7.0 / 3.0), 0},
		{math.Sqrt2, 0, 4.0 / math.Sqrt(7.0), -2.0 * math.Sqrt(3.0/7.0), math.Sqrt2},
	})
}

func TestSPDsolve01(tst *testing.T) {

	//verbose()
	chk.PrintTitle("TestSPDsolve 01")

	a := NewMatrixSlice([][]float64{
		{2, 1, 1, 3, 2},
		{1, 2, 2, 1, 1},
		{1, 2, 9, 1, 5},
		{3, 1, 1, 7, 1},
		{2, 1, 5, 1, 8},
	})
	b := []float64{-2, 4, 3, -5, 1}
	x := make([]float64, 5)
	SolveRealLinSysSPD(x, a, b)
	checkResid(tst, a, x, b, 1e-14)
	chk.Vector(tst, "x = inv(a) * b", 1e-13, x, []float64{
		-629.0 / 98.0,
		+237.0 / 49.0,
		-53.0 / 49.0,
		+62.0 / 49.0,
		+23.0 / 14.0,
	})
}

func TestSPDsolve02(tst *testing.T) {

	//verbose()
	chk.PrintTitle("TestSPDsolve 02")

	a := NewMatrixSlice([][]float64{
		{2, 1, 1, 3, 2},
		{1, 2, 2, 1, 1},
		{1, 2, 9, 1, 5},
		{3, 1, 1, 7, 1},
		{2, 1, 5, 1, 8},
	})
	b := []float64{-2, 4, 3, -5, 1}
	B := []float64{24, 29, 110, 12, 102}
	x := make([]float64, 5)
	X := make([]float64, 5)
	SolveTwoRealLinSysSPD(x, X, a, b, B)
	checkResid(tst, a, x, b, 1e-14)
	checkResid(tst, a, X, B, 1e-14)
	chk.Vector(tst, "x = inv(a) * b", 1e-13, x, []float64{
		-629.0 / 98.0,
		+237.0 / 49.0,
		-53.0 / 49.0,
		+62.0 / 49.0,
		+23.0 / 14.0,
	})
	chk.Vector(tst, "X = inv(a) * B", 1e-13, X, []float64{0, 4, 7, -1, 8})
}
