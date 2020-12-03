package hooks

import (
	"bytes"
	"text/template"
	"time"

	"github.com/pkg/errors"
)

type hookData struct {
	ID         string
	Version    string
	CreatedAt  string
	GoFishPath string
}

const goFishID = "# Hook created by go-fish"

// renderScript generates a shell script to be used for the git hooks.
func renderScript(goFishPath, version string) (string, error) {
	tmpl, err := template.New("hook").Parse(hookTemplate)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse hook script template")
	}

	buf := &bytes.Buffer{}
	hook := hookData{
		ID:         goFishID,
		Version:    version,
		CreatedAt:  time.Now().Format("Jan 2, 2006 at 3:04pm (MST)"),
		GoFishPath: goFishPath,
	}

	err = tmpl.Execute(buf, hook)
	if err != nil {
		return "", errors.Wrap(err, "failed to generate hook script")
	}

	return buf.String(), nil
}

const hookTemplate = `#!/bin/sh
{{.ID}}
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
	if [ "${GOFISH_DEBUG}" = "true" ] || [ "${GOFISH_DEBUG}" = "1" ]; then
		"$gofishPath" --verbose run $hookName "$gitParams"
	else
    	"$gofishPath" run $hookName "$gitParams"
	fi
else
    echo "Can't find go-fish, skipping $hookName hook"
fi
`
