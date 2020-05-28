package modules

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"

	"github.com/bwmarrin/discordgo"
	"github.com/jadech32/discordbot/pkg/helpers"
	log "github.com/sirupsen/logrus"
)

// Convert ...
type Convert struct{}

// Run converts Float value from currency to anotherk
func (c Convert) Run(m *discordgo.MessageCreate, s *discordgo.Session, args []string) error {
	if len(args) < 3 {
		return nil
	}
	// See if currencies are valid ones - handle error from API
	notFloat := false
	amount := args[0]
	from := strings.ToUpper(args[1])
	to := strings.ToUpper(args[2])
	URL := fmt.Sprintf("https://api.exchangeratesapi.io/latest?symbols=%s&base=%s", to, from)
	resp, err := http.Get(URL)
	if err != nil {
		log.Errorf("Failed Currency Request: %s", err)
	}

	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		notFloat = true
	}

	embedBody := helpers.EmbedTemplate()

	bytesBody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	result := gjson.ParseBytes(bytesBody)

	if resp.StatusCode != http.StatusOK {
		// Wrong currency code
		msg := result.Get("error").String()
		if notFloat {
			msg = fmt.Sprintf("%s is not a valid number", amount)
		}
		embedBody.SetDescription(msg)
	} else {
		// HTTP OK 200
		toRate := result.Get("rates").Get(strings.ToUpper(to)).Float()
		result := amountFloat * toRate
		msg := fmt.Sprintf("%.2f %s = %.2f %s", amountFloat, strings.ToUpper(from), result, strings.ToUpper(to))
		embedBody.SetDescription(msg)
	}

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, embedBody.MessageEmbed)
	if err != nil {
		return err
	}

	return nil
}
