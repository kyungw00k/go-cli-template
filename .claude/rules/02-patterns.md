# Key Patterns

## Adding a New Command

1. Create `internal/cli/<command>.go`
2. Define `var <name>Cmd = &cobra.Command{...}`
3. In `init()`, register: `rootCmd.AddCommand(<name>Cmd)` or `parentCmd.AddCommand(<name>Cmd)`
4. Add i18n keys to `internal/i18n/messages.go` (both ko and en maps)

## Adding a Subcommand

1. Create parent: `internal/cli/<parent>_root.go` (GroupID, RunE → cmd.Help())
2. Create sub: `internal/cli/<parent>_<sub>.go`
3. Register: `<parent>Cmd.AddCommand(<sub>Cmd)`

## Adding Table Columns

1. Define columns in the command: `[]output.TableColumn{{Header: i18n.T(key), Key: "json_field"}}`
2. Add i18n header keys to messages.go
3. Update `output.AllColumns` in formatter.go for CSV full export

## Adding i18n Messages

1. Add key constant in `internal/i18n/messages.go`
2. Add translation to both `ko` and `en` maps
3. Use: `i18n.T(i18n.KeyName)` or `i18n.Tf(i18n.KeyName, args...)`

## i18n Key Naming Conventions

| Prefix | Purpose | Example |
|--------|---------|---------|
| `Msg` | Command Short/Long | `MsgRootShort` |
| `Hdr` | Table header | `HdrName` |
| `Err` | Error message | `ErrNoResults` |
| `Flag` | Flag usage | `FlagOutputUsage` |
| `Group` | Command group | `GroupCache` |
