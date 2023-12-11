package usecases

import (
	"akim/internal/controller/http/tools"
	"akim/internal/domain/dto"
	"akim/internal/domain/model"
	"akim/utility/tools/builder"
	"akim/utility/tools/fuzzyLogic"
	"akim/utility/tools/validateDates/intervalSplit"
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
)

type NoSQLDatabase interface {
	FindByInterval(ctx context.Context, artifact *model.FuzzyArchitecturalArtifact) ([]model.FuzzyArchitecturalArtifact, error)
}

type SqlDataBaseForInterval interface {
	ForFuzzyArtifactFind(
		sqlVekStart, sqlDecadeStart, sqlYearStart, sqlVekEnd, sqlDecadeEnd, sqlYearEnd int16) ([]model.Artifact, error)
}

type FuzzyArtifactUseCase struct {
	repo    NoSQLDatabase
	sqlRepo SqlDataBaseForInterval
}

func NewFuzzyArtifactUseCase(nosqlRepository NoSQLDatabase, sqlRepo SqlDataBaseForInterval) *FuzzyArtifactUseCase {
	return &FuzzyArtifactUseCase{repo: nosqlRepository, sqlRepo: sqlRepo}
}

func (f FuzzyArtifactUseCase) FuzzySqlArtifactFind(wg *sync.WaitGroup, ctx context.Context, input *dto.FuzzyArtifactInput) ([]dto.FuzzyArtefactOutput, error) {
	defer wg.Done()

	fuzzyArtefact := model.NewFuzzyArtefact()
	fuzzyArtefact.IntervalStart = input.IntervalStart
	fuzzyArtefact.IntervalEnd = input.IntervalEnd

	sqlVekStart, sqlDecadeStart, sqlYearStart := intervalSplit.Split(fuzzyArtefact.IntervalStart)
	sqlVekEnd, sqlDecadeEnd, sqlYearEnd := intervalSplit.Split(fuzzyArtefact.IntervalEnd)

	sqlResults, err := f.sqlRepo.ForFuzzyArtifactFind(sqlVekStart, sqlDecadeStart, sqlYearStart, sqlVekEnd, sqlDecadeEnd, sqlYearEnd)
	if err != nil {
		return nil, err
	}

	var results []dto.FuzzyArtefactOutput

	for _, v := range sqlResults {
		buildingAge, err := tools.DataMaker(v.Century, v.Decade, v.Age)
		dtoOutputAge := strconv.Itoa(buildingAge)

		photoPath, err := tools.LoadPhoto(v.Name)
		if err != nil {
			log.Println("trouble with photo:", err)
		}
		results = append(results, dto.FuzzyArtefactOutput{
			BuildingName:        v.Name,
			ImagePath:           photoPath,
			BuildingDescription: v.Description,
			BuildingAge:         fmt.Sprintf("%s Год", dtoOutputAge),
			Membership:          1 / (float64(fuzzyArtefact.IntervalEnd) - float64(fuzzyArtefact.IntervalStart)),
		})
	}
	return results, nil
}

func (f FuzzyArtifactUseCase) FindByInterval(wg *sync.WaitGroup, ctx context.Context, input *dto.FuzzyArtifactInput) ([]dto.FuzzyArtefactOutput, error) {
	defer wg.Done()

	fuzzyArtefact := model.NewFuzzyArtefact()
	fuzzyArtefact.IntervalStart = input.IntervalStart
	fuzzyArtefact.IntervalEnd = input.IntervalEnd

	results, err := f.repo.FindByInterval(ctx, fuzzyArtefact)
	if err != nil {
		return nil, err
	}

	for k := range results {
		results[k].Membership = fuzzyLogic.CalculateMembership(&results[k], input.IntervalStart, input.IntervalEnd)
	}

	var res []dto.FuzzyArtefactOutput

	for _, v := range results {
		photo, err := tools.LoadPhoto(v.Name)
		if err != nil {
			log.Println("error with LoadPhoto:", err)
		}

		res = append(res, dto.FuzzyArtefactOutput{
			BuildingName:        v.Name,
			ImagePath:           photo,
			BuildingDescription: v.Description,
			BuildingAge:         builder.BuildingAge(v.AdditionalInfo, v.IntervalStart, v.IntervalEnd),
			Membership:          v.Membership,
		})
	}

	return res, nil
}
