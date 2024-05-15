// Code generated by ffakes v0.0.9 DO NOT EDIT.

package nearearth

import (
	"testing"
)

var (
	_ ImpactService = (*FakeImpactService)(nil)
)

type FakeImpactService struct {
	t           *testing.T
	GetAllCount int
	FGetAll     []func() (Impacts, error)
}

type GetAllFunc = func() (Impacts, error)
type ImpactServiceOption func(f *FakeImpactService)

func OnGetAll(fn ...GetAllFunc) ImpactServiceOption {
	return func(f *FakeImpactService) {
		f.FGetAll = append(f.FGetAll, fn...)
	}
}

func (f *FakeImpactService) OnGetAll(fns ...GetAllFunc) {
	f.FGetAll = append(f.FGetAll, fns...)
}

func NewFakeImpactService(t *testing.T, opts ...ImpactServiceOption) *FakeImpactService {
	f := &FakeImpactService{t: t}
	for _, opt := range opts {
		opt(f)
	}
	t.Cleanup(func() {
		if f.GetAllCount != len(f.FGetAll) {
			t.Fatalf("expected GetAll to be called %d times but got %d", len(f.FGetAll), f.GetAllCount)
		}
	})
	return f
}

func (fake *FakeImpactService) GetAll() (Impacts, error) {
	var idx = fake.GetAllCount
	if fake.GetAllCount >= len(fake.FGetAll) {
		idx = len(fake.FGetAll) - 1
	}
	if len(fake.FGetAll) != 0 {
		o1, o2 := fake.FGetAll[idx]()
		fake.GetAllCount++
		return o1, o2
	}
	return Impacts{}, nil
}
