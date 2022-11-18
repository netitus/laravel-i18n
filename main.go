package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/netitus/laravel-i18n/git"
	"github.com/netitus/laravel-i18n/slices"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

var transFilePath string
var fromCommit string
var toCommit string

var customRegexp string

var translations map[string]string

func init() {
	flag.StringVar(&transFilePath, `t`, ``, `Translations file`)
	flag.StringVar(&fromCommit, `c`, `master`, `Starting commit hash`)
	flag.StringVar(&toCommit, `toc`, ``, `Ending commit hash`)
	flag.StringVar(&customRegexp, `pat`, ``, `Custom regex pattern`)
	flag.Parse()

	if transFilePath == `` {
		log.Fatalf(`No translations file passed`)
	}

	_, err := os.Stat(transFilePath)
	if os.IsNotExist(err) {
		log.Fatalf(`Translations file does not exist`)
	}
}

func main() {
	jfh, _ := os.Open(transFilePath)
	defer jfh.Close()

	jb, _ := io.ReadAll(jfh)
	if err := json.Unmarshal(jb, &translations); err != nil {
		log.Fatalf(`Error parsing translations JSON: %s`, err.Error())
	}

	var diff slices.StringSlice = git.GitDiff(fromCommit, toCommit)
	if customRegexp == `` {
		diff.Filter(func(val string) bool {
			return strings.Index(val, `__(`) > -1 || strings.Index(val, `trans(`) > -1 // || strings.Index(val, `@@`) == 0
		})
	}

	tsp1 := regexp.MustCompile("__\\([\"']([^'\"]+)\\)?") // double escaped version of  __\(["']([\w\s'?!;:.-]*)["'][),]
	tsp2 := regexp.MustCompile("trans\\([\"']([^'\"]+)\\)?")
	tsp3 := regexp.MustCompile(`$-^`) // cost efficient zero-matching pattern
	if customRegexp != `` {
		var err error
		tsp3, err = regexp.Compile(customRegexp)
		if err != nil {
			log.Fatalf(`Error compiling custom regular expression: %v`, err.Error())
		}
	}

	transStrings := make(slices.StringSlice, 0)
	for _, v := range diff {
		if found := tsp1.FindAllStringSubmatch(v, -1); found != nil {
			for _, f := range found {
				transStrings = append(transStrings, f[1])
			}
		}
		if found := tsp2.FindAllStringSubmatch(v, -1); found != nil {
			for _, f := range found {
				transStrings = append(transStrings, f[1])
			}
		}
		if customRegexp != `` {
			if found := tsp3.FindAllStringSubmatch(v, -1); found != nil {
				for _, f := range found {
					transStrings = append(transStrings, f[1])
				}
			}
		}
	}

	transStrings.Unique()
	newTranslations := make([]string, 0)
	for _, t := range transStrings {
		if _, exists := translations[t]; !exists {
			newTranslations = append(newTranslations, t)
		}
	}

	if len(newTranslations) > 0 {
		fmt.Printf("Found the following new translations since %s:\n\n\t", strings.TrimSpace(fromCommit))
		sort.Strings(newTranslations)

		fmt.Printf("%s\n\nRemember to check if you have the latest translations file!\n", strings.Join(newTranslations, "\n\t"))
	} else {
		fmt.Println(`No new translations found`)
	}

}
