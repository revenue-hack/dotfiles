#!/bin/sh
pyenv install 3.11.2
pyenv global 3.11.2
brew install neovim
# LSP
pip install python-language-server

exec $SHELL -l

