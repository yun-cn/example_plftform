package constants

// RoomInitializeQuery graphql query
const RoomInitializeQuery = `query RoomInitialize($roomId: ID!, $roomUid: String, $roomRentType: RentType!, $canRentType: RentType!) {
	  room(uid: $roomUid, rentType: $roomRentType) {
	    __typename
	    canRentRoom(rentType: $canRentType)
	    sponsoredPromotions {
	      __typename
	      results {
	        __typename
	        ...roomSponsoredPromotion
	      }
    }
  }
  reputations(page: 1, perPage: 3, roomId: $roomId, fromType: USER) {
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
  stayRoomTypeFields {
	    __typename
	    results {
	      __typename
	      ...stayRoomTypeFragment
	    }
  }
}fragment roomSponsoredPromotion on SponsoredPromotion {
	  __typename
	  name
	  link
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
}fragment stayRoomTypeFragment on StayRoomTypeField {
	  __typename
	  id
	  name
	  nameText
}`
