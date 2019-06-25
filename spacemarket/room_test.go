package spacemarket

import (
	"testing"

	"github.com/yanshiyason/noonde_platform/spacemarket/types"
)

// URLFor returns correct URL depending on RentType
func TestURLFor(t *testing.T) {
	tables := []struct {
		id       int
		rt       types.RentType
		expected string
	}{
		{10, types.RTDay, BaseURL + "/rooms/10.json?rent_type=1"},
		{20, types.RTNight, BaseURL + "/rooms/20.json?rent_type=2"},
	}

	for _, table := range tables {
		url, _ := URLFor(table.id, table.rt)
		if url.String() != table.expected {
			t.Errorf("Expected %s to equal %s.", url, table.expected)
		}
	}
}
