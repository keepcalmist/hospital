package doctors

import (
	"context"
)

type Doctor struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	LastName   string `json:"last_name"`
	HospitalId int    `json:"hospital_id"`
}

type Repository interface {
	SelectAll(ctx context.Context) ([]Doctor, error)
}
