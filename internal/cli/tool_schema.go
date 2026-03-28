package cli

import (
	"encoding/json"
	"fmt"

	"{{MODULE_PATH}}/internal/i18n"
	"github.com/spf13/cobra"
)

type toolParam struct {
	Type        string   `json:"type"`
	Description string   `json:"description,omitempty"`
	Enum        []string `json:"enum,omitempty"`
	Default     any      `json:"default,omitempty"`
}

type toolSchema struct {
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Parameters  *toolSchemaParameters `json:"parameters,omitempty"`
}

type toolSchemaParameters struct {
	Type       string               `json:"type"`
	Properties map[string]toolParam `json:"properties"`
	Required   []string             `json:"required,omitempty"`
}

// TODO: Update schemas for your commands
var schemas = []toolSchema{
	{
		Name:        "{{PROJECT_NAME}}_search",
		Description: "{{DESCRIPTION_EN}}",
		Parameters: &toolSchemaParameters{
			Type: "object",
			Properties: map[string]toolParam{
				"keyword": {Type: "string", Description: "Search keyword"},
				"output":  {Type: "string", Description: "Output format", Enum: []string{"json", "jsonl", "csv", "table"}, Default: "json"},
			},
			Required: []string{"keyword"},
		},
	},
	{
		Name:        "{{PROJECT_NAME}}_cache_clear",
		Description: "Clear the local search cache.",
	},
	{
		Name:        "{{PROJECT_NAME}}_cache_stats",
		Description: "Show cache statistics.",
	},
}

var toolSchemaCmd = &cobra.Command{
	Use:     "tool-schema [command]",
	Short:   i18n.T(i18n.MsgToolSchemaShort),
	GroupID: "util",
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var out interface{}

		if len(args) == 1 {
			name := "{{PROJECT_NAME}}_" + args[0]
			for _, s := range schemas {
				if s.Name == name {
					out = s
					break
				}
			}
			if out == nil {
				return fmt.Errorf("unknown command: %s", args[0])
			}
		} else {
			out = schemas
		}

		b, err := json.MarshalIndent(out, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(b))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(toolSchemaCmd)
}
