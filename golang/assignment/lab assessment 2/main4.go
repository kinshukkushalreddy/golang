collection := client.Database("mydb").Collection("mycollection")
filter := bson.D{{"status", "active"}}
fieldName := "age"
values, err := collection.Distinct(context.Background(), fieldName, filter)
if err != nil {
	log.Fatal(err)
}
fmt.Println(values)

