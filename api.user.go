package tg

/*
User
This object represents a Telegram user or bot.

Field	Type	Description
id	in	Unique identifier for this user or bot. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
is_bot	Boolean	True, if this user is a bot
first_name	string	User's or bot's first name
last_name	String	Optional. User's or bot's last name
username	String	Optional. User's or bot's username
language_code	String	Optional. IETF language tag of the user's language
can_join_groups	Boolean	Optional. True, if the bot can be invited to groups. Returned only in getMe.
can_read_all_group_messages	Boolean	Optional. True, if privacy mode is disabled for the bot. Returned only in getMe.
supports_inline_queries	Boolean	Optional. True, if the bot supports inline queries. Returned only in getMe.
*/
type TUser struct {
	ID                      int    `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name"`
	Username                string `json:"username"`
	LanguageCode            string `json:"language_code"`
	CanJoinGroups           bool   `json:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`
}
