package output

import (
	"os"

	"golang.org/x/term"
)

type TableColumn struct {
	Header string
	Key    string
}

type Formatter interface {
	Format(data interface{}) error
}

func IsTTY() bool {
	return term.IsTerminal(int(os.Stdout.Fd()))
}

// AllColumns is the full set of columns for CSV export.
// TODO: Update with your actual data fields.
var AllColumns = []TableColumn{
	{Header: "id", Key: "id"},
	{Header: "name", Key: "name"},
}

func NewFormatter(format string, columns []TableColumn) Formatter {
	switch format {
	case "table":
		return &TableFormatter{Columns: columns}
	case "json":
		return &JSONFormatter{}
	case "jsonl":
		return &JSONLFormatter{}
	case "csv":
		return &CSVFormatter{Columns: AllColumns}
	case "auto":
		if IsTTY() {
			return &TableFormatter{Columns: columns}
		}
		return &JSONFormatter{}
	default:
		return &JSONFormatter{}
	}
}
