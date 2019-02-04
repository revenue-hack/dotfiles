# Path to your oh-my-zsh installation.

# The next line updates PATH for the Google Cloud SDK.
if [ -f '~/google-cloud-sdk/path.zsh.inc' ]; then source '~/google-cloud-sdk/path.zsh.inc'; fi

# The next line enables shell command completion for gcloud.
if [ -f '~/google-cloud-sdk/completion.zsh.inc' ]; then source '~/google-cloud-sdk/completion.zsh.inc'; fi

if [ -d $HOME/.oh-my-zsh ] ; then
  export ZSH=$HOME/.oh-my-zsh
fi
export PATH=$HOME/bin:$PATH
export NODE_PATH="/usr/local/lib/node_modules"
export PATH=/usr/local/apache-maven-3.5.0/bin:$PATH
export CATALINA_HOME="/Applications/apache-tomcat-8.0.41"
export PATH="$CATALINA_HOME/bin:$CATALINA_HOME/lib:$PATH"
export PATH=/usr/local/bin:$PATH
export PATH=$HOME/.rbenv/shims:$PATH
export PATH=$HOME/.pyenv/shims:$PATH
export PATH=$HOME/.phpenv/shims:$PATH
export PATH=$HOME/.ndenv/shims:$PATH
export PATH=$HOME/google-cloud-sdk/bin:$PATH
export PATH=/usr/local/opt/mysql@5.7/bin:$PATH
# Set name of the theme to load.
# Look in ~/.oh-my-zsh/themes/
# Optionally, if you set this to "random", it'll load a random theme each
# time that oh-my-zsh is loaded.
ZSH_THEME="gallois"
if [ -d .anyenv/envs/rbenv/bin ] ; then
  export PATH="$PATH:$HOME/.anyenv/envs/rbenv/bin"
  eval "$(rbenv init - zsh)"
fi
if [ -d .anyenv/envs/phpenv/bin ] ; then
  export PATH="$PATH:$HOME/.anyenv/envs/phpenv/bin"
  eval "$(rbenv init - zsh)"
fi
if [ -d .anyenv/envs/pyenv/bin ] ; then
  export PATH="$PATH:$HOME/.anyenv/envs/pyenv/bin"
  eval "$(rbenv init - zsh)"
fi
if [ -d "$HOME/.anyenv" ] ; then
  export ANYENV_ROOT="$HOME/.anyenv"
  export PATH="$HOME/.anyenv/bin:$PATH"
  eval "$(anyenv init -)"
  eval "$(pyenv virtualenv-init -)"
  # tmux対応
  for D in `\ls $HOME/.anyenv/envs`
  do
    export PATH="$HOME/.anyenv/envs/$D/shims:$PATH"
  done
fi
if [ -d $HOME/.goenv ] ; then
  export PATH="$HOME/.goenv/bin:$PATH"
  eval "$(goenv init -)"
fi

export GOPATH=$HOME/go
export GOROOT=$HOME/go
#export CLASSPATH=/Applications/Eclipse_4.6.2.app/Contents/workspace/Yasui_Ozeki/build/classes
# Uncomment the following line to use case-sensitive completion.
# CASE_SENSITIVE="true"

# Uncomment the following line to use hyphen-insensitive completion. Case
# sensitive completion must be off. _ and - will be interchangeable.
# HYPHEN_INSENSITIVE="true"

# Uncomment the following line to disable bi-weekly auto-update checks.
# DISABLE_AUTO_UPDATE="true"

# Uncomment the following line to change how often to auto-update (in days).
# export UPDATE_ZSH_DAYS=13

# Uncomment the following line to disable colors in ls.
# DISABLE_LS_COLORS="true"

# Uncomment the following line to disable auto-setting terminal title.
# DISABLE_AUTO_TITLE="true"

# Uncomment the following line to enable command auto-correction.
# ENABLE_CORRECTION="true"

# Uncomment the following line to display red dots whilst waiting for completion.
# COMPLETION_WAITING_DOTS="true"

# Uncomment the following line if you want to disable marking untracked files
# under VCS as dirty. This makes repository status check for large repositories
# much, much faster.
# DISABLE_UNTRACKED_FILES_DIRTY="true"

# Uncomment the following line if you want to change the command execution time
# stamp shown in the history command output.
# The optional three formats: "mm/dd/yyyy"|"dd.mm.yyyy"|"yyyy-mm-dd"
# HIST_STAMPS="mm/dd/yyyy"

# Would you like to use another custom folder than $ZSH/custom?
# ZSH_CUSTOM=/path/to/new-custom-folder

# Which plugins would you like to load? (plugins can be found in ~/.oh-my-zsh/plugins/*)
# Custom plugins may be added to ~/.oh-my-zsh/custom/plugins/
# Example format: plugins=(rails git textmate ruby lighthouse)
# Add wisely, as too many plugins slow down shell startup.
plugins=(git)

# User configuration

# export MANPATH="/usr/local/man:$MANPATH"

if [ -f $ZSH/oh-my-zsh.sh ] ; then
  source $ZSH/oh-my-zsh.sh
fi

# You may need to manually set your language environment
# export LANG=en_US.UTF-8

# Preferred editor for local and remote sessions
# if [[ -n $SSH_CONNECTION ]]; then
#   export EDITOR='vim'
# else
#   export EDITOR='mvim'
# fi

# Compilation flags
# export ARCHFLAGS="-arch x86_64"

# ssh
# export SSH_KEY_PATH="~/.ssh/dsa_id"

# Set personal aliases, overriding those provided by oh-my-zsh libs,
# plugins, and themes. Aliases can be placed here, though oh-my-zsh
# users are encouraged to define aliases within the ZSH_CUSTOM folder.
# For a full list of active aliases, run `alias`.
# alias
alias grn='grep -r -n'
alias e='emacsclient -nw -a ""'
#
# Example aliases
 alias zshconfig="mate ~/.zshrc"
 alias ohmyzsh="mate ~/.oh-my-zsh"

#THIS MUST BE AT THE END OF THE FILE FOR SDKMAN TO WORK!!!
export SDKMAN_DIR="~/.sdkman"
[[ -s "$SDKMAN_DIR/bin/sdkman-init.sh" ]] && source "$SDKMAN_DIR/bin/sdkman-init.sh"

# GO MODULE
export GO111MODULE=on
# nvim
export XDG_CONFIG_HOME=$HOME/.config

bindkey '^]' peco-src
function peco-src() {
  local src=$( ghq list --full-path | peco --query "$LBUFFER")
  if [ -n "$src" ]; then
    BUFFER="cd $src"
    zle accept-line
  fi
  zle -R -c
}
zle -N peco-src

