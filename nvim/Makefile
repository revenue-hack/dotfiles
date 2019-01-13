.PHONY: setup setup-vim-go run

all: main

main: setup setup-vim-go run

setup:
	brew install neovim/neovim/neovim
	unlink ~/.vimrc
	ln -s ${PWD} ~/.vimrc
	unlink ~/.config/nvim
	ln -s ${PWD} ~/.config/nvim

# install gometalinter for vim-go plugin
setup-vim-go:
	brew tap alecthomas/homebrew-tap
	brew install gometalinter
	cd ${GOPATH}/src/github.com/nsf/gocode && go install

run:
	nvim

