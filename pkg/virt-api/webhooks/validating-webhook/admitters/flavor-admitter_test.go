package admitters

import (
	"encoding/json"
	"net/http"

	apiflavor "kubevirt.io/api/flavor"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	admissionv1 "k8s.io/api/admission/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	flavorv1alpha1 "kubevirt.io/api/flavor/v1alpha1"
)

var _ = Describe("Validating Flavor Admitter", func() {
	var (
		admitter  *FlavorAdmitter
		flavorObj *flavorv1alpha1.VirtualMachineFlavor
	)

	BeforeEach(func() {
		admitter = &FlavorAdmitter{}

		flavorObj = &flavorv1alpha1.VirtualMachineFlavor{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-name",
				Namespace: "test-namespace",
			},
		}
	})

	It("should accept valid flavor", func() {
		ar := createFlavorAdmissionReview(flavorObj)
		response := admitter.Admit(ar)
		Expect(response.Allowed).To(BeTrue(), "Expected flavor to be allowed.")
	})

	It("should reject unsupported version", func() {
		ar := createFlavorAdmissionReview(flavorObj)
		ar.Request.Resource.Version = "unsupportedversion"
		response := admitter.Admit(ar)

		Expect(response.Allowed).To(BeFalse(), "Expected flavor to not be allowed")
		Expect(response.Result.Code).To(Equal(int32(http.StatusBadRequest)), "Expected error 400: BadRequest")
	})
})

var _ = Describe("Validating ClusterFlavor Admitter", func() {
	var (
		admitter         *ClusterFlavorAdmitter
		clusterFlavorObj *flavorv1alpha1.VirtualMachineClusterFlavor
	)

	BeforeEach(func() {
		admitter = &ClusterFlavorAdmitter{}

		clusterFlavorObj = &flavorv1alpha1.VirtualMachineClusterFlavor{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "test-name",
				Namespace: "test-namespace",
			},
		}
	})

	It("should accept valid flavor", func() {
		ar := createClusterFlavorAdmissionReview(clusterFlavorObj)
		response := admitter.Admit(ar)
		Expect(response.Allowed).To(BeTrue(), "Expected flavor to be allowed.")
	})

	It("should reject unsupported version", func() {
		ar := createClusterFlavorAdmissionReview(clusterFlavorObj)
		ar.Request.Resource.Version = "unsupportedversion"
		response := admitter.Admit(ar)

		Expect(response.Allowed).To(BeFalse(), "Expected flavor to not be allowed")
		Expect(response.Result.Code).To(Equal(int32(http.StatusBadRequest)), "Expected error 400: BadRequest")
	})
})

func createFlavorAdmissionReview(flavor *flavorv1alpha1.VirtualMachineFlavor) *admissionv1.AdmissionReview {
	bytes, err := json.Marshal(flavor)
	ExpectWithOffset(1, err).ToNot(HaveOccurred(), "Could not JSON encode flavor: %v", flavor)

	return &admissionv1.AdmissionReview{
		Request: &admissionv1.AdmissionRequest{
			Operation: admissionv1.Create,
			Resource: metav1.GroupVersionResource{
				Group:    flavorv1alpha1.SchemeGroupVersion.Group,
				Version:  flavorv1alpha1.SchemeGroupVersion.Version,
				Resource: apiflavor.PluralResourceName,
			},
			Object: runtime.RawExtension{
				Raw: bytes,
			},
		},
	}
}

func createClusterFlavorAdmissionReview(clusterFlavor *flavorv1alpha1.VirtualMachineClusterFlavor) *admissionv1.AdmissionReview {
	bytes, err := json.Marshal(clusterFlavor)
	ExpectWithOffset(1, err).ToNot(HaveOccurred(), "Could not JSON encode flavor: %v", clusterFlavor)

	return &admissionv1.AdmissionReview{
		Request: &admissionv1.AdmissionRequest{
			Operation: admissionv1.Create,
			Resource: metav1.GroupVersionResource{
				Group:    flavorv1alpha1.SchemeGroupVersion.Group,
				Version:  flavorv1alpha1.SchemeGroupVersion.Version,
				Resource: apiflavor.ClusterPluralResourceName,
			},
			Object: runtime.RawExtension{
				Raw: bytes,
			},
		},
	}
}
