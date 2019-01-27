#!/usr/bin/env bash

# Living in https://gist.github.com/rubencaro/633cd90065d399d5fe1b56e46440d2bb
# Loosely based on https://github.com/tartley/rerun2

ignore_secs=0.25
clear='false'
verbose='false'
ignore_until=$(date +%s.%N)
excludes='\.git|\.coverage|\.cache|tmp|env'
cmd=""

function usage {
    echo "Rerun a given command every time filesystem changes are detected."
    echo ""
    echo "Usage: $(basename $0) [OPTIONS] COMMAND"
    echo ""
    echo "  -c, --clear             Clear the screen before each execution of COMMAND."
    echo "  -v, --verbose           Print details about this script's execution."
    echo "  -h, --help              Display this help and exit."
    echo "  -e, --exclude REGEXP    Don't fire for paths that match given REGEXP."
    echo "                          Can be given several times."
    echo ""
    echo "Run the given COMMAND, and then every time filesystem changes are"
    echo "detected in or below the current directory, run COMMAND again."
    echo ""
    echo "Given REGEXPs will be concatenated using OR and passed to 'inotifywait'."
    echo "Any path matching any of those REGEXPs will not be monitored for changes."
    echo "Default excludes are '$excludes'"
    echo ""
    echo "COMMAND can only be a simple command, ie. \"executable arg arg...\"."
    echo "For compound commands, use:"
    echo ""
    echo "    $(basename $0) bash -c \"ls -l | grep ^d\""
    echo ""
}

while [ $# -gt 0 ]; do
    case "$1" in
        -c|--clear) clear='true';;
        -v|--verbose) verbose='true';;
        -h|--help) usage; exit;;
        -e|--exclude) # add to default excludes
            excludes="$excludes|$2"; shift;;
        *) # the rest is the actual cmd
            cmd="$@"; break;;
    esac
    shift
done

function maybe_echo {
    [ $verbose = "true" ] && echo $1
}

function execute {
    [ $clear = "true" ] && clear
    maybe_echo "Excludes: $excludes"
    maybe_echo "Running: $cmd"
    (sleep $ignore_secs && $cmd) &
}

execute

inotifywait --quiet --recursive --monitor --format "%e %w%f" \
    --event modify --event move --event create --event delete \
    --exclude "($excludes)" \
    . | while read changed
do
    # ignore events that are too close in time, fire only once
    if [ $(echo "$(date +%s.%N) > $ignore_until" | bc) -eq 1 ] ; then
        ignore_until=$(echo "$(date +%s.%N) + $ignore_secs" | bc)
        maybe_echo "Changed: $changed"
        execute
    fi
done