#!/bin/sh
if type "git" > /dev/null 2>&1; then
  git config --global user.name "revenue-hack"
  git config --global user.email K.odeveloper10@gmail.com
  git config --global push.default current
  git config --global core.excludesfile ~/dotfiles/.gitignore_global
  git config --global core.ignorecase false
  echo "success git config"
else
  echo "git command not exist"
fi
