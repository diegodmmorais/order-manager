package gateway

type AddressInputMapper struct {
	Street  string `bson:"street,omitempty"`
	Number  string `bson:"number,omitempty"`
	Zipcode string `bson:"zip_code,omitempty"`
	City    string `bson:"city,omitempty"`
}
