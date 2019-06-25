package elasticsearch_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	noonde "github.com/yanshiyason/noonde_platform"

	"github.com/olivere/elastic"

	"github.com/yanshiyason/noonde_platform/elasticsearch"
)

const testIndex = "noonde_test"
const testListingsIndex = "noonde_test_listings"
const testRequestsIndex = "noonde_test_requests"
const testPlaceSuggestionsIndex = "noonde_test_place_suggestions"

const elasticSearchURL = "http://127.0.0.1:9200"

func setup(t *testing.T) {
	// setup elasticsearch
	tearDown(t)

	path := filepath.Join("../scripts", "elastic_search_create_index.sh")
	out, err := exec.Command("/bin/bash", path, testIndex).CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(out))
}

func tearDown(t *testing.T) {
	path := filepath.Join("../scripts", "elastic_search_drop_index.sh")
	out, err := exec.Command("/bin/bash", path, testIndex).CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(out))
}

func defaultClient(t *testing.T) *elastic.Client {
	ctx := context.TODO()

	// Obtain a client and connect to the default Elasticsearch installation
	// on 127.0.0.1:9200.
	client, err := elastic.NewClient()
	if err != nil {
		t.Fatal(err)
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping(elasticSearchURL).Do(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// Use the IndexExists service to check if a specified index exists.
	_, err = client.IndexExists(testIndex).Do(ctx)
	if err != nil {
		t.Fatal(err)
	}

	return client
}

type Key string

func (k Key) DocID() string {
	return string(k)
}

func Test_GetRequest(t *testing.T) {
	setup(t)

	doc := map[string]interface{}{
		"hello": 1,
	}
	search := elasticsearch.NewDefaultService()
	search.ListingsIndex = testListingsIndex
	search.RequestsIndex = testRequestsIndex

	key := Key("spacemarket://location:東京,page:1,number_of_guests:2")

	err := search.IndexRequest(key, doc)
	if err != nil {
		t.Fatalf(`got error: %s`, err)
	}

	time.Sleep(1 * time.Second)

	response, err := search.GetRequest(key)
	if err != nil {
		t.Errorf(`got error: %s`, err)
	}

	j, err := json.Marshal(&response)
	if err != nil {
		t.Fatalf(`got error: %s`, err)
	}

	jsonDoc := `{"hello":1}`

	if string(j) != jsonDoc {
		t.Errorf(`
		got:      %s
		expected: %s
		`, string(j), jsonDoc)
	}
}

func Test_GetListing(t *testing.T) {
	setup(t)
	search := elasticsearch.NewDefaultService()
	search.ListingsIndex = testListingsIndex
	search.RequestsIndex = testRequestsIndex

	listingId := 5000
	payload := struct {
		ID int `json:"id"`
	}{
		ID: listingId,
	}
	now, _ := time.Parse(time.RFC3339, "2019-03-10T00:00:00+00:00")
	results := &noonde.IndexedListingDetails{
		IDOnPlatform: fmt.Sprintf("%d", listingId),
		Data:         payload,
		RefreshedAt:  &now,
	}

	docID := Key("spacemarket://5000")
	err := search.IndexListing(docID, results)
	if err != nil {
		fmt.Printf("error indexing search results: %+v\n", err)
		return
	}

	time.Sleep(1 * time.Second)

	response, err := search.GetListing(docID)
	if err != nil {
		t.Errorf(`got error: %s`, err)
	}

	j, err := json.Marshal(&response)
	if err != nil {
		t.Fatalf(`got error: %s`, err)
	}

	jsonDoc := `{"id_on_platform":"5000","data":{"id":5000},"refreshed_at":"2019-03-10T00:00:00Z"}`

	if string(j) != jsonDoc {
		t.Errorf(`
		got:      %s
		expected: %s
		`, string(j), jsonDoc)
	}
}
