package sanity

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/kubernetes/pkg/cloudprovider"
	"k8s.io/api/core/v1"
)

// SanityContext holds the variables that each test can depend on. It
// gets initialized before each test block runs.
type SanityContext struct {
	instance cloudprovider.Instances

	node v1.Node
	Config  *Config
}

type Config struct {
	Node *v1.Node
	NodeAddress []v1.NodeAddress
}


func Test(t *testing.T, reqConfig *Config)  {

	sc := &SanityContext{
		Config: reqConfig,
	}

	registerTestsInGinkgo(sc)
	RegisterFailHandler(Fail)

	RunSpecs(t, "Cloud Controller Manager Test Suite")
}


func GinkgoTest(reqConfig *Config) {
	sc := &SanityContext{
		Config: reqConfig,
	}

	registerTestsInGinkgo(sc)
}