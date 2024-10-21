package kickcom

import (
	"context"
	"net/http"
)

// GetChatroomRulesV2 returns the rules of the chat.
// The function is not tested.
func (k *Kick) GetChatroomRulesV2(
	ctx context.Context,
	channelSlug string,
) (*ChatroomRulesV2, error) {
	return Request[ChatroomRulesV2](
		ctx,
		k,
		http.MethodGet,
		RouteChatRoomGetRules,
		RouteVars{"channel": channelSlug},
		nil,
		NoBody,
	)
}

// ChatroomRulesV2 is a response to GetChatroomRulesV2.
type ChatroomRulesV2 struct {
	Status struct {
		Error   bool   `json:"error"`
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"status"`
	Data struct {
		Rules string `json:"rules"`
	} `json:"data"`
}
