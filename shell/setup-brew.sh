#!/bin/sh
if !(type "brew" > /dev/null 2>&1) ; then
  /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
fi
brew doctor
brew install git fzf the_silver_searcher pwgen rename

brew install jq tmux tfenv protobuf
tfenv install 0.12.19

brew cask install bloomrpc

# PHPENV用
#brew install autoconf bison bzip2 curl icu4c libedit libjpeg libiconv libpng libxml2 libzip openssl re2c tidy-html5 zlib

