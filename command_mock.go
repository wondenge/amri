package amri

import (
	"github.com/wondenge/kamili"
)

// MockCommand is an implementation of Command that can be used for tests.
// It is publicly exported from this package in case you want to use it
// externally.
type MockCommand struct {
	// Settable
	HelpText     string
	RunResult    int
	SynopsisText string

	// Set by the command
	RunCalled bool
	RunArgs   []string
}

func (c *MockCommand) Help() string {
	return c.HelpText
}

func (c *MockCommand) Run(args []string) int {
	c.RunCalled = true
	c.RunArgs = args

	return c.RunResult
}

func (c *MockCommand) Synopsis() string {
	return c.SynopsisText
}

// MockCommandAutocomplete is an implementation of CommandAutocomplete.
type MockCommandAutocomplete struct {
	MockCommand

	// Settable
	AutocompleteArgsValue  kamili.Predictor
	AutocompleteFlagsValue kamili.Flags
}

func (c *MockCommandAutocomplete) AutocompleteArgs() kamili.Predictor {
	return c.AutocompleteArgsValue
}

func (c *MockCommandAutocomplete) AutocompleteFlags() kamili.Flags {
	return c.AutocompleteFlagsValue
}

// MockCommandHelpTemplate is an implementation of CommandHelpTemplate.
type MockCommandHelpTemplate struct {
	MockCommand

	// Settable
	HelpTemplateText string
}

func (c *MockCommandHelpTemplate) HelpTemplate() string {
	return c.HelpTemplateText
}
