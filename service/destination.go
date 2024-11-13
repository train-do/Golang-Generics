package service

import (
	"fmt"
	"math"

	"github.com/train-do/Golang-Generics/model"
	"github.com/train-do/Golang-Generics/repository"
)

type ServiceDestination struct {
	Repo *repository.RepoDestination
}

func NewServiceDestination(repo *repository.RepoDestination) *ServiceDestination {
	return &ServiceDestination{repo}
}
func (s *ServiceDestination) GetAll(qp model.QueryParams) (model.Response, error) {
	destinations, totalItem, err := s.Repo.FindAll(qp)
	if err != nil {
		return model.Response{}, err
	}
	if qp.Page == 0 {
		qp.Page = 1
	}
	response := model.Response{
		Page:       qp.Page,
		Limit:      6,
		TotalItem:  totalItem,
		TotalPages: int(math.Ceil(float64(totalItem) / float64(6))),
		Data:       destinations,
	}
	fmt.Println(destinations, totalItem)
	return response, nil
}
