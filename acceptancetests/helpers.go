package acceptancetests

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

const (
	baseBinName = "temp-testbinary"
)

func LaunchTestProgram(port string) (cleanup func(), sendInterrupt func() error, err error) {
	binPath, err := buildBinary()
	if err != nil {
		return nil, nil, err
	}

	sendInterrupt, kill, err := runServer(binPath, port)
	cleanup = func() {
		if kill != nil {
			kill()
		}
		_ = os.Remove(binPath)
	}
	if err != nil {
		cleanup() // even though it's not listening correctly, the program could still be running
		return nil, nil, err
	}
	return cleanup, sendInterrupt, nil
}

func runServer(binPath string, port string) (sendInterrupt func() error, kill func(), err error) {
	runBinCommand := exec.Command(binPath)
	if err := runBinCommand.Start(); err != nil {
		return nil, nil, fmt.Errorf("cannot run binary: %s", err)
	}

	kill = func() {
		_ = runBinCommand.Process.Kill()
	}

	sendInterrupt = func() error {
		return runBinCommand.Process.Signal(syscall.SIGTERM)
	}
	err = waitForServerListening(port)
	return
}

func waitForServerListening(port string) error {
	for i := 0; i < 30; i++ {
		conn, _ := net.Dial("tcp", "localhost:"+port)
		if conn != nil {
			_ = conn.Close()
			return nil
		}
		time.Sleep(100 * time.Millisecond)
	}
	return fmt.Errorf("nothing seems to be listening on localhost:%s", port)
}

func buildBinary() (string, error) {
	rootDir, err := getRootDir()
	if err != nil {
		return "", err
	}
	serverDir := rootDir + "/acceptancetests/server"
	binName := randomString(10) + "-" + baseBinName
	binPath := serverDir + "/" + binName
	buildCommand := exec.Command("go", "build", "-o", binName, ".")
	buildCommand.Dir = serverDir
	if err := buildCommand.Run(); err != nil {
		return "", fmt.Errorf("cannot build binary %s: %s", binPath, err)
	}
	if err := os.Chmod(binPath, 0755); err != nil {
		return "", fmt.Errorf("failed to make binary executable: %s", err)
	}
	return binPath, nil
}

func getRootDir() (string, error) {
	moduleRoot := exec.Command("go", "list", "-m", "-f", "{{.Dir}}")
	output, err := moduleRoot.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func randomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, length)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
