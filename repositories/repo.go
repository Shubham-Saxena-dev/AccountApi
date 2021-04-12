package repositories

import (
	"GolangProject/apis"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories interface {
	GetAccounts() ([]apis.AccountCreateRequest, error)
	GetAccount(Id string) (apis.AccountCreateRequest, error)
	CreateAccount(request apis.AccountCreateRequest) (apis.AccountCreateRequest, error)
	DeleteAccount(Id string) error
	UpdateAccount(Id string, request apis.AccountUpdateRequest) (apis.AccountCreateRequest, error)
}

type repositories struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewAccountRepositories(collection *mongo.Collection, ctx context.Context) Repositories {
	return &repositories{
		collection: collection,
		ctx:        ctx,
	}
}

func (r *repositories) GetAccounts() ([]apis.AccountCreateRequest, error) {
	cursor, err := r.collection.Find(r.ctx, bson.D{})
	defer cursor.Close(r.ctx)
	if err != nil {
		return []apis.AccountCreateRequest{}, err
	}
	var accounts []apis.AccountCreateRequest

	if cursor.All(r.ctx, &accounts); err != nil {
		return []apis.AccountCreateRequest{}, err
	}

	return accounts, nil
}

func (r *repositories) GetAccount(Id string) (apis.AccountCreateRequest, error) {
	account := apis.AccountCreateRequest{}
	uid, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return apis.EmptyCreateAccount, err
	}
	err = r.collection.FindOne(r.ctx, bson.M{
		"_id": uid,
	}).Decode(&account)
	if err != nil {
		return apis.EmptyCreateAccount, err
	}
	return account, nil
}

func (r *repositories) CreateAccount(request apis.AccountCreateRequest) (apis.AccountCreateRequest, error) {
	request.Id = primitive.NewObjectID()
	_, err := r.collection.InsertOne(r.ctx, request)

	if err != nil {
		return apis.EmptyCreateAccount, err
	}
	return apis.AccountCreateRequest{
		Id: request.Id,
	}, nil
}

func (r *repositories) DeleteAccount(Id string) error {
	uid, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return err
	}
	_, err = r.collection.DeleteOne(r.ctx, bson.M{
		"_id": uid,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *repositories) UpdateAccount(Id string, request apis.AccountUpdateRequest) (apis.AccountCreateRequest, error) {
	uid, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return apis.EmptyCreateAccount, err
	}
	_, err = r.collection.UpdateOne(r.ctx, bson.M{
		"_id": uid,
	}, bson.M{
		"$set": request,
	})
	if err != nil {
		return apis.EmptyCreateAccount, err
	}
	return r.GetAccount(Id)
}
