package events

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func ReactionAdd(s *discordgo.Session, event *discordgo.MessageReactionAdd) {
	if event.GuildID == "487730100972814337" && event.ChannelID == "487767664526098442" {
		guild, _ := s.Guild(event.GuildID)
		if id := getRoleID(event.Emoji.ID); id != nil {
			if member, err := s.GuildMember(guild.ID, event.UserID); err == nil {
				if hasRole := memberHasRole(member, *id); !hasRole {
					if err := s.GuildMemberRoleAdd(guild.ID, member.User.ID, *id); err != nil {
						role := getRoleById(guild, *id)
						log.Printf("Failed to add role %s (%s) to user %s (%s): %v", role.Name, role.ID, member.User.String(), member.User.ID, err)
					}
				}
			}
		}
	}
}

func ReactionRemove(s *discordgo.Session, event *discordgo.MessageReactionRemove) {
	if event.GuildID == "487730100972814337" && event.ChannelID == "487767664526098442" {
		guild, _ := s.Guild(event.GuildID)
		if id := getRoleID(event.Emoji.ID); id != nil {
			if member, err := s.GuildMember(guild.ID, event.UserID); err == nil {
				if hasRole := memberHasRole(member, *id); hasRole {
					if err := s.GuildMemberRoleRemove(guild.ID, member.User.ID, *id); err != nil {
						role := getRoleById(guild, *id)
						log.Printf("Failed to remove role %s (%s) to user %s (%s): %v", role.Name, role.ID, member.User.String(), member.User.ID, err)
					}
				}
			}
		}
	}
}
