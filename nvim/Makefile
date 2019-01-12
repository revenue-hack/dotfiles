.PHONY: setup

all: main

main: setup

setup:
	/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
	brew doctor
	brew install neovim/neovim/neovim
	unlink ~/.vimrc
	ln -s ${PWD} ~/.vimrc
	unlink ~/.config/nvim
	ln -s ${PWD} ~/.config/nvim
	nvim

