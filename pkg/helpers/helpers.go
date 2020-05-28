package helpers

import (
	"fmt"
	"time"

	discord "github.com/jadech32/discordbot/pkg/embed"
)

// FooterText returns the text for the Embed Footer
func FooterText() string {
	return fmt.Sprintf("Discord Bot | %s", time.Now().Format("Jan 2 2006 15:04:05"))
}

// EmbedTemplate returns the basic template for an Embed
func EmbedTemplate() *discord.Embed {
	return discord.NewEmbed().
		SetFooterText(FooterText()).
		SetColor(16291071)
}

// TimeIn returns the time in UTC if the name is "" or "UTC".
// It returns the local time if the name is "Local".
// Otherwise, the name is taken to be a location name in
// the IANA Time Zone database, such as "Africa/Lagos".
func TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}
