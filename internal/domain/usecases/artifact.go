package usecases

import (
	"akim/internal/controller/http/tools"
	"akim/internal/domain/dto"
	"akim/internal/domain/model"
	"akim/utility/tools/validateDates/relationalValidation"
	"context"
	"fmt"
	"log"
	"strconv"
)

type SQLDatabase interface {
	FindByCentury(century string) ([]model.Artifact, error)
	FindAllInfo(century, decade, age string) ([]model.Artifact, error)
	FindByDecade(century, decade string) ([]model.Artifact, error)
	FuzzyFindByCentury(century string, start, end int) ([]model.Artifact, error)
	FuzzyFindByDecade(century, decade string, startYear, endYear int) ([]model.Artifact, error)
}

type ArtifactUseCase struct {
	repo SQLDatabase
}

func NewRepository(sqlRepository SQLDatabase) *ArtifactUseCase {
	return &ArtifactUseCase{repo: sqlRepository}
}

func (auc *ArtifactUseCase) FindByDecade(ctx context.Context, input *dto.ArtifactInput) ([]dto.ArtifactOutput, error) {
	// TODO: правильная обработка ошибки

	toModel := model.NewArtefact(input.Century, input.Decade, input.Year)

	artifactResults, err := auc.repo.FindByDecade(toModel.Century, toModel.Decade)
	if err != nil {
		return nil, err
	}
	var results []dto.ArtifactOutput

	for _, v := range artifactResults {
		buildingAge, err := tools.DataMaker(v.Century, v.Decade, v.Age)
		dtoOutputAge := strconv.Itoa(buildingAge)
		photoPath, err := tools.LoadPhoto(v.Name)
		if err != nil {
			log.Print("photo error:", err)
		}

		results = append(results, dto.ArtifactOutput{
			BuildingName:        v.Name,
			ImagePath:           photoPath,
			BuildingDescription: v.Description,
			BuildingAge:         fmt.Sprintf("%s Год", dtoOutputAge),
		})
	}

	return results, err
}

func (auc *ArtifactUseCase) FindAllInfo(ctx context.Context, input *dto.ArtifactInput) ([]dto.ArtifactOutput, error) {
	// TODO: правильная обработка ошибки

	toModel := model.NewArtefact(input.Century, input.Decade, input.Year)

	artifactResults, err := auc.repo.FindAllInfo(toModel.Century, toModel.Decade, toModel.Age)
	if err != nil {
		return nil, err
	}
	var results []dto.ArtifactOutput

	for _, v := range artifactResults {
		buildingAge, err := tools.DataMaker(v.Century, v.Decade, v.Age)
		dtoOutputAge := strconv.Itoa(buildingAge)
		photoPath, err := tools.LoadPhoto(v.Name)
		if err != nil {
			log.Print("photo error:", err)
		}
		results = append(results, dto.ArtifactOutput{
			BuildingName:        v.Name,
			ImagePath:           photoPath,
			BuildingDescription: v.Description,
			BuildingAge:         fmt.Sprintf("%s Год", dtoOutputAge),
		})
	}

	return results, err
}

func (auc *ArtifactUseCase) FuzzyFindByCentury(ctx context.Context, input *dto.ArtifactInput) ([]dto.ArtifactOutput, error) {
	validateCentury, err := relationalValidation.ValidateCentury(input.CenturyInterval)
	if err != nil {
		return nil, err

	}
	start := validateCentury[0]
	end := validateCentury[len(validateCentury)-1]

	toModel := model.NewArtefact(input.Century, input.Decade, input.Year)

	artifactResults, err := auc.repo.FuzzyFindByCentury(toModel.Century, start, end)
	if err != nil {
		return nil, err
	}
	var results []dto.ArtifactOutput

	for _, v := range artifactResults {
		buildingAge, err := tools.DataMaker(v.Century, v.Decade, v.Age)
		dtoOutputAge := strconv.Itoa(buildingAge)
		photoPath, err := tools.LoadPhoto(v.Name)
		if err != nil {
			log.Print("photo error:", err)
		}

		results = append(results, dto.ArtifactOutput{
			BuildingName:        v.Name,
			ImagePath:           photoPath,
			BuildingDescription: v.Description,
			BuildingAge:         fmt.Sprintf("%s Год", dtoOutputAge),
		})
	}

	return results, err
}
func (auc *ArtifactUseCase) FuzzyFindByDecade(ctx context.Context, input *dto.ArtifactInput) ([]dto.ArtifactOutput, error) {

	toModel := model.NewArtefact(input.Century, input.Decade, input.Year)

	validateDecade, err := relationalValidation.ValidateDecade(input.DecadeInterval)
	if err != nil {
		return nil, err
	}

	startYear := validateDecade[0]
	endYear := validateDecade[len(validateDecade)-1]

	artifactResults, err := auc.repo.FuzzyFindByDecade(toModel.Century, toModel.Decade, startYear, endYear)
	if err != nil {
		return nil, err
	}

	var results []dto.ArtifactOutput

	for _, v := range artifactResults {
		buildingAge, err := tools.DataMaker(v.Century, v.Decade, v.Age)
		dtoOutputAge := strconv.Itoa(buildingAge)
		photoPath, err := tools.LoadPhoto(v.Name)
		if err != nil {
			log.Print("photo error:", err)
		}
		results = append(results, dto.ArtifactOutput{
			BuildingName:        v.Name,
			ImagePath:           photoPath,
			BuildingDescription: v.Description,
			BuildingAge:         fmt.Sprintf("%s Год", dtoOutputAge),
		})
	}
	return results, err
}
func (auc *ArtifactUseCase) FindByCentury(ctx context.Context, input *dto.ArtifactInput) ([]dto.ArtifactOutput, error) {

	toModel := model.NewArtefact(input.Century, input.Decade, input.Year)

	artifactResults, err := auc.repo.FindByCentury(toModel.Century)
	if err != nil {
		return nil, err
	}
	var results []dto.ArtifactOutput

	for _, v := range artifactResults {
		buildingAge, err := tools.DataMaker(v.Century, v.Decade, v.Age)
		dtoOutputAge := strconv.Itoa(buildingAge)
		photoPath, err := tools.LoadPhoto(v.Name)
		if err != nil {
			log.Print("photo error:", err)
		}
		results = append(results, dto.ArtifactOutput{
			BuildingName:        v.Name,
			ImagePath:           photoPath,
			BuildingDescription: v.Description,
			BuildingAge:         fmt.Sprintf("%s Год", dtoOutputAge),
		})
	}
	return results, err
}
