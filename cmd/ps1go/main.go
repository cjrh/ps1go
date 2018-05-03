/*
Been testing this commandline manually, with and without
being in a branch or virtualenv:

$ ../ps1go '{{.Virtualenv}} -- {{.Branch}} $'


Shell path: "G:\Programs\Git\bin\sh.exe" -login -i

See example of running multiple subprocesses in goroutines:
https://gist.github.com/proudlygeek/4a9355bad16a62025a46

*/

package main

import (
	"gopkg.in/urfave/cli.v1"
	"os"
	"log"
	"html/template"
	"bytes"
	"os/exec"
	"strings"

	"path/filepath"
	"fmt"
)

func runCommand(cmdline string) (string) {
	elements := strings.Fields(cmdline)
	cmd := exec.Command(elements[0], elements[1:]...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		// TODO: should have a debug mode and log out here.
		return ""
	}
	return strings.TrimSpace(out.String())
}

func gitHash() (string) {
	cmdLine := "git rev-parse --short HEAD"
	return runCommand(cmdLine)
}

func gitBranch() (string) {
	cmdLine := "git symbolic-ref --quiet --short HEAD"
	return runCommand(cmdLine)
}

func gitBranchOrHash() (string) {
	out := gitBranch()
	if out == "" {
		out = gitHash()
	}
	//yellow := color.New(color.FgYellow).SprintFunc()
	//color.NoColor = false
	//out = color.YellowString(out)
	//os.Stdout.WriteString(out)
	//fmt.Print(out)
	return out
}

func virtualenv() (string) {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if pair[0] == "VIRTUAL_ENV" {
			basename := filepath.Base(pair[1])
			return basename
		}
	}
	return ""
}

func generate(input string) (string) {
	/*
	Ideas for other fields:

	- date
	- time
	- quotes file
	- news headline??
	- mercurial/svn/fossil/ branch info
	- uncommitted changes
	- difference between origin and upstream commits
	- summary of 'ls' info? e.g. number of files and dirs
	- shortened version of path, e.g. 1 or 2 chars for each level?
	 */

	m := map[string]string{
		"Branch":     gitBranchOrHash(),
		"Virtualenv": virtualenv(),

		"Reset": "\x1b[0m",

		"Bold":       "\x1b[1m",
		"Dim":        "\x1b[2m",
		"Underlined": "\x1b[4m",
		"Blink":      "\x1b[5m",
		"Reverse":    "\x1b[7m",
		"Hidden":     "\x1b[8m",

		"Default":        "\x1b[39m",
		"Black":          "\x1b[30m",
		"Red":            "\x1b[31m",
		"Green":          "\x1b[32m",
		"Yellow":         "\x1b[33m",
		"Blue":           "\x1b[34m",
		"Magenta":        "\x1b[35m",
		"Cyan":           "\x1b[36m",
		"Light_gray	": "\x1b[37m",
		"Dark_gray":      "\x1b[90m",
		"Light_red":      "\x1b[91m",
		"Light_green":    "\x1b[92m",
		"Light_yellow":   "\x1b[93m",
		"Light_blue":     "\x1b[94m",
		"Light_magenta":  "\x1b[95m",
		"Light_cyan":     "\x1b[96m",
		"White":          "\x1b[97m",

		// Backgrounds
		"bgDefault":       "\x1b[49m",
		"bgBlack":         "\x1b[40m",
		"bgRed":           "\x1b[41m",
		"bgGreen":         "\x1b[42m",
		"bgYellow":        "\x1b[43m",
		"bgBlue":          "\x1b[44m",
		"bgMagenta":       "\x1b[45m",
		"bgCyan":          "\x1b[46m",
		"bgLight_gray":    "\x1b[47m",
		"bgDark_gray":     "\x1b[100m",
		"bgLight_red":     "\x1b[101m",
		"bgLight_green":   "\x1b[102m",
		"bgLight_yellow":  "\x1b[103m",
		"bgLight_blue":    "\x1b[104m",
		"bgLight_magenta": "\x1b[105m",
		"bgLight_cyan":    "\x1b[106m",
		"bgWhite":         "\x1b[107m",
	}

	var msg bytes.Buffer
	// Force a reset at the end.
	tmpl, err := template.New("ps1").Parse(input + m["Reset"])
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(&msg, m)
	if err != nil {
		panic(err)
	}

	return msg.String()
}

func main() {
	fmt.Println("\x1b[31;1mHello, World!\x1b[0m")
	app := cli.NewApp()
	app.Name = "ps1go"
	app.Usage = "PS1 prompt generator"
	app.Action = func(c *cli.Context) {
		prompt := generate(c.Args().Get(0))
		// The STRING \x1b must be replaced with the BYTES 0x1B, which is ESC.
		// This allows the user to specify formatting information manually.
		prompt = strings.Replace(prompt, `\x1b`, "\x1b", -1)
		fmt.Print(prompt)
		//os.Stdout.Write([]byte(prompt))
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
