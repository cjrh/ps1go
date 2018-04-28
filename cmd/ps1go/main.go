/*
Been testing this commandline manually, with and without
being in a branch or virtualenv:

$ ../ps1go '{{.Virtualenv}} -- {{.Branch}} $'

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
	// TODO: colorable by mattn does not properly support cygwin-type terms yet.
	//"github.com/fatih/color"
	"fmt"
	"path/filepath"
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

	type Params struct {
		Branch string
		Virtualenv string
	}

	params := &Params{
		Branch:     gitBranchOrHash(),
		Virtualenv: virtualenv(),
	}

	var msg bytes.Buffer
	tmpl, err := template.New("ps1").Parse(input)
	if err != nil { panic(err) }
	err = tmpl.Execute(&msg, params)
	if err != nil { panic(err)}

	return msg.String()
}

func main() {
	app := cli.NewApp()
	app.Name = "ps1go"
	app.Usage = "PS1 prompt generator"
	app.Action = func (c *cli.Context) {
		prompt := generate(c.Args().Get(0))
		fmt.Print(prompt)
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
