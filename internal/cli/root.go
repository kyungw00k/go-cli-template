package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"{{MODULE_PATH}}/api"
	"{{MODULE_PATH}}/cache"
	"{{MODULE_PATH}}/internal/i18n"
	"{{MODULE_PATH}}/internal/output"
	"github.com/spf13/cobra"
)

var Version = "dev"

var flagOutput string

var rootCmd = &cobra.Command{
	Use:     "{{PROJECT_NAME}} <keyword>",
	Short:   i18n.T(i18n.MsgRootShort),
	Long:    i18n.T(i18n.MsgRootLong),
	Version: Version,
	Args:    cobra.ArbitraryArgs,
	Example: `  {{PROJECT_NAME}} hello
  {{PROJECT_NAME}} hello -o json
  {{PROJECT_NAME}} hello -o csv > results.csv`,
	RunE: runSearch,
}

func init() {
	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true

	rootCmd.AddGroup(
		&cobra.Group{ID: "cache", Title: i18n.T(i18n.GroupCache)},
		&cobra.Group{ID: "util", Title: i18n.T(i18n.GroupUtil)},
	)

	pf := rootCmd.PersistentFlags()
	pf.StringVarP(&flagOutput, "output", "o", "auto", i18n.T(i18n.FlagOutputUsage))

	rootCmd.SuggestionsMinimumDistance = 2
}

func Execute() error {
	return rootCmd.Execute()
}

func runSearch(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return cmd.Help()
	}

	keyword := strings.Join(args, " ")

	// Cache-first lookup
	c, err := cache.Open()
	if err != nil {
		return err
	}
	defer c.Close()

	var results []api.Result

	if data, ok := c.Get(keyword); ok {
		if err := json.Unmarshal(data, &results); err != nil {
			results = nil
		}
	}

	if results == nil {
		client := api.NewClient()
		results, err = client.Search(cmd.Context(), keyword)
		if err != nil {
			return err
		}

		if err := c.Set(keyword, results); err != nil {
			fmt.Fprintln(os.Stderr, "warning: failed to cache results:", err)
		}
	}

	if len(results) == 0 {
		fmt.Fprintln(os.Stderr, i18n.T(i18n.ErrNoResults))
		return nil
	}

	// TODO: Customize columns for your data
	columns := []output.TableColumn{
		{Header: i18n.T(i18n.HdrID), Key: "id"},
		{Header: i18n.T(i18n.HdrName), Key: "name"},
	}

	formatter := output.NewFormatter(flagOutput, columns)
	return formatter.Format(results)
}
