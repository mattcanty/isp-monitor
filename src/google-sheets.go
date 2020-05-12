package main

import (
	"net/http"

	"google.golang.org/api/sheets/v4"
)

func appendSpreadsheet(conf googleSpreadSheet, client *http.Client) error {
	sheetsService, err := sheets.New(client)
	if err != nil {
		return err
	}

	rb := &sheets.ValueRange{
		Values: values,
	}

	_, err = sheetsService.Spreadsheets.Values.Append(conf.ID, conf.Range, rb).ValueInputOption("RAW").InsertDataOption("INSERT_ROWS").Do()

	return err
}
