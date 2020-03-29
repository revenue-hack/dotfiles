#!/bin/sh
if [ ! -e "$HOME/.anyenv" ] ; then
  git clone https://github.com/riywo/anyenv ~/.anyenv
  git clone https://github.com/yyuu/pyenv-virtualenv.git ~/.anyenv/envs/pyenv/plugins/pyenv-virtualenv
  anyenv install --init
fi
if [ ! -e "$HOME/.anyenv/envs/pyenv" ] ; then
  anyenv install pyenv
fi
if [ ! -e "$HOME/.anyenv/envs/nodenv" ] ; then
  anyenv install nodenv
  ndenv install v12.10.0
  # TypescriptのLSP
  npm install -g typescript typescript-language-server
  # PHPのLSP
  npm -g install intelephense
  # VueのLSP
  npm i -g vls
  npm i -g eslint
  npm i -g eslint-loader
  npm i -g eslint-plugin-vue
  # Dockerfile LSP
  npm install -g dockerfile-language-server-nodejs
  # Bash LSP
  npm i -g bash-language-server
  # Vim LSP
  npm install -g vim-language-server
  # SQL LSP
  npm i -g sql-language-server
fi
if [ ! -e "$HOME/.anyenv/envs/phpenv" ] ; then
  anyenv install phpenv
fi
if [ ! -e "$HOME/.anyenv/envs/rbenv" ] ; then
  anyenv install rbenv
  rbenv install 2.6.0
  gem install solargraph
fi

exec $SHELL -l

