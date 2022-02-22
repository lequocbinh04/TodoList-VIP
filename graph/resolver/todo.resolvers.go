package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"TodoList/database"
	"TodoList/db_model"
	"TodoList/graph/generated"
	"TodoList/graph/model"
	"TodoList/helpers"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, text string) (*model.Todo, error) {
	c, err := helpers.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	tmpCurrentUser, _ := c.Get("logged_address")
	currentAddress := tmpCurrentUser.(string)
	addressCollection := database.ConnectCollection("address")
	creator := &model.Creator{}
	err = addressCollection.FindOne(ctx, bson.M{"address": currentAddress}).Decode(creator)
	if err != nil {
		return nil, err
	}
	newNote := db_model.Note{
		ID:        primitive.NewObjectID(),
		Text:      text,
		Done:      false,
		CreatedBy: currentAddress,
		CreatedAt: time.Now(),
	}
	noteCollection := database.ConnectCollection("notes")
	_, err = noteCollection.InsertOne(ctx, newNote)
	if err != nil {
		return nil, err
	}
	creator.AmountToken += 1
	_, err = addressCollection.UpdateOne(ctx, bson.M{"address": currentAddress}, bson.M{"$set": creator})
	if err != nil {
		return nil, err
	}
	return &model.Todo{
		ID:      newNote.ID.Hex(),
		Text:    newNote.Text,
		Done:    newNote.Done,
		Creator: creator,
	}, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (bool, error) {
	c, err := helpers.GinContextFromContext(ctx)
	if err != nil {
		return false, err
	}
	tmpCurrentUser, _ := c.Get("logged_address")
	currentAddress := tmpCurrentUser.(string)

	noteCollection := database.ConnectCollection("notes")
	// delete note with id and created_by currentAddress
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID, "created_by": currentAddress}
	// check note with filter is exists
	count, err := noteCollection.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	_, err = noteCollection.DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, state bool, text string, id string) (*model.Todo, error) {
	c, err := helpers.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	tmpCurrentUser, _ := c.Get("logged_address")
	currentAddress := tmpCurrentUser.(string)
	noteCollection := database.ConnectCollection("notes")
    objID, _ := primitive.ObjectIDFromHex(id)
    filter := bson.M{"_id": objID, "created_by": currentAddress}
    count, err := noteCollection.CountDocuments(ctx, filter)
    if err != nil {
        return nil, err
    }
    if count == 0 {
        return nil, fmt.Errorf("not found")
    }
    if text != "" {
        _, err = noteCollection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"text": text, "done": state}})
        if err != nil {
            return nil, err
        }
    } else {
        _, err = noteCollection.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"done": state}})
        if err != nil {
            return nil, err
        }
    }
    var todo db_model.Note
    err = noteCollection.FindOne(ctx, filter).Decode(&todo)
    if err != nil {
        return nil, err
    }
        
    creator := &model.Creator{}
    addressCollection := database.ConnectCollection("address")
    err = addressCollection.FindOne(ctx, bson.M{"address": currentAddress}).Decode(creator)
    if err != nil {
        return nil, err
    }
    return &model.Todo{
        ID:      id,
        Text:    todo.Text,
        Done:    todo.Done,
        Creator: creator,
    }, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	c, err := helpers.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	tmpCurrentUser, _ := c.Get("logged_address")
	currentAddress := tmpCurrentUser.(string)
	noteCollection := database.ConnectCollection("notes")
	// find not create by currentAddress
	filter := bson.M{"created_by": currentAddress}
	cursor, err := noteCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var todos []*model.Todo
	for cursor.Next(ctx) {
		var todo db_model.Note
		err := cursor.Decode(&todo)
		if err != nil {
			return nil, err
		}
		creator := &model.Creator{}
		addressCollection := database.ConnectCollection("address")
		err = addressCollection.FindOne(ctx, bson.M{"address": todo.CreatedBy}).Decode(creator)
		if err != nil {
			return nil, err
		}
		todos = append(todos, &model.Todo{
			ID:      todo.ID.Hex(),
			Text:    todo.Text,
			Done:    todo.Done,
			Creator: creator,
		})
	}
	return todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) DoneTodo(ctx context.Context, id string) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}
