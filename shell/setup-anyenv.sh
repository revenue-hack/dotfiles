#!/bin/sh
if !(type "anyenv" > /dev/null 2>&1); then
  git clone https://github.com/riywo/anyenv ~/.anyenv
  git clone https://github.com/yyuu/pyenv-virtualenv.git ~/.anyenv/envs/pyenv/plugins/pyenv-virtualenv
  exec $SHELL -l
  anyenv install --init
  source ~/.zshrc
fi

if !(type "pyenv" > /dev/null 2>&1); then
  anyenv install pyenv
fi

if !(type "ndenv" > /dev/null 2>&1); then
  anyenv install nodenv
fi
if type "ndenv" > /dev/null 2>&1 ; then
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

if !(type "phpenv" > /dev/null 2>&1); then
  anyenv install phpenv
fi

if !(type "rbenv" > /dev/null 2>&1); then
  anyenv install rbenv
fi

if type "rbenv" > /dev/null 2>&1 ; then
  rbenv install 2.6.0
  exec $SHELL -l
  gem install solargraph
fi

exec $SHELL -l

