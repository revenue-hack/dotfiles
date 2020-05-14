#!/bin/sh

if !(type "anyenv" > /dev/null 2>&1); then
  git clone https://github.com/riywo/anyenv ~/.anyenv
fi


if [ -d "$HOME/.anyenv" ] ; then
  export ANYENV_ROOT="$HOME/.anyenv"
  export PATH="$HOME/.anyenv/bin:$PATH"
  eval "$(anyenv init - zsh)"
fi

if !(type "pyenv" > /dev/null 2>&1); then
  anyenv install pyenv
fi

if !(type "nodenv" > /dev/null 2>&1); then
  anyenv install nodenv
fi

if [ -d $HOME/.anyenv/envs/nodenv/bin ] ; then
  export PATH="$PATH:$HOME/.anyenv/envs/nodenv/bin"
  eval "$(nodenv init - zsh)"
fi

if type "nodenv" > /dev/null 2>&1 ; then
  nodenv install 12.10.0
  nodenv global 12.10.0

  if type "npm" > /dev/null 2>&1 ; then
    ./install-npm.sh
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

if [ -d $HOME/.anyenv/envs/rbenv/bin ] ; then
  export PATH="$PATH:$HOME/.anyenv/envs/rbenv/bin"
  eval "$(rbenv init - zsh)"
fi

if type "rbenv" > /dev/null 2>&1 ; then
   CONFIGURE_OPTS="--with-zlib-dir=$(brew --prefix zlib) --with-bz2=$(brew --prefix bzip2) --with-curl=$(brew --prefix curl) --with-iconv=$(brew --prefix libiconv) --with-libedit=$(brew --prefix libedit) --with-readline=$(brew --prefix readline) --with-tidy=$(brew --prefix tidy-html5)" phpenv install 7.1.33
  source ~/.zshrc > /dev/null 2>&1
fi

if type "rbenv" > /dev/null 2>&1 ; then
  rbenv install 2.6.0
  source ~/.zshrc > /dev/null 2>&1
  gem install solargraph
else
  echo "rbenv not found"
fi

source ~/.zshrc > /dev/null 2>&1

