package constants

// SearchRoomsQuery graphql query
const SearchRoomsQuery = `query searchRooms($page: Int, $perPage: Int, $withRecommend: Boolean, $eventType: Int, $geocode: String, $location: String, $state: String, $priceType: PriceType!, $startedAt: String, $endedAt: String, $startedTime: String, $endedTime: String, $minCapacity: Int, $maxCapacity: Int, $minPrice: Int, $maxPrice: Int, $hasDirectReservationPlans: Boolean, $hasTodayReservationPlans: Boolean, $hasLastMinuteDiscountPlans: Boolean, $sponsoredPromotionIds: String, $keyword: String, $amenities: String) {
	searchRooms(page: $page, perPage: $perPage, withRecommend: $withRecommend, eventType: $eventType, geocode: $geocode, location: $location, state: $state, priceType: $priceType, startedAt: $startedAt, endedAt: $endedAt, startedTime: $startedTime, endedTime: $endedTime, minCapacity: $minCapacity, maxCapacity: $maxCapacity, minPrice: $minPrice, maxPrice: $maxPrice, hasDirectReservationPlans: $hasDirectReservationPlans, hasTodayReservationPlans: $hasTodayReservationPlans, hasLastMinuteDiscountPlans: $hasLastMinuteDiscountPlans, sponsoredPromotionIds: $sponsoredPromotionIds, keyword: $keyword, amenities: $amenities) {
	  __typename
  pageInfo {
		__typename
	totalCount
  }
  results {
		__typename
	...room
  }
}
}fragment room on SearchRoom {
	__typename
id
uid
name
hasLastMinuteDiscountPlans
hasDirectReservationPlans
capacity
totalReputationCount
totalReputationScore
stateText
city
access
prices {
	  __typename
  ...priceFragment
}
thumbnails {
	  __typename
  url
}
isInquiryOnly
ownerRank
latitude
longitude
isFavorite
sponsoredPromotions {
	  __typename
  name
}
availablePlanCount
plans {
	  __typename
  name
  price {
		__typename
	...priceFragment
  }
  directReservationAccepted
  isLastMinuteDiscount
}
isCancelFree
}fragment priceFragment on SearchPrice {
	__typename
minText
minUnitText
maxText
maxUnitText
}`
