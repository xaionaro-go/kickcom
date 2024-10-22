package kickcom

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// ChatMessageSenderV2 is a representation of the sender of a chat
// message in API v2.
type ChatMessageSenderV2 struct {
	ID       int        `json:"id"`
	Slug     string     `json:"slug"`
	Username string     `json:"username"`
	Identity IdentityV2 `json:"identity"`
}

// ChatMessageSenderV2 is a representation of a chat message in API v2.
type ChatMessageV2 struct {
	ID        string              `json:"id"`
	ChatID    int                 `json:"chat_id"`
	UserID    int                 `json:"user_id"`
	Content   string              `json:"content"`
	Type      string              `json:"type"`
	Metadata  any                 `json:"metadata"`
	CreatedAt time.Time           `json:"created_at"`
	Sender    ChatMessageSenderV2 `json:"sender"`
}

// ChatMessageSenderV2 is response in API v2 on a request to provide
// chat messages.
type ChatMessagesV2Reply struct {
	Status struct {
		Error   bool   `json:"error"`
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"status"`
	Data struct {
		Messages      []ChatMessageV2 `json:"messages"`
		Cursor        string          `json:"cursor"`
		PinnedMessage any             `json:"pinned_message"`
	} `json:"data"`
}

// GetChatMessagesV2 requests a list of messages from the channel using API v2.
func (k *Kick) GetChatMessagesV2(
	ctx context.Context,
	channelID uint64,
	cursor uint64,
) (*ChatMessagesV2Reply, error) {
	uriValues := url.Values{}
	if cursor != 0 {
		// this does not work for an unknown reason:
		// if you set a correct value of a cursor, you just always get an empty response.
		uriValues.Set("cursor", fmt.Sprintf("%d", cursor))
	}
	return Request[ChatMessagesV2Reply](
		ctx,
		k,
		http.MethodGet,
		RouteChatHistoryChannelMessages,
		RouteVars{"channelId": channelID},
		uriValues,
		NoBody,
	)
}

/*
func (k *Kick) TBD_SendChatMessageV2(
	ctx context.Context,
) {
	return Request[SendChatMessageReply](
		ctx,
		k,
		http.MethodPost,
		fmt.Sprintf("api/v2/messages/send/%d", chatRoomID),
		SendChatMessageRequest{
		},
	)
}
*/

// DeleteChatMessageReply is a response to a request to delete a chat
// message.
type DeleteChatMessageReply struct{}

// TBDDeleteChatMessage is a To-Be-Developed (TBD) function to delete
// a chat message. DO NOT USE THIS FUNCTION.
func (k *Kick) TBDDeleteChatMessage(
	ctx context.Context,
	chatRoomID uint64,
	messageID uint64,
) (*DeleteChatMessageReply, error) {
	return Request[DeleteChatMessageReply](
		ctx,
		k,
		http.MethodPost,
		RouteDeleteChatMessage,
		RouteVars{
			"chatroomId": chatRoomID,
			"messageId":  messageID,
		},
		nil,
		NoBody,
	)
}
