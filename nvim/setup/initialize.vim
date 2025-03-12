" neovim python
let g:python_host_prog=$PYENV_ROOT.'/shims/python2'
let g:python3_host_prog=$PYENV_ROOT.'/shims/python3'
let g:ruby_host_prog=$RBENV_ROOT.'/versions/3.0.0/bin/ruby'

set modifiable
set write
colorscheme molokai
syntax on
set t_Co=256
let g:molokai_original = 1
let g:rehash256 = 1
set background=dark
set number
set expandtab
set noignorecase
set autoindent
set smartindent
set encoding=utf-8
set fileencodings=utf-8
" □や○文字が崩れる問題を解決
set ambiwidth=double
"タブの可視化
set list
set listchars=tab:»-,trail:-,eol:↲,extends:»,precedes:«,nbsp:%
"改行文字とタブ文字の色設定（NonTextが改行、SpecialKeyがタブ）
hi NonText guibg=NONE guifg=DarkGreen
"hi SpecialKey guibg=NONE guifg=Gray40
" バックスペースを使えるようにする
set backspace=indent,eol,start
" 隠しファイルを表示する
let NERDTreeShowHidden = 1
set noswapfile
" カーソルが何行目の何列目に置かれているかを表示する
set ruler
" コマンドラインに使われる画面上の行数
set cmdheight=2
" エディタウィンドウの末尾から2行目にステータスラインを常時表示させる
set laststatus=2
" インクリメンタルサーチ. １文字入力毎に検索を行う
set incsearch
" 検索パターンに大文字小文字を区別しない
set ignorecase
" 検索パターンに大文字を含んでいたら大文字小文字を区別する
set smartcase
" 検索結果をハイライト
set hlsearch
" コマンドモードの補完
set wildmenu
" 保存するコマンド履歴の数
set history=5000
" ウインドウのタイトルバーにファイルのパス情報等を表示する
set title
" コマンドラインモードで<Tab>キーによるファイル名補完を有効にする
set wildmenu
" 入力中のコマンドを表示する
set showcmd
" クリップボードへのコピー
set clipboard=unnamedplus
