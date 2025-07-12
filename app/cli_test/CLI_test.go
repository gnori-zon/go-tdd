package cli_test

import (
	"fmt"
	"github.com/gnori-zon/go-tdd/app"
	cli2 "github.com/gnori-zon/go-tdd/app/cli"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	cases := []struct {
		wantName string
	}{
		{wantName: "Chris"},
		{wantName: "Cleo"},
	}

	for _, testCase := range cases {
		t.Run(testCase.wantName, func(t *testing.T) {
			in := strings.NewReader(fmt.Sprintf("%s wins\n", testCase.wantName))
			playerStore := &poker.StubPlayerStore{}
			cli := cli2.NewCLI(playerStore, in)
			cli.PlayPoker()

			poker.AssertSavedWin(t, playerStore.Wins, 1, testCase.wantName)
		})
	}
}
