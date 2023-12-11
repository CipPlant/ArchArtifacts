package model

type FuzzyArchitecturalArtifact struct {
	Membership     float64
	Name           string `bson:"name"`
	Description    string `bson:"description"`
	AdditionalInfo string `bson:"addition"`
	IntervalStart  int16  `bson:"start_date"`
	IntervalEnd    int16  `bson:"interval_end"`
}

func NewFuzzyArtefact() *FuzzyArchitecturalArtifact {
	return &FuzzyArchitecturalArtifact{}
}
