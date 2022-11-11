package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

const range_RekapJob = "Rekap Job!A2:Q"

func main() {
	data, err := ioutil.ReadFile("secret/client_secret.json")
	checkError(err)
	conf, err := google.JWTConfigFromJSON(data, sheets.SpreadsheetsScope)
	checkError(err)

	client := conf.Client(context.TODO())
	srv, err := sheets.New(client)
	checkError(err)

	spreadsheetID := "1eftlgXStFGK5iojtDGSPO-UkSofAYa7ACOmB0th2t7I"
	err = readSpreadsheet(srv, spreadsheetID)
	checkError(err)

}

func readSpreadsheet(srv *sheets.Service, spreadsheetID string) error {
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, range_RekapJob).Do()
	if err != nil {
		return err
	}

	if len(resp.Values) == 0 {
		return fmt.Errorf("no data found.")
	}

	for index, row := range resp.Values {
		fmt.Printf("%v %s, %s\n", index, row[0], row[3])
	}

	return nil
}
