SELECT count(*) from public.objects
WHERE ST_Contains(
	ST_GeomFromText('POLYGON((
		37.568742 55.776336,
		37.563248 55.715140,
		37.679978 55.719791,
		37.668305 55.799550,
		37.568742 55.776336
		))'
	),
	ST_Point(object_point_lng, object_point_lat)
);