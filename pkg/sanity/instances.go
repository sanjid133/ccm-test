package sanity


import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
//	"k8s.io/api/core/v1"
//	"k8s.io/kubernetes/pkg/cloudprovider"
	"context"
	"k8s.io/kubernetes/pkg/cloudprovider"
)

var _ = DescribeSanity("Instances", func(sc *SanityContext) {


	Describe("NodeAddresses", func() {
		BeforeEach(func() {
			_, err := sc.instance.NodeAddresses(context.TODO(), "some-fake-node")
			if err.Error() == cloudprovider.NotImplemented.Error() {
				Skip("NodeAddress not implemented")
			}
		})
		It("should return node addresses", func() {
			nodeAddress, err := sc.instance.NodeAddresses(context.TODO(), sc.Config.Node.Name)
			By("checking successful response")
			Expect(err).NotTo(HaveOccurred())
			Expect(nodeAddress).NotTo(BeEmpty())
			Expect(nodeAddress).To(Equal(sc.Config.NodeAddress))
		})

		It("should not return node address", func() {
			nodeAddress, err := sc.instance.NodeAddresses(context.TODO(), "some-fake-node")
			By("checking successful response")
			Expect(err).To(HaveOccurred())
			Expect(nodeAddress).To(BeEmpty())
		})
	})

})
