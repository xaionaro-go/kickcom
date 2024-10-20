package kickcom

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type ChatMessageSenderV2 struct {
	ID       int      `json:"id"`
	Slug     string   `json:"slug"`
	Username string   `json:"username"`
	Identity Identity `json:"identity"`
}

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

func (k *Kick) GetChatMessagesV2(
	ctx context.Context,
	channelID uint64,
	cursor uint64,
) (*ChatMessagesV2Reply, error) {
	uriValues := url.Values{}
	if cursor != 0 {
		uriValues.Set("cursor", fmt.Sprintf("%d", cursor))
	}
	return Request[ChatMessagesV2Reply](
		ctx,
		k,
		http.MethodGet,
		fmt.Sprintf("api/v2/channels/%d/messages", channelID),
		uriValues,
		NoBody,
	)
}

/*
func (k *Kick) SendChatMessageV2(
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

type DeleteChatMessageReply struct{}

func (k *Kick) DeleteChatMessage(
	ctx context.Context,
	chatRoomID uint64,
	messageID uint64,
) (*DeleteChatMessageReply, error) {
	return Request[DeleteChatMessageReply](
		ctx,
		k,
		http.MethodPost,
		fmt.Sprintf("api/v2/chatrooms/%d/messages/%d", chatRoomID, messageID),
		nil,
		NoBody,
	)
}
