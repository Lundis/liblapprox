package base

import (
	. "code.google.com/p/liblundis/lmath/util/cont"
)
type Basis interface {
	Function() Function
	GetBases(bases_out []float64)
	SetBases(bases_in []float64)
	Get(deg int) float64
	Set(deg int, v float64)
	String() string
}

type BasisImpl struct {
	V []float64
}

type BasisImplConverter func(*BasisImpl) Basis

func NewBasisImpl(degree int) *BasisImpl {
	b := new(BasisImpl)
	b.V = make([]float64, degree + 1)
	return b
}

func (self BasisImpl) GetBases(bases_out []float64) {
	for i := range self.V {
		bases_out[i] = self.V[i]
	}
}

func (self BasisImpl) SetBases(bases_in []float64) {
	for i := range self.V {
		self.V[i] = bases_in[i]
	}
}

func (self BasisImpl) Get(deg int) float64 {
	return self.V[deg]
}

func (self BasisImpl) Set(deg int, v float64) {
	self.V[deg] = v
}