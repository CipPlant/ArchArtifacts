package mongoFilter

import "go.mongodb.org/mongo-driver/bson"

func SearchFiler(a, b int16) bson.M {
	filter := bson.M{
		"$or": []bson.M{
			{
				"start_date":   bson.M{"$gte": a},
				"interval_end": bson.M{"$lte": b},
			},
			{
				"start_date":   bson.M{"$lt": a},
				"interval_end": bson.M{"$gte": a, "$lte": b},
			},
			{
				"start_date":   bson.M{"$lte": b, "$gte": a},
				"interval_end": bson.M{"$gt": b},
			},
			{
				"start_date":   bson.M{"$lt": a},
				"interval_end": bson.M{"$gt": b},
			},
			{
				"start_date": bson.M{"$gte": a, "$lte": b},
			},
		},
	}
	return filter
}
