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

