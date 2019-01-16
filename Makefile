.PHONY: setup setup-brew setup-ghq setup-anyenv setup-neovim

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

setup-anyenv:
	sh shell/setup-anyenv.sh

setup-neovim:
	sh shell/setup-nvim.sh

