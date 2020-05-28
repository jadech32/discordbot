package discord

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/bwmarrin/discordgo"
	"github.com/jadech32/discordbot/pkg/modules"
)

// Discord ..
// Every module will need to pass around this struct to add Message
type Discord struct {
	client *discordgo.Session
	prefix string
	// Map of commandType and their associated messageIDs for said type
	modules *modules.Modules
}

// NewDiscord returns a Discord Client
// Default prefix: !
func NewDiscord(token string, modules *modules.Modules) (*Discord, error) {
	tokenStr := fmt.Sprintf("Bot %s", token)
	cl, err := discordgo.New(tokenStr)
	if err != nil {
		return nil, errors.WithMessage(err, "Unable to Create Discord Session")
	}
	d := &Discord{
		client:  cl,
		modules: modules,
		prefix:  "!",
	}

	return d, nil
}

// InitModules registers a handler to the bot.
// Only when the prefix is present
func (d *Discord) InitModules() {
	prefixLen := len(d.prefix)
	d.client.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if len(m.Content) < prefixLen {
			return
		}

		// If prefix is not existant in incoming message
		if string(m.Content[0:prefixLen]) != d.prefix {
			return
		}

		commandAndArgs := strings.TrimSpace(m.Content[prefixLen:])
		commandAndArgsArray := strings.Split(commandAndArgs, " ")
		argsOnly := commandAndArgsArray[1:]
		if strings.Join(commandAndArgsArray, "") == "" {
			return
		}

		commandName := commandAndArgsArray[0]
		// If the specific command doesn't exist
		if _, ok := d.modules.ModuleStore[commandName]; !ok {
			return
		}

		// Send args to the Command
		err := d.modules.ModuleStore[commandName].Run(m, s, argsOnly)

		if err != nil {
			log.Errorf("Error for %s Command: %v", commandName, err)
			return
		}

		log.Infof("%+v", d.modules.MessageStore)
	})
}

// InitHandlers adds all handlers to the discord session
func (d *Discord) InitHandlers() {
	for _, handler := range d.modules.HandlerStore {
		d.client.AddHandler(handler)
	}
}

// SetPrefix overrides the default prefix of "!". Must be used before InitHandler()
func (d *Discord) SetPrefix(prefix string) {
	d.prefix = prefix
}

// Start opens a websocket connection to Discord and begins listening
func (d *Discord) Start() {
	d.client.Open()
}

// Close shuts down the websocket connection to Discord
func (d *Discord) Close() {
	d.client.Close()
}
