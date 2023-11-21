package api

import (
	"akim/internal/domain/dto"
	"akim/internal/domain/model"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

type ArtefactUsecase interface {
	FindAllInfo(context.Context, *dto.ArtifactInput) ([]dto.ArtifactOutput, error)
	FindByDecade(context.Context, *dto.ArtifactInput) ([]dto.ArtifactOutput, error)
	FuzzyFindByCentury(context.Context, *dto.ArtifactInput) ([]dto.ArtifactOutput, error)
	FuzzyFindByDecade(context.Context, *dto.ArtifactInput) ([]dto.ArtifactOutput, error)
	FindByCentury(context.Context, *dto.ArtifactInput) ([]dto.ArtifactOutput, error)
}

func FindHandler(uc ArtefactUsecase) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		century := r.FormValue("vek")
		decade := r.FormValue("desyat")
		year := r.FormValue("god")
		interval1 := r.FormValue("interval1")
		interval2 := r.FormValue("interval2")

		if century == "" {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		input := &dto.ArtifactInput{
			Century:         century,
			Decade:          decade,
			Year:            year,
			CenturyInterval: interval1,
			DecadeInterval:  interval2,
		}
		ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancelFunc()

		var results []dto.ArtifactOutput

		if input.CenturyInterval != "" && input.Decade == "" {
			results, err = uc.FuzzyFindByCentury(ctx, input)
		} else if input.Decade != "" && input.DecadeInterval == "" && input.Year == "" {
			results, err = uc.FindByDecade(ctx, input)
		} else if input.DecadeInterval != "" && input.Year == "" {
			results, err = uc.FuzzyFindByDecade(ctx, input)
		} else if input.Year != "" {
			results, err = uc.FindAllInfo(ctx, input)
		} else {
			results, err = uc.FindByCentury(ctx, input)
		}

		if err != nil {
			switch {
			case errors.Is(err, model.ErrNoResults):
				rw.WriteHeader(http.StatusNoContent)
				return
			default:
				log.Print("error FindHandler.")
				http.Error(rw, "internal error", http.StatusInternalServerError)
				return
			}
		}

		rw.Header().Set("Content-Type", "application/json")

		err = json.NewEncoder(rw).Encode(results)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	}
}
