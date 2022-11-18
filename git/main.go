package git

import (
	"github.com/netitus/laravel-i18n/slices"
	"log"
	"os/exec"
	"strings"
)

func GitDiff(from, to string) []string {
	var diff string
	if to != `` {
		diff = CmdToString(`git`, `diff`, from, to, `--` /*, `|`, `grep`, `^+`*/)
	} else {
		diff = CmdToString(`git`, `diff`, from, `--` /*, `|`, `grep`, `^+`*/)
	}
	var lines slices.StringSlice = strings.Split(diff, "\n")
	lines.Filter(func(val string) bool {
		return strings.Index(val, `+`) == 0 || strings.Index(val, `@@`) == 0 // file context
	})

	return lines
}

func CmdToString(cmd string, args ...string) string {
	b, err := exec.Command(cmd, args...).Output()
	if err != nil {
		log.Fatalf(`Failed executing syscall; %s`, err.Error())
	}

	return string(b)
}
