#!/bin/sh
if [ ! -d "$HOME/.goenv" ] ; then
  git clone https://github.com/syndbg/goenv.git ~/.goenv
fi

source ~/.zshrc > /dev/null 2>&1

if type "goenv" > /dev/null 2>&1 ; then
  goenv install 1.14.1 && goenv global 1.14.1
  # goのLSPサーバ
  go get golang.org/x/tools/gopls@latest

  # terraform LSP
  git clone https://github.com/juliosueiras/terraform-lsp.git
  cd terraform-lsp
  GO111MODULE=on go mod download
  make
  make copy
  cd ../ && rm -rf terraform-lsp
else
  echo "goenv not found"
fi

