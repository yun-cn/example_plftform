package constants

// {"query":"",
// "variables":{"id":"6149"}}

// SpaceDetailQuery graphql query
const SpaceDetailQuery = `query SpaceDetail($id: ID!) {
	  rooms(ids: [$id]) {
	    __typename
	    ...spaceDetail
	  }
}fragment spaceDetail on Room {
	  __typename
	  ...space
	  uid
	  bookingBeforeLimitDays
	  minBookingHours
	  siteUrl
	  spaceCategory {
	    __typename
	    title
	  }
  seoDescription
  usages {
	    __typename
	    id
	    title
	  }
  summaryBusinessDays {
	    __typename
	    days
	    businessHours
	  }
  summaryPrice
  summaryPriceTable {
	    __typename
	    timetable {
	      __typename
	      time
	      price
	    }
    days
  }
  attentions {
	    __typename
	    id
	    isOfficial
	    title
	    priority
	    description
	  }
  notice
  building {
	    __typename
	    address
	    id
	    lat
	    lon
	    prefecture {
	      __typename
	      id
	    }
    title
  }
  review: review {
	    __typename
	    ...review
	  }
  cancelPolicies {
	    __typename
	    title
	    percentage
	  }
  otherEquipmentNames
  freeEquipments: equipmentRooms {
	    __typename
	    ...freeEquipment
	  }
  chargedEquipments: chargedEquipmentRooms {
	    __typename
	    ...chargedEquipment
	  }
}fragment space on Room {
	  __typename
	  images: roomImages {
	    __typename
	    spaceId: roomId
	    id
	    filePath
	  }
  id
  title
  friendlyTitle
  isAnyAvailable
  isOrderApprove
  seoDescription
  square
  reviewCount
  averagePoint
  spaceType: roomType
  capacity
  summaryPrice
  summaryMinPrice
  summaryMaxPrice
  spaceUrl: siteUrl
  building {
	    __typename
	    summaryAccesses: summaryAccess {
	      __typename
	      access
	      line
	      station
	    }
    lat
    lon
    parentArea {
	      __typename
	      id
	      title
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
	}fragment freeEquipment on EquipmentRoom {
	  __typename
	  title
	  description
	  countDescription
	}fragment chargedEquipment on EquipmentRoom {
	  __typename
	  ...freeEquipment
	  summaryPrice
	}`
