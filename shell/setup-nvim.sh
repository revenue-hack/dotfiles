#!/bin/sh
brew install neovim
CFLAGS="-I$(xcrun --show-sdk-path)/usr/include" pyenv install -s 2.7
pyenv virtualenv 2.7 neovim-2
pyenv shell neovim-2 && pip install neovim
CFLAGS="-I$(xcrun --show-sdk-path)/usr/include" pyenv install -s 2.7.15
CFLAGS="-I$(xcrun --show-sdk-path)/usr/include" pyenv install -s 3.6.5
pyenv global 3.6.5 2.7.15
pyenv rehash
# LSP
pyenv global 3.6.5
pip install python-language-server

CFLAGS="-I$(xcrun --show-sdk-path)/usr/include" pyenv install -s 3.6.1
pyenv virtualenv 3.6.1 neovim-3
pyenv shell neovim-3 && pip install neovim

nodenv install -s 10.15 && npm i -g neovim

rbenv install -s 2.6.0 && rbenv global 2.6.0 && gem install neovim

exec $SHELL -l

