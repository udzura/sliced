package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"

	cli "github.com/urfave/cli/v2"
)

const VERSION = "0.1.1"

func actionSliced(c *cli.Context) error {
	total := c.Int("n")
	if total < 1 {
		return fmt.Errorf("-n should be specified or more than 0")
	}
	index := c.Int("i")
	if total <= index {
		return fmt.Errorf("-i must be less than -n")
	}

	stream := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		stream = append(stream, strings.TrimSuffix(scanner.Text(), "\n"))
	}

	seed := c.Int64("seed")
	if seed < 1 {
		sha := c.String("commit-sha")
		if sha == "" {
			return fmt.Errorf("Either seed number or sha is required")
		}

		seed_, err := strconv.ParseInt(sha[0:15], 16, 64)
		if err != nil {
			return fmt.Errorf("parsing seed failed: %w", err)
		}
		seed = seed_
	}

	rand.Seed(seed)

	indexes := make([]int, len(stream))
	for i := range stream {
		indexes[i] = i
	}
	rand.Shuffle(len(indexes), func(i1, i2 int) { indexes[i1], indexes[i2] = indexes[i2], indexes[i1] })

	start := math.Ceil(float64(len(stream)) / float64(total) * float64(index))
	end := math.Ceil(float64(len(stream)) / float64(total) * float64(index+1))
	log.Printf("start: %v end:%v total:%v", start, end, len(stream))
	part := indexes[int(start):int(end)]

	result := make([]string, 0)
	for _, i := range part {
		result = append(result, stream[i])
	}

	for _, ln := range result {
		if os.Getenv("DEBUG") == "1" {
			fmt.Fprintf(os.Stderr, "output: %s\n", ln)
		}
		fmt.Printf("%s\n", ln)
	}

	return nil
}

func main() {
	app := &cli.App{
		Name:    "sliced",
		Usage:   "Stream to slice generator cli",
		Version: VERSION,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     "n",
				Usage:    "total number of slices",
				Required: true,
				EnvVars:  []string{"SLICED_TOTAL_SLICE"},
			},
			&cli.IntFlag{
				Name:     "i",
				Usage:    "index of slice to fetch/ starts with 0",
				Required: true,
				EnvVars:  []string{"SLICED_CURRENT_SLICE"},
			},
			&cli.Int64Flag{
				Name:    "seed",
				Aliases: []string{"S"},
				Usage:   "Fixed seed to use generating random",
				EnvVars: []string{"SLICED_SEED"},
			},
			&cli.StringFlag{
				Name:    "commit-sha",
				Aliases: []string{"C"},
				Usage:   "Fixed seed to use generating random, commit-SHA formed",
				EnvVars: []string{"SLICED_COMMIT_SHA", "COMMIT_SHA"},
			},
		},
		Action: actionSliced,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
