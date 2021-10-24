package sportsobjectsdetailed

import (
	"github.com/techpotion/leaders2021-backend/gen/pb"
)

// ObjectDetailedToObject is a dummy converter from SportsObjectDetailed structure to SportsObject
func ObjectDetailedToObject(in *pb.SportsObjectDetailed) *pb.SportsObject {
	return &pb.SportsObject{
		ObjectId:                     in.ObjectId,
		ObjectName:                   in.ObjectName,
		ObjectPoint:                  in.ObjectPoint,
		DepartmentalOrganizationId:   in.DepartmentalOrganizationId,
		DepartmentalOrganizationName: in.DepartmentalOrganizationName,
		Availability:                 in.Availability,
	}
}

// UniqueObjectsFromDetailed returnes unique SportsObject from SportsObjectDetailed list
func UniqueObjectsFromDetailed(objects []*pb.SportsObjectDetailed) []*pb.SportsObject {
	var result []*pb.SportsObject
	idsHashMap := make(map[uint32]bool)
	for _, object := range objects {
		if _, ok := idsHashMap[object.ObjectId]; !ok {
			idsHashMap[object.ObjectId] = true
			result = append(result, ObjectDetailedToObject(object))
		}
	}
	return result
}
