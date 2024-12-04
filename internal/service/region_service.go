pakcage service

import (
	"encoding/json"
	"evermos-project/internal/domain"
	"fmt"
	"net/http"
)

const baseURL = "https://emsifa.github.io/api-wilayah-indonesia/api"

type RegionService struct {}

func NewRegionService() *RegionService {
	return &RegionService{}
}

func (s *RegionService) GetProvinces() ([]domain.Province, error) {
	resp, err := http.Get(fmt.Sprintf("%s/provinces.json", baseURL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var provinces []domain.Province
	if err := json.NewDecoder(resp.Body).Decode(&provinces); err != nil {
		return nil, err
	}

	return provinces, nil
}

func (s *RegionService) GetCities(provinceID string) ([]domain.City, error) {
	resp, err := http.Get(fmt.Sprintf("%s/regencies/%s.json", baseURL, provinceID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var cities []domain.City
	if err := json.NewDecoder(resp.Body).Decode(&cities); err != nil {
		return nil, err
	}

	return cities, nil
}

func (s *RegionService) ValidateProvinceID(provinceID string) bool {
    provinces, err := s.GetProvinces()
    if err != nil {
        return false
    }

    for _, p := range provinces {
        if p.ID == provinceID {
            return true
        }
    }
    return false
}

func (s *RegionService) ValidateCityID(provinceID, cityID string) bool {
    cities, err := s.GetCities(provinceID)
    if err != nil {
        return false
    }

    for _, c := range cities {
        if c.ID == cityID {
            return true
        }
    }
    return false
}