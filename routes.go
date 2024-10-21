package kickcom

// Route is a selector of a specific endpoint.
type Route string

const (
	// RouteChannelsShow is a route to the endpoint to get channel info.
	RouteChannelsShow = Route("channels.show")

	// RouteChatHistoryChannelMessages is a route to the endpoint to get channel messages history.
	RouteChatHistoryChannelMessages = Route("chat-history.channel-messages")

	// RouteDeleteChatMessage is a route to the endpoint to delete a chat message.
	RouteDeleteChatMessage = Route("delete.chatmessage")

	// RouteChatRoomGetRules is a route to the endpoint to get chat room rules.
	RouteChatRoomGetRules = Route("chatroom.getRules")

	// RouteChatRoomShow is a route to the endpoint to get chat room info using API v1.
	RouteChatRoomShow = Route("chatroom.show")

	// RouteChatRoomShow is a route to the endpoint to get chat room info using API v2.
	RouteChannelChatRoom = Route("channel.chatroom")

	// RouteKickTokenCreate is a route to the endpoint to create kick token (for authentication).
	RouteKickTokenCreate = Route("kick.token.create")

	// RouteSwaggerGetDocs is a route to the endpoint to get swagger docs. IT DOES NOT WORK.
	RouteSwaggerGetDocs = Route("swagger.getDocs")
)
