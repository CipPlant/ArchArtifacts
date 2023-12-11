package dto

type FuzzyArtifactInput struct {
	IntervalStart int16 `json:"intervalStart"`
	IntervalEnd   int16 `json:"intervalEnd"`
}

type FuzzyArtefactOutput struct {
	BuildingName        string  `json:"buildingName"`
	ImagePath           string  `json:"imagePath"`
	BuildingDescription string  `json:"buildingDescription"`
	BuildingAge         string  `json:"buildingAge"`
	Membership          float64 `json:"-"`
}
