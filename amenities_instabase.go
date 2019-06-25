package noonde

import "fmt"

var amenitiesInstabaseIds map[Amenity]int

func init() {
	amenitiesInstabaseIds = map[Amenity]int{
		AmenityProjector:         5,
		AmenityWhiteboard:        4,
		AmenityWifi:              1,
		AmenityKitchenFacilities: 10,
		AmenityMirrors:           15,
		AmenityParking:           21,
		AmenityAirConditioner:    20,
		AmenityDvdBrPlayer:       14,
		AmenityTv:                19,
		AmenitySoundSystem:       17,
		AmenityMonitor:           7,
		AmenityChairs:            3,
		AmenityTables:            2,
		// Only instabase: Not on spacemarket: Think how to handle... for now ignore.
		//
		// AmenityPiano
		// AmenityComputer
		// AmenityMusicalInstruments
		// AmenityMicSet
		// AmenityWiredMicSet
		// AmenityWiredLAN
		// AmenityOtherAmenities
		//
	}

}

// ToInstabaseID convert an amenity to it's ID on instabase.
func (a Amenity) ToInstabaseID() int {
	if id, ok := amenitiesInstabaseIds[a]; ok {
		return id
	}
	return 0
}

// ToInstabaseIDs convert a collection of amenities to their instabase IDs.
func (aa Amenities) ToInstabaseIDs() []int {
	var results []int
	for _, a := range aa {
		if id, ok := amenitiesInstabaseIds[a]; ok {
			results = append(results, id)
		}
	}

	return results
}

// InstabaseToNoondeAmenity translates instabase language to noonde language.
func InstabaseToNoondeAmenity(a string) Amenity {
	amenities := map[string]Amenity{
		"エレベーター":      AmenityElevator,
		"トイレ":         AmenityToilet,
		"監視カメラ":       AmenitySurveillanceCamera,
		"電源・コンセント":    AmenityPowerEquipment, // Maybe wrong?
		"DVDプレイヤー":    AmenityDvdBrPlayer,
		"アンプ・スピーカー":   AmenitySoundSystem,
		"エアコン":        AmenityAirConditioner,
		"キッチン":        AmenityKitchenFacilities,
		"スクリーン":       AmenityProjector, // Maybe wrong?
		"スポットライト":     AmenityLightingEquipment,
		"スライドプロジェクター": AmenityProjector,
		"テーブル":        AmenityTables,
		"テレビ":         AmenityTv,
		"バスタブ、シャワー":   AmenityBathtub,
		"プロジェクター":     AmenityProjector,
		"ホワイトボード":     AmenityWhiteboard,
		"モニター":        AmenityMonitor,
		"椅子":          AmenityChairs,
		"無線LAN":       AmenityWifi,
		"鏡":           AmenityMirrors,
		"駐車場":         AmenityParking,
		// パソコン                                // Don't have in spacemarket ..
		// ピアノ                                 // Don't have in spacemarket ..
		// フロアマット                              // Don't have in spacemarket ..
		// マイクセット                              // Don't have in spacemarket ..
		// レーザー光線機・アップライト照明・デザイン照明・スポットライト照明など // Don't have in spacemarket ..
		// レッドカーペット                            // Don't have in spacemarket ..
		// 卓球、ビリヤード、カラオケ                       // Don't have in spacemarket ..
		// 子供用イス                               // Don't have in spacemarket ..
		// 子供用カトラリー                            // Don't have in spacemarket ..
		// 座布団　                                // Don't have in spacemarket ..
		// 有線LAN                               // Don't have in spacemarket ..
		// 楽器                                  // Don't have in spacemarket ..
		// 演台                                  // Don't have in spacemarket ..
	}

	if amenity, ok := amenities[a]; ok {
		return amenity
	}

	fmt.Printf("Found new amenity: %s", a)
	// TODO
	return AmenityTables
}
