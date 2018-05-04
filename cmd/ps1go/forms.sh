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
V1='[{{.Green}}{{.Virtualenv}}{{.Reset}}]'
V0='{{.Dim}}{{.Dark_gray}}(no venv){{.Reset}}'
VENV="{{if .Virtualenv}}$V1 {{else}}$V0 {{end}}"

echo "Conditional with color decoration (activated) BASH MULTILINE"
source venvfake/Scripts/activate
../ps1go "$VENV${BRANCH} \$ "
deactivate
echo -e '\n'

echo "Conditional with color decoration (not activated) BASH MULTILINE"
../ps1go "$VENV${BRANCH} \$(jobs) \$ "
echo -e '\n'
