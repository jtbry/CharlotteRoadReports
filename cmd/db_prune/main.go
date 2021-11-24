package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"time"

	"cloud.google.com/go/storage"
	"github.com/joho/godotenv"
	"github.com/jtbry/CharlotteRoadReports/pkg/api"
	"google.golang.org/api/option"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Create pgsql storage / db connection
	godotenv.Load()
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	var rowCount int64
	var rowMax int64 = 10000
	db.Table("incidents").Count(&rowCount)

	if rowCount >= rowMax {
		difference := rowCount - (rowMax - 1000)
		fmt.Printf("Cleaning %d rows... (%d/%d)\n", difference, rowCount, rowMax)
		trash := make([]api.Incident, 0)
		db.Order("start_timestamp").Limit(int(difference)).Find(&trash)
		err := storeIncidentBackup(trash)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		} else {
			tx := db.Delete(&trash)
			if tx.Error != nil {
				fmt.Fprintf(os.Stderr, "%s\n", tx.Error)
				os.Exit(1)
			} else {
				fmt.Printf("%d rows deleted.\n", tx.RowsAffected)
			}
		}
	} else {
		fmt.Printf("Only %d rows\n", rowCount)
	}
}

func storeIncidentBackup(incidents []api.Incident) error {
	// Set up Cloud Storage
	fmt.Printf("Storing %d incidents in the cloud\n", len(incidents))
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsJSON([]byte(os.Getenv("GCP_JSON"))))
	if err != nil {
		return err
	}
	defer client.Close()

	bucket := client.Bucket("cltrr-archive")
	object := bucket.Object(time.Now().Format(time.RFC1123))

	// Write to cloud storage
	writer := object.NewWriter(ctx)
	if err := writeIncidentsAsCsv(writer, incidents); err != nil {
		return err
	}

	// Cleanup
	if err := writer.Close(); err != nil {
		return err
	}

	// Set object meta data
	metadata := storage.ObjectAttrsToUpdate{
		Metadata: map[string]string{
			"oldestIncident": incidents[0].StartTimestamp.Format(time.RFC1123),
			"newestIncident": incidents[len(incidents)-1].StartTimestamp.Format(time.RFC1123),
			"rows":           strconv.Itoa(len(incidents)),
		},
	}
	if _, err := object.Update(ctx, metadata); err != nil {
		return err
	}

	return nil
}

func writeIncidentsAsCsv(writer io.Writer, incidents []api.Incident) error {
	// Write header
	e := reflect.ValueOf(&api.Incident{}).Elem()
	for i := 0; i < e.NumField(); i++ {
		if i != 0 {
			fmt.Fprint(writer, ",")
		}
		fmt.Fprint(writer, e.Type().Field(i).Name)
	}
	fmt.Fprintf(writer, "\n")

	// Write objects
	for _, incident := range incidents {
		e := reflect.ValueOf(&incident).Elem()
		for i := 0; i < e.NumField(); i++ {
			if i != 0 {
				fmt.Fprint(writer, ",")
			}
			fmt.Fprintf(writer, "\"%s\"", fmt.Sprint(e.Field(i).Interface()))
		}
		fmt.Fprintf(writer, "\n")
	}
	return nil
}
