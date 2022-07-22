package proverbs

import (
	"fmt"
	"os"
	"strings"
)

type proverb struct {
	n    int
	line string
}

func newProverb(n int, line string) *proverb {
	return &proverb{n: n, line: line}
}

func load(path string) ([]*proverb, error) {
	if path == "" {
		path = "proverbs/proverbs.txt"
	}
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(f), "\n")
	var proverbs []*proverb
	for i := range lines {
		proverbs = append(proverbs, newProverb(i+1, lines[i]))
	}
	return proverbs, nil
}

func (p proverb) String() string {
	return fmt.Sprintf("%d. %q\n", p.n, p.line)
}
