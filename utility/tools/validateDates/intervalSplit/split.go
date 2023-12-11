package intervalSplit

func Split(interval int16) (century, decade, year int16) {
	century = interval / 100
	decade = (interval - century*100) / 10
	year = interval - century*100 - decade*10
	return century + 1, decade + 1, year
}
