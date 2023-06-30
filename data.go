package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gopkg.in/guregu/null.v3"
	"math"
	"sort"
)

type Result struct {
	Id          null.String `json:"id"`
	Name        null.String `json:"name"`
	Website     null.String `json:"website"`
	Coordinates null.String `json:"coordinates"`
	Description null.String `json:"description"`
	Rating      null.Float  `json:"rating"`
	Distance    float64     `json:"distance"`
}

func getdata(db *sqlx.DB, long, lat, distance float64, circle bool) ([]Result, error) {
	queryResult := []Result{}
	//if we want a square radius
	if circle != true {
		geom := fmt.Sprintf("SRID=4326;POINT(%f %f)", long, lat)

		err := db.Select(&queryResult, `SELECT *, ST_Distance(coordinates::geography,ST_SetSRID(ST_MakePoint($1, $2),4326)::geography) AS distance
    									FROM my_table
    									WHERE ST_Intersects(coordinates::geography,
    									    ST_Envelope(ST_Buffer((ST_GeomFromEWKT($4))::geography,
    									        $3)::geometry)
    									    );`,
			long, lat, distance, geom)
		if err != nil {
			return nil, err
		}

		//if we want circle
	} else {
		err := db.Select(&queryResult, `SELECT *, ST_Distance(coordinates::geography,ST_SetSRID(ST_MakePoint($1, $2),4326)::geography) AS distance 
							FROM my_table
							WHERE
							ST_DWithin(coordinates::geography,ST_SetSRID(ST_MakePoint($1, $2),4326)::geography,$3)`,
			long, lat, distance)
		if err != nil {
			return nil, err
		}
	}
	//sort slice in ascertaining order by distance
	sort.Slice(queryResult, func(i, j int) bool {
		var a = queryResult[i].Distance
		var b = queryResult[j].Distance
		//but if distance < 50 sort descending by rating
		if math.Abs(a-b) < 50 {
			return queryResult[i].Rating.Float64 > queryResult[j].Rating.Float64
		}
		return queryResult[i].Distance < queryResult[j].Distance

	})

	return queryResult, nil
}
