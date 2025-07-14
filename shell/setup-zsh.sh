#!/bin/sh
if [ ! -d "$HOME/.oh-my-zsh" ] ; then
  curl -L https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh | sh
  rm -rf ~/.zshrc
  ln -s ~/dotfiles/.zshrc ~/.zshrc
  touch ~/dotfiles/.zshrc.local
  ln -s ~/dotfiles/.zshrc.local ~/.zshrc.local
  # claude
  ln -sf ~/dotfiles/claude/CLAUDE.md ~/.claude/CLAUDE.md
  # gitconfig
  ln -sf ~/dotfiles/.gitconfig ~/.gitconfig
fi
#exec $SHELL
#source ~/.zshrc > /dev/null

