package fake

import (
	"context"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/kubernetes/pkg/cloudprovider"
	"github.com/pharmer/pharmer/cloud"
//	fakecloud "k8s.io/kubernetes/pkg/cloudprovider/providers/fake"
)

type instances struct {
}

func newInstances() cloudprovider.Instances {
	return &instances{}
}

func (i *instances) NodeAddresses(_ context.Context, name types.NodeName) ([]v1.NodeAddress, error) {
	return fake.Addresses, nil
}



func (i *instances) NodeAddressesByProviderID(_ context.Context, providerID string) ([]v1.NodeAddress, error) {
	return fake.Addresses, nil
}


func (i *instances) InstanceID(_ context.Context, nodeName types.NodeName) (string, error) {
	return fake.InstanceID(context.TODO(), nodeName)
}

func (i *instances) InstanceType(ctx context.Context, nodeName types.NodeName) (string, error) {
	return fake.InstanceType(ctx, nodeName)
}

func (i *instances) InstanceTypeByProviderID(ctx context.Context, providerID string) (string, error) {
	return fake.InstanceTypeByProviderID(ctx, providerID)
}

func (i *instances) AddSSHKeyToAllInstances(_ context.Context, user string, keyData []byte) error {
	return cloud.ErrNotImplemented
}

func (i *instances) CurrentNodeName(_ context.Context, hostname string) (types.NodeName, error) {
	return types.NodeName(hostname), nil
}

func (i *instances) InstanceExistsByProviderID(ctx context.Context, providerID string) (bool, error) {
	return fake.InstanceExistsByProviderID(ctx, providerID)
}

func (i *instances) InstanceShutdownByProviderID(ctx context.Context, providerID string) (bool, error) {
	return false, cloudprovider.NotImplemented
}

