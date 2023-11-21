package dto

type FuzzyArtifactInput struct {
	IntervalStart int `json:"intervalStart"`
	IntervalEnd   int `json:"intervalEnd"`
}

type FuzzyArtefactOutput struct {
	BuildingName        string  `json:"buildingName"`
	ImagePath           string  `json:"imagePath"`
	BuildingDescription string  `json:"buildingDescription"`
	BuildingAge         string  `json:"buildingAge"`
	Membership          float64 `json:"-"`
}
