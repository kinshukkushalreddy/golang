collection := client.Database("mydb").Collection("mycollection")
filter := bson.D{{"age", bson.D{{"$gt", 25}}}}
cur, err := collection.Find(context.Background(), filter)
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

