package elasticsearch_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/yanshiyason/noonde_platform/elasticsearch"
)

func Test_GetPlaceSuggestions(t *testing.T) {
	setup(t)

	search := elasticsearch.NewDefaultService()
	search.PlaceSuggestionsIndex = testPlaceSuggestionsIndex

	results := []map[string]interface{}{
		{"name": "下関市", "ward_id": 35201, "search_id": 0, "prefecture_id": 0},
		{"name": "品川駅", "ward_id": 0, "search_id": 1130103, "prefecture_id": 0},
		{"name": "静岡県", "ward_id": 0, "search_id": 0, "prefecture_id": 22},
	}

	doc := map[string]interface{}{
		"results": results,
	}

	docID := Key("shi")

	err := search.IndexPlaceSuggestion(docID, doc)
	if err != nil {
		t.Fatalf(`got error: %s`, err)
	}

	time.Sleep(1 * time.Second)

	response, err := search.GetPlaceSuggestion(docID)
	if err != nil {
		t.Errorf(`got error: %s`, err)
	}

	j, err := json.Marshal(&response)
	if err != nil {
		t.Fatalf(`got error: %s`, err)
	}

	jsonDoc := `{"results":[{"name":"下関市","prefecture_id":0,"search_id":0,"ward_id":35201},{"name":"品川駅","prefecture_id":0,"search_id":1130103,"ward_id":0},{"name":"静岡県","prefecture_id":22,"search_id":0,"ward_id":0}]}`

	if string(j) != jsonDoc {
		t.Errorf(`
		got:      %s
		expected: %s
		`, string(j), jsonDoc)
	}
}
