#!/bin/sh
if [ ! -d "$HOME/.goenv" ] ; then
  git clone https://github.com/syndbg/goenv.git ~/.goenv
fi

if [ -d "$HOME/.goenv" ] ; then
  export GOENV_ROOT="$HOME/.goenv"
  export PATH="$GOENV_ROOT/bin:$PATH"
  eval "$(goenv init -)"
fi

if type "goenv" > /dev/null 2>&1 ; then
  goenv install 1.18.0 && goenv global 1.18.0
  # goのLSPサーバ
  go install golang.org/x/tools/gopls

  # gPRC, ProtocolBuffers
  go install  github.com/golang/protobuf/protoc-gen-go
  go install google.golang.org/grpc
  go install honnef.co/go/tools/cmd/staticcheck
  go install github.com/go-delve/delve/cmd/dlv@latest
  # terraform LSP
  git clone https://github.com/juliosueiras/terraform-lsp.git
  cd terraform-lsp
  GO111MODULE=on go mod download
  make
  cp ./terraform-lsp /usr/local/bin
  cd ../ && rm -rf terraform-lsp
  # VimなどのLSP
  go get github.com/mattn/efm-langserver
else
  echo "goenv not found"
fi

