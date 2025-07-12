package cli

import (
	"bufio"
	"github.com/gnori-zon/go-tdd/app"
	"io"
	"strings"
)

type CLI struct {
	playerStore poker.PlayerStore
	scanner     *bufio.Scanner
}

func NewCLI(playerStore poker.PlayerStore, in io.Reader) *CLI {
	return &CLI{playerStore, bufio.NewScanner(in)}
}

func (cli *CLI) PlayPoker() {
	line := cli.readLine()
	cli.playerStore.SaveWin(extractWinner(line))
}

func (cli *CLI) readLine() string {
	cli.scanner.Scan()
	return cli.scanner.Text()
}

func extractWinner(text string) string {
	return strings.Replace(text, " wins", "", 1)
}
