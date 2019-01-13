.PHONY: setup setup-brew setup-ghq

all: main

main: setup

setup: setup-brew setup-ghq

setup-brew:
	/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
	brew doctor

setup-ghq:
	brew install ghq
	brew install peco
	git config --global ghq.root ${GOPATH}/src

