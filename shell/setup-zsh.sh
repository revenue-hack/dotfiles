#!/bin/sh
if [! -d $HOME/.oh-my-zsh ] ; then
  brew install zsh
  curl -L https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh | sh
  ln -s ~/dotfiles/.zshrc ~/.zshrc
  sudo chsh -s /usr/local/bin/zsh
  source ~/.zshrc > /dev/null
fi

