.PHONY: setup setup-vim-go run

all: main

main: setup setup-vim-go run

setup:
	brew install neovim/neovim/neovim
	unlink ~/.vimrc
	ln -s ${PWD}/init.vim ~/.vimrc
	cd ~/.config && unlink nvim &2> /dev/null
	ln -s ${PWD} ~/.config/nvim
run:
	nvim

