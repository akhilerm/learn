package NDMRefactTest

import (
	"fmt"
	"github.com/akhilerm/TEst/NDMRefactTest/client"
	"github.com/akhilerm/TEst/NDMRefactTest/filter"
	"github.com/openebs/node-disk-manager/db/kubernetes"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
)

func main() {
	c, err := client.New()
	if err != nil {
		// log or panic
	}

	requirements := make(labels.Requirements, 0)
	requirements = append(requirements,
		*GenerateLabelRequirement(kubernetes.KubernetesHostNameLabel, "hostname"),
		*GenerateLabelRequirement("openebs.io/managed", "true"),
	)

	selector := labels.NewSelector()
	// selector := labels.Everything()
	// selector := labels.Nothing()

	selector = selector.Add(requirements...)

	// various cases

	// 1. Listing with a selector having hostname and some random label
	// lists as the block device api resource itself
	bdAPIList, err := c.ListBlockDevicesAPI(selector)

	// a list of filters to be applied
	filtersList := []filter.Func{
		filter.WithBlockDeviceUnclaimed(),
	}

	// another list of filters to be applied
	selectionList := []filter.Func{
		filter.WithCapacity(10737418240),
	}

	// filter with the filters list
	filteredList := filter.Filter(bdAPIList, filtersList...)

	// select a single device with another set of filters
	selectedDevice, err := filter.Select(filteredList, selectionList...)
	if err != nil {
		// lof error
	}

	fmt.Println("Selected Device", selectedDevice)

	// lists as the internal block device struct used by NDM.
	bdList, err := c.ListBlockDevices(selector, filter.BaseFilters()...)
	if err != nil {
		//log
	}
	fmt.Println("BlockDevice list", bdList)

}

func GenerateLabelRequirement(key, value string) *labels.Requirement {
	req, err := labels.NewRequirement(
		key,
		selection.Equals,
		[]string{value})
	if err != nil {
		// log error
		// actually this error should never happen.
		// because error occurs for Equals only when there is no value
	}
	return req
}
