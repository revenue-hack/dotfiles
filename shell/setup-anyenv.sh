#!/bin/sh
if [ ! -e "$HOME/.anyenv" ] ; then
  git clone https://github.com/riywo/anyenv ~/.anyenv
fi
if [ ! -e "$HOME/.anyenv/envs/pyenv" ] ; then
  anyenv install pyenv
fi
if [ ! -e "$HOME/.anyenv/envs/ndenv" ] ; then
  anyenv install ndenv
fi
if [ ! -e "$HOME/.anyenv/envs/phpenv" ] ; then
  anyenv install phpenv
fi
if [ ! -e "$HOME/.anyenv/envs/rbenv" ] ; then
  anyenv install rbenv
fi

exec $SHELL -l

