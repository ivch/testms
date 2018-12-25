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
		name, err := dec.Token()
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
				Name:    fmt.Sprint(name),
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

//dec := json.NewDecoder(f)
//
//t, err := dec.Token()
//if err != nil {
//	log.Fatal(err)
//}
//
//fmt.Printf("%T: %v\n", t, t)
//
//i := 0
//for {
//	fmt.Println(">>>>>>>")
//	key, err := dec.Token()
//	if err == io.EOF {
//		break
//	}
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("%T: %v\n", key, key)
//
//	if dec.More() {
//		fmt.Println("=========")
//		i++
//		var m Port
//		//var k string
//		//var v Port
//		////	//	var m map[string]Port
//		////	//	// decode an array value (Message)
//		err := dec.Decode(&m)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		fmt.Println(m)
//
//		//req := api.SaveRequest{
//		//	Name:    fmt.Sprint(key),
//		//	City:    m.City,
//		//	Country: m.Country,
//		//	Coords:  m.Coordinates,
//		//}
//		//
//		//res, err := c.Save(ctx, &req)
//		//if err != nil {
//		//	log.Fatalf("Save failed: %v", err)
//		//}
//		//log.Printf("Create result: <%+v>\n\n", res)
//
//		//err = dec.Decode(&v)
//		//if err != nil {
//		//	log.Fatal(err)
//		//}
//		//
//		//fmt.Println(k, v)
//		//fmt.Printf(" (more)")
//	}
//
//	if i > 1 {
//		break
//	}
//	fmt.Printf("\n")
//}

// while the array contains values
//PrintMemUsage()
//for dec.More() {
//	var m interface{}
//	//	var m map[string]Port
//	//	// decode an array value (Message)
//	err := dec.Decode(&m)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println(m)
//	//	PrintMemUsage()
//	t, _ = dec.Token()
//	fmt.Printf("%T: %v\n", t, t)
//	time.Sleep(500 * time.Millisecond)
//}
//
//// read closing bracket
//t, err = dec.Token()
//if err != nil {
//	log.Fatal(err)
//}
//PrintMemUsage()
//fmt.Printf("%T: %v\n", t, t)
