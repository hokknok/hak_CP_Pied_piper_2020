package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Point struct {
	X float64
	Y float64
}

type PointEx struct {
	ID int
	Px float64
	Py float64
}

type Poligon struct {
	ID    int
	Erlid string
	P1x   float64
	P1y   float64
	P2x   float64
	P2y   float64
	P3x   float64
	P3y   float64
	P4x   float64
	P4y   float64
}

type MapPtoP struct {
	Polymap  Poligon
	datePmap dateP
}

type dateP struct {
	Id                      int
	Zid                     string
	Ts                      string
	Customers_cnt_long_sum  int
	Customers_cnt_work_sum  int
	Customers_cnt_loc_sum   int
	Customers_cnt_total_sum int
}

var db *sql.DB

func init() {
	var err error
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	db, err = sql.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

}

func GetDB() *sql.DB {
	return db
}

func (p Poligon) getPoint() [4]Point {
	pp := [4]Point{
		Point{p.P1x, p.P1y},
		Point{p.P2x, p.P2y},
		Point{p.P3x, p.P3y},
		Point{p.P4x, p.P4y},
	}
	return pp
}

func getAllPoint() []*PointEx {
	rows, err := db.Query("SELECT * FROM PointEx;")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	bks := make([]*PointEx, 0)
	for rows.Next() {
		bk := new(PointEx)
		err := rows.Scan(&bk.ID, &bk.Px, &bk.Py)
		if err != nil {
			log.Fatal(err)
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return bks
}

func (p *Poligon) getAll() []*Poligon {
	rows, err := db.Query("SELECT * FROM Poligon;")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	bks := make([]*Poligon, 0)
	for rows.Next() {
		bk := new(Poligon)
		err := rows.Scan(&bk.ID, &bk.Erlid, &bk.P1x, &bk.P1y, &bk.P2x, &bk.P2y, &bk.P3x, &bk.P3y, &bk.P4x, &bk.P4y)
		if err != nil {
			log.Fatal(err)
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return bks
}
