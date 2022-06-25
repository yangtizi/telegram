package tg

/*Chat
This object represents a chat.

Field	Type	Description
id	Integer	Unique identifier for this chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
type	String	Type of chat, can be either “private”, “group”, “supergroup” or “channel”
title	String	Optional. Title, for supergroups, channels and group chats
username	String	Optional. Username, for private chats, supergroups and channels if available
first_name	String	Optional. First name of the other party in a private chat
last_name	String	Optional. Last name of the other party in a private chat
photo	ChatPhoto	Optional. Chat photo. Returned only in getChat.
bio	String	Optional. Bio of the other party in a private chat. Returned only in getChat.
has_private_forwards	True	Optional. True, if privacy settings of the other party in the private chat allows to use tg://user?id=<user_id> links only in chats with the user. Returned only in getChat.
description	String	Optional. Description, for groups, supergroups and channel chats. Returned only in getChat.
invite_link	String	Optional. Primary invite link, for groups, supergroups and channel chats. Returned only in getChat.
pinned_message	Message	Optional. The most recent pinned message (by sending date). Returned only in getChat.
permissions	ChatPermissions	Optional. Default chat member permissions, for groups and supergroups. Returned only in getChat.
slow_mode_delay	Integer	Optional. For supergroups, the minimum allowed delay between consecutive messages sent by each unpriviledged user; in seconds. Returned only in getChat.
message_auto_delete_time	Integer	Optional. The time after which all messages sent to the chat will be automatically deleted; in seconds. Returned only in getChat.
has_protected_content	True	Optional. True, if messages from the chat can't be forwarded to other chats. Returned only in getChat.
sticker_set_name	String	Optional. For supergroups, name of group sticker set. Returned only in getChat.
can_set_sticker_set	True	Optional. True, if the bot can change the group sticker set. Returned only in getChat.
linked_chat_id	Integer	Optional. Unique identifier for the linked chat, i.e. the discussion group identifier for a channel and vice versa; for supergroups and channel chats. This identifier may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier. Returned only in getChat.
location	ChatLocation	Optional. For supergroups, the location to which the supergroup is connected. Returned only in getChat.
*/

type TChat struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	// Photo                 ChatPhoto       `json:"photo"`
	Bio                string    `json:"bio"`
	HasPrivateForwards bool      `json:"has_private_forwards"`
	Description        string    `json:"description"`
	InviteLink         string    `json:"invite_link"`
	PinnedMessage      *TMessage `json:"pinned_message"`
	// Permissions           ChatPermissions `json:"permissions"`
	SlowModeDelay         int    `json:"slow_mode_delay"`
	MessageAutoDeleteTime int    `json:"message_auto_delete_time"`
	HasProtectedContent   bool   `json:"has_protected_content"`
	StickerSetName        string `json:"sticker_set_name"`
	CanSetStickerSet      bool   `json:"can_set_sticker_set"`
	LinkedChatId          int    `json:"linked_chat_id"`
	// Location              *TChatLocation `json:"location"`
}
