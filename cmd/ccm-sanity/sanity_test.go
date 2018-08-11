package sanity


import (
	"fmt"
	"testing"

	"github.com/kubernetes-csi/csi-test/pkg/sanity"
)

const (
	prefix string = "csi."
	cloudProvider string = "fake"
)

var (
	VERSION = "(dev)"
	version bool
	config  sanity.Config
)

func init() {

}

func TestSanity(t *testing.T) {
	if version {
		fmt.Printf("Version = %s\n", VERSION)
		return
	}
	if len(config.Address) == 0 {
		t.Fatalf("--%sendpoint must be provided with an CSI endpoint", prefix)
	}
	sanity.Test(t, &config)
}

