package service

import (
	"context"
	"database/sql"

	api "github.com/ivch/testms/api/v1"
)

type service struct {
	db *sql.DB
}

//New creates PortDomainServiceServer instance
func New(db *sql.DB) api.PortDomainServiceServer {
	return &service{
		db: db,
	}
}

//Save receives port data and stores it to database
//in case if port is already in db - updates
func (s *service) Save(ctx context.Context, r *api.SaveRequest) (*api.SaveResponse, error) {
	stmt, err := s.db.Prepare("INSERT INTO ports (id, name, city, country, lat, lng) " +
		"VALUES (?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE " +
		"city = ?, country = ?, lat = ?, lng = ?")
	if err != nil {
		return nil, err
	}

	if _, err := stmt.ExecContext(ctx, r.GetId(), r.GetName(), r.GetCity(), r.GetCountry(),
		r.Coords.GetLat(), r.Coords.GetLng(), r.GetCity(), r.GetCountry(),
		r.Coords.GetLat(), r.Coords.GetLng()); err != nil {
		return nil, err
	}

	return &api.SaveResponse{Name: r.Name}, nil
}

//Get retrieves data about port
func (s *service) Get(ctx context.Context, r *api.GetRequest) (*api.GetResponse, error) {
	var res api.GetResponse
	res.Coords = &api.Coordinates{}

	err := s.db.QueryRowContext(ctx, "SELECT id, name, city, country, lat, lng FROM ports WHERE id = ?", r.GetId()).
		Scan(&res.Id, &res.Name, &res.City, &res.Country, &res.Coords.Lat, &res.Coords.Lng)

	return &res, err
}
