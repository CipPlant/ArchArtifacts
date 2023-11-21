package tools

import (
	"akim/internal/domain/model"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strconv"
)

func DataMaker(century, decade, age string) (int, error) {
	centuryInt, err := strconv.Atoi(century)
	if err != nil {
		return 0, fmt.Errorf("wrong century")
	}
	centuryInt = centuryInt*100 - 100

	decadeInt, err := strconv.Atoi(decade)
	if err != nil {
		return 0, fmt.Errorf("wrong decade")
	}
	decadeInt = decadeInt*10 - 10
	if decadeInt < 0 {
		decadeInt = 0
	}
	ageInt, err := strconv.Atoi(age)
	if err != nil {
		return 0, fmt.Errorf("wrong age")
	}
	return centuryInt + decadeInt + ageInt, nil
}

func LoadPhoto(s string) (string, error) {
	photoPath := s + ".jpg"
	dirPath := "../static/photos/" + photoPath

	image1Data, err := ioutil.ReadFile(dirPath)
	if err != nil {
		return "", model.ErrNoSuchPhoto
	}
	image1Base64 := base64.StdEncoding.EncodeToString(image1Data)
	return image1Base64, nil
}
