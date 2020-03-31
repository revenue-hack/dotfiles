.PHONY: setup setup-brew setup-ghq setup-anyenv setup-neovim setup-gitignore-global

all: main

main: setup

setup: setup-brew setup-git setup-zsh setup-ghq setup-anyenv setup-neovim setup-goenv setup-gitignore-global

setup-git:
	./shell/setup-git.sh

setup-brew:
	./shell/setup-brew.sh

setup-zsh:
	./shell/setup-zsh.sh

setup-ghq:
	brew install ghq
	brew install peco
	git config --global ghq.root ${GOPATH}/src
	source ~/.zshrc

setup-goenv:
	./shell/setup-goenv.sh

setup-anyenv:
	./shell/setup-anyenv.sh

setup-neovim:
	./shell/setup-nvim.sh

setup-gitignore-global:
	git config --global core.excludesfile ${PWD}/.gitignore_global

