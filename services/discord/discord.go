package discord

import (
	"log"

	"opdevelopers.live/user/config"
	"opdevelopers.live/user/constants"
	"opdevelopers.live/user/helpers/util"
)

func SendErrorMessageToDiscord(message string) {
	config := config.Get()
	payload := map[string]interface{}{
		"content": message,
	}
	res, err := util.CallAPI(constants.DISCORD_URL+config.DiscordWebHookId+"/"+config.DiscordWebHookToken+"?wait="+config.DiscordShouldWaitForResponse, "POST", constants.CONTENT_TYPE_XML_ENCODED, payload, nil)

	if err != nil {
		log.Println(err)
	}

	log.Println(res)
}
