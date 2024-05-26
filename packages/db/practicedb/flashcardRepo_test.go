package practicedb_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"blinders/packages/db/practicedb"
	dbutils "blinders/packages/db/utils"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoTestURL = "mongodb://localhost:27017"
	mongoTestDB  = "blinder-test"
)

func TestFlashCardsRepo(t *testing.T) {
	repo := GetTestFlashCardsRepo(t)

	var (
		userID       = primitive.NewObjectID()
		collectionID = primitive.NewObjectID()
	)

	card := &practicedb.FlashCard{
		RawModel: dbutils.RawModel{
			ID: primitive.NewObjectID(),
		},
		UserID:       userID,
		CollectionID: collectionID,
		FrontText:    "Front of the card",
		BackText:     "Back of the card",
	}

	insertedCard, err := repo.Insert(card)
	assert.Nil(t, err)

	foundCard, err := repo.GetByID(insertedCard.ID)
	assert.Nil(t, err)
	assert.Equal(t, *insertedCard, *foundCard)

	foundWithUserID, err := repo.GetFlashCardByUserID(userID)
	assert.Nil(t, err)
	assert.Contains(t, foundWithUserID, *insertedCard)

	foundWithCollectionID, err := repo.GetFlashCardsByCollectionID(collectionID)
	assert.Nil(t, err)
	assert.Contains(t, foundWithCollectionID, *insertedCard)

	collection, err := repo.GetFlashCardCollectionByID(collectionID)
	assert.Nil(t, err)
	assert.Contains(t, collection.FlashCards, *insertedCard)
	assert.Equal(t, collection.ID, collectionID)

	newCard := &practicedb.FlashCard{
		RawModel: dbutils.RawModel{
			ID: insertedCard.ID,
		},
		UserID:       insertedCard.UserID,
		CollectionID: insertedCard.CollectionID,
		FrontText:    "Updated front of the card",
		BackText:     "Updated back of the card",
	}
	err = repo.UpdateFlashCard(insertedCard.ID, newCard)
	assert.Nil(t, err)

	err = repo.UpdateFlashCard(primitive.NilObjectID, newCard)
	assert.NotNil(t, err)

	updatedCard, err := repo.GetByID(insertedCard.ID)
	assert.Nil(t, err)
	assert.Equal(t, *newCard, *updatedCard)

	err = repo.DeleteFlashCardByID(insertedCard.ID)
	assert.Nil(t, err)

	deletedCard, err := repo.GetByID(insertedCard.ID)
	assert.NotNil(t, err)
	assert.Nil(t, deletedCard)
	CleanRepo(t, repo.Collection)
}

func TestGetFlashCardCollectionsByUserID(t *testing.T) {
	repo := GetTestFlashCardsRepo(t) // Assuming you have a function to create a new repo
	userID := primitive.NewObjectID()

	collectionsID := []primitive.ObjectID{
		primitive.NewObjectID(),
		primitive.NewObjectID(),
		primitive.NewObjectID(),
	}

	cards := []practicedb.FlashCard{}

	for i := 0; i < 100; i++ {
		card := &practicedb.FlashCard{
			UserID:       userID,
			CollectionID: collectionsID[i%len(collectionsID)],
			FrontText:    fmt.Sprintf("sample front text %d", i),
			BackText:     fmt.Sprintf("sample back text %d", i),
		}
		insertedCard, err := repo.InsertRaw(card)
		assert.NoError(t, err)
		assert.NotNil(t, insertedCard)
		assert.NotNil(t, insertedCard.ID)
		cards = append(cards, *insertedCard)
	}

	result, err := repo.GetFlashCardCollectionsByUserID(userID)
	assert.NoError(t, err)

	// Verify that the returned collections are the ones we used
	assert.Equal(t, len(collectionsID), len(result))
	for _, collection := range result {
		assert.Contains(t, collectionsID, collection.ID)
		assert.Equal(t, userID, collection.UserID)
		for _, card := range collection.FlashCards {
			assert.Contains(t, cards, card)
		}
	}

	// verify that the card belongs to correct collection
	for _, card := range cards {
		for _, collection := range result {
			if collection.ID == card.CollectionID {
				assert.Contains(t, collection.FlashCards, card)
			} else {
				assert.NotContains(t, collection.FlashCards, card)
			}
		}
	}

	deleteCollection := collectionsID[0]
	// verify that delete collection works
	err = repo.DeleteCardCollectionByID(deleteCollection)
	assert.Nil(t, err)

	// verify that delete not existed collection works
	err = repo.DeleteCardCollectionByID(deleteCollection)
	assert.NotNil(t, err)

	// verify that the collection is deleted
	collection, err := repo.GetFlashCardCollectionByID(deleteCollection)
	assert.NotNil(t, err)
	assert.Nil(t, collection)

	collectionID := collectionsID[1]
	collectionCards, err := repo.GetFlashCardsByCollectionID(collectionID)
	assert.Nil(t, err)

	for _, card := range collectionCards {
		assert.Contains(t, cards, card)
		assert.Equal(t, card.CollectionID, collectionID)
	}
	CleanRepo(t, repo.Collection)
}

func GetTestFlashCardsRepo(t *testing.T) *practicedb.FlashCardsRepo {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoTestURL))
	assert.Nil(t, err)

	db := client.Database(mongoTestDB)

	return practicedb.NewFlashCardRepo(db)
}
