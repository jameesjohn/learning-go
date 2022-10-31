package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/emicklei/go-restful"
	"github.com/jameesjohn/railAPI/dbutils"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// TrainResource is the model for holding rail information
type TrainResource struct {
	ID              int
	DriverName      string
	OperatingStatus bool
}

// StationResource is the model for station information (locations)
type StationResource struct {
	ID          int
	Name        string
	OpeningTime time.Time
	ClosingTime time.Time
}

// ScheduleResource links both trains and stations
type ScheduleResource struct {
	ID          int
	TrainId     int
	StationId   int
	ArrivalTime time.Time
}

func (t *TrainResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/v1/trains").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/{train-id}").To((t.getTrain)))
	ws.Route(ws.POST("").To((t.createTrain)))
}

func main() {
	// Connect to Database.
	db, err := sql.Open("sqlite3", "./railapi.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}
	// Create tables
	dbutils.Initialize(db)
}
