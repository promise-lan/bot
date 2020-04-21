package events

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

var (
	emojiToRoleMap = map[string]string{
		"488170878446862338": "487767936430374914", // NodeJS
		"488172791074521089": "487767930956808217", // Python
		"487785622644195347": "487767940523884544", // C/C++
		"487785622573023232": "487768950847832066", // C#/.NET
		"487785624443682846": "487768951439360000", // Java
		"488173744658055170": "487768952647188507", // Ruby
		"487785626972848138": "487769155622273024", // Swift
		"487785624250875915": "487769324178636810", // Rust
		"487785623692902400": "487769141147598848", // Elixir
		"487785623780982805": "487768954199080961", // Erlang
		"487785622681944064": "487768953653821450", // Perl
		"702238456130174977": "702237493621162114", // Golang
		"487786316621152277": "487767938225405962", // GET TO THE CHOPPER
	}
)

func getEmojiById(guild *discordgo.Guild, id string) *discordgo.Emoji {
	for _, emoji := range guild.Emojis {
		if emoji.ID == id {
			return emoji
		}
	}
	return nil
}

func getRoleById(guild *discordgo.Guild, id string) *discordgo.Role {
	for _, role := range guild.Roles {
		if role.ID == id {
			return role
		}
	}
	return nil
}

func getRoleID(emoji string) *string {
	for emojiID, roleID := range emojiToRoleMap {
		if emojiID == emoji {
			return &roleID
		}
	}
	return nil
}

func memberHasRole(member *discordgo.Member, role string) bool {
	for _, id := range member.Roles {
		if id == role {
			return true
		}
	}
	return false
}

func Ready(s *discordgo.Session, event *discordgo.Ready) {
	log.Printf("Client is ready as %s (%s)", event.User.String(), event.User.ID)
	guild, _ := s.Guild("487730100972814337")
	for emojiId, roleId := range emojiToRoleMap {
		emoji := getEmojiById(guild, emojiId)
		role := getRoleById(guild, roleId)
		log.Printf("Emoji[%s[%s]] -> Role[%s[%s]]", emoji.Name, emoji.ID, role.Name, role.ID)
	}
}
