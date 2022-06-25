package tg

// Update
// This object represents an incoming update.
// At most one of the optional parameters can be present in any given update.

// Field	Type	Description
// update_id	Integer	The update's unique identifier. Update identifiers start from a certain positive number and increase sequentially. This ID becomes especially handy if you're using Webhooks, since it allows you to ignore repeated updates or to restore the correct update sequence, should they get out of order. If there are no new updates for at least a week, then identifier of the next update will be chosen randomly instead of sequentially.
// message	Message	Optional. New incoming message of any kind — text, photo, sticker, etc.
// edited_message	Message	Optional. New version of a message that is known to the bot and was edited
// channel_post	Message	Optional. New incoming channel post of any kind — text, photo, sticker, etc.
// edited_channel_post	Message	Optional. New version of a channel post that is known to the bot and was edited
// inline_query	InlineQuery	Optional. New incoming inline query
// chosen_inline_result	ChosenInlineResult	Optional. The result of an inline query that was chosen by a user and sent to their chat partner. Please see our documentation on the feedback collecting for details on how to enable these updates for your bot.
// callback_query	CallbackQuery	Optional. New incoming callback query
// shipping_query	ShippingQuery	Optional. New incoming shipping query. Only for invoices with flexible price
// pre_checkout_query	PreCheckoutQuery	Optional. New incoming pre-checkout query. Contains full information about checkout
// poll	Poll	Optional. New poll state. Bots receive only updates about stopped polls and polls, which are sent by the bot
// poll_answer	PollAnswer	Optional. A user changed their answer in a non-anonymous poll. Bots receive new votes only in polls that were sent by the bot itself.
// my_chat_member	ChatMemberUpdated	Optional. The bot's chat member status was updated in a chat. For private chats, this update is received only when the bot is blocked or unblocked by the user.
// chat_member	ChatMemberUpdated	Optional. A chat member's status was updated in a chat. The bot must be an administrator in the chat and must explicitly specify “chat_member” in the list of allowed_updates to receive these updates.
// chat_join_request	ChatJoinRequest	Optional. A request to join the chat has been sent. The bot must have the can_invite_users administrator right in the chat to receive these updates.

type TUpdate struct {
	UpdateID          int       `json:"update_id"` // 更新的唯一标识符。更新标识符从某个正数开始，并按顺序递增。如果您使用的是Webhook，则此ID变得特别方便，因为它允许您忽略重复的更新或恢复正确的更新序列，如果它们出现故障。如果至少一周内没有新更新，则将随机选择下一次更新的标识符，而不是按顺序选择。
	Message           *TMessage `json:"message"`
	EditedMessage     *TMessage `json:"edited_message"`
	ChannelPost       *TMessage `json:"channel_post"`
	EditedChannelPost *TMessage `json:"edited_channel_post"`
	// InlineQuery        *TInlineQuery        `json:"inline_query"`
	// ChosenInlineResult *TChosenInlineResult `json:"chosen_inline_result"`
	// CallbackQuery      *TCallbackQuery      `json:"callback_query"`
	// ShippingQuer       *TShippingQuery      `json:"shipping_query"`
	// PreCheckoutQuery   *TPreCheckoutQuery   `json:"pre_checkout_query"`
	// Poll               *TPoll               `json:"poll"`
	// PollAnswer         *TPollAnswer         `json:"poll_answer"`
	// MyChatMember       *TChatMemberUpdated  `json:"my_chat_member"`
	// ChatMember         *TChatMemberUpdated  `json:"chat_member"`
	// ChatJoinRequest    *TChatJoinRequest    `json:"chat_join_request"`
}
