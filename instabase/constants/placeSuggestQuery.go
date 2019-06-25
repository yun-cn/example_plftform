package constants

// PlaceSuggestQuery graphql query
const PlaceSuggestQuery = `query PlaceSuggest($keyword: String!) {
	  placeSearch: areaSearch(query: $keyword) {
	    __typename
	    ...SuggestResult
	  }
}fragment SuggestResult on AreaSearch {
	  __typename
	  name
	  modelType
	  modelId
	}`
