#!/bin/sh

if !(type "anyenv" > /dev/null 2>&1); then
  git clone https://github.com/riywo/anyenv ~/.anyenv
fi


if [ -d "$HOME/.anyenv" ] ; then
  export ANYENV_ROOT="$HOME/.anyenv"
  export PATH="$HOME/.anyenv/bin:$PATH"
  eval "$(anyenv init - zsh)"
  anyenv install --init
fi

if !(type "pyenv" > /dev/null 2>&1); then
  anyenv install pyenv > /dev/null
  exec $SHELL -l
fi

if !(type "volta" > /dev/null 2>&1); then
  curl https://get.volta.sh | bash
fi

if [ -d $HOME/.volta/bin/volta ] ; then
  exec $SHELL -l
  volta install node
fi

if !(type "phpenv" > /dev/null 2>&1); then
  anyenv install phpenv > /dev/null
  exec $SHELL -l
fi

if !(type "rbenv" > /dev/null 2>&1); then
  anyenv install rbenv > /dev/null
  exec $SHELL -l
fi

if [ -d $HOME/.anyenv/envs/rbenv/bin ] ; then
  export PATH="$PATH:$HOME/.anyenv/envs/rbenv/bin"
  eval "$(rbenv init - zsh)"
fi

#if type "phpenv" > /dev/null 2>&1 ; then
#   CONFIGURE_OPTS="--with-zlib-dir=$(brew --prefix zlib) --with-bz2=$(brew --prefix bzip2) --with-curl=$(brew --prefix curl) --with-iconv=$(brew --prefix libiconv) --with-libedit=$(brew --prefix libedit) --with-readline=$(brew --prefix readline) --with-tidy=$(brew --prefix tidy-html5)" phpenv install 7.1.33
#  source ~/.zshrc > /dev/null 2>&1
#fi

#if type "rbenv" > /dev/null 2>&1 ; then
#  rbenv install 3.0.0
#  source ~/.zshrc > /dev/null 2>&1
#  gem install solargraph
#else
#  echo "rbenv not found"
#fi

exec $SHELL -l

