package fftw

// #include "fftw3.h"
import "C"

type DftR2C1DPlan struct {
	plan
	Real        []float64
	Complex     []complex128
	n, cmplxLen int
}

func NewDftR2C1DPlan(n int, flag PlanFlag) (plan DftR2C1DPlan) {
	plan.n = n
	plan.cmplxLen = int(n>>1 + 1)
	plan.Real = make([]float64, n)
	plan.Complex = make([]complex128, plan.cmplxLen)
	plan.plan.p = C.fftw_plan_dft_r2c_1d(
		C.int(n),
		(*C.double)(&plan.Real[0]),
		(*C.fftw_complex)(&plan.Complex[0]),
		C.uint(flag),
	)

	return
}

func (p DftR2C1DPlan) ExecuteNewArray(in []float64) (out []complex128) {
	// zero padding
	var in1 []float64
	in1 = make([]float64, p.n)
	copy(in1, in)

	out = make([]complex128, p.cmplxLen)
	C.fftw_execute_dft_r2c(p.p, (*C.double)(&in1[0]), (*C.fftw_complex)(&out[0]))
	return
}
