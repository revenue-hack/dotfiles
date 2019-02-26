#!/bin/sh
if [ ! -d "$HOME/.goenv" ] ; then
  git clone https://github.com/syndbg/goenv.git ~/.goenv
fi
exec $SHELL -l
goenv install 1.11.5 && goenv global 1.11.5

