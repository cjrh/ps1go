# ps1go

PS1 prompt generator

# Quickstart

Download the `ps1go` executable for your platform, place it on your shell path, and set your `PS1` environment
variable to call the executable with your customized template.

For example, here's an example that works with a Python 
virtual environment:

```bash
# .bash_profile
VENV_ACTIVE='[{{.Green}}{{.Virtualenv}}{{.Default}}]'
VENV_NOTACTIVE='{{.Dim}}{{.Dark_gray}}(no venv){{.Default}}'
VENV="{{if .Virtualenv}}$VENV_ACTIVE {{else}}$VENV_NOTACTIVE {{end}}"

export PS1="\$(ps1go "$VENV \w \$ ")"
```

The `.Virtualenv` template parameter is built into `ps1go`. It automatically detects whether a 
Python virtual environment has been activated, and if so, sets the value of `{{.Virtualenv}}` to 
be the name of the activated virtualenv.

The if-statement syntax is all built into the _Go_ `text/template` package, and has its own
language. Refer to that documentation to see how to use it.

# Background

In your `.bash_profile`, you can customise your shell prompt by setting (or changing) an environment variable
called `PS1`. You can display your current setting with the following:

```
$ echo $PS1
```

Here are a few defaults for different systems:

```
# Debian
PS1='${debian_chroot:+($debian_chroot)}\u@\h:\w\$ '

# Centos/Fedora
[ "$PS1" = "\\s-\\v\\\$ " ] && PS1="[\u@\h \W]\\$ "

# Gentoo (/etc/bash/bashrc)
if [[ ${EUID} == 0 ]] ; then
    PS1='\[\033[01;31m\]\h\[\033[01;34m\] \W \$\[\033[00m\] '
else
    PS1='\[\033[01;32m\]\u@\h\[\033[01;34m\] \w \$\[\033[00m\] '
fi
```

The _Gentoo_ one looks very complicated because of all those
[ANSI color codes](https://stackoverflow.com/a/33206814). You
can see more of those in the default [Git bash for Windows](https://gitforwindows.org/) `PS1` 
prompt setting:

```
$ echo $PS1
\[\033]0;$TITLEPREFIX:$PWD\007\]\n\[\033[32m\]\u@\h \[\033[35m\]$MSYSTEM \[\033[33m\]\w\[\033[36m\]`
__git_ps1`\[\033[0m\]\n$
```

This is a pretty complicated prompt because, in addition of having color codes, it also runs a _command_
`__git_ps1` in the course of generating the prompt, to give additional information about the status
of a git repo that you might happen to be inside.

We do the same thing, but you call `ps1go` instead of `__git_ps1` (or any other executable).

# Extended example

```bash

# Dislay current git branch - built into ps1go
BRANCH='{{.Red}}{{.Branch}}{{.Default}}'

# Virtualenv support, building up a template from smaller pieces
V1='[{{.Green}}{{.Virtualenv}}{{.Default}}]'
V0='{{.Dim}}{{.Dark_gray}}(no venv){{.Default}}'
VENV="{{if .Virtualenv}}$V1 {{else}}$V0 {{end}}"

# Current user, with own color - this uses built-in PS1 format code \u
USER="{{.Light_blue}}\u{{.Default}}"

jobscount () {
    # Example of calling a bash function from your PS1.
    local count=$(eval 'jobs -p | wc -l')
    # Conditionally emit template text if the background jobs count is
    # greater than zero
    if [ $count -gt 0 ]
    then
        echo " {{.Blue}}jobs:$count{{.Default}}"
    fi
}

fpath () {
    # Current directory, with a color, also uses built-in PS1 code \w
    echo "{{.Yellow}}\w{{.Default}}"
}

export PS1="\$(ps1go \"$USER:\h:$(fpath) $VENV$BRANCH$(jobswrite) \n$ \")"
```

Links
-----

https://www.gnu.org/software/bash/manual/html_node/Controlling-the-Prompt.html#Controlling-the-Prompt
