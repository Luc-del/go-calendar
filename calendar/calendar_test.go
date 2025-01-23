package calendar_test

import (
	"fmt"
	"go-calendar/calendar"
	"go-calendar/calendar/event"
	"go-calendar/calendar/reminder"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func loadTestData(t *testing.T) string {
	split := strings.Split(t.Name(), "/")
	file := fmt.Sprintf("testdata/%s.ics", strings.ReplaceAll(split[len(split)-1], " ", "_"))
	data, err := os.ReadFile(file)
	require.NoError(t, err, file)

	return string(data)
}

func replace(data string) string {
	var res = data
	for _, line := range strings.Split(data, "\n") {
		if strings.HasPrefix(line, "UID:") {
			res = strings.ReplaceAll(res, line, "UID:__UID__")
		}
		if strings.HasPrefix(line, "DTSTAMP:") {
			res = strings.ReplaceAll(res, line, "DTSTAMP:__DTSTAMP__")
		}
	}
	return res
}

func TestRender(t *testing.T) {
	loc, err := time.LoadLocation("Europe/Paris")
	require.NoError(t, err)

	start := time.Date(1993, 7, 10, 20, 5, 0, 0, loc)

	t.Run("without reminders", func(t *testing.T) {
		got := calendar.Render(event.NewRenderer("my title", start, time.Hour))
		assert.Equal(t, loadTestData(t), replace(got))
	})

	t.Run("with reminders", func(t *testing.T) {
		got := calendar.Render(
			event.NewRenderer("my title", start, 30*time.Minute,
				event.WithDescription("event with reminders"),
				event.WithLocation("somewhere there"),
				event.WithReminder(func() string {
					return reminder.Render(reminder.RelativeTrigger{
						Timing:   reminder.Before,
						Duration: 30 * time.Minute,
					},
						reminder.WithDescription("reminder 1"),
					)
				}),
				event.WithReminder(func() string {
					return reminder.Render(reminder.RelativeTrigger{
						Timing:   reminder.After,
						Duration: 5 * time.Minute,
					},
						reminder.WithDescription("reminder 2"),
					)
				}),
			),
			event.NewRenderer("second title", start.Add(time.Hour), time.Minute,
				event.WithDescription("event without reminder"),
			),
			event.NewRenderer("third title", start.Add(24*time.Hour), time.Hour,
				event.WithDescription("event with a single reminder"),
				event.WithReminder(func() string {
					return reminder.Render(reminder.RelativeTrigger{
						Timing:   reminder.Before,
						Duration: time.Hour,
					},
						reminder.WithDescription("single reminder"),
					)
				}),
			),
		)
		assert.Equal(t, loadTestData(t), replace(got))
	})
}
