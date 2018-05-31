package schemas

type ThresholdState struct {
	DatacenterId int `json:"datacenterId,omitempty"`
	DatacenterAbbreviation string `json:"datacenterAbbreviation,omitempty"`
	Warning int `json:"warning"`
	Critical int `json:"critical"`
}

type DatacenterThresholds map[string]ThresholdState

type Thresholds struct {
	Offering Offering `json:"offering"`
	DatacenterThresholds DatacenterThresholds `json:"datacenterThresholds"`
}

type Offering struct {
	Href string `json:"href"`
	OfferingId int `json:"offeringId"`
}

type ThresholdsUpdate struct {
	Offering Offering `json:"offering"`
	DatacenterThresholds []ThresholdState `json:"datacenterThresholds"`
}
