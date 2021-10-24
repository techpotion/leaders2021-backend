package sportsobjects

import "github.com/techpotion/leaders2021-backend/gen/pb"

// UniqueObjects returns unique SportsObject from SportsObject list
func UniqueObjects(objects []*pb.SportsObject) []*pb.SportsObject {
	var result []*pb.SportsObject
	idsHashMap := make(map[uint32]bool)
	for _, object := range objects {
		if _, ok := idsHashMap[object.ObjectId]; !ok {
			idsHashMap[object.ObjectId] = true
			result = append(result, object)
		}
	}
	return result
}
