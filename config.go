package main

import (
	"os"
	"path/filepath"
)

type configuration struct {
	GoogleSpreadsheet googleSpreadSheet
	GoogleAuth        googleAuth
}

type googleSpreadSheet struct {
	ID    string
	Range string
}

type googleAuth struct {
	CredentialsPath string
	TokenPath       string
}

func configure() (configuration, error) {
	conf := configuration{}
	home, err := os.UserHomeDir()

	if err != nil {
		return conf, err
	}

	defaulConfigRoot := filepath.Join(home, ".config", "isp-monitor")
	configRoot := getEnv("CONFIG_ROOT", defaulConfigRoot)
	googleSheetsAuthDir := filepath.Join(configRoot, "google-sheets")

	googleAuth := googleAuth{
		CredentialsPath: filepath.Join(googleSheetsAuthDir, "credentials.json"),
		TokenPath:       filepath.Join(googleSheetsAuthDir, "token.json"),
	}

	googleSpreadSheet := googleSpreadSheet{
		ID:    "1rpUobs799LuviC5pGjk3sOd3KxfuM_M5BkzUO_OS1S0",
		Range: "A1",
	}

	conf.GoogleSpreadsheet = googleSpreadSheet
	conf.GoogleAuth = googleAuth

	return conf, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
