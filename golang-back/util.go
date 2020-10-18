package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jonas-p/go-shp"
	geojson "github.com/paulmach/go.geojson"
)

func UploadShp() {
	shape, err := shp.Open("./data/fishnet/fishnet.shp")
	if err != nil {
		log.Fatal(err)
	}
	defer shape.Close()

	// upload (DBF)
	fields := shape.Fields()

	// loop get data for shp
	for shape.Next() {
		n, p := shape.Shape()
		// get attributes
		for k, f := range fields {
			val := shape.ReadAttribute(n, k)
			if f.String() == "cell_id" {
				fmt.Println(val, p.BBox()) // Poligon cell_id
			}
		}
	}

}

func UploadCsv() {
	files, err := filepath.Glob("./data/point/*.csv")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		update(f)
	}
}

func update(filen string) {
	file, err := os.Open(filen)
	if err != nil {
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	// reader.FieldsPerRecord = 3
	reader.Comment = '#'

	for {
		record, e := reader.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		if record[2] != "" {
			// fmt.Println(record[0], record[2]) //tsp
			fmt.Println(record[2])
		}
	}
}

func convert_geojson(parr []*PointEx) []byte {
	fc := geojson.NewFeatureCollection()
	for _, v := range parr {
		fc.AddFeature(geojson.NewPointFeature([]float64{v.Px, v.Py}))
	}

	rawJSON, err := fc.MarshalJSON()
	if err != nil {
		fmt.Printf("error: %v", err)
		return []byte("none")
	}
	return rawJSON
}
