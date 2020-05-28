package modules

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jadech32/discordbot/pkg/helpers"
)

// Help ..
type Help struct{}

// Run returns a discordgo.MessageEmbed for the default help message
func (h Help) Run(m *discordgo.MessageCreate, s *discordgo.Session, args []string) error {
	emb := helpers.EmbedTemplate().
		AddField("Utility Commands", "help - Display this message again.").
		AddField("Exchange", "exchange - Convert currencies e.g. `exchange 1 USD CAD`").
		MessageEmbed

	_, err := s.ChannelMessageSendEmbed(m.ChannelID, emb)
	if err != nil {
		return err
	}

	return nil
}
