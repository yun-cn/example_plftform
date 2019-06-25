package constants

// ReviewsQuery graphql query
const ReviewsQuery = `query Reviews($spaceId: ID!, $after: String) {
	  reviews(roomId: $spaceId, first: 20, after: $after) {
	    __typename
	    edges {
	      __typename
	      node {
	        __typename
	        ...review
	      }
      cursor
    }
    pageInfo {
	      __typename
	      endCursor
	      hasNextPage
	      hasPreviousPage
	    }
  }
}fragment review on Review {
	  __typename
	  id
	  title
	  comment
	  point
	  usage
	  age
	  gender
	  createdAt
	}`
