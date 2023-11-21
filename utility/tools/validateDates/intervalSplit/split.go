package intervalSplit

func Split(interval int) (century, decade, year int) {
	century = interval / 100
	decade = (interval - century*100) / 10
	year = interval - century*100 - int(decade)*10
	return century + 1, decade + 1, year
}
