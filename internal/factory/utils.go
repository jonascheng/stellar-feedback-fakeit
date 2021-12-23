package factory

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func CounterfeitTimestamp(counterfeitHour bool, counterfeitDay bool, counterfeitMonth bool) time.Time {
	timestamp := time.Now().UTC()

	if counterfeitHour || counterfeitDay || counterfeitMonth {
		// prevent generate a future datetime
		timestamp = timestamp.AddDate(-1, 0, 0)
	}

	if counterfeitHour {
		hour := time.Duration(gofakeit.Number(0, 23)) * time.Hour
		timestamp = timestamp.Add(hour)
	}

	if counterfeitDay {
		day := gofakeit.Number(1, 31)
		timestamp = timestamp.AddDate(0, 0, day)
	}

	if counterfeitMonth {
		month := gofakeit.Number(1, 12)
		timestamp = timestamp.AddDate(0, month, 0)
	}

	return timestamp
}
