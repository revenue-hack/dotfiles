if [ -d $HOME/.oh-my-zsh ] ; then
  export ZSH=$HOME/.oh-my-zsh
fi
if [ -f $ZSH/oh-my-zsh.sh ] ; then
  source $ZSH/oh-my-zsh.sh
fi

eval "$(starship init zsh)"

export PATH=$HOME/bin:$PATH

ZSH_THEME="gallois"

export NODE_PATH="/usr/local/lib/node_modules"
export PATH=/usr/local/apache-maven-3.5.0/bin:$PATH
export CATALINA_HOME="/Applications/apache-tomcat-8.0.41"
export PATH="$CATALINA_HOME/bin:$CATALINA_HOME/lib:$PATH"
export PATH=/usr/local/bin:$PATH
export PATH=$HOME/google-cloud-sdk/bin:$PATH
export VOLTA_HOME="$HOME/.volta"
export PATH="$VOLTA_HOME/bin:$PATH"
export PATH="$HOME/.rd/bin:$PATH"

if [ -d $HOME/.pyenv/shims ] ; then
  export PATH=$HOME/.pyenv/shims:$PATH
fi

if [ -d $HOME/.rbenv/shims ] ; then
  export PATH=$HOME/.rbenv/shims:$PATH
fi

if [ -d $HOME/.phpenv/shims ] ; then
  export PATH=$HOME/.phpenv/shims:$PATH
fi

if [ -d $HOME/.ndenv/shims ] ; then
  export PATH=$HOME/.ndenv/shims:$PATH
fi

if [ -d $HOME/.anyenv/envs/nodenv/bin ] ; then
  export PATH="$PATH:$HOME/.anyenv/envs/nodenv/bin"
  eval "$(nodenv init - zsh)"
fi

if [ -d $HOME/.anyenv/envs/rbenv/bin ] ; then
  export PATH="$PATH:$HOME/.anyenv/envs/rbenv/bin"
  eval "$(rbenv init - zsh)"
fi
if [ -d $HOME/.anyenv/envs/phpenv/bin ] ; then
  export PATH="$PATH:$HOME/.anyenv/envs/phpenv/bin"
  eval "$(phpenv init - zsh)"
fi
if [ -d $HOME/.anyenv/envs/pyenv/bin ] ; then
  export PATH="$PATH:$HOME/.anyenv/envs/pyenv/bin"
  eval "$(pyenv init - zsh)"
  #eval "$(pyenv virtualenv-init - zsh)"
fi
if [ -d "$HOME/.anyenv" ] ; then
  export ANYENV_ROOT="$HOME/.anyenv"
  export PATH="$HOME/.anyenv/bin:$PATH"
  eval "$(anyenv init - zsh)"
  # tmux対応
  for D in `\ls $HOME/.anyenv/envs`
  do
    export PATH="$HOME/.anyenv/envs/$D/shims:$PATH"
  done
fi

export EDITOR=nvim
if type "direnv" > /dev/null 2>&1 ; then
  eval "$(direnv hook zsh)"
fi

if [ -d $HOME/.goenv/bin ] ; then
  export GOENV_ROOT="$HOME/.goenv"
  export PATH="$GOENV_ROOT/bin:$PATH"
  eval "$(goenv init -)"
fi

export GOPATH=$HOME/go
export PATH=$GOROOT/bin:$PATH
export PATH=$GOPATH/bin:$PATH

export PATH=$PATH:$HOME/pear/bin

# PHPENVのため
#if [ -d /usr/local/opt/bison@2.7 ] ; then
#  export PATH="/usr/local/opt/bison@2.7/bin:$PATH"
#fi
#if [ -d /usr/local/opt/libxml2 ] ; then
#  export PATH="/usr/local/opt/libxml2/bin:$PATH"
#fi
#if [ -d /usr/local/opt/openssl@1.1 ] ; then
#  export PATH="/usr/local/opt/openssl@1.1/bin:$PATH"
#fi

if [ -d /usr/local/opt/php@7.2 ] ; then
  export PATH="/usr/local/opt/php@7.2/bin:$PATH"
fi

# Which plugins would you like to load? (plugins can be found in ~/.oh-my-zsh/plugins/*)
# Custom plugins may be added to ~/.oh-my-zsh/custom/plugins/
# Example format: plugins=(rails git textmate ruby lighthouse)
# Add wisely, as too many plugins slow down shell startup.
plugins=(git)

# export MANPATH="/usr/local/man:$MANPATH"

# alias
alias grn='grep -r -n'
alias e='emacsclient -nw -a ""'
alias dis="docker images --format '{{.Size}}\t{{.Repository}}\t{{.Tag}}' | sort -r"
alias gll='gll'
alias zshconfig="mate ~/.zshrc"
alias ohmyzsh="mate ~/.oh-my-zsh"
alias gcurl='curl -H "Authorization: Bearer $(gcloud auth print-identity-token)"'
alias ext='exa --long --tree'
alias exla='exa -la'
alias n='nvim'
alias psw='procs --watch --watch-interval 5'
alias ps='procs'
alias chf='chflags'
alias chfru='chflags -R uchg'
alias chfrnou='chflags -R nouchg'
alias chfnou='chflags nouchg'

## git pull 時に --set-upstream-to しろというエラーが出た時に自動処理させる
function gll() {
  ## カレントブランチ名
  local current_branch_name=$(git rev-parse --abbrev-ref @)
  ## リモートブランチを指定して git pull する
  git branch --set-upstream-to="origin/$current_branch_name" "$current_branch_name"
  git pull origin "$current_branch_name"
}

#THIS MUST BE AT THE END OF THE FILE FOR SDKMAN TO WORK!!!
export SDKMAN_DIR="~/.sdkman"
[[ -s "$SDKMAN_DIR/bin/sdkman-init.sh" ]] && source "$SDKMAN_DIR/bin/sdkman-init.sh"

export DOCKER_BUILDKIT=1
export GO111MODULE=auto
# nvim
export XDG_CONFIG_HOME=$HOME/.config

bindkey '^]' peco-src
function peco-src() {
  #local src=$( ghq list --full-path | peco --query "$LBUFFER")
  local src=$( find $(ghq root)/*/*/* -type d -prune | sed -e 's#'$(ghq root)'/##' | peco --query "$LBUFFER")
  if [ -n "$src" ]; then
    BUFFER="cd $GOPATH/src/$src"
    zle accept-line
  fi
  zle -R -c
}
zle -N peco-src

checkout-fzf-gitbranch() {
  local GIT_BRANCH=$(git branch -vv | grep -v HEAD | fzf +m)
  if [ -n "$GIT_BRANCH" ]; then
    #git checkout $(echo "$GIT_BRANCH" | sed "s/.* //" | sed "s#remotes/[^/]*/##")
    git checkout $(echo "$GIT_BRANCH" | awk '{print $1}')
  fi
  zle accept-line
}
zle -N checkout-fzf-gitbranch
bindkey '^[' checkout-fzf-gitbranch

export GOOGLE_APPLICATION_CREDENTIALS=~/.gcloud/key.json

# tabtab source for serverless package
# uninstall by removing these lines or running `tabtab uninstall serverless`
[[ -f /Users/ko1014/.anyenv/envs/ndenv/versions/10.15/lib/node_modules/serverless/node_modules/tabtab/.completions/serverless.zsh ]] && . /Users/ko1014/.anyenv/envs/ndenv/versions/10.15/lib/node_modules/serverless/node_modules/tabtab/.completions/serverless.zsh
# tabtab source for sls package
# uninstall by removing these lines or running `tabtab uninstall sls`
[[ -f /Users/ko1014/.anyenv/envs/ndenv/versions/10.15/lib/node_modules/serverless/node_modules/tabtab/.completions/sls.zsh ]] && . /Users/ko1014/.anyenv/envs/ndenv/versions/10.15/lib/node_modules/serverless/node_modules/tabtab/.completions/sls.zsh
# tabtab source for slss package
# uninstall by removing these lines or running `tabtab uninstall slss`
[[ -f /Users/ko1014/.anyenv/envs/ndenv/versions/10.15/lib/node_modules/serverless/node_modules/tabtab/.completions/slss.zsh ]] && . /Users/ko1014/.anyenv/envs/ndenv/versions/10.15/lib/node_modules/serverless/node_modules/tabtab/.completions/slss.zsh


[ -f ~/.fzf.zsh ] && source ~/.fzf.zsh
## phpenv用
export PATH="/usr/local/opt/bison@2.7/bin:$PATH"

export NVIM_COC_LOG_FILE=~/.cache/tmp/coc.log
# >>> conda initialize >>>
# !! Contents within this block are managed by 'conda init' !!
__conda_setup="$('~/opt/anaconda3/bin/conda' 'shell.zsh' 'hook' 2> /dev/null)"
if [ $? -eq 0 ]; then
    eval "$__conda_setup"
else
    if [ -f "~/opt/anaconda3/etc/profile.d/conda.sh" ]; then
        . "~/opt/anaconda3/etc/profile.d/conda.sh"
    else
        export PATH="~/opt/anaconda3/bin:$PATH"
    fi
fi
unset __conda_setup
# <<< conda initialize <<<

