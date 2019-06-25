package elasticsearch

// Query is an elastic search query
type Query struct {
	Type string
	// IDQuery find a document by id.
	IDQuery     string
	SearchQuery string
	TagsQuery   string
	Filter      string
	Start       interface{}
	End         interface{}
	From        int
	Size        int
	SortField   string
	SortAsc     bool
	Asc         bool
	Pretty      bool
}
