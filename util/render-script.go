package util

import (
	"bytes"
	"text/template"
	"time"
)

type hook struct {
	Version    string
	CreatedAt  string
	GoFishPath string
}

// RenderScript generates a shell script to be used for the git hooks.
func RenderScript(goFishPath string) (string, error) {
	hook := hook{
		Version,
		time.Now().Format("Jan 2, 2006 at 3:04pm (MST)"),
		goFishPath,
	}

	tmpl, err := template.New("hook").Parse(hookTemplate)

	if err != nil {
		return "", err
	}

	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, hook)

	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

const hookTemplate = `#! /bin/sh

# Hook created by go-fish
# Version: {{.Version}}
# Created At: {{.CreatedAt}}

gofishPath="{{.GoFishPath}}"
hookName="$(basename "$0")"
gitParams="$*"

debug() {
    if [ "${GOFISH_DEBUG}" = "true" ] || [ "${GOFISH_DEBUG}" = "1" ]; then
        echo "go-fish:debug $1"
    fi
}

debug "$hookName hook started"

if [ "$\{GOFISH_SKIP_HOOKS}" = "true" ] || [ "$\{GOFISH_SKIP_HOOKS}" = "1" ]; then
    debug "GOFISH_SKIP_HOOKS is set to $\{GOFISH_SKIP_HOOKS}, skipping hook"
    exit 0
fi

if [ -f "$gofishPath" ]; then
    "$gofishPath" run $hookName "$gitParams"
else
    echo "Can't find go-fish, skipping $hookName hook"
fi
`
