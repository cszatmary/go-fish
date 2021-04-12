#!/bin/sh
if [ -z "$gofish_skip_init" ]; then
  debug () {
    if [ -n "$GOFISH_DEBUG" ]; then
      echo "go-fish (debug): $1"
    fi
  }

  readonly hook_name="$(basename "$0")"
  debug "running $hook_name hook..."

  if [ -n "$GOFISH_SKIP" ]; then
    debug "GOFISH_SKIP is set, skipping hook"
    exit 0
  fi

  export readonly gofish_skip_init=1
  sh -e "$0" "$@"
  code="$?"

  if [ $code != 0 ]; then
    echo "go-fish (error): $hook_name hook exited with code $code"
  fi
  exit $code
fi
