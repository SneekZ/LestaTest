package mystem

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
)

func useMystem(input string) (string, string, error) {
	exePath := ".\\mystem\\mystem.exe"

	cmd := exec.Command(exePath, "-nl")

	var stdout, stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	cmd.Stdin = bytes.NewBufferString(input)

	err := cmd.Run()
	if err != nil {
		return "", "", err
	}

	return stdout.String(), stderr.String(), nil
}

func parseMystemOut(input string) []string {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var words = []string{}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		parts := strings.Split(line, "|")
		word := strings.ToLower(strings.Trim(parts[0], " ?"))
		words = append(words, word)
	}

	return words
}

func Literalize(input string) []string {
	// Утилита mystem используется для приведения слов к стандартной форме

	stdout, _, _ := useMystem(input)

	literalized := parseMystemOut(stdout)

	return literalized
}