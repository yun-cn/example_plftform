package constants

// SearchStayRoomsQuery graphql query
const SearchStayRoomsQuery = `query searchStayRooms($page: Int, $perPage: Int, $withRecommend: Boolean, $eventType: Int, $geocode: String, $location: String, $state: String, $startedAt: String, $endedAt: String, $minCapacity: Int, $minPrice: Int, $maxPrice: Int, $hasDirectReservationPlans: Boolean, $hasTodayReservationPlans: Boolean, $hasLastMinuteDiscountPlans: Boolean, $hasWeeklyDiscountPlans: Boolean, $hasMonthlyDiscountPlans: Boolean, $sponsoredPromotionIds: String, $stayRoomTypes: [StayRoomType], $keyword: String, $amenities: String) {
	searchStayRooms(page: $page, perPage: $perPage, withRecommend: $withRecommend, eventType: $eventType, geocode: $geocode, location: $location, state: $state, startedAt: $startedAt, endedAt: $endedAt, minCapacity: $minCapacity, minPrice: $minPrice, maxPrice: $maxPrice, hasDirectReservationPlans: $hasDirectReservationPlans, hasTodayReservationPlans: $hasTodayReservationPlans, hasLastMinuteDiscountPlans: $hasLastMinuteDiscountPlans, hasWeeklyDiscountPlans: $hasWeeklyDiscountPlans, hasMonthlyDiscountPlans: $hasMonthlyDiscountPlans, sponsoredPromotionIds: $sponsoredPromotionIds, stayRoomTypes: $stayRoomTypes, keyword: $keyword, amenities: $amenities) {
	  __typename
  pageInfo {
		__typename
	totalCount
  }
  results {
		__typename
	...stayRoom
  }
}
}fragment stayRoom on SearchStayRoom {
	__typename
id
uid
name
hasLastMinuteDiscountPlans
hasDirectReservationPlans
stayCapacity
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
