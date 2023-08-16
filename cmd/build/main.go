package main

import (
	"flag"
	"fmt"
	"os"

	"go.sazak.io/cv-go/pkg/builder"

	"golang.org/x/exp/slog"
)

var (
	outputPath = flag.String("out", builder.DefaultOutputPath, fmt.Sprintf("Output path. Defaults to %q", builder.DefaultOutputPath))
	configPath = flag.String("config", builder.DefaultConfigPath, fmt.Sprintf("Config YAML file path. Defaults to %q", builder.DefaultConfigPath))
)

func main() {
	flag.Parse()

	opts := []builder.Option{}
	if *outputPath != "" {
		opts = append(opts, builder.WithOutputPath(*outputPath))
	}
	if *configPath != "" {
		opts = append(opts, builder.WithConfigPath(*configPath))
	}

	b, err := builder.New(opts...)
	if err != nil {
		slog.Error("Failed to initialize CV builder: %w", err)
		os.Exit(1)
	}
	path, err := b.Build()
	if err != nil {
		slog.Error("Failed to build the CV: %w", err)
		os.Exit(1)
	}

	slog.Info("Successfully built the CV and saved to %s!", path)
}
