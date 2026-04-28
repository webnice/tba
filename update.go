package tgbotapi

import "context"

// Update это ответ на обновление от функции getUpdates.
type Update struct {
	// ctx is either the client or server context. It should only
	// be modified via copying the whole Request using Clone or WithContext.
	// It is unexported to prevent people from using Context wrong
	// and mutating the contexts held by callers of the same request.
	ctx context.Context
	// UpdateID is the update's unique identifier.
	// Update identifiers start from a certain positive number and increase
	// sequentially.
	// This ID becomes especially handy if you're using Webhooks,
	// since it allows you to ignore repeated updates or to restore
	// the correct update sequence, should they get out of order.
	// If there are no new updates for at least a week, then identifier
	// of the next update will be chosen randomly instead of sequentially.
	UpdateID int64 `json:"update_id"`
	// Message new incoming message of any kind — text, photo, sticker, etc.
	//
	// optional
	Message *Message `json:"message,omitempty"`
	// EditedMessage new version of a message that is known to the bot and was
	// edited
	//
	// optional
	EditedMessage *Message `json:"edited_message,omitempty"`
	// ChannelPost new version of a message that is known to the bot and was
	// edited
	//
	// optional
	ChannelPost *Message `json:"channel_post,omitempty"`
	// EditedChannelPost new incoming channel post of any kind — text, photo,
	// sticker, etc.
	//
	// optional
	EditedChannelPost *Message `json:"edited_channel_post,omitempty"`
	// BusinessConnection the bot was connected to or disconnected from a
	// business account, or a user edited an existing connection with the bot
	//
	// optional
	BusinessConnection *BusinessConnection `json:"business_connection,omitempty"`
	// BusinessMessage is a new non-service message from a
	// connected business account
	//
	// optional
	BusinessMessage *Message `json:"business_message,omitempty"`
	// EditedBusinessMessage is a new version of a message from a
	// connected business account
	//
	// optional
	EditedBusinessMessage *Message `json:"edited_business_message,omitempty"`
	// DeletedBusinessMessages are the messages were deleted from a
	// connected business account
	//
	// optional
	DeletedBusinessMessages *BusinessMessagesDeleted `json:"deleted_business_messages,omitempty"`
	// MessageReaction is a reaction to a message was changed by a user.
	//
	// optional
	MessageReaction *MessageReactionUpdated `json:"message_reaction,omitempty"`
	// MessageReactionCount reactions to a message with anonymous reactions were changed.
	//
	// optional
	MessageReactionCount *MessageReactionCountUpdated `json:"message_reaction_count,omitempty"`
	// InlineQuery new incoming inline query
	//
	// optional
	InlineQuery *InlineQuery `json:"inline_query,omitempty"`
	// ChosenInlineResult is the result of an inline query
	// that was chosen by a user and sent to their chat partner.
	// Please see our documentation on the feedback collecting
	// for details on how to enable these updates for your bot.
	//
	// optional
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	// CallbackQuery new incoming callback query
	//
	// optional
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`
	// ShippingQuery new incoming shipping query. Only for invoices with
	// flexible price
	//
	// optional
	ShippingQuery *ShippingQuery `json:"shipping_query,omitempty"`
	// PreCheckoutQuery new incoming pre-checkout query. Contains full
	// information about checkout
	//
	// optional
	PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query,omitempty"`
	// PurchasedPaidMedia is user purchased paid media with a non-empty
	// payload sent by the bot in a non-channel chat
	//
	// optional
	PurchasedPaidMedia *PaidMediaPurchased `json:"purchased_paid_media,omitempty"`
	// Pool new poll state. Bots receive only updates about stopped polls and
	// polls, which are sent by the bot
	//
	// optional
	Poll *Poll `json:"poll,omitempty"`
	// PollAnswer user changed their answer in a non-anonymous poll. Bots
	// receive new votes only in polls that were sent by the bot itself.
	//
	// optional
	PollAnswer *PollAnswer `json:"poll_answer,omitempty"`
	// MyChatMember is the bot's chat member status was updated in a chat. For
	// private chats, this update is received only when the bot is blocked or
	// unblocked by the user.
	//
	// optional
	MyChatMember *ChatMemberUpdated `json:"my_chat_member,omitempty"`
	// ChatMember is a chat member's status was updated in a chat. The bot must
	// be an administrator in the chat and must explicitly specify "chat_member"
	// in the list of allowed_updates to receive these updates.
	//
	// optional
	ChatMember *ChatMemberUpdated `json:"chat_member,omitempty"`
	// ChatJoinRequest is a request to join the chat has been sent. The bot must
	// have the can_invite_users administrator right in the chat to receive
	// these updates.
	//
	// optional
	ChatJoinRequest *ChatJoinRequest `json:"chat_join_request,omitempty"`
	// ChatBoostUpdated represents a boost added to a chat or changed.
	//
	// optional
	ChatBoost *ChatBoostUpdated `json:"chat_boost,omitempty"`
	// RemovedChatBoost represents a boost removed from a chat.
	//
	// optional
	RemovedChatBoost *ChatBoostRemoved `json:"removed_chat_boost,omitempty"`
	// ManagedBot A new bot was created to be managed by the bot, or token or owner of a managed bot was changed.
	//
	// optional
	ManagedBot *ManagedBotUpdated `json:"managed_bot,omitempty"`
}

// Context Возвращает контекст. Чтобы изменить контекст, используйте [Update.WithContext].
//
// Возвращаемый контекст всегда не равен нулю. По умолчанию используется фоновый контекст.
func (upd *Update) Context() context.Context {
	if upd.ctx != nil {
		return upd.ctx
	}
	return context.Background()
}

// WithContext возвращает не полную копию u с изменённым контекстом на ctx.
//
// Предоставленный ctx должен быть отличным от нуля, если передан nil объект, это вызовет панику.
func (upd *Update) WithContext(ctx context.Context) (u2 *Update) {
	if ctx == nil {
		panic("nil context")
	}
	u2 = new(Update)
	*u2 = *upd
	u2.ctx = ctx
	return
}

// SentFrom возвращает пользователя, отправившего обновление.
//
// Может быть равно нулю, если Telegram не предоставил информацию о пользователе в объекте обновления.
func (upd *Update) SentFrom() *User {
	switch {
	case upd.Message != nil:
		return upd.Message.From
	case upd.EditedMessage != nil:
		return upd.EditedMessage.From
	case upd.InlineQuery != nil:
		return upd.InlineQuery.From
	case upd.ChosenInlineResult != nil:
		return upd.ChosenInlineResult.From
	case upd.CallbackQuery != nil:
		return upd.CallbackQuery.From
	case upd.ShippingQuery != nil:
		return upd.ShippingQuery.From
	case upd.PreCheckoutQuery != nil:
		return upd.PreCheckoutQuery.From
	default:
		return nil
	}
}

// CallbackData возвращает данные запроса обратного вызова, если они существуют.
func (upd *Update) CallbackData() string {
	if upd.CallbackQuery != nil {
		return upd.CallbackQuery.Data
	}
	return ""
}

// FromChat возвращает чат, в котором произошло обновление.
func (upd *Update) FromChat() *Chat {
	switch {
	case upd.Message != nil:
		return &upd.Message.Chat
	case upd.EditedMessage != nil:
		return &upd.EditedMessage.Chat
	case upd.ChannelPost != nil:
		return &upd.ChannelPost.Chat
	case upd.EditedChannelPost != nil:
		return &upd.EditedChannelPost.Chat
	case upd.CallbackQuery != nil && upd.CallbackQuery.Message != nil:
		return &upd.CallbackQuery.Message.Chat
	default:
		return nil
	}
}
