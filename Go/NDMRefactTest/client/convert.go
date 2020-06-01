package client

import (
	"github.com/openebs/node-disk-manager/blockdevice"
	apis "github.com/openebs/node-disk-manager/pkg/apis/openebs/v1alpha1"
)

func convertBlockDeviceAPIListToBlockDeviceList(in *apis.BlockDeviceList, out *[]blockdevice.BlockDevice) error {
	return nil
}

