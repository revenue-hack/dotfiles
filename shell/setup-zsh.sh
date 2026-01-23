#!/bin/sh
if [ ! -d "$HOME/.oh-my-zsh" ] ; then
  curl -L https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh | sh
  rm -rf ~/.zshrc
  ln -s ~/dotfiles/.zshrc ~/.zshrc
  touch ~/dotfiles/.zshrc.local
  ln -s ~/dotfiles/.zshrc.local ~/.zshrc.local
  # claude
  ln -sf ~/dotfiles/claude/CLAUDE.md ~/.claude/CLAUDE.md
  ln -sf ~/dotfiles/claude/settings.json ~/.claude/settings.json
  ln -sf ~/dotfiles/claude/agents/* ~/.claude/agents/
  ln -sf ~/dotfiles/claude/skills/* ~/.claude/skills/
  # gitconfig
  ln -sf ~/dotfiles/.gitconfig ~/.gitconfig
  # gemini
  ln -sf ~/dotfiles/gemini/settings.json ~/.gemini/settings.json
fi
#exec $SHELL
#source ~/.zshrc > /dev/null

