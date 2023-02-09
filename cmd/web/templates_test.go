package main

import (
	"testing"
	"time"

	"github.com/t1mtw0/snippetbox/internal/assert"
)

func TestHumanDate(t *testing.T) {
	t.Run("UTC", func(t *testing.T) {
		hd := humanDate(time.Date(2022, 3, 17, 10, 15, 0, 0, time.UTC))
		assert.Equal(t, hd, "17 Mar 2022 at 10:15")
	})

	t.Run("Empty", func(t *testing.T) {
		hd := humanDate(time.Time{})
		assert.Equal(t, hd, "")
	})

	t.Run("CET", func(t *testing.T) {
		hd := humanDate(time.Date(2022, 3, 17, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)))
		assert.Equal(t, hd, "17 Mar 2022 at 09:15")
	})
}
