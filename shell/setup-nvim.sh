#!/bin/sh
pyenv install -s 2.7
pyenv virtualenv 2.7 neovim-2
pyenv shell neovim-2 && pip install neovim

pyenv install -s 3.6.1
pyenv virtualenv 3.6.1 neovim-3
pyenv shell neovim-3 && pip install neovim

ndenv install -s 10.15 && npm i -g neovim

rbenv install -s 2.6.0 && gem install neovim

