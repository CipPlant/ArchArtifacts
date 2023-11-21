package api

import (
	"akim/internal/domain/dto"
	"akim/internal/domain/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"
)

type FuzzyUseCase interface {
	FindByInterval(ctx context.Context, input *dto.FuzzyArtifactInput) ([]dto.FuzzyArtefactOutput, error)
	FuzzySqlArtifactFind(ctx context.Context, input *dto.FuzzyArtifactInput) ([]dto.FuzzyArtefactOutput, error)
}

func FuzzyHandler(us FuzzyUseCase) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
		}
		input := &dto.FuzzyArtifactInput{}
		startStr := r.FormValue("startValue")
		intValStart, err := strconv.Atoi(startStr)
		if err != nil {
			http.Error(rw,
				"Ошибка в запросе. Неправильно заданы входные значения.",
				http.StatusBadRequest)
			return
		}
		input.IntervalStart = intValStart
		endStr := r.FormValue("endValue")
		intValEnd, err := strconv.Atoi(endStr)
		if err != nil {
			http.Error(rw,
				"Ошибка в запросе. Неправильно заданы входные значения.",
				http.StatusBadRequest)
			return
		}
		input.IntervalEnd = intValEnd

		if input.IntervalStart > input.IntervalEnd {
			http.Error(rw, fmt.Sprintf(
				"Ошибка в запросе. Начальное значение больше конечного значения"),
				http.StatusBadRequest)
			return
		} else if input.IntervalEnd > time.Now().Year() {
			http.Error(rw, fmt.Sprintf(
				"Ошибка в запросе. Конечное значение года не может быть больше %d", time.Now().Year()),
				http.StatusBadRequest)
			return
		}

		ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancelFunc()

		sqlResult, sqlErr := us.FuzzySqlArtifactFind(ctx, input)
		mongoResult, mongoErr := us.FindByInterval(ctx, input)

		if sqlErr != nil || mongoErr != nil {
			switch {
			case errors.Is(sqlErr, model.ErrNoResults) && errors.Is(mongoErr, model.ErrNoResults):
				rw.WriteHeader(http.StatusNoContent)
				return
			case sqlErr == nil && errors.Is(mongoErr, model.ErrNoResults) || mongoErr == nil && errors.Is(sqlErr, model.ErrNoResults):
			default:
				log.Print(err, "h.uc.FindByInterval")
				http.Error(rw, "internal error", http.StatusInternalServerError)
				return
			}
		}

		for _, v := range sqlResult {
			mongoResult = append(mongoResult, v)
		}

		sort.Slice(mongoResult, func(i, j int) bool {
			return mongoResult[i].Membership > mongoResult[j].Membership
		})

		for _, v := range mongoResult {
			fmt.Println(v.BuildingName, v.Membership)
		}

		rw.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(rw).Encode(mongoResult)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	}
}
