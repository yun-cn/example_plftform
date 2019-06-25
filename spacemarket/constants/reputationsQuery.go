package constants

// ReputationsQuery graphql query
const ReputationsQuery = `query Reputations($page: Int, $perPage: Int, $rentType: RentType, $roomId: ID!) {
	  reputations(page: $page, perPage: $perPage, rentType: $rentType, roomId: $roomId, fromType: USER) {
	    __typename
	    pageInfo {
	      __typename
	      totalCount
	    }
    results {
	      __typename
	      ...reputationFragment
	    }
  }
}fragment reputationFragment on Reputation {
	  __typename
	  description
	  from {
	    __typename
	    ... on User {
	      id
	      name
	      profileImage
	      username
	    }
  }
  reservation {
	    __typename
	    startedAt
	    eventTypeText
	    numberOfGuests
	  }
}`
