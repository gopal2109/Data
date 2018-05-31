package cli

import (
	"fmt"
	"rpis/backend"
	"rpis/api/models"
)

type InitThresholds struct {}

func (it InitThresholds) Usage() {
	fmt.Printf("rpis initthresholds [help]")
}

func (it InitThresholds) Run(args ...string) {
	
	offering := models.Offering{
		Href: "product/catalog/offering/111",
		OfferingId: 111,
	}
	datacenterthresholds := make(models.DatacenterThresholds, 0)
	datacenterthresholds["DFW3"] = models.ThresholdState{111, "DFW3", 60, 50}
	datacenterthresholds["IAD3"] = models.ThresholdState{111, "DFW3", 60, 50}
	datacenterthresholds["LON5"] = models.ThresholdState{111, "DFW3", 60, 50}

	t := models.Thresholds{DatacenterThresholds:datacenterthresholds, Offering:offering}
	_ = backend.GetDB().C("Thresholds").Insert(&t)

	
	offering = models.Offering{
		Href: "product/catalog/offering/222",
		OfferingId: 222,
	}
	datacenterthresholds = make(models.DatacenterThresholds, 0)
	datacenterthresholds["DFW3"] = models.ThresholdState{222, "DFW3",20, 10}
	datacenterthresholds["IAD3"] = models.ThresholdState{222, "DFW3", 100, 50}
	datacenterthresholds["LON5"] = models.ThresholdState{222, "DFW3", 60, 50}

	t = models.Thresholds{DatacenterThresholds:datacenterthresholds, Offering:offering}
		
	_ = backend.GetDB().C("Thresholds").Insert(&t)

	offering = models.Offering{
		Href: "product/catalog/offering/333",
		OfferingId: 333,
	}
	datacenterthresholds = make(models.DatacenterThresholds, 0)
	datacenterthresholds["DFW3"] = models.ThresholdState{333, "DFW3", 200, 120}
	datacenterthresholds["IAD3"] = models.ThresholdState{333, "DFW3", 15, 5}
	datacenterthresholds["LON5"] = models.ThresholdState{333, "DFW3", 60, 50}

	t = models.Thresholds{DatacenterThresholds:datacenterthresholds, Offering:offering}
		
	_ = backend.GetDB().C("Thresholds").Insert(&t)
}
