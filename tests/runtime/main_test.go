package tests

import (
	"flag"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

var opts = godog.Options{
	Format: "progress",
	Output: colors.Colored(os.Stdout),
}

func TestMain(m *testing.M) {
	flag.Parse()

	opts.Paths = flag.Args()

	status := godog.TestSuite{
		Name:                "runtime",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()

	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}

func InitializeScenario(_ *godog.ScenarioContext) {
}
