collection := client.Database("mydb").Collection("mycollection")
filter := bson.D{{"name", "John Doe"}}
var result bson.M
err := collection.FindOne(context.Background(), filter).Decode(&result)
if err != nil {
	log.Fatal(err)
}
fmt.Println(result)

