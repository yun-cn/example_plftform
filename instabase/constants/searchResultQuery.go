package constants

// usage ids:
// uchiawase: 1
// lecon: 2
// party: 4
// photoshoot: 10
// therapy 6
// counseling 13
// studygroup 5
// boardgame 12
// seminar 3
// yoga 7
// mom party 14
// interview/test 11
// workshop 15
// showing film 16
// workplace 8

// {
//     "bottomRightLon": null,
//     "prefectureId": null,
//     "capacityIds": [
//         3
//     ],
//     "fromDateMonth": 3,
//     "toTime": "14:00",
//     "wardId": null,
//     "page": 1,
//     "conditionIds": [
//         9,
//         36,
//         25,
//         78,
//         12,
//         16,
//         79,
//         72,
//         17,
//         30,
//         68
//     ],
//     "categoryIds": [
//         1,
//         2,
//         4,
//         10,
//         3,
//         6,
//         8,
//         9,
//         7,
//         12
//     ],
//     "equipmentIds": [
//         5,
//         4,
//         1,
//         21,
//         15,
//         11,
//         10,
//         14,
//         16,
//         20,
//         19,
//         8,
//         17,
//         12,
//         13,
//         7,
//         18,
//         2,
//         3,
//         9
//     ],
//     "perPage": null,
//     "bottomRightLat": null,
//     "fromDateYear": 2019,
//     "fromDateDay": 14,
//     "orderBy": null,
//     "usageIds": [
//         1,
//         3,
//         2,
//         7,
//         4,
//         14,
//         10,
//         11,
//         6,
//         15,
//         13,
//         16,
//         5,
//         8,
//         12
//     ],
//     "topLeftLat": null,
//     "areaId": null,
//     "topLeftLon": null,
//     "fromTime": "12:00",
//     "stationId": 1130208
// }

// $prefectureId: Int = null,
// $wardId: Int = null,
// $areaId: Int = null,
// $categoryIds: [Int!] = null,
// $usageIds: [Int!] = null,
// $capacityIds: [Int!] = null,
// $equipmentIds: [Int!] = null,
// $conditionIds: [Int!] = null,
// $stationId: Int = null,
// $topLeftLat: Float = null,
// $topLeftLon: Float = null,
// $bottomRightLat: Float = null,
// $bottomRightLon: Float = null,
// $fromDateYear: Int = null,
// $fromDateMonth: Int = null,
// $fromDateDay: Int = null,
// $fromTime: String = null,
// $toTime: String = null,
// $orderBy: String = null,
// $page: Int = null,
// $perPage: Int = null

// SearchResultQuery graphql query
const SearchResultQuery = `query SearchResult($prefectureId: Int = null, $wardId: Int = null, $areaId: Int = null, $categoryIds: [Int!] = null, $usageIds: [Int!] = null, $capacityIds: [Int!] = null, $equipmentIds: [Int!] = null, $conditionIds: [Int!] = null, $stationId: Int = null, $topLeftLat: Float = null, $topLeftLon: Float = null, $bottomRightLat: Float = null, $bottomRightLon: Float = null, $fromDateYear: Int = null, $fromDateMonth: Int = null, $fromDateDay: Int = null, $fromTime: String = null, $toTime: String = null, $orderBy: String = null, $page: Int = null, $perPage: Int = null) {
	  spaces: search(prefectureId: $prefectureId, wardId: $wardId, areaId: $areaId, cat: $categoryIds, u: $usageIds, c: $capacityIds, e: $equipmentIds, f: $conditionIds, s: $stationId, topLeftLat: $topLeftLat, topLeftLon: $topLeftLon, bottomRightLat: $bottomRightLat, bottomRightLon: $bottomRightLon, fromDateYear: $fromDateYear, fromDateMonth: $fromDateMonth, fromDateDay: $fromDateDay, fromTime: $fromTime, toTime: $toTime, orderBy: $orderBy, page: $page, perPage: $perPage) {
	    __typename
	    ...space
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
}`
