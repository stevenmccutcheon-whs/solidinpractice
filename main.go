package main

import (
	"log/slog"
	"os"

	"github.com/smccutcheon-whs/solidinpractice/nearearth"
)

func main() {
	// Set up the environment
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))
	slog.Info("Starting up")

	// Create the file service
	fileService, err := nearearth.NewImpactFileService()
	if err != nil {
		slog.Error("error creating file service", "error", err)
		os.Exit(1)
	}
	fileCount, err := nearearth.Count(fileService)
	if err != nil {
		slog.Error("error counting near earth impacts", "error", err)
		os.Exit(1)
	}
	slog.Info("counted near earth impacts from file", "count", fileCount)

	/// Create the http service
	httpService := nearearth.NewImpactHTTPService()
	httpCount, err := nearearth.Count(httpService)
	if err != nil {
		slog.Error("error counting near earth impacts", "error", err)
		os.Exit(1)
	}

	// Create the nil service
	slog.Info("counted near earth impacts from http", "count", httpCount)
	nilCount, err := nearearth.Count(nil)
	if err != nil {
		slog.Error("error counting near earth impacts", "error", err)
		os.Exit(1)
	}
	slog.Info("counted near earth impacts from nil", "count", nilCount)
}
