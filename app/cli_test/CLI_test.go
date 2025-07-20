package cli_test

import (
	"bytes"
	"fmt"
	poker "github.com/gnori-zon/go-tdd/app"
	cli2 "github.com/gnori-zon/go-tdd/app/cli"
	"io"
	"strings"
	"testing"
)

var dummyStdOut = &bytes.Buffer{}

func TestCLI(t *testing.T) {
	t.Run("should store winner", func(t *testing.T) {
		cases := []struct {
			wantName string
		}{
			{wantName: "Chris"},
			{wantName: "Cleo"},
		}

		for _, testCase := range cases {
			t.Run(testCase.wantName, func(t *testing.T) {
				userWinsLine := fmt.Sprintf("%s wins", testCase.wantName)
				in := userInput("2", userWinsLine)
				game := &poker.SpyGame{}
				cli := cli2.NewCLI(in, dummyStdOut, game)

				cli.PlayPoker()

				poker.AssertGameStartWith(t, game, 2)
				poker.AssertGameFinishWith(t, game, testCase.wantName)
			})
		}
	})

	t.Run("it prompts the user to enter count players", func(t *testing.T) {
		stdOut := &bytes.Buffer{}
		in := userInput("7", "tom wins")
		game := &poker.SpyGame{}
		cli := cli2.NewCLI(in, stdOut, game)

		cli.PlayPoker()

		assertMessagesSentToUser(t, stdOut, cli2.PlayerPrompt)
		poker.AssertGameStartWith(t, game, 7)
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		in := userInput("Press")
		stdOut := &bytes.Buffer{}
		game := &poker.SpyGame{}
		cli := cli2.NewCLI(in, stdOut, game)

		cli.PlayPoker()

		assertMessagesSentToUser(t, stdOut, cli2.PlayerPrompt, cli2.BadPlayerInputErrMsg)

		if game.StartCalled {
			t.Errorf("game should not have started")
		}
	})

	t.Run("it prints and error when a bad string for finish lene and does not finish game", func(t *testing.T) {
		in := userInput("5", "Lloyd is a killer")
		stdOut := &bytes.Buffer{}
		game := &poker.SpyGame{}
		cli := cli2.NewCLI(in, stdOut, game)

		cli.PlayPoker()

		assertMessagesSentToUser(t, stdOut, cli2.PlayerPrompt, cli2.BadFormatWinRecordErrMsg)
		poker.AssertGameStartWith(t, game, 5)

		if game.FinishCalled {
			t.Errorf("game should not have finished")
		}
	})
}

func userInput(lines ...string) io.Reader {
	return strings.NewReader(strings.Join(lines, "\n"))
}

func assertMessagesSentToUser(t testing.TB, stdOut *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdOut.String()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
