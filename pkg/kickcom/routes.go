package kickcom

type Route string

const (
	RouteChannelsShow               = Route("channels.show")
	RouteChatHistoryChannelMessages = Route("chat-history.channel-messages")
	RouteDeleteChatMessage          = Route("delete.chatmessage")
	RouteChatRoomGetRules           = Route("chatroom.getRules")
	RouteChatRoomShow               = Route("chatroom.show")
	RouteChannelChatRoom            = Route("channel.chatroom")
	RouteKickTokenCreate            = Route("kick.token.create")
)
