package repository

import (
	"akim/internal/config"
	"akim/internal/domain/model"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(config *config.Config) *MySQLRepository {

	cfg := mysql.Config{
		User:   config.UserName,
		Passwd: config.Password,
		DBName: config.DbName,
		Addr:   fmt.Sprintf("%s:%d", config.MySQL.Host, config.MySQL.Port),
	}
	open, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal("error with open to db", err)
	}
	err = open.Ping()
	if err != nil {
		log.Fatal("error with ping to db", err)
	}
	return &MySQLRepository{
		db: open,
	}
}

func (r *MySQLRepository) FindByCentury(century string) ([]model.Artifact, error) {
	prepare, err := r.db.Prepare(
		"SELECT Название, Век, Десятилетие, Год, Описание" +
			" FROM Buildings" +
			" WHERE Век = ?")

	if err != nil {
		return nil, errors.New(fmt.Sprintf("error with query in mysql: %v", err))
	}

	rows, err := prepare.Query(century)
	defer rows.Close()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error with query in mysql: %v", err))
	}
	var results []model.Artifact
	for rows.Next() {
		var result model.Artifact
		err = rows.Scan(&result.Name, &result.Century, &result.Decade, &result.Age, &result.Description)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error with marshalling to struct: %v", err))
		}
		results = append(results, result)
	}
	if len(results) == 0 {
		return nil, model.ErrNoResults
	}
	return results, nil
}

func (r *MySQLRepository) FindAllInfo(century, decade, age string) ([]model.Artifact, error) {

	prepare, err := r.db.Prepare(
		"SELECT Название, Век, Десятилетие, Год, Описание" +
			" FROM Buildings" +
			" WHERE Век = ?" +
			" AND Десятилетие = ?" +
			" AND Год = ?")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error with query in mysql: %v", err))
	}

	rows, err := prepare.Query(century, decade, age)
	defer rows.Close()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error with query in mysql: %v", err))
	}

	var results []model.Artifact
	for rows.Next() {
		var result model.Artifact
		err = rows.Scan(&result.Name, &result.Century, &result.Decade, &result.Age, &result.Description)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error with marshalling to struct: %v", err))
		}
		results = append(results, result)
	}
	if len(results) == 0 {
		return nil, model.ErrNoResults
	}
	return results, nil
}

func (r *MySQLRepository) FindByDecade(century, decade string) ([]model.Artifact, error) {

	prepare, err := r.db.Prepare(
		"SELECT Название, Век, Десятилетие, Год, Описание" +
			" FROM Buildings" +
			" WHERE Век = ?" +
			" AND Десятилетие = ?")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error with query in mysql: %v", err))
	}
	rows, err := prepare.Query(century, decade)
	defer rows.Close()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error with query in mysql: %v", err))
	}

	var results []model.Artifact
	for rows.Next() {

		var result model.Artifact

		err = rows.Scan(&result.Name, &result.Century, &result.Decade, &result.Age, &result.Description)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error with marshalling to struct: %v", err))
		}
		results = append(results, result)
	}
	if len(results) == 0 {
		return nil, model.ErrNoResults
	}
	return results, nil
}

func (r *MySQLRepository) FuzzyFindByCentury(century string, start, end int) ([]model.Artifact, error) {

	prepare, err := r.db.Prepare(
		"SELECT Название, Век, Десятилетие, Год, Описание" +
			" FROM Buildings" +
			" WHERE Век = ?" +
			" AND Десятилетие BETWEEN ? AND ?")

	if err != nil {
		return nil, errors.New(fmt.Sprintf("error with query in mysql: %v", err))
	}

	rows, err := prepare.Query(century, start, end)
	defer rows.Close()

	if err != nil {
		return nil, errors.New(fmt.Sprintf("error with query in mysql: %v", err))
	}
	var results []model.Artifact
	for rows.Next() {
		var result model.Artifact

		err = rows.Scan(&result.Name, &result.Century, &result.Decade, &result.Age, &result.Description)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error with marshalling to struct: %v", err))
		}
		results = append(results, result)
	}
	if len(results) == 0 {
		return nil, model.ErrNoResults
	}
	return results, nil
}

func (r *MySQLRepository) FuzzyFindByDecade(century, decade string, startYear, endYear int) ([]model.Artifact, error) {

	prepare, err := r.db.Prepare(
		"SELECT Название, Век, Десятилетие, Год, Описание" +
			" FROM Buildings" +
			" WHERE Век = ?" +
			" AND Десятилетие = ?" +
			" AND Год BETWEEN ? AND ?")

	if err != nil {
		return nil, errors.New(fmt.Sprintf("error with query in mysql: %v", err))
	}

	rows, err := prepare.Query(century, decade, startYear, endYear)
	defer rows.Close()

	if err != nil {
		return nil, errors.New(fmt.Sprintf("error with query in mysql: %v", err))
	}
	var results []model.Artifact
	for rows.Next() {
		var result model.Artifact

		err = rows.Scan(&result.Name, &result.Century, &result.Decade, &result.Age, &result.Description)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error with marshalling to struct: %v", err))
		}
		results = append(results, result)
	}
	if len(results) == 0 {
		return nil, model.ErrNoResults
	}
	return results, nil
}

func (r *MySQLRepository) ForFuzzyArtifactFind(
	sqlVekStart, sqlDecadeStart, sqlYearStart, sqlVekEnd, sqlDecadeEnd, sqlYearEnd int16) ([]model.Artifact, error) {
	prepare, err := r.db.Prepare(
		"SELECT Название, Век, Десятилетие, Год, Описание" +
			" FROM Buildings" +
			" WHERE (Век, Десятилетие, Год) >= (?, ?, ?)" +
			" AND " +
			" (Век, Десятилетие, Год) <= (?, ?, ?)")

	if err != nil {
		return nil, errors.New(fmt.Sprintf("error with query in mysql: %v", err))
	}
	rows, err := prepare.Query(sqlVekStart, sqlDecadeStart, sqlYearStart, sqlVekEnd, sqlDecadeEnd, sqlYearEnd)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error with query in mysql: %v", err))
	}
	var results []model.Artifact

	for rows.Next() {
		var result model.Artifact
		err = rows.Scan(&result.Name, &result.Century, &result.Decade, &result.Age, &result.Description)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error with marshalling to struct: %v", err))
		}
		results = append(results, result)
	}

	if len(results) == 0 {
		return nil, model.ErrNoResults
	}
	return results, nil
}
