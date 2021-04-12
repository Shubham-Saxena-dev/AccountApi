package apis

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var EmptyCreateAccount = AccountCreateRequest{}

type AccountCreateRequest struct {
	Id      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name    string             `json:"name" validate:"required,min=2,max=100"`
	Email   string             `json:"email" validate:"required,email"`
	Zipcode int32              `json:"zip_code" validate:"required,min=2,max=10"`
}

type AccountUpdateRequest struct {
	Name    string `json:"name" validate:"required,min=2,max=100"`
	Zipcode int32  `json:"zip_code" validate:"required,min=2,max=10"`
}
