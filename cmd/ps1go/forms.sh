#!/usr/bin/env bash
# source this inside the faketesting dir

BRANCH='{{.Red}}{{.Branch}}{{.Default}}'

echo "Basic"
../ps1go.exe "{{.Virtualenv}} -- ${BRANCH} $"
echo -e '\n'

echo "Conditional not activated"
../ps1go "{{if .Virtualenv}}{{.}} -- {{end}}${BRANCH} $"
echo -e '\n'

echo "Conditional (activated)"
source venvfake/Scripts/activate
../ps1go "{{if .Virtualenv}}{{.Virtualenv}} -- {{end}}${BRANCH} $"
deactivate
echo -e '\n'

echo "Conditional not activated"
../ps1go "{{if .Virtualenv}}{{.Virtualenv}} -- {{end}}${BRANCH} $"
echo -e '\n'

echo "Conditional with decoration (activated)"
source venvfake/Scripts/activate
../ps1go "{{if .Virtualenv}}[{{.Virtualenv}}] {{end}}${BRANCH} $"
deactivate
echo -e '\n'

echo "Conditional with decoration (not activated)"
../ps1go "{{if .Virtualenv}}[{{.Virtualenv}}] {{end}}${BRANCH} $"
echo -e '\n'

echo "Conditional with color decoration (activated)"
source venvfake/Scripts/activate
../ps1go "{{if .Virtualenv}}[{{.Green}}{{.Virtualenv}}{{.Reset}}] {{else}}{{.Dim}}{{.Dark_gray}}(no venv){{.Reset}} {{end}}${BRANCH} $"
deactivate
echo -e '\n'

echo "Conditional with color decoration (not activated)"
../ps1go "{{if .Virtualenv}}[{{.Green}}{{.Virtualenv}}{{.Reset}}] {{else}}{{.Dim}}{{.Dark_gray}}(no venv){{.Reset}} {{end}}${BRANCH} $"
echo -e '\n'

# Using template pieces
V1='[{{.Green}}{{.Virtualenv}}{{.Default}}]'
V0='{{.Dim}}{{.Dark_gray}}(no venv){{.Default}}'
#V0='{{.Dim}}(no venv){{.Default}}'
VENV="{{if .Virtualenv}}$V1 {{else}}$V0 {{end}}"
USER="{{.Light_blue}}\u{{.Default}}"

../ps1go "$VENV${BRANCH} $"
echo -e '\n'

jobswrite () {
    # Conditionally emit template text if the background jobs count is
    # greater than zero
    local count=$(eval 'jobs -p | wc -l')
#    eval '>&2 echo "count was $count "'
    if [ $count -gt 0 ]
    then
        echo " {{.Blue}}jobs:$count{{.Default}}"
    fi
}

fpath () {
    echo "{{.Yellow}}\w{{.Default}}"
}

echo "Conditional with color decoration (activated) BASH MULTILINE"
source venvfake/Scripts/activate
../ps1go "$VENV${BRANCH} \$ "
deactivate
echo -e '\n'

echo "Conditional with color decoration (not activated) BASH MULTILINE"
#vim <(../ps1go "$VENV${BRANCH} {{.Blue}}$(jobs -p | wc -l | grep -v ^0 | xargs echo ){{.Reset}} \$ ")
out=$(../ps1go "$VENV${BRANCH} {{.Blue}}$(jobs -p | wc -l | grep -v ^0 | xargs echo ){{.Reset}} \$ ")
echo $out
echo -e '\n'

PS1BAK=$PS1

echo "Conditional with color decoration (not activated) jobs NO BG"
PS1="$(../ps1go "$VENV${BRANCH}$(jobswrite) \$ ")"
printf '%s\n' "${PS1@P}"
echo -e '\n'

sleep 0.2 &
echo "Conditional with color decoration (not activated) jobs BG"
PS1="$(../ps1go "\j $VENV${BRANCH}$(jobswrite) \$ ")"
printf '%s\n' "${PS1@P}"
echo -e '\n'
wait


echo "Conditional with color decoration (not activated) jobs NO BG"
PS1="$(../ps1go "\j $USER:$(fpath) $VENV$BRANCH$(jobswrite) \n$ ")"
printf '%s\n' "${PS1@P}"
echo -e '\n'

echo "Conditional with color decoration (not activated) jobs NO BG"

../ps1go "{{.Light_yellow}}$USER:\h:$(fpath) $VENV$BRANCH$(jobswrite) {{.Cyan}}\d {{.Underlined}}{{.Blue}}\@{{.Default}}\n$ "
echo -e '\n'
PS1="\$(../ps1go \"{{.Light_yellow}}$USER:\h:$(fpath) $VENV$BRANCH$(jobswrite) {{.Cyan}}\d {{.Underlined}}{{.Blue}}\@{{.Default}}\n$ \")"
printf '%s\n' "${PS1@P}"
echo -e '\n'

#PS1=$PS1BAK
echo
