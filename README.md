# ps1go

PS1 prompt generator

# Quickstart

Download the `ps1go` executable for your platform, place it on your shell path, and set your `PS1` environment
variable to call the executable with your customized template.

For example:

TBD

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
