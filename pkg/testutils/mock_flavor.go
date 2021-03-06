package testutils

import (
	k8sfield "k8s.io/apimachinery/pkg/util/validation/field"

	v1 "kubevirt.io/api/core/v1"
	flavorv1alpha1 "kubevirt.io/api/flavor/v1alpha1"
	"kubevirt.io/kubevirt/pkg/flavor"
)

type MockFlavorMethods struct {
	FindFlavorSpecFunc func(vm *v1.VirtualMachine) (*flavorv1alpha1.VirtualMachineFlavorSpec, error)
	ApplyToVmiFunc     func(field *k8sfield.Path, flavorspec *flavorv1alpha1.VirtualMachineFlavorSpec, vmiSpec *v1.VirtualMachineInstanceSpec) flavor.Conflicts
}

var _ flavor.Methods = &MockFlavorMethods{}

func (m *MockFlavorMethods) FindFlavorSpec(vm *v1.VirtualMachine) (*flavorv1alpha1.VirtualMachineFlavorSpec, error) {
	return m.FindFlavorSpecFunc(vm)
}

func (m *MockFlavorMethods) ApplyToVmi(field *k8sfield.Path, flavorspec *flavorv1alpha1.VirtualMachineFlavorSpec, vmiSpec *v1.VirtualMachineInstanceSpec) flavor.Conflicts {
	return m.ApplyToVmiFunc(field, flavorspec, vmiSpec)
}

func NewMockFlavorMethods() *MockFlavorMethods {
	return &MockFlavorMethods{
		FindFlavorSpecFunc: func(_ *v1.VirtualMachine) (*flavorv1alpha1.VirtualMachineFlavorSpec, error) {
			return nil, nil
		},
		ApplyToVmiFunc: func(_ *k8sfield.Path, _ *flavorv1alpha1.VirtualMachineFlavorSpec, _ *v1.VirtualMachineInstanceSpec) flavor.Conflicts {
			return nil
		},
	}
}
