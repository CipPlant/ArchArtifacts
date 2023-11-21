package dto

type ArtifactInput struct {
	Century         string `json:"vek"`
	Decade          string `json:"desyat"`
	Year            string `json:"god"`
	CenturyInterval string `json:"interval1"`
	DecadeInterval  string `json:"interval2"`
}

type ArtifactOutput struct {
	BuildingName        string `json:"buildingName"`
	ImagePath           string `json:"imagePath"`
	BuildingDescription string `json:"buildingDescription"`
	BuildingAge         string `json:"buildingAge"`
}
