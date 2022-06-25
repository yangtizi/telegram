package tg

/*
Message
This object represents a message.

Field	Type	Description
message_id	Integer	Unique message identifier inside this chat
from	User	Optional. Sender of the message; empty for messages sent to channels. For backward compatibility, the field contains a fake sender user in non-channel chats, if the message was sent on behalf of a chat.
sender_chat	Chat	Optional. Sender of the message, sent on behalf of a chat. For example, the channel itself for channel posts, the supergroup itself for messages from anonymous group administrators, the linked channel for messages automatically forwarded to the discussion group. For backward compatibility, the field from contains a fake sender user in non-channel chats, if the message was sent on behalf of a chat.
date	Integer	Date the message was sent in Unix time
chat	Chat	Conversation the message belongs to
forward_from	User	Optional. For forwarded messages, sender of the original message
forward_from_chat	Chat	Optional. For messages forwarded from channels or from anonymous administrators, information about the original sender chat
forward_from_message_id	Integer	Optional. For messages forwarded from channels, identifier of the original message in the channel
forward_signature	String	Optional. For forwarded messages that were originally sent in channels or by an anonymous chat administrator, signature of the message sender if present
forward_sender_name	String	Optional. Sender's name for messages forwarded from users who disallow adding a link to their account in forwarded messages
forward_date	Integer	Optional. For forwarded messages, date the original message was sent in Unix time
is_automatic_forward	True	Optional. True, if the message is a channel post that was automatically forwarded to the connected discussion group
reply_to_message	Message	Optional. For replies, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
via_bot	User	Optional. Bot through which the message was sent
edit_date	Integer	Optional. Date the message was last edited in Unix time
has_protected_content	True	Optional. True, if the message can't be forwarded
media_group_id	String	Optional. The unique identifier of a media message group this message belongs to
author_signature	String	Optional. Signature of the post author for messages in channels, or the custom title of an anonymous group administrator
text	String	Optional. For text messages, the actual UTF-8 text of the message, 0-4096 characters
entities	Array of MessageEntity	Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
animation	Animation	Optional. Message is an animation, information about the animation. For backward compatibility, when this field is set, the document field will also be set
audio	Audio	Optional. Message is an audio file, information about the file
document	Document	Optional. Message is a general file, information about the file
photo	Array of PhotoSize	Optional. Message is a photo, available sizes of the photo
sticker	Sticker	Optional. Message is a sticker, information about the sticker
video	Video	Optional. Message is a video, information about the video
video_note	VideoNote	Optional. Message is a video note, information about the video message
voice	Voice	Optional. Message is a voice message, information about the file
caption	String	Optional. Caption for the animation, audio, document, photo, video or voice, 0-1024 characters
caption_entities	Array of MessageEntity	Optional. For messages with a caption, special entities like usernames, URLs, bot commands, etc. that appear in the caption
contact	Contact	Optional. Message is a shared contact, information about the contact
dice	Dice	Optional. Message is a dice with random value
game	Game	Optional. Message is a game, information about the game. More about games »
poll	Poll	Optional. Message is a native poll, information about the poll
venue	Venue	Optional. Message is a venue, information about the venue. For backward compatibility, when this field is set, the location field will also be set
location	Location	Optional. Message is a shared location, information about the location
new_chat_members	Array of User	Optional. New members that were added to the group or supergroup and information about them (the bot itself may be one of these members)
left_chat_member	User	Optional. A member was removed from the group, information about them (this member may be the bot itself)
new_chat_title	String	Optional. A chat title was changed to this value
new_chat_photo	Array of PhotoSize	Optional. A chat photo was change to this value
delete_chat_photo	True	Optional. Service message: the chat photo was deleted
group_chat_created	True	Optional. Service message: the group has been created
supergroup_chat_created	True	Optional. Service message: the supergroup has been created. This field can't be received in a message coming through updates, because bot can't be a member of a supergroup when it is created. It can only be found in reply_to_message if someone replies to a very first message in a directly created supergroup.
channel_chat_created	True	Optional. Service message: the channel has been created. This field can't be received in a message coming through updates, because bot can't be a member of a channel when it is created. It can only be found in reply_to_message if someone replies to a very first message in a channel.
message_auto_delete_timer_changed	MessageAutoDeleteTimerChanged	Optional. Service message: auto-delete timer settings changed in the chat
migrate_to_chat_id	Integer	Optional. The group has been migrated to a supergroup with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
migrate_from_chat_id	Integer	Optional. The supergroup has been migrated from a group with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
pinned_message	Message	Optional. Specified message was pinned. Note that the Message object in this field will not contain further reply_to_message fields even if it is itself a reply.
invoice	Invoice	Optional. Message is an invoice for a payment, information about the invoice. More about payments »
successful_payment	SuccessfulPayment	Optional. Message is a service message about a successful payment, information about the payment. More about payments »
connected_website	String	Optional. The domain name of the website on which the user has logged in. More about Telegram Login »
passport_data	PassportData	Optional. Telegram Passport data
proximity_alert_triggered	ProximityAlertTriggered	Optional. Service message. A user in the chat triggered another user's proximity alert while sharing Live Location.
video_chat_scheduled	VideoChatScheduled	Optional. Service message: video chat scheduled
video_chat_started	VideoChatStarted	Optional. Service message: video chat started
video_chat_ended	VideoChatEnded	Optional. Service message: video chat ended
video_chat_participants_invited	VideoChatParticipantsInvited	Optional. Service message: new participants invited to a video chat
web_app_data	WebAppData	Optional. Service message: data sent by a Web App
reply_markup	InlineKeyboardMarkup	Optional. Inline keyboard attached to the message. login_url buttons are represented as ordinary url buttons.
*/
type TMessage struct {
	Message_id             int       `json:"message_id"`
	From                   *TUser    `json:"from"`
	SenderChat             *TChat    `json:"sender_chat"`
	Date                   int       `json:"date"`
	Chat                   *TChat    `json:"chat"`
	ForwardFrom            *TUser    `json:"forward_from"`
	ForwardFrom_chat       *TChat    `json:"forward_from_chat"`
	ForwardFrom_message_id int       `json:"forward_from_message_id"`
	ForwardSignature       string    `json:"forward_signature"`
	ForwardSender_name     string    `json:"forward_sender_name"`
	ForwardDate            int       `json:"forward_date"`
	IsAutomaticForward     bool      `json:"is_automatic_forward"`
	ReplyToMessage         *TMessage `json:"reply_to_message"`
	ViaBot                 *TUser    `json:"via_bot"`
	EditDate               int       `json:"edit_date"`
	HasProtected_content   bool      `json:"has_protected_content"`
	MediaGroupID           string    `json:"media_group_id"`
	AuthorSignature        string    `json:"author_signature"`
	Text                   string    `json:"text"`
	// Entities                      *TArray                         `json:"entities"`
	// Animation                     *TAnimation                     `json:"animation"`
	// Audio                         *TAudio                         `json:"audio"`
	// Document                      *TDocument                      `json:"document"`
	// Photo                         *TArray                         `json:"photo"`
	// Sticker                       *TSticker                       `json:"sticker"`
	// Video                         *TVideo                         `json:"video"`
	// VideoNote                     *TVideoNote                     `json:"video_note"`
	// Voice                         *TVoice                         `json:"voice"`
	// Caption                       string                          `json:"caption"`
	// CaptionEntities               *TArray                         `json:"caption_entities"`
	// Contact                       *TContact                       `json:"contact"`
	// Dice                          *TDice                          `json:"dice"`
	// Game                          *TGame                          `json:"game"`
	// Poll                          *TPoll                          `json:"poll"`
	// Venue                         *TVenue                         `json:"venue"`
	// Location                      *TLocation                      `json:"location"`
	// NewChatMembers                *TArray                         `json:"new_chat_members"`
	// LeftChatMember                *TUser                          `json:"left_chat_member"`
	// NewChatTitle                  string                          `json:"new_chat_title"`
	// NewChatPhoto                  *TArray                         `json:"new_chat_photo"`
	// DeleteChatPhoto               bool                          `json:"delete_chat_photo"`
	// GroupChatCreated              bool                          `json:"group_chat_created"`
	// SupergroupChatCreated         bool                          `json:"supergroup_chat_created"`
	// ChannelChatCreated            bool                          `json:"channel_chat_created"`
	// MessageAutoDeleteTimerChanged *TMessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed"`
	// MigrateToChatID               int                             `json:"migrate_to_chat_id"`
	// MigrateFromChatID             int                             `json:"migrate_from_chat_id"`
	// PinnedMessage                 *TMessage                       `json:"pinned_message"`
	// Invoice                       *TInvoice                       `json:"invoice"`
	// SuccessfulPayment             *TSuccessfulPayment             `json:"successful_payment"`
	// ConnectedWebsite              string                          `json:"connected_website"`
	// PassportData                  *TPassportData                  `json:"passport_data"`
	// ProximityAlertTriggered       *TProximityAlertTriggered       `json:"proximity_alert_triggered"`
	// VideoChatScheduled            *TVideoChatScheduled            `json:"video_chat_scheduled"`
	// VideoChatStarted              *TVideoChatStarted              `json:"video_chat_started"`
	// VideoChatEnded                *TVideoChatEnded                `json:"video_chat_ended"`
	// VideoChatParticipants_invited *TVideoChatParticipantsInvited  `json:"video_chat_participants_invited"`
	// WebAppData                    *TWebAppData                    `json:"web_app_data"`
	// ReplyMarkup                   *TInlineKeyboardMarkup          `json:"reply_markup"`
}
