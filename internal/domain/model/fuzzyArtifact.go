package model

type FuzzyArchitecturalArtifact struct {
	Name           string `bson:"name"`
	IntervalStart  int    `bson:"start_date"`
	IntervalEnd    int    `bson:"interval_end"`
	Description    string `bson:"description"`
	AdditionalInfo string `bson:"addition"`
	Membership     float64
}

func NewFuzzyArtefact() *FuzzyArchitecturalArtifact {
	return &FuzzyArchitecturalArtifact{}
}
