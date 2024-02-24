package wschat

import (
	"testing"

	"blinders/packages/db"
	"blinders/packages/db/models"
	"blinders/packages/session"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	if app == nil {
		InitApp(
			session.NewManager(redis.NewClient(&redis.Options{Addr: "localhost:6379"})),
			db.NewMongoManager("mongodb://localhost:27017", "blinders"),
		)
	}
}

func TestSendMessageFailedWithWrongPayload(t *testing.T) {
	_, err := HandleSendMessage(
		primitive.NewObjectID().Hex(),
		primitive.NewObjectID().Hex(),
		UserSendMessagePayload{
			ChatEvent:      ChatEvent{Type: UserSendMessage},
			Content:        "hello world",
			ConversationID: "wrongID",
			ResolveID:      "resolveID",
		})

	assert.NotNil(t, err)

	_, err = HandleSendMessage(
		primitive.NewObjectID().Hex(),
		primitive.NewObjectID().Hex(),
		UserSendMessagePayload{
			ChatEvent:      ChatEvent{Type: UserSendMessage},
			Content:        "hello world",
			ConversationID: primitive.NewObjectID().Hex(),
			ResolveID:      "resolveID",
		})

	assert.NotNil(t, err)
}

func TestSendMessageFailedWithConversationNotFound(t *testing.T) {
	_, err := HandleSendMessage(
		primitive.NewObjectID().Hex(),
		primitive.NewObjectID().Hex(),
		UserSendMessagePayload{
			ChatEvent:      ChatEvent{Type: UserSendMessage},
			Content:        "hello world",
			ConversationID: primitive.NewObjectID().Hex(),
		})

	assert.NotNil(t, err)
}

func TestSendMessageFailedWithUserIsNotMember(t *testing.T) {
	conversation, _ := app.DB.Conversations.InsertNewConversation(models.Conversation{})
	_, err := HandleSendMessage(
		primitive.NewObjectID().Hex(),
		primitive.NewObjectID().Hex(),
		UserSendMessagePayload{
			ChatEvent:      ChatEvent{Type: UserSendMessage},
			Content:        "hello world",
			ConversationID: conversation.ID.Hex(),
		})

	assert.NotNil(t, err)
}

func TestSendMessageWithNoError(t *testing.T) {
	user, _ := app.DB.Users.InsertNewRawUser(models.User{})
	conversation, _ := app.DB.Conversations.InsertNewRawConversation(models.Conversation{
		Members: []models.Member{{UserID: user.ID}},
	})
	_, err := HandleSendMessage(
		user.ID.Hex(),
		primitive.NewObjectID().Hex(),
		UserSendMessagePayload{
			ChatEvent:      ChatEvent{Type: UserSendMessage},
			Content:        "hello world",
			ConversationID: conversation.ID.Hex(),
		})

	assert.Nil(t, err)
}

func TestSendMessageFailedWithInvalidMessageToReply(t *testing.T) {
	_, err := HandleSendMessage(
		primitive.NewObjectID().Hex(),
		primitive.NewObjectID().Hex(),
		UserSendMessagePayload{
			ChatEvent:      ChatEvent{Type: UserSendMessage},
			Content:        "hello world",
			ConversationID: primitive.NewObjectID().Hex(),
			ReplyTo:        primitive.NewObjectID().Hex(),
		})

	assert.NotNil(t, err)
}

func TestSendMessageWithValidMessageToReply(t *testing.T) {
	user, _ := app.DB.Users.InsertNewRawUser(models.User{})
	conversation, _ := app.DB.Conversations.InsertNewRawConversation(models.Conversation{
		Members: []models.Member{{UserID: user.ID}},
	})
	message, _ := app.DB.Messages.InsertNewRawMessage(models.Message{
		ConversationID: conversation.ID,
	})

	_, err := HandleSendMessage(
		user.ID.Hex(),
		primitive.NewObjectID().Hex(),
		UserSendMessagePayload{
			ChatEvent:      ChatEvent{Type: UserSendMessage},
			Content:        "hello world",
			ConversationID: conversation.ID.Hex(),
			ReplyTo:        message.ID.Hex(),
		})

	assert.Nil(t, err)
}

func TestSendMessageSuccess(t *testing.T) {
	user, _ := app.DB.Users.InsertNewRawUser(models.User{})
	conversation, _ := app.DB.Conversations.InsertNewRawConversation(models.Conversation{
		Members: []models.Member{{UserID: user.ID}},
	})
	_, err := HandleSendMessage(
		user.ID.Hex(),
		primitive.NewObjectID().Hex(),
		UserSendMessagePayload{
			ChatEvent:      ChatEvent{Type: UserSendMessage},
			Content:        "hello world",
			ConversationID: conversation.ID.Hex(),
		})

	assert.Nil(t, err)
}

func TestSendMessageWithDistribution(t *testing.T) {
	sender, _ := app.DB.Users.InsertNewRawUser(models.User{})
	recipient1, _ := app.DB.Users.InsertNewRawUser(models.User{})
	recipient2, _ := app.DB.Users.InsertNewRawUser(models.User{})
	conversation, _ := app.DB.Conversations.InsertNewRawConversation(
		models.Conversation{
			Members: []models.Member{
				{UserID: sender.ID},
				{UserID: recipient1.ID},
				{UserID: recipient2.ID},
			},
		})

	sConnID := primitive.NewObjectID().Hex()
	r1connID := primitive.NewObjectID().Hex()
	r2connID := primitive.NewObjectID().Hex()
	_ = app.Session.AddSession(recipient1.ID.Hex(), r1connID)
	_ = app.Session.AddSession(recipient2.ID.Hex(), r2connID)

	resolveID := primitive.NewObjectID().Hex()
	content := "hello world"
	dCh, err := HandleSendMessage(
		sender.ID.Hex(),
		sConnID,
		UserSendMessagePayload{
			ResolveID:      resolveID,
			ChatEvent:      ChatEvent{Type: UserSendMessage},
			Content:        content,
			ConversationID: conversation.ID.Hex(),
		})

	assert.Nil(t, err)

	expectedMap := map[string]bool{}
	for {
		de := <-dCh
		if de == nil {
			break
		}
		expectedMap[de.ConnectionID] = true
		switch de.ConnectionID {
		case sConnID:
			payload := de.Payload.(ServerAckSendMessagePayload)
			assert.Equal(t, ServerAckSendMessage, payload.ChatEvent.Type)
			assert.Equal(t, conversation.ID, payload.Message.ConversationID)
			assert.Equal(t, "", payload.Error.Error)
			assert.Equal(t, content, payload.Message.Content)
			assert.Equal(t, resolveID, payload.ResolveID)
		case r1connID:
			payload := de.Payload.(ServerSendMessagePayload)
			assert.Equal(t, ServerSendMessage, payload.ChatEvent.Type)
			assert.Equal(t, conversation.ID, payload.Message.ConversationID)
			assert.Equal(t, content, payload.Message.Content)
		case r2connID:
			payload := de.Payload.(ServerSendMessagePayload)
			assert.Equal(t, ServerSendMessage, payload.ChatEvent.Type)
			assert.Equal(t, conversation.ID, payload.Message.ConversationID)
			assert.Equal(t, content, payload.Message.Content)
		}

	}

	assert.True(t, expectedMap[r1connID])
	assert.True(t, expectedMap[r2connID])
	assert.True(t, expectedMap[sConnID])
	assert.Equal(t, 3, len(expectedMap))
}

func TestSendMessageWithDistributionWithOfflineRecipient(t *testing.T) {
	sender, _ := app.DB.Users.InsertNewRawUser(models.User{})
	recipient1, _ := app.DB.Users.InsertNewRawUser(models.User{})
	recipient2, _ := app.DB.Users.InsertNewRawUser(models.User{})
	conversation, _ := app.DB.Conversations.InsertNewRawConversation(
		models.Conversation{
			Members: []models.Member{
				{UserID: sender.ID},
				{UserID: recipient1.ID},
				{UserID: recipient2.ID},
			},
		})

	sConnID := primitive.NewObjectID().Hex()
	r1connID := primitive.NewObjectID().Hex()
	_ = app.Session.AddSession(recipient1.ID.Hex(), r1connID)

	resolveID := primitive.NewObjectID().Hex()
	content := "hello world"
	dCh, err := HandleSendMessage(
		sender.ID.Hex(),
		sConnID,
		UserSendMessagePayload{
			ResolveID:      resolveID,
			ChatEvent:      ChatEvent{Type: UserSendMessage},
			Content:        content,
			ConversationID: conversation.ID.Hex(),
		})

	assert.Nil(t, err)

	expectedMap := map[string]bool{}
	for {
		de := <-dCh
		if de == nil {
			break
		}
		expectedMap[de.ConnectionID] = true

	}

	assert.True(t, expectedMap[r1connID])
	assert.True(t, expectedMap[sConnID])
	assert.Equal(t, 2, len(expectedMap))
}

func TestSendMessageWithDistributionWithMultipleSessionsPerUser(t *testing.T) {
	sender, _ := app.DB.Users.InsertNewRawUser(models.User{})
	recipient1, _ := app.DB.Users.InsertNewRawUser(models.User{})
	recipient2, _ := app.DB.Users.InsertNewRawUser(models.User{})
	conversation, _ := app.DB.Conversations.InsertNewRawConversation(
		models.Conversation{
			Members: []models.Member{
				{UserID: sender.ID},
				{UserID: recipient1.ID},
				{UserID: recipient2.ID},
			},
		})

	sConnID := primitive.NewObjectID().Hex()
	r1connID := primitive.NewObjectID().Hex()
	r1connID2 := primitive.NewObjectID().Hex()
	_ = app.Session.AddSession(recipient1.ID.Hex(), r1connID)
	_ = app.Session.AddSession(recipient1.ID.Hex(), r1connID2)

	resolveID := primitive.NewObjectID().Hex()
	content := "hello world"
	dCh, err := HandleSendMessage(
		sender.ID.Hex(),
		sConnID,
		UserSendMessagePayload{
			ResolveID:      resolveID,
			ChatEvent:      ChatEvent{Type: UserSendMessage},
			Content:        content,
			ConversationID: conversation.ID.Hex(),
		})

	assert.Nil(t, err)

	expectedMap := map[string]bool{}
	for {
		de := <-dCh
		if de == nil {
			break
		}
		expectedMap[de.ConnectionID] = true

	}

	assert.True(t, expectedMap[r1connID])
	assert.True(t, expectedMap[r1connID2])
	assert.True(t, expectedMap[sConnID])
	assert.Equal(t, 3, len(expectedMap))
}

func TestSendMessageWithDistributionWithStoredMessage(t *testing.T) {
	sender, _ := app.DB.Users.InsertNewRawUser(models.User{})
	recipient1, _ := app.DB.Users.InsertNewRawUser(models.User{})
	recipient2, _ := app.DB.Users.InsertNewRawUser(models.User{})
	conversation, _ := app.DB.Conversations.InsertNewRawConversation(
		models.Conversation{
			Members: []models.Member{
				{UserID: sender.ID},
				{UserID: recipient1.ID},
				{UserID: recipient2.ID},
			},
		})

	sConnID := primitive.NewObjectID().Hex()
	r1connID := primitive.NewObjectID().Hex()
	_ = app.Session.AddSession(recipient1.ID.Hex(), r1connID)

	resolveID := primitive.NewObjectID().Hex()
	content := "hello world"
	dCh, err := HandleSendMessage(
		sender.ID.Hex(),
		sConnID,
		UserSendMessagePayload{
			ResolveID:      resolveID,
			ChatEvent:      ChatEvent{Type: UserSendMessage},
			Content:        content,
			ConversationID: conversation.ID.Hex(),
		})

	assert.Nil(t, err)

	var message1 models.Message
	var message2 models.Message
	for {
		de := <-dCh
		if de == nil {
			break
		}
		if de.ConnectionID == sConnID {
			message1 = de.Payload.(ServerAckSendMessagePayload).Message
		} else if de.ConnectionID == r1connID {
			message2 = de.Payload.(ServerSendMessagePayload).Message
		}

	}
	storedMessage, err := app.DB.Messages.GetMessageByID(message1.ID)
	assert.Nil(t, err)
	assert.Equal(t, storedMessage, message1)
	assert.Equal(t, storedMessage, message2)
}
