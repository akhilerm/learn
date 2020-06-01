package filter

import (
	"fmt"
	apis "github.com/openebs/node-disk-manager/pkg/apis/openebs/v1alpha1"
	"github.com/openebs/node-disk-manager/pkg/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Config struct{}

type Func func(blockDeviceList *apis.BlockDeviceList) (filteredBlockDeviceList *apis.BlockDeviceList)

func WithBlockDeviceClaimState(claimState apis.DeviceClaimState) Func {
	return func(blockDeviceList *apis.BlockDeviceList) *apis.BlockDeviceList {
		filteredBDList := &apis.BlockDeviceList{
			TypeMeta: metav1.TypeMeta{
				Kind:       "BlockDevice",
				APIVersion: "openebs.io/v1alpha1",
			},
		}
		for _, bd := range blockDeviceList.Items {
			if bd.Status.ClaimState == claimState {
				filteredBDList.Items = append(filteredBDList.Items, bd)
			}
		}
		return filteredBDList
	}
}

func WithBlockDeviceUnclaimed() Func {
	return WithBlockDeviceClaimState(apis.BlockDeviceUnclaimed)
}

func WithReconcileEnabled() Func {
	// can be replaced with an annotation level filter with all operators.
	return func(blockDeviceList *apis.BlockDeviceList) *apis.BlockDeviceList {
		filteredBDList := blockDeviceList
		for _, bd := range blockDeviceList.Items {
			if isReconcileEnabled(bd.Annotations) {
				filteredBDList.Items = append(filteredBDList.Items, bd)
			}
		}
		return filteredBDList
	}
}

func WithCapacity(bytes uint64) Func {
	return func(blockDeviceList *apis.BlockDeviceList) *apis.BlockDeviceList {
		filteredBDList := blockDeviceList
		for _, bd := range blockDeviceList.Items {
			if bd.Spec.Capacity.Storage >= bytes {
				filteredBDList.Items = append(filteredBDList.Items, bd)
			}
		}
	}
}

func WithAnnotation(key, operator, value string) Func {
	return func(blockDeviceList *apis.BlockDeviceList) *apis.BlockDeviceList {
		// can add annotation filter logic here
		return blockDeviceList
	}
}

func BaseFilters() []Func {
	return []Func{
		WithReconcileEnabled(),
	}
}

// Filter out blockdevices and get a list of block devices
func Filter(bdList *apis.BlockDeviceList, funcs ...Func) *apis.BlockDeviceList {
	filteredList := bdList
	for _, f := range funcs {
		filteredList = f(filteredList)
	}
	if len(filteredList.Items) == 0 {
		// log no items left after filtering
	}
	return filteredList
}

// Select one block device after applying all the filters
func Select(bdList *apis.BlockDeviceList, funcs ...Func) (*apis.BlockDevice, error) {
	filteredList := Filter(bdList, funcs...)

	if len(filteredList.Items) == 0 {
		// log no items left
		return nil, fmt.Errorf("no item left to select")
	}
	return &filteredList.Items[0], nil
}

func isReconcileEnabled(annotation map[string]string) bool {
	WithReconcileEnabled().Name()
	return annotation["openebs.io/reconcile"] != "false"

}

func (f *Func) Name() string {

}