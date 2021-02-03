package discord

import (
	"github.com/thisisibrahimd/rich-go/client"
)

func SetRichPresence(clientid string, activity client.Activity) {
	err := client.Login(clientid)
	if err != nil {
		panic(err)
	}

	err = client.SetActivity(activity)
	if err != nil {
		panic(err)
	}
}
