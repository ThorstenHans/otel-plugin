package cmd

// To add support for an additional observability stack, do the following:
// 1. Add a const with the name of the stack
// 2. Update the switch branch of GetComposeFileName()
// 3. Add Stack specific Compose file to the `./otel-config` folder
// 4. Add the Stack specific Compose file to the list of assets in `./spin-pluginify.toml`
type Stack int

const (
	Aspire Stack = iota
	Default
)

func GetStackByFlags(aspire bool) Stack {
	if aspire {
		return Aspire
	}
	return Default
}

func (s Stack) GetComposeFileName() string {
	switch s {
	case Aspire:
		return "compose.aspire.yaml"
	default:
		return "compose.yaml"

	}
}
