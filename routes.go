package kickcom

// Route is a selector of a specific endpoint.
type Route string

const (
	// RouteSubcategoriesAll is a route to the endpoint to get all subcategories using API v1.
	RouteSubcategoriesAll = Route("subcategories.all")

	// RouteChannelsShow is a route to the endpoint to get channel info using API v1.
	RouteChannelsShow = Route("channels.show")

	// RouteChannelLivestream is a route to the endpoint to get channel livestream info using API v2.
	RouteChannelLivestream = Route("channel.livestream")

	// RouteChatHistoryChannelMessages is a route to the endpoint to get channel messages history.
	RouteChatHistoryChannelMessages = Route("chat-history.channel-messages")

	// RouteDeleteChatMessage is a route to the endpoint to delete a chat message.
	RouteDeleteChatMessage = Route("delete.chatmessage")

	// RouteChatRoomGetRules is a route to the endpoint to get chat room rules using API v2.
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
