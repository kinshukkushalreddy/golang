collection := client.Database("mydb").Collection("mycollection")
pipeline := bson.A{
	bson.D{{"$match", bson.D{{"status", "active"}}}},
	bson.D{{"$group", bson.D{{"_id", "$age"}, {"count", bson.D{{"$sum", 1}}}}}},
}
cur, err := collection.Aggregate(context.Background(), pipeline)
if err != nil {
	log.Fatal(err)
}
defer cur.Close(context.Background())
for cur.Next(context.Background()) {
	var result bson.M
	err := cur.Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
if err := cur.Err(); err != nil {
	log.Fatal(err)
}

