package fftw

// #include "fftw3.h"
import "C"

type DftC2R1DPlan struct {
	plan
	Real        []float64
	Complex     []complex128
	n, cmplxLen int
}

func NewDftC2R1DPlan(n int, flag PlanFlag) (plan DftC2R1DPlan) {
	plan.n = n
	plan.cmplxLen = int(n>>1 + 1)
	plan.Real = make([]float64, n)
	plan.Complex = make([]complex128, plan.cmplxLen)
	plan.plan.p = C.fftw_plan_dft_c2r_1d(
		C.int(n),
		(*C.fftw_complex)(&plan.Complex[0]),
		(*C.double)(&plan.Real[0]),
		C.uint(flag),
	)

	return
}

func (p DftC2R1DPlan) ExecuteNewArray(in []complex128) (out []float64) {
	var in1 []complex128
	in1 = make([]complex128, p.n)
	copy(in1, in)

	out = make([]float64, p.n)
	C.fftw_execute_dft_c2r(p.p, (*C.fftw_complex)(&in1[0]), (*C.double)(&out[0]))
	return
}
