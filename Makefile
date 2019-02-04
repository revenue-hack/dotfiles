.PHONY: setup setup-brew setup-ghq setup-anyenv setup-neovim setup-gitignore-global

all: main

main: setup

setup: setup-brew setup-zsh setup-ghq setup-anyenv setup-neovim setup-goenv setup-gitignore-global

setup-brew:
	sh shell/setup-brew.sh

setup-zsh:
	sh shell/setup-zsh.sh

setup-ghq:
	brew install ghq
	brew install peco
	git config --global ghq.root ${GOPATH}/src

setup-goenv:
	sh shell/setup-goenv.sh

setup-anyenv:
	sh shell/setup-anyenv.sh

setup-neovim:
	sh shell/setup-nvim.sh

setup-gitignore-global:
	git config --global core.excludesfile ${PWD}/.gitignore_global

