#!/usr/bin/env bash
color_red='\e[0;31m'
color_cyan='\e[0;36m'
color_none='\e[0m'

function error {
    echo -e "${color_red}$1${color_none}"
}

function fail {
    if [ -n "$1" ]; then
        error "$1"
    fi
    exit 1
}

function assert_success {
    echo -e "${color_cyan}$1${color_none}"
    eval $1
    if [ $? -ne 0 ]; then
        fail "$2"
    fi
}

# $1 = directory
function updateOrigin() {
    if [ -d ${1} -a -d ${1}/.git ]; then
        # echo ${1}
        current_origin=$(git -C ${1} remote -v | grep origin | head -n1 | awk '{print $2}')
        if ! [ -z "$current_origin" ]; then
            # echo $current_origin
            new_origin=$(echo "$current_origin" | sed 's/stash/bitbucket/')
            new_origin=$(echo "$new_origin" | sed 's/.is.com/.xant.tech/')
            # echo $new_origin
            if [ "$current_origin" != "$new_origin" ]; then
                echo "${1} $current_origin"
                # echo "git -C ${1} remote set-url origin $new_origin"
                assert_success "git -C ${1} remote set-url origin $new_origin" "Failed to set $1 origin to $new_origin"
            fi
        fi
    fi
}

for dir in ~/Documents/Coding/Go/src/bitbucket.xant.tech/super/*
do
    updateOrigin "${dir}"
done
