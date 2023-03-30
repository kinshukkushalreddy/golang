collection := client.Database("mydb").Collection("mycollection")
filter := bson.D{{"status", "active"}}
count, err := collection.CountDocuments(context.Background(), filter)
if err != nil {
	log.Fatal(err)
}
fmt.Println(count)

