package fftw

//  #include "fftw3.h"
import "C"

func InitThreads() {
	C.fftw_init_threads()
}

func PlanWithNThreads(n int) {
	C.fftw_plan_with_nthreads(C.int(n))
}

func CleanupThreads() {
	C.fftw_cleanup_threads()
}
