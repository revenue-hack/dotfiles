#!/bin/sh
if [ ! -d "$HOME/.goenv" ] ; then
  git clone https://github.com/syndbg/goenv.git ~/.goenv
fi
exec $SHELL -l
goenv install 1.14.1 && goenv global 1.14.1
# goのLSPサーバ
go get golang.org/x/tools/gopls@latest

