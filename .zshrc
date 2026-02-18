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
export PATH="/Users/ko1014/.local/bin:$PATH"
export PATH="$HOME/.config/composer/vendor/bin:$PATH"


if [ -d $HOME/.pyenv/shims ] ; then
  export PATH=$HOME/.pyenv/shims:$PATH
fi

export RBENV_ROOT="$HOME/.rbenv"
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

#if [ -d $HOME/.goenv/bin ] ; then
  export GOENV_ROOT="$HOME/.goenv"
  export PATH="$GOENV_ROOT/bin:$PATH"
  eval "$(goenv init -)"
#fi

export GOENV_DISABLE_GOPATH=1
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
alias claude-notify='~/claude/scripts/notify_claude_complete.sh'

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
# nvim
export XDG_CONFIG_HOME=$HOME/.config

bindkey '^]' peco-src
function peco-src() {
  #local src=$( ghq list --full-path | peco --query "$LBUFFER")
  local src=$( find $(ghq root)/*/*/* -type d -prune | sed -e 's#'$(ghq root)'/##' | peco --query "$LBUFFER")
  if [ -n "$src" ]; then
    BUFFER="cd ~/go/src/$src"
    zle accept-line
  fi
  zle -R -c
}
zle -N peco-src

checkout-fzf-gitbranch() {
  # 現在のディレクトリがgit worktree内かどうかをチェック
  local IS_WORKTREE=$(git worktree list 2>/dev/null | grep -c "$(pwd)")
  # gitリポジトリのルートディレクトリを取得
  local GIT_ROOT=$(git rev-parse --show-toplevel 2>/dev/null)
  # メインのworktree（.gitがあるディレクトリ）のパスを取得
  local MAIN_WORKTREE=$(git worktree list | head -n1 | awk '{print $1}')
  
  # ブランチ一覧を取得（HEADを除外し、先頭の*やスペースを削除）
  local BRANCHES=$(git branch -vv | grep -v HEAD | sed 's/^[ *]*//')
  # worktree一覧を取得（メインworktreeを除外）
  local WORKTREES=$(git worktree list | tail -n +2 | awk '{print "[worktree] " $1 " " $3}' 2>/dev/null)
  
  # ブランチとworktreeを結合してfzfで選択
  local SELECTION=$(
    {
      # 各ブランチに[branch]タグを付与
      echo "$BRANCHES" | while IFS= read -r line; do
        [ -n "$line" ] && echo "[branch] $line"
      done
      # worktree一覧を追加
      echo "$WORKTREES"
    } | fzf +m
  )
  
  if [ -n "$SELECTION" ]; then
    if echo "$SELECTION" | grep -q '^\[branch\]'; then
      # ブランチ名を抽出（[branch]タグと先頭の+記号を削除し、最初の単語を取得）
      local BRANCH=$(echo "$SELECTION" | sed 's/^\[branch\] //' | sed 's/^+//' | awk '{print $1}')
      
      # ブランチ選択時は常にメインのgitディレクトリでcheckoutを実行
      if [ "$IS_WORKTREE" -gt 0 ] && [ "$MAIN_WORKTREE" != "$(pwd)" ]; then
        # worktree内にいる場合は、メインディレクトリに移動してからcheckout
        BUFFER="cd $MAIN_WORKTREE && git checkout $BRANCH"
      else
        # 既にメインディレクトリにいる場合は、その場でcheckout
        git checkout $BRANCH
      fi
    elif echo "$SELECTION" | grep -q '^\[worktree\]'; then
      # worktreeのパスを抽出して移動
      local WORKTREE_PATH=$(echo "$SELECTION" | sed 's/\[worktree\] //' | awk '{print $1}')
      BUFFER="cd $WORKTREE_PATH"
    fi
  fi
  zle accept-line
}
zle -N checkout-fzf-gitbranch
bindkey '^[' checkout-fzf-gitbranch
bindkey '^r' fzf-history-widget


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

export GOENV_ROOT=$HOME/.goenv
export PATH=$GOENV_ROOT/bin:$PATH
eval "$(goenv init -)"

# ~/.zshrc
# Load personal/local config if it exists
if [ -f ~/.zshrc.local ]; then
  source ~/.zshrc.local
fi

[[ "$TERM_PROGRAM" == "kiro" ]] && . "$(kiro --locate-shell-integration-path zsh)"
