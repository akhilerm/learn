package client

import (
	"context"
	"github.com/akhilerm/TEst/NDMRefactTest/filter"
	"github.com/openebs/node-disk-manager/blockdevice"
	apis "github.com/openebs/node-disk-manager/pkg/apis/openebs/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (c *Client) ListBlockDevicesAPI(selector labels.Selector) (*apis.BlockDeviceList, error) {
	listBlockDevice := &apis.BlockDeviceList{
		TypeMeta: v1.TypeMeta{
			Kind:       "BlockDevice",
			APIVersion: "openebs.io/v1alpha1",
		},
	}

	mls := client.MatchingLabelsSelector{Selector: selector}
	ns := client.InNamespace(c.namespace)

	listOptions := []client.ListOption{
		mls,
		ns,
	}

	err := c.client.List(context.TODO(), listBlockDevice, listOptions...)
	if err != nil {
		return nil, err
	}

	return listBlockDevice, nil
}

func (c *Client) ListBlockDevices(selector labels.Selector, filters ...filter.Func) ([]blockdevice.BlockDevice, error) {
	bdAPIList, err := c.ListBlockDevicesAPI(selector)
	if err != nil {
		return nil, err
	}

	filteredBDList := filter.Filter(bdAPIList, filters...)

	bdList := make([]blockdevice.BlockDevice, 0)
	err = convertBlockDeviceAPIListToBlockDeviceList(filteredBDList, &bdList)
	if err != nil {
		return nil, err
	}
	return bdList, nil
}

