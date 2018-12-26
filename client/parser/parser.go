package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	api "github.com/ivch/testms/api/v1"
)

type parser struct {
	file *os.File
	data chan<- api.SaveRequest
}

type port struct {
	Name    string
	City    string
	Country string
	Coord   []float32 `json:"coordinates"`
}

//Start start parsing of given file and send parsed data to channel shared with service
func Start(filepath string, ch chan<- api.SaveRequest) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}

	p := parser{
		file: f,
		data: ch,
	}

	go p.run()

	return nil
}

func (p *parser) run() {
	defer close(p.data)

	dec := json.NewDecoder(p.file)
	_, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	for {
		id, err := dec.Token()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if dec.More() {
			var entity port
			err := dec.Decode(&entity)
			if err != nil {
				log.Fatal(err)
			}

			req := api.SaveRequest{
				Id:      fmt.Sprint(id),
				Name:    entity.Name,
				City:    entity.City,
				Country: entity.Country,
			}

			if len(entity.Coord) == 2 {
				coord := api.Coordinates{
					Lat: entity.Coord[0],
					Lng: entity.Coord[1],
				}
				req.Coords = &coord
			}

			p.data <- req
		}
	}
}
