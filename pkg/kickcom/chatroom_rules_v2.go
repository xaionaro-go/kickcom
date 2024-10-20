package kickcom

import (
	"context"
	"fmt"
	"net/http"
)

func (k *Kick) GetChatroomRulesV2(
	ctx context.Context,
	channelSlug string,
) (*ChatroomRulesV2, error) {
	return Request[ChatroomRulesV2](
		ctx,
		k,
		http.MethodGet,
		fmt.Sprintf("/api/v2/channels/%s/chatroom/rules", channelSlug),
		nil,
		NoBody,
	)
}

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
