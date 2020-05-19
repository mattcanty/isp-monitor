package main

import (
	"net/http"

	"google.golang.org/api/sheets/v4"
)

func appendSpreadsheet(conf googleSpreadSheet, client *http.Client, values [][]interface{}) error {
	sheetsService, err := sheets.New(client)
	if err != nil {
		return err
	}

	requestBody := &sheets.ValueRange{
		Values: values,
	}

	_, err = sheetsService.Spreadsheets.Values.Append(conf.ID, conf.Range, requestBody).ValueInputOption("RAW").InsertDataOption("INSERT_ROWS").Do()

	return err
}
