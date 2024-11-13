package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/train-do/Golang-Generics/model"
)

func GenerateQuery(qp model.QueryParams, args *[]any) (string, string, string) {
	page := ``
	sort := ``
	search := `where 1=1`
	count := 0
	if qp.SearchPlace != "" {
		count++
		search += fmt.Sprintf(` and d."name" ilike $%d`, count)
		*args = append(*args, qp.SearchPlace)
	}
	if qp.SearchDate != "" {
		count++
		search += fmt.Sprintf(` and s."date"=$%d`, count)
		*args = append(*args, qp.SearchDate)
	}
	if qp.SearchPrice != 0 {
		count++
		search += fmt.Sprintf(` and ds."price" <= $%d`, count)
		*args = append(*args, qp.SearchPrice)
	}
	if qp.SortDate {
		sort += `order by s."date"`
	} else if qp.SortPrice != "" {
		if qp.SortPrice == "asc" {
			sort += `order by ds."price" asc`
		} else if qp.SortPrice == "desc" {
			sort += `order by ds."price" desc`
		}
	} else if qp.SortName {
		sort += `order by d."name"`
	}
	if qp.Page == 0 || qp.Page == 1 {
		page = `limit 6 offset 0;`
	} else {
		page = fmt.Sprintf(`limit 6 offset %d;`, ((qp.Page - 1) * 6))
	}
	return page, sort, search
}

func SetResponse(w http.ResponseWriter, data model.Response, statusCode int, message string) model.Response {
	w.WriteHeader(statusCode)
	data.StatusCode = statusCode
	data.Message = message
	return data
}
func ToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return num
}
func ToBool(str string) bool {
	return strings.ToLower(str) == "true"
}
