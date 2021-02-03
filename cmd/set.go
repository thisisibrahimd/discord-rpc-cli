package cmd

import (
	"discord-rpc-cli/discord"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/thisisibrahimd/rich-go/client"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set [clientid]",
	Short: "Set a rich presence session",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		state, err := cmd.Flags().GetString("state")
		details, err := cmd.Flags().GetString("details")
		lgimg, err := cmd.Flags().GetString("lgimg")
		lgtxt, err := cmd.Flags().GetString("lgtxt")
		smimg, err := cmd.Flags().GetString("smimg")
		smtxt, err := cmd.Flags().GetString("smtxt")
		if err != nil {
			fmt.Errorf("could not parse flags")
		}

		now := time.Now()
		activity := client.Activity{
			State:      state,
			Details:    details,
			LargeImage: lgimg,
			LargeText:  lgtxt,
			SmallImage: smimg,
			SmallText:  smtxt,
			Timestamps: &client.Timestamps{
				Start: &now,
			},
		}

		discord.SetRichPresence(args[0], activity)

		time.Sleep(time.Hour * 24)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	setCmd.PersistentFlags().String("state", "", "The user's current party status")
	setCmd.PersistentFlags().String("details", "", "What the player is currently doing")
	setCmd.PersistentFlags().String("lgimg", "", "The id for a large asset of the activity")
	setCmd.PersistentFlags().String("lgtxt", "", "Text displayed when hovering over the large image of the activity")
	setCmd.PersistentFlags().String("smimg", "", "The id for a small asset of the activity")
	setCmd.PersistentFlags().String("smtxt", "", "Text displayed when hovering over the small image of the activity")
}
