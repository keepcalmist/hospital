package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/keepcalmist/hospital/pkg/doctors"
)

type DoctorRepo struct {
	db *sqlx.DB
}

func NewDoctorRepo(db *sqlx.DB) doctors.Repository {
	return &DoctorRepo{db: db}
}

func (r *DoctorRepo) SelectAll(ctx context.Context) ([]doctors.Doctor, error) {
	dcts := []doctors.Doctor{}
	err := r.db.SelectContext(ctx, dcts, "SELECT * FROM doctors order by id")
	if err != nil {
		return nil, err
	}
	return dcts, nil
}
