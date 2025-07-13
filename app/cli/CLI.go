package cli

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		bufio.NewScanner(in),
		out,
		game,
	}
}

const PlayerPrompt = "Please enter the number of players: "

const BadPlayerInputErrMsg = "you're so silly"
const BadFormatWinRecordErrMsg = "bad format win record"

func (cli *CLI) PlayPoker() {
	_, _ = fmt.Fprint(cli.out, PlayerPrompt)

	countPlayersLine := cli.readLine()
	numberOfPlayers, err := strconv.Atoi(strings.Trim(countPlayersLine, "\n"))
	if err != nil {
		_, _ = fmt.Fprint(cli.out, BadPlayerInputErrMsg)
		return
	}
	cli.game.Start(numberOfPlayers)

	winnerLine := cli.readLine()
	winner, ok := extractWinner(winnerLine)
	if !ok {
		_, _ = fmt.Fprint(cli.out, BadFormatWinRecordErrMsg)
		return
	}
	cli.game.Finish(winner)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(text string) (string, bool) {
	words := strings.Split(text, " ")
	if len(words) != 2 {
		return "", false
	}
	if strings.TrimSpace(words[1]) != "wins" {
		return "", false
	}
	return strings.TrimSpace(words[0]), true
}
