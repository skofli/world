/*
Copyright © 2022 Ben Overmyer <ben@overmyer.net>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/ironarachne/world/config"
	"os"

	"github.com/ironarachne/world/pkg/heraldry"
	"github.com/ironarachne/world/pkg/random"
	"github.com/spf13/cobra"
)

// heraldryCmd represents the heraldry command
var heraldryCmd = &cobra.Command{
	Use:   "heraldry",
	Short: "Generate fantasy heraldry",
	Long:  `This command generates fantasy heraldry.`,
	Run: func(cmd *cobra.Command, args []string) {
		configFile, _ := cmd.Flags().GetString("config")
		config.LoadConfig(configFile)
		randomSeed := random.SeedString()
		random.SeedFromString(randomSeed)
		fieldType := ""
		chargeTag := ""

		var o heraldry.Device
		var err error

		o, err = heraldry.GenerateByParameters(fieldType, chargeTag)
		if err != nil {
			fmt.Println(fmt.Errorf("Failed to generate heraldry: %w", err))
			os.Exit(1)
		}

		sd := o.Simplify()

		res, err := json.Marshal(sd)
		if err != nil {
			fmt.Println(fmt.Errorf("Failed to generate heraldry: %w", err))
			os.Exit(1)
		}
		fmt.Print(string(res))
	},
}

func init() {
	rootCmd.AddCommand(heraldryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// heraldryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// heraldryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
