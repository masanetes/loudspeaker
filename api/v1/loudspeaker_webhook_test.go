package v1

import (
	"bytes"
	"context"
	"errors"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func mutateTest(before string, after string) {
	ctx := context.Background()

	y, err := os.ReadFile(before)
	Expect(err).NotTo(HaveOccurred())
	d := yaml.NewYAMLOrJSONDecoder(bytes.NewReader(y), 4096)
	beforeLoudspeaker := &Loudspeaker{}
	err = d.Decode(beforeLoudspeaker)
	Expect(err).NotTo(HaveOccurred())

	err = k8sClient.Create(ctx, beforeLoudspeaker)
	Expect(err).NotTo(HaveOccurred())

	ret := &Loudspeaker{}
	err = k8sClient.Get(ctx, types.NamespacedName{Name: beforeLoudspeaker.GetName(), Namespace: beforeLoudspeaker.GetNamespace()}, ret)
	Expect(err).NotTo(HaveOccurred())

	y, err = os.ReadFile(after)
	Expect(err).NotTo(HaveOccurred())
	d = yaml.NewYAMLOrJSONDecoder(bytes.NewReader(y), 4096)
	afterLoudspeaker := &Loudspeaker{}
	err = d.Decode(afterLoudspeaker)
	Expect(err).NotTo(HaveOccurred())

	Expect(ret.Spec).Should(Equal(afterLoudspeaker.Spec))
}

func validateTest(file string, valid bool) {
	ctx := context.Background()
	y, err := os.ReadFile(file)
	Expect(err).NotTo(HaveOccurred())
	d := yaml.NewYAMLOrJSONDecoder(bytes.NewReader(y), 4096)
	loudspeaker := &Loudspeaker{}
	err = d.Decode(loudspeaker)
	Expect(err).NotTo(HaveOccurred())

	err = k8sClient.Create(ctx, loudspeaker)

	if valid {
		Expect(err).NotTo(HaveOccurred(), "Loudspeaker: %v", loudspeaker)
	} else {
		Expect(err).To(HaveOccurred(), "Loudspeaker: %v", loudspeaker)
		statusErr := &k8serrors.StatusError{}
		Expect(errors.As(err, &statusErr)).To(BeTrue())
		expected := loudspeaker.Annotations["message"]
		Expect(statusErr.ErrStatus.Message).To(ContainSubstring(expected))
	}
}

var _ = Describe("Loudspeaker Admission Webhook", func() {
	Context("Mutating Webhook", func() {
		It("should mutate a Loudspeaker", func() {
			mutateTest(filepath.Join("testdata", "mutating", "before.yaml"), filepath.Join("testdata", "mutating", "after.yaml"))
		})
	})
	Context("Validating Webhook", func() {
		It("valid Loudspeaker", func() {
			validateTest(filepath.Join("testdata", "validating", "valid.yaml"), true)
		})
		It("invalid Loudspeaker: Unsupported .spec.listeners[*].type", func() {
			validateTest(filepath.Join("testdata", "validating", "unsupported-listener-type.yaml"), false)
		})
		It("invalid Loudspeaker: Duplicated .spec.listeners[*].credentials", func() {
			validateTest(filepath.Join("testdata", "validating", "duplicate-credentials.yaml"), false)
		})
		It("invalid Loudspeaker: Without .spec.listeners", func() {
			validateTest(filepath.Join("testdata", "validating", "without-listeners.yaml"), false)
		})
		//It("invalid Loudspeaker: Without .spec.listeners[*].credentials", func() {
		//	validateTest(filepath.Join("testdata", "validating", "without-credentials.yaml"), false)
		//})
		It("invalid Loudspeaker: Without .spec.listeners[*].subscribes", func() {
			validateTest(filepath.Join("testdata", "validating", "without-subscribes.yaml"), false)
		})
		//It("invalid Loudspeaker: Without .spec.listeners[*].subscribes[*].namespace", func() {
		//	validateTest(filepath.Join("testdata", "validating", "without-namespace.yaml"), false)
		//})
	})
})
