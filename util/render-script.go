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
	RootDir    string
}

// RenderScript generates a shell script to be used for the git hooks.
func RenderScript(goFishPath, rootDir string) (string, error) {
	hook := hook{
		Version,
		time.Now().Format("Jan 2, 2006 at 3:04pm (MST)"),
		goFishPath,
		rootDir,
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
rootDir="{{.RootDir}}"
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
	if [ "${GOFISH_DEBUG}" = "true" ] || [ "${GOFISH_DEBUG}" = "1" ]; then
		"$gofishPath" -v -p "$rootDir" run $hookName "$gitParams"
	else
    	"$gofishPath" -p "$rootDir" run $hookName "$gitParams"
	fi
else
    echo "Can't find go-fish, skipping $hookName hook"
fi
`
