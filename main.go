package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go-calendar/calendar"
	"go-calendar/calendar/event"
	"go-calendar/calendar/reminder"
	"go-calendar/model"
	"go-calendar/parser"
	"os"
	"time"
)

func main() {
	inputPath := flag.String("input", "./events.json", "json file containing calendar events data")
	outputPath := flag.String("output", "./calendar.ics", "file to export ics calendar")
	flag.Parse()

	data, err := os.ReadFile(*inputPath)
	if err != nil {
		panic(err)
	}

	var cfg parser.Events
	if err := json.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}

	events := make([]model.Renderer, len(cfg.Events))
	for i, e := range cfg.Events {
		opts := make([]event.Option, len(e.Reminders)+2)
		opts[0] = event.WithDescription(e.Description)
		opts[1] = event.WithLocation(e.Location)
		for i, r := range e.Reminders {
			if r.Type != reminder.Notification {
				panic(fmt.Sprintf("reminder type not handled: %q", r.Type))
			}
			if r.TriggerType != reminder.Relative {
				panic(fmt.Sprintf("reminder trigger type not handled: %q", r.TriggerType))
			}

			opts[i+2] = event.WithReminder(
				reminder.NewRenderer(
					reminder.RelativeTrigger{
						Timing:   reminder.Timing(r.AfterEvent),
						Duration: time.Duration(r.Duration),
					},
					reminder.WithDescription(r.Description),
				),
			)
		}
		events[i] = event.NewRenderer(
			e.Title,
			e.StartDate,
			time.Duration(e.Duration),
			opts...,
		)
	}

	ics := calendar.Render(events...)
	if err := os.WriteFile(*outputPath, []byte(ics), 0650); err != nil {
		panic(err)
	}
}
