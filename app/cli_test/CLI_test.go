package cli_test

import (
	"bytes"
	"fmt"
	cli2 "github.com/gnori-zon/go-tdd/app/cli"
	"io"
	"strings"
	"testing"
)

type SpyGame struct {
	startCalled  bool
	finishCalled bool
	startWith    int
	finishWith   string
}

func (s *SpyGame) Start(numberOfPlayers int) {
	s.startCalled = true
	s.startWith = numberOfPlayers
}

func (s *SpyGame) Finish(winner string) {
	s.finishCalled = true
	s.finishWith = winner
}

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
				game := &SpyGame{}
				cli := cli2.NewCLI(in, dummyStdOut, game)

				cli.PlayPoker()

				assertGameStartWith(t, game, 2)
				assertGameFinishWith(t, game, testCase.wantName)
			})
		}
	})

	t.Run("it prompts the user to enter count players", func(t *testing.T) {
		stdOut := &bytes.Buffer{}
		in := userInput("7")
		game := &SpyGame{}
		cli := cli2.NewCLI(in, stdOut, game)

		cli.PlayPoker()

		assertMessagesSentToUser(t, stdOut, cli2.PlayerPrompt)
		assertGameStartWith(t, game, 7)
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		in := userInput("Press")
		stdOut := &bytes.Buffer{}
		game := &SpyGame{}
		cli := cli2.NewCLI(in, stdOut, game)

		cli.PlayPoker()

		assertMessagesSentToUser(t, stdOut, cli2.PlayerPrompt, cli2.BadPlayerInputErrMsg)

		if game.startCalled {
			t.Errorf("game should not have started")
		}
	})

	t.Run("it prints and error when a bad string for finish lene and does not finish game", func(t *testing.T) {
		in := userInput("5", "Lloyd is a killer")
		stdOut := &bytes.Buffer{}
		game := &SpyGame{}
		cli := cli2.NewCLI(in, stdOut, game)

		cli.PlayPoker()

		assertMessagesSentToUser(t, stdOut, cli2.PlayerPrompt, cli2.BadFormatWinRecordErrMsg)
		assertGameStartWith(t, game, 5)

		if game.finishCalled {
			t.Errorf("game should not have finished")
		}
	})
}

func userInput(lines ...string) io.Reader {
	return strings.NewReader(strings.Join(lines, "\n"))
}

func assertGameStartWith(t *testing.T, game *SpyGame, wantNumberOfPlayers int) {
	if game.startWith != wantNumberOfPlayers {
		t.Errorf("want start with %d, but got %d", wantNumberOfPlayers, game.startWith)
	}
}

func assertGameFinishWith(t testing.TB, game *SpyGame, wantName string) {
	t.Helper()
	if game.finishWith != wantName {
		t.Errorf("want finish game with %q, but got %q", wantName, game.finishWith)
	}
}

func assertMessagesSentToUser(t testing.TB, stdOut *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdOut.String()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
