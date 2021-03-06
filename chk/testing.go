// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chk

import (
	"math"
	"math/cmplx"
	"testing"
)

// Scalar compares two scalars
func Scalar(tst *testing.T, msg string, tol, res, correct float64) {
	CheckAndPrint(tst, msg, tol, math.Abs(res-correct))
}

// ScalarC compares two scalars (complex version)
func ScalarC(tst *testing.T, msg string, tolNorm float64, res, correct complex128) {
	CheckAndPrint(tst, msg, tolNorm, cmplx.Abs(res-correct))
}

// AnaNum compares analytical versus numerical values
func AnaNum(tst *testing.T, msg string, tol, ana, num float64, verbose bool) {
	e := PrintAnaNum(msg, tol, ana, num, verbose)
	if e != nil {
		tst.Errorf("%v", e.Error())
	}
}

// String compares two strings
func String(tst *testing.T, str, correct string) {
	if str != correct {
		PrintFail("error %q != %q\n", str, correct)
		tst.Errorf("string failed with: %q != %q", str, correct)
		return
	}
	PrintOk("%s == %s", str, correct)
}

// Int compares two ints
func Int(tst *testing.T, msg string, val, correct int) {
	if val != correct {
		PrintFail("%s: error %d != %d\n", msg, val, correct)
		tst.Errorf("%s failed with: %d != %d", msg, val, correct)
		return
	}
	PrintOk("%s: %d == %d", msg, val, correct)
}

// Int32 compares two int32
func Int32(tst *testing.T, msg string, val, correct int32) {
	if val != correct {
		PrintFail("%s: error %d != %d\n", msg, val, correct)
		tst.Errorf("%s failed with: %d != %d", msg, val, correct)
		return
	}
	PrintOk("%s: %d == %d", msg, val, correct)
}

// Int64 compares two int64
func Int64(tst *testing.T, msg string, val, correct int64) {
	if val != correct {
		PrintFail("%s: error %d != %d\n", msg, val, correct)
		tst.Errorf("%s failed with: %d != %d", msg, val, correct)
		return
	}
	PrintOk("%s: %d == %d", msg, val, correct)
}

// Ints compares two slices of integers
func Ints(tst *testing.T, msg string, a, b []int) {
	if len(a) != len(b) {
		PrintFail("%s error len(a)=%d != len(b)=%d\n", msg, len(a), len(b))
		tst.Errorf("%s failed: slices have different lengths: %v != %v", msg, a, b)
		return
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			PrintFail("%s error %d != %d\n", msg, a[i], b[i])
			tst.Errorf("%s failed: slices are different: %dth comp %v != %v\n%v != \n%v", msg, i, a[i], b[i], a, b)
			return
		}
	}
	PrintOk(msg)
}

// Int32s compares two slices of 32 integers
func Int32s(tst *testing.T, msg string, a, b []int32) {
	if len(a) != len(b) {
		PrintFail("%s error len(a)=%d != len(b)=%d\n", msg, len(a), len(b))
		tst.Errorf("%s failed: slices have different lengths: %v != %v", msg, a, b)
		return
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			PrintFail("%s error %d != %d\n", msg, a[i], b[i])
			tst.Errorf("%s failed: slices are different: %dth comp %v != %v\n%v != \n%v", msg, i, a[i], b[i], a, b)
			return
		}
	}
	PrintOk(msg)
}

// Int64s compares two slices of 64 integers
func Int64s(tst *testing.T, msg string, a, b []int64) {
	if len(a) != len(b) {
		PrintFail("%s error len(a)=%d != len(b)=%d\n", msg, len(a), len(b))
		tst.Errorf("%s failed: slices have different lengths: %v != %v", msg, a, b)
		return
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			PrintFail("%s error %d != %d\n", msg, a[i], b[i])
			tst.Errorf("%s failed: slices are different: %dth comp %v != %v\n%v != \n%v", msg, i, a[i], b[i], a, b)
			return
		}
	}
	PrintOk(msg)
}

// Bools compare two slices of bools
func Bools(tst *testing.T, msg string, a, b []bool) {
	if len(a) != len(b) {
		PrintFail("%s error len(%q)=%d != len(%q)=%d\n", msg, a, len(a), b, len(b))
		tst.Errorf("%s failed: slices have different lengths: %v != %v", msg, a, b)
		return
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			PrintFail("%s error %v != %v\n", msg, a[i], b[i])
			tst.Errorf("%s failed: slices are different: %dth comp %v != %v\n%v != \n%v", msg, i, a[i], b[i], a, b)
			return
		}
	}
	PrintOk(msg)
}

// Strings compare two slices of strings
func Strings(tst *testing.T, msg string, a, b []string) {
	if len(a) != len(b) {
		PrintFail("%s error len(%q)=%d != len(%q)=%d\n", msg, a, len(a), b, len(b))
		tst.Errorf("%s failed: slices have different lengths: %v != %v", msg, a, b)
		return
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			PrintFail("%s error %v != %v\n", msg, a[i], b[i])
			tst.Errorf("%s failed: slices are different: %dth comp %v != %v\n%v != \n%v", msg, i, a[i], b[i], a, b)
			return
		}
	}
	PrintOk(msg)
}

// Matrix compares two matrices
func Matrix(tst *testing.T, msg string, tol float64, res, correct [][]float64) {
	zero := false
	if len(correct) == 0 {
		zero = true
	} else {
		if len(res) != len(correct) {
			PrintFail("%s error\n", msg)
			tst.Errorf("%s failed: res and correct matrices have different lengths. %d != %d", msg, len(res), len(correct))
			return
		}
	}
	var diff, maxdiff float64
	for i := 0; i < len(res); i++ {
		if !zero {
			if len(res[i]) != len(correct[i]) {
				PrintFail("%s  error\n", msg)
				tst.Errorf("%s failed: matrices have different number of columns", msg)
			}
		}
		for j := 0; j < len(res[i]); j++ {
			if math.IsNaN(res[i][j]) {
				tst.Errorf("%s failed: NaN detected => %v", msg, res[i][j])
			}
			if zero {
				diff = math.Abs(res[i][j])
			} else {
				diff = math.Abs(res[i][j] - correct[i][j])
			}
			if diff > maxdiff {
				maxdiff = diff
			}
		}
	}
	CheckAndPrint(tst, msg, tol, maxdiff)
}

// Vector compares two vectors
func Vector(tst *testing.T, msg string, tol float64, res, correct []float64) {
	zero := false
	if len(correct) == 0 {
		zero = true
	} else {
		if len(res) != len(correct) {
			PrintFail("%s error\n", msg)
			tst.Errorf("%s failed: res and correct vectors have different lengths. %d != %d", msg, len(res), len(correct))
			return
		}
	}
	var diff, maxdiff float64
	for i := 0; i < len(res); i++ {
		if math.IsNaN(res[i]) {
			tst.Errorf("%s failed: NaN detected => %v", msg, res[i])
		}
		if zero {
			diff = math.Abs(res[i])
		} else {
			diff = math.Abs(res[i] - correct[i])
		}
		if diff > maxdiff {
			maxdiff = diff
		}
	}
	CheckAndPrint(tst, msg, tol, maxdiff)
}

// MatrixC compares two matrices of complex nummbers
func MatrixC(tst *testing.T, msg string, tol float64, res, correct [][]complex128) {
	zero := false
	if len(correct) == 0 {
		zero = true
	} else {
		if len(res) != len(correct) {
			PrintFail("%s error\n", msg)
			tst.Errorf("%s failed: res and correct matrices have different lengths. %d != %d", msg, len(res), len(correct))
			return
		}
	}
	var diff, maxdiff float64
	var diffz, maxdiffz float64
	for i := 0; i < len(res); i++ {
		if !zero {
			if len(res[i]) != len(correct[i]) {
				PrintFail("%s  error\n", msg)
				tst.Errorf("%s failed: matrices have different number of columns", msg)
			}
		}
		for j := 0; j < len(res[i]); j++ {
			if zero {
				diff = math.Abs(real(res[i][j]))
				diffz = math.Abs(imag(res[i][j]))
			} else {
				diff = math.Abs(real(res[i][j]) - real(correct[i][j]))
				diffz = math.Abs(imag(res[i][j]) - imag(correct[i][j]))
			}
			if diff > maxdiff {
				maxdiff = diff
			}
			if diffz > maxdiffz {
				maxdiffz = diffz
			}
		}
	}
	if maxdiff > tol || maxdiffz > tol {
		PrintFail("%s error |maxdiff| = %g,  |maxdiffz| = %g\n", msg, maxdiff, maxdiffz)
		tst.Errorf("%s failed with |maxdiff| = %g,  |maxdiffz| = %g", msg, maxdiff, maxdiffz)
		return
	}
	PrintOk(msg)
}

// VectorC compares two vectors of complex nummbers
func VectorC(tst *testing.T, msg string, tol float64, res, correct []complex128) {
	zero := false
	if len(correct) == 0 {
		zero = true
	} else {
		if len(res) != len(correct) {
			PrintFail("%s error\n", msg)
			tst.Errorf("%s failed: res and correct matrices have different lengths. %d != %d", msg, len(res), len(correct))
			return
		}
	}
	var diff, maxdiff float64
	var diffz, maxdiffz float64
	for i := 0; i < len(res); i++ {
		if zero {
			diff = math.Abs(real(res[i]))
			diffz = math.Abs(imag(res[i]))
		} else {
			diff = math.Abs(real(res[i]) - real(correct[i]))
			diffz = math.Abs(imag(res[i]) - imag(correct[i]))
		}
		if diff > maxdiff {
			maxdiff = diff
		}
		if diffz > maxdiffz {
			maxdiffz = diffz
		}
	}
	if maxdiff > tol || maxdiffz > tol {
		PrintFail("%s error |maxdiff| = %g,  |maxdiffz| = %g\n", msg, maxdiff, maxdiffz)
		tst.Errorf("%s failed with |maxdiff| = %g,  |maxdiffz| = %g", msg, maxdiff, maxdiffz)
		return
	}
	PrintOk(msg)
}

// StrMat compares matrices of strings
func StrMat(tst *testing.T, msg string, res, correct [][]string) {
	empty := false
	if len(correct) == 0 {
		empty = true
	} else {
		if len(res) != len(correct) {
			PrintFail("%s error len(res)=%d != len(correct)=%d\n", msg, len(res), len(correct))
			tst.Errorf("%s failed: res and correct matrices have different lengths. %d != %d", msg, len(res), len(correct))
			return
		}
	}
	for i := 0; i < len(res); i++ {
		if !empty {
			if len(res[i]) != len(correct[i]) {
				PrintFail("%s error len(res[%d])=%d != len(correct[%d])=%d\n", msg, i, len(res[i]), i, len(correct[i]))
				tst.Errorf("%s failed: string matrices have different number of columns", msg)
				return
			}
		}
		for j := 0; j < len(res[i]); j++ {
			var c string
			if !empty {
				c = correct[i][j]
			}
			if res[i][j] != c {
				PrintFail("%s error [%d,%d] %v != %v\n", msg, i, j, res[i][j], c)
				tst.Errorf("%s failed: different str matrices:\n [%d,%d] item is wrong: %v != %v", msg, i, j, res[i][j], c)
				return
			}
		}
	}
	PrintOk(msg)
}

// IntMat compares matrices of ints
func IntMat(tst *testing.T, msg string, res, correct [][]int) {
	zero := false
	if len(correct) == 0 {
		zero = true
	} else {
		if len(res) != len(correct) {
			PrintFail("%s error len(res)=%d != len(correct)=%d\n", msg, len(res), len(correct))
			tst.Errorf("%s failed: res and correct matrices have different lengths. %d != %d", msg, len(res), len(correct))
			return
		}
	}
	for i := 0; i < len(res); i++ {
		if !zero {
			if len(res[i]) != len(correct[i]) {
				PrintFail("%s error len(res[%d])=%d != len(correct[%d])=%d\n", msg, i, len(res[i]), i, len(correct[i]))
				tst.Errorf("%s failed: matrices have different number of columns", msg)
				return
			}
		}
		for j := 0; j < len(res[i]); j++ {
			var c int
			if !zero {
				c = correct[i][j]
			}
			if res[i][j] != c {
				PrintFail("%s error [%d,%d] %v != %v\n", msg, i, j, res[i][j], c)
				tst.Errorf("%s failed: different int matrices:\n [%d,%d] item is wrong: %v != %v", msg, i, j, res[i][j], c)
				return
			}
		}
	}
	PrintOk(msg)
}

// Deep3 compares two deep3 slices
func Deep3(tst *testing.T, msg string, tol float64, a, b [][][]float64) {
	if len(a) != len(b) {
		PrintFail("%s error len(a)=%d != len(b)=%d\n", msg, len(a), len(b))
		tst.Errorf("%s failed: slices have different lengths: %v != %v", msg, a, b)
		return
	}
	for i := 0; i < len(a); i++ {
		if len(a[i]) != len(b[i]) {
			PrintFail("%s error len(a[%d])=%d != len(b[%d])=%d\n", msg, i, len(a[i]), i, len(b[i]))
			tst.Errorf("%s failed: subslices have different lengths", msg)
			return
		}
		for j := 0; j < len(a[i]); j++ {
			if len(a[i][j]) != len(b[i][j]) {
				PrintFail("%s error len(a[%d][%d])=%d != len(b[%d][%d])=%d\n", msg, i, j, len(a[i][j]), i, j, len(b[i][j]))
				tst.Errorf("%s failed: subsubslices have different lengths", msg)
				return
			}
			for k := 0; k < len(a[i][j]); k++ {
				if math.Abs(a[i][j][k]-b[i][j][k]) > tol {
					PrintFail("%s error %v != %v\n", msg, a[i][j][k], b[i][j][k])
					tst.Errorf("%s failed: slices are different: %d,%d,%d component %v != %v\n%v != \n%v", msg, i, j, k, a[i][j][k], b[i][j][k], a, b)
					return
				}
			}
		}
	}
	PrintOk(msg)
}

// Deep4 compares two deep4 slices
func Deep4(tst *testing.T, msg string, tol float64, a, b [][][][]float64) {
	if len(a) != len(b) {
		PrintFail("%s error len(a)=%d != len(b)=%d\n", msg, len(a), len(b))
		tst.Errorf("%s failed: slices have different lengths: %v != %v", msg, a, b)
		return
	}
	for i := 0; i < len(a); i++ {
		if len(a[i]) != len(b[i]) {
			PrintFail("%s error len(a[%d])=%d != len(b[%d])=%d\n", msg, i, len(a[i]), i, len(b[i]))
			tst.Errorf("%s failed: subslices have different lengths", msg)
			return
		}
		for j := 0; j < len(a[i]); j++ {
			if len(a[i][j]) != len(b[i][j]) {
				PrintFail("%s error len(a[%d][%d])=%d != len(b[%d][%d])=%d\n", msg, i, j, len(a[i][j]), i, j, len(b[i][j]))
				tst.Errorf("%s failed: subsubslices have different lengths", msg)
				return
			}
			for k := 0; k < len(a[i][j]); k++ {
				if len(a[i][j][k]) != len(b[i][j][k]) {
					PrintFail("%s error len(a[%d][%d][%d])=%d != len(b[%d][%d][%d])=%d\n", msg, i, j, k, len(a[i][j][k]), i, j, k, len(b[i][j][k]))
					tst.Errorf("%s failed: subsubsubslices have different lengths", msg)
					return
				}
				for l := 0; l < len(a[i][j][k]); l++ {
					if math.Abs(a[i][j][k][l]-b[i][j][k][l]) > tol {
						PrintFail("%s error %v != %v\n", msg, a[i][j][k][l], b[i][j][k][l])
						tst.Errorf("%s failed: slices are different: %d,%d,%d,%d component %v != %v\n%v != \n%v", msg, i, j, k, l, a[i][j][k][l], b[i][j][k][l], a, b)
						return
					}
				}
			}
		}
	}
	PrintOk(msg)
}
