package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"flag"

	"github.com/hashicorp/go-version"
)

func main() {
	flagReverse := flag.Bool("r", false, "reverse the result of comparisons")
	flagStrict := flag.Bool("s", false, "error if any line is not a semver (ignore by default)")
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	var versions []*version.Version
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		ver, err := version.NewVersion(line)
		if err != nil {
			if *flagStrict {
				fmt.Fprintln(os.Stderr, fmt.Errorf("%s is not a semver: %v", line, err))
				os.Exit(1)
			}
			continue
		}
		versions = append(versions, ver)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	var collection sort.Interface = version.Collection(versions)
	if *flagReverse {
		collection = sort.Reverse(collection)
	}
	sort.Sort(collection)
	for _, ver := range versions {
		fmt.Println(ver.Original())
	}
}
