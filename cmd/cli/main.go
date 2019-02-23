package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dohernandez/travels-budget/internal/platform/bootstrap"
	"github.com/dohernandez/travels-budget/internal/platform/config"
	handler "github.com/dohernandez/travels-budget/internal/platform/handler/cli"
	"github.com/dohernandez/travels-budget/internal/platform/intput"
	"github.com/dohernandez/travels-budget/pkg/version"
	"github.com/urfave/cli"
)

var binaryName = "travels-budget"

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Println(version.Info().String())
		return
	}

	ctx, cancelCtx := context.WithCancel(context.TODO())
	defer cancelCtx()

	conf, err := config.LoadEnv()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	app := cli.NewApp()
	app.Version = version.Info().Version
	app.Name = binaryName
	app.Usage = "To build a personalized itinerary planner based on budget and number of days " +
		"to inspire travelers to seek out the memorable things to do."
	app.UsageText = fmt.Sprintf("%s command [arguments]", binaryName)

	c := bootstrap.NewContainer(conf)

	app.Commands = []cli.Command{
		{
			Name:   "personalized",
			Usage:  "Build a personalized itinerary planner based on budget and number of days to inspire travelers to seek out the memorable things to do",
			Action: handler.NewPersonalizeScheduleHandler(ctx, c),
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  intput.BudgetKey + ", b",
					Usage: "budget willing to spend",
				},
				cli.StringFlag{
					Name:  intput.DaysKey + ", d",
					Usage: "days willing to use",
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
