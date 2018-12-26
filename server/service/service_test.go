package service

import (
	"context"
	"errors"
	"testing"

	api "github.com/ivch/testms/api/v1"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestService_Save(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	srv := New(db)

	coords := api.Coordinates{
		Lat: 1.1,
		Lng: 1.1,
	}
	r := api.SaveRequest{
		Id:      "TTEST",
		Name:    "Name",
		City:    "City",
		Country: "Country",
		Coords:  &coords,
	}

	stmt := mock.ExpectPrepare(`INSERT INTO ports \(id, name, city, country, lat, lng\) VALUES \(\?, \?, \?, \?, \?, \?\) ON DUPLICATE KEY UPDATE city = \?, country = \?, lat = \?, lng = \?`)
	stmt.ExpectExec().WithArgs(r.GetId(), r.GetName(), r.GetCity(), r.GetCountry(),
		r.Coords.GetLat(), r.Coords.GetLng(), r.GetCity(), r.GetCountry(),
		r.Coords.GetLat(), r.Coords.GetLng()).WillReturnResult(sqlmock.NewResult(1, 1))

	if _, err := srv.Save(context.Background(), &r); err != nil {
		t.Error(err)
	}

	r.Id = "testtt"
	stmt = mock.ExpectPrepare(`INSERT INTO ports \(id, name, city, country, lat, lng\) VALUES \(\?, \?, \?, \?, \?, \?\) ON DUPLICATE KEY UPDATE city = \?, country = \?, lat = \?, lng = \?`)
	stmt.ExpectExec().WithArgs(r.GetId(), r.GetName(), r.GetCity(), r.GetCountry(),
		r.Coords.GetLat(), r.Coords.GetLng(), r.GetCity(), r.GetCountry(),
		r.Coords.GetLat(), r.Coords.GetLng()).WillReturnError(errors.New("some error"))

	if _, err := srv.Save(context.Background(), &r); err == nil {
		t.Errorf("unexpeceted error %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestService_Get(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	srv := New(db)

	r := api.GetRequest{Id: "TESTT"}

	rows := sqlmock.NewRows([]string{"id", "name", "city", "country", "lat", "lng"}).
		AddRow("TESTT", "name", "city", "country", 1.1, 1.1)
	mock.ExpectQuery(`SELECT id, name, city, country, lat, lng FROM ports WHERE id = \?`).WithArgs(r.GetId()).WillReturnRows(rows)

	if _, err := srv.Get(context.Background(), &r); err != nil {
		t.Error(err)
	}
}
