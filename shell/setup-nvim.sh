#!/bin/sh
brew install neovim
CFLAGS="-I$(xcrun --show-sdk-path)/usr/include" pyenv install -s 3.6.5
pip3 install pynvim
pyenv global 3.6.5
pyenv rehash
# LSP
pyenv global 3.6.5
pip install python-language-server

nodenv install -s 14.15.3 && npm i -g neovim

rbenv install -s 3.0.0 && rbenv global 3.0.0 && gem install neovim

exec $SHELL -l

