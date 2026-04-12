package tgbotapi

// Адрес и шаблон API официального облака телеграм.
const defaultEndpointApi = "https://api.telegram.org/bot%s/%s"

// Адрес и шаблон точки доступа к файлам.
const defaultEndpointFile = "https://api.telegram.org/file/bot%s/%s"

// Постоянные значения для действий в чате.
const (
	ChatTyping          = "typing"
	ChatUploadPhoto     = "upload_photo"
	ChatRecordVideo     = "record_video"
	ChatUploadVideo     = "upload_video"
	ChatRecordVoice     = "record_voice"
	ChatUploadVoice     = "upload_voice"
	ChatUploadDocument  = "upload_document"
	ChatChooseSticker   = "choose_sticker"
	ChatFindLocation    = "find_location"
	ChatRecordVideoNote = "record_video_note"
	ChatUploadVideoNote = "upload_video_note"
)

// Постоянные значения для ParseMode в MessageConfig.
const (
	ModeMarkdown   = "Markdown"
	ModeMarkdownV2 = "MarkdownV2"
	ModeHTML       = "HTML"
)

// Постоянные значения для типов обновлений.
const (
	// UpdateTypeMessage is new incoming message of any kind — text, photo, sticker, etc.
	UpdateTypeMessage = "message"

	// UpdateTypeEditedMessage is new version of a message that is known to the bot and was edited
	UpdateTypeEditedMessage = "edited_message"

	// UpdateTypeChannelPost is new incoming channel post of any kind — text, photo, sticker, etc.
	UpdateTypeChannelPost = "channel_post"

	// UpdateTypeEditedChannelPost is new version of a channel post that is known to the bot and was edited
	UpdateTypeEditedChannelPost = "edited_channel_post"

	// UpdateTypeBusinessConnection is the bot was connected to or disconnected from a business account,
	// or a user edited an existing connection with the bot
	UpdateTypeBusinessConnection = "business_connection"

	// UpdateTypeBusinessMessage is a new non-service message from a connected business account
	UpdateTypeBusinessMessage = "business_message"

	// UpdateTypeEditedBusinessMessage is a new version of a message from a connected business account
	UpdateTypeEditedBusinessMessage = "edited_business_message"

	// UpdateTypeDeletedBusinessMessages are the messages were deleted from a connected business account
	UpdateTypeDeletedBusinessMessages = "deleted_business_messages"

	// UpdateTypeMessageReactionis is a reaction to a message was changed by a user
	UpdateTypeMessageReaction = "message_reaction"

	// UpdateTypeMessageReactionCount are reactions to a message with anonymous reactions were changed
	UpdateTypeMessageReactionCount = "message_reaction_count"

	// UpdateTypeInlineQuery is new incoming inline query
	UpdateTypeInlineQuery = "inline_query"

	// UpdateTypeChosenInlineResult i the result of an inline query that was chosen by a user and sent to their
	// chat partner. Please see the documentation on the feedback collecting for
	// details on how to enable these updates for your bot.
	UpdateTypeChosenInlineResult = "chosen_inline_result"

	// UpdateTypeCallbackQuery is new incoming callback query
	UpdateTypeCallbackQuery = "callback_query"

	// UpdateTypeShippingQuery is new incoming shipping query. Only for invoices with flexible price
	UpdateTypeShippingQuery = "shipping_query"

	// UpdateTypePreCheckoutQuery is new incoming pre-checkout query. Contains full information about checkout
	UpdateTypePreCheckoutQuery = "pre_checkout_query"

	// UpdateTypePurchasedPaidMedia is a user purchased paid media with a non-empty payload
	// sent by the bot in a non-channel chat
	UpdateTypePurchasedPaidMedia = "purchased_paid_media"

	// UpdateTypePoll is new poll state. Bots receive only updates about stopped polls and polls
	// which are sent by the bot
	UpdateTypePoll = "poll"

	// UpdateTypePollAnswer is when user changed their answer in a non-anonymous poll. Bots receive new votes
	// only in polls that were sent by the bot itself.
	UpdateTypePollAnswer = "poll_answer"

	// UpdateTypeMyChatMember is when the bot's chat member status was updated in a chat. For private chats, this
	// update is received only when the bot is blocked or unblocked by the user.
	UpdateTypeMyChatMember = "my_chat_member"

	// UpdateTypeChatMember is when the bot must be an administrator in the chat and must explicitly specify
	// this update in the list of allowed_updates to receive these updates.
	UpdateTypeChatMember = "chat_member"

	// UpdateTypeChatJoinRequest is request to join the chat has been sent.
	// The bot must have the can_invite_users administrator right in the chat to receive these updates.
	UpdateTypeChatJoinRequest = "chat_join_request"

	// UpdateTypeChatBoost is chat boost was added or changed.
	// The bot must be an administrator in the chat to receive these updates.
	UpdateTypeChatBoost = "chat_boost"

	// UpdateTypeRemovedChatBoost is boost was removed from a chat.
	// The bot must be an administrator in the chat to receive these updates.
	UpdateTypeRemovedChatBoost = "removed_chat_boost"

	// UpdateTypeManagedBot Updates about the creation of managed bots and the change of their token.
	UpdateTypeManagedBot = "managed_bot"
)
