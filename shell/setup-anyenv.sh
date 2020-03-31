#!/bin/zsh
if !(type "anyenv" > /dev/null 2>&1); then
  git clone https://github.com/riywo/anyenv ~/.anyenv
  $SHELL -l
  anyenv install --init
fi

source ~/.zshrc

if !(type "pyenv" > /dev/null 2>&1); then
  anyenv install pyenv
fi

if !(type "nodenv" > /dev/null 2>&1); then
  anyenv install nodenv
fi
if type "nodenv" > /dev/null 2>&1 ; then
  nodenv install 12.10.0
  nodenv global 12.10.0
  source ~/.zshrc
  if type "npm" > /dev/null 2>&1 ; then
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
  else
    echo "npm not found"
  fi
else
  echo "nodenv not found"
fi

if !(type "phpenv" > /dev/null 2>&1); then
  anyenv install phpenv
fi

if !(type "rbenv" > /dev/null 2>&1); then
  anyenv install rbenv
fi

if type "rbenv" > /dev/null 2>&1 ; then
  rbenv install 2.6.0
  source ~/.zshrc
  gem install solargraph
else
  echo "rbenv not found"
fi

$SHELL -l

