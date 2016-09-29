package services_test

import (
	. "kubevirt/core/pkg/virt-controller/services"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/go-kit/kit/log"
	"os"
	"kubevirt/core/pkg/virt-controller/entities"
	"bytes"
)

var _ = Describe("Template", func() {

	logger := log.NewLogfmtLogger(os.Stderr)

	Describe("Rendering", func() {
		Context("with correct parameters", func() {
			It("should work", func() {
				svc, err := NewTemplateService(logger, "../templates/manifest-template.yaml", "kubevirt", "virt-launcher")
				Expect(err).To(BeNil())

				buffer := new(bytes.Buffer)

				err = svc.RenderManifest(&entities.VM{Name:"testvm"}, buffer)

				Expect(err).To(BeNil())
				Expect(buffer.String()).To(ContainSubstring("image: kubevirt/virt-launcher"))
				Expect(buffer.String()).To(ContainSubstring("domain: testvm"))
				Expect(buffer.String()).To(ContainSubstring("name: virt-launcher-testvm"))
				Expect(buffer.String()).To(ContainSubstring("/etc/vdsm/dom/testvm.xml"))
			})
		})
		Context("with wrong template path", func() {
			It("should fail", func() {
				_, err := NewTemplateService(logger, "templates/manifest-template.yaml", "kubevirt", "virt-launcher")
				Expect(err).To(Not(BeNil()))
			})
		})
	})

})
