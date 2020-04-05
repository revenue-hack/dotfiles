#!/bin/sh
if !(type "brew" > /dev/null 2>&1) ; then
  /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
fi
brew doctor
brew install fzf
brew install the_silver_searcher

brew install jq tmux

brew install php@7.2
curl -s http://getcomposer.org/installer | php
# PHPENV用
#brew install re2c
#brew install bison@2.7
#brew install openssl
#brew install libxml2
#brew link --force openssl
#brew link --force libxml2
#brew install mcrypt
#brew install libjpeg
#brew install libpng
#brew install icu4c

