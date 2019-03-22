package main

import (
	"testing"
	"time"

	"github.com/uzimaru0000/eps/packages"
)

func TestCacheCheck(t *testing.T) {
	now := time.Now()
	oneWeekAgo := now.AddDate(0, 0, -7)

	result := packages.CacheCheck(oneWeekAgo, now)
	if !result {
		t.Fatalf("Fail")
	}
}
