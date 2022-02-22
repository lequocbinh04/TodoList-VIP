package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"TodoList/database"
	"TodoList/graph/generated"
	"TodoList/graph/model"
	"TodoList/helpers"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *queryResolver) Account(ctx context.Context) (*model.Creator, error) {
	c, err := helpers.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	tmpCurrentUser, _ := c.Get("logged_address")
	currentAddress := tmpCurrentUser.(string)
	addressCollection := database.ConnectCollection("address")

	cnt, err := addressCollection.CountDocuments(ctx, bson.M{"address": currentAddress})
	if err != nil {
		return nil, err
	}
	if cnt == 0 {
		defaultCreator := &model.Creator{
			Address:     currentAddress,
			AmountToken: 0,
		}
		_, err := addressCollection.InsertOne(ctx, defaultCreator)
		if err != nil {
			return nil, err
		}
		return defaultCreator, nil
	}
	// find the creator and return
	creator := &model.Creator{}
	err = addressCollection.FindOne(ctx, bson.M{"address": currentAddress}).Decode(creator)
	if err != nil {
		return nil, err
	}
	return creator, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
