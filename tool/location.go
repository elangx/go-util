package tool

import "math"

// EarthRadius 地球半径，单位：米
const earthRadius = 6371 * 1000

// HaversineDistance 根据经纬度计算两点之间的距离，以米为单位
func HaversineDistance(lat1, lng1, lat2, lng2 float64) float64 {
	// 将经纬度转换为弧度
	var lat1Rad, lon1Rad, lat2Rad, lon2Rad float64
	lat1Rad = lat1 * math.Pi / 180
	lon1Rad = lng1 * math.Pi / 180
	lat2Rad = lat2 * math.Pi / 180
	lon2Rad = lng2 * math.Pi / 180

	// 计算两点间的经度差和纬度差
	var deltaLat, deltaLon float64
	deltaLat = lat2Rad - lat1Rad
	deltaLon = lon2Rad - lon1Rad

	// 应用Haversine公式计算距离
	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// 计算距离，单位：米
	distance := earthRadius * c
	return distance
}
