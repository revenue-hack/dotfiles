#!/bin/sh
if [ ! -d "$HOME/.oh-my-zsh" ] ; then
  curl -L https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh | sh
  rm -rf ~/.zshrc
  ln -s ~/dotfiles/.zshrc ~/.zshrc
fi
source ~/.zshrc > /dev/null

