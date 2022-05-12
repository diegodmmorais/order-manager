package order

type AddressInputMapper struct {
	Street  string `bson:"street"`
	Number  string `bson:"number"`
	Zipcode string `bson:"zip_code"`
	City    string `bson:"city"`
}
