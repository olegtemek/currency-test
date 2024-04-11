package service

import (
	"encoding/xml"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/olegtemek/currency-task/internal/model"
)

type repository interface {
	Save(currencies []*model.Currency, date string)
	Get(date string, code string) ([]*model.Currency, error)
}

type Service struct {
	log  *slog.Logger
	repo repository
}

var (
	BASE_API_URL = "https://nationalbank.kz/rss/get_rates.cfm"
)

func New(log *slog.Logger, repo repository) *Service {
	return &Service{
		log:  log,
		repo: repo,
	}
}

func (s *Service) Save(date string) error {
	url := fmt.Sprintf("%s?fdate=%s", BASE_API_URL, date)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("cannot send http request: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("cannot read http response: %s", err)
	}

	rates := &model.Rates{}
	err = xml.Unmarshal(body, &rates)
	if err != nil {
		return fmt.Errorf("error parsing XML: %s", err)
	}

	currencies := make([]*model.Currency, len(rates.Items))
	for i, item := range rates.Items {
		currencies[i] = &item
	}

	go s.repo.Save(currencies, date)

	return nil
}

func (s *Service) Get(date string, code string) ([]*model.Currency, error) {

	currencies, err := s.repo.Get(date, code)
	if err != nil {
		return nil, err
	}

	return currencies, nil
}
