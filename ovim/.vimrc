set nocompatible               " Be iMproved
filetype off                   " Required!
if &compatible
  set nocompatible
endif
if has('vim_starting')
  set runtimepath+=~/.vim/bundle/neobundle.vim/
endif
set number
colorscheme molokai
syntax on
let g:molokai_original = 1
let g:rehash256 = 1
set background=dark
set t_Co=256
"タブ関連(phpは4スペ)
if has("autocmd")
  autocmd VimEnter * execute 'NERDTree'
  "fileタイプの検索を有効にする
  filetype plugin on
  " fileタイプに合わせてインデントを利用する
  filetype indent on
  autocmd FileType php setlocal sw=4 sts=4 ts=4 et
  autocmd FileType html setlocal sw=2 sts=2 ts=2 et
  autocmd FileType javascript  setlocal sw=2 sts=2 ts=2 et
  autocmd FileType sql setlocal sw=2 sts=2 ts=2 et
  autocmd FileType twig setlocal sw=2 sts=2 ts=2 et
  autocmd FileType xml setlocal sw=2 sts=2 ts=2 et
  autocmd FileType yaml setlocal sw=2 sts=2 ts=2 et
  autocmd FileType zsh setlocal sw=2 sts=2 ts=2 et
  autocmd FileType vim setlocal sw=2 sts=2 ts=2 et
  autocmd FileType css setlocal sw=2 sts=2 ts=2 et
  autocmd FileType scss setlocal sw=2 sts=2 ts=2 et
  autocmd FileType java setlocal sw=2 sts=2 ts=2 et
  autocmd FileType sass setlocal sw=2 sts=2 ts=2 et
  autocmd FileType jade setlocal sw=2 sts=2 ts=2 et
  autocmd FileType ruby setlocal sw=2 sts=2 ts=2 et
  autocmd FileType vcl setlocal sw=2 sts=2 ts=2 et
  autocmd FileType swift setlocal sw=2 sts=2 ts=2 et
  autocmd FileType vue syntax sync fromstart
  autocmd FileType vue setlocal sw=2 sts=2 ts=2 et
  autocmd FileType json setlocal sw=2 sts=2 ts=2 et
  autocmd FileType python setlocal sw=4 sts=4 ts=4 et
  autocmd FileType vue syntax sync fromstart
  autocmd BufRead,BufNewFile *.ts set filetype=typescript
  autocmd FileType typescript setlocal sw=2 sts=2 ts=2 et
  syntax enable
endif
let g:vue_disable_pre_processors=1
set expandtab
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
" バッファ一覧
noremap <C-P> :Unite buffer<CR>
" ファイル一覧
noremap <C-N> :Unite -buffer-name=file file<CR>
" 最近使ったファイルの一覧
noremap <C-Z> :Unite file_mru<CR>


"括弧自動挿入
imap { {}<LEFT>
imap [ []<LEFT>
imap ( ()<LEFT>

" phpcs,phpmd
"execute pathogen#infect()
"set statusline+=%#warningmsg#
"set statusline+=%{SyntasticStatuslineFlag()}
"set statusline+=%*

"let g:syntastic_always_populate_loc_list = 1
"let g:syntastic_auto_loc_list = 1
"let g:syntastic_check_on_open = 1
"let g:syntastic_check_on_wq = 0
"let g:syntastic_php_checkers = ['php', 'phpcs', 'phpmd']
"let g:syntastic_php_phpcs_args='--standard=psr2'
"nerdtreeのalias
nnoremap <silent><C-e> :NERDTreeToggle<CR>
"ファイル移動のalias
noremap <C-l> <C-w>l
noremap <C-j> <C-w>j
noremap <C-h> <C-w>h
noremap <C-k> <C-w>k
" 0レジスタ貼り付け
nnoremap <C-p> "0p
" Dのヤンク
nnoremap Y y$
let mapleader = ","
" php-cs
let g:php_cs_fixer_path = "php"
" psr0 psr1 psr2 allを指定
let g:php_cs_fixer_level = "psr2"
" default sf20 sf21の指定symfonyの指定などの構造確認？が出来る
let g:php_cs_fixer_config = "default"
" phpコマンドの場所
let g:php_cs_fixer_php_path = "/usr/bin/phpcs"
" フィルター（http://cs.sensiolabs.org/ここにしていされているやつが使える）
let g:php_cs_fixer_fixers_list = "psr0"
let g:php_cs_fixer_enable_default_mapping = 1
let g:php_cs_fixer_verbose = 0
nnoremap <C-a> :call PhpCsFixerFixDirectory()<CR>
nnoremap <silent><leader>pcf :call PhpCsFixerFixFile()<CR>
" スワップファイルは使わない(ときどき面倒な警告が出るだけで役に立ったことがない)
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

call neobundle#begin(expand('~/.vim/bundle/'))
  NeoBundleFetch 'Shougo/neobundle.vim'

  filetype plugin indent on
  if neobundle#exists_not_installed_bundles()
    echomsg 'Not installed bundles : ' .
          \ string(neobundle#get_not_installed_bundle_names())
    echomsg 'Please execute ":NeoBundleInstall" command.'
  else
    NeoBundle 'tpope/vim-fugitive'
    NeoBundle 'posva/vim-vue'
    " 行末の半角スペースを可視化
    NeoBundle 'bronson/vim-trailing-whitespace'
    "unite.vim ファイル開きを簡略化
    NeoBundle 'Shougo/unite.vim'
    "Unite.vimで最近使ったファイルを表示できるようにする
    NeoBundle 'Shougo/neomru.vim'
    " コメントON/OFFを手軽に実行Ctrl+ハイフン
    NeoBundle 'tomtom/tcomment_vim'
    " typescript
    NeoBundle 'typescript-vim'
    " editconfig
    "NeoBundle 'editorconfig/editorconfig-vim'
    "nerdtree
    NeoBundle 'scrooloose/nerdtree'
    " インデントの可視化
    NeoBundle 'Yggdroot/indentLine'
    " gitを使う
    NeoBundle 'cohama/agit.vim'
    "html自動補完Ctrl+y,
    NeoBundle 'mattn/emmet-vim'
    " html5補完機能
    NeoBundle 'Shougo/neocomplcache'
    " html5のコードをシンタックス表示する
    NeoBundle 'hail2u/vim-css3-syntax'
    " Varnish vcl
    NeoBundle 'pld-linux/vim-syntax-vcl'
    " php-cs-filter
    NeoBundle 'stephpy/vim-php-cs-fixer'
    " jadeのsyntaxハイライト
    NeoBundle 'smerrill/vcl-vim-plugin'
    NeoBundle 'digitaltoad/vim-jade'
    " vue
    NeoBundle 'posva/vim-vue'
    " ブラウザを自動更新する
    NeoBundle 'tell-k/vim-browsereload-mac'
    NeoBundle 'tomasr/molokai'
    NeoBundle 'Shougo/unite.vim'
    NeoBundle 'ujihisa/unite-colorscheme'
    if has('lua') " lua機能が有効になっている場合・・・・・・①
      " コードの自動補完
      NeoBundle 'Shougo/neocomplete.vim'
      " スニペットの補完機能
      NeoBundle "Shougo/neosnippet.vim"
      " スニペット集
      NeoBundle 'Shougo/neosnippet-snippets'
    endif
  endif
call neobundle#end()

"----------------------------------------------------------
" neocomplete・neosnippetの設定
"----------------------------------------------------------
if neobundle#is_installed('neocomplete.vim')
  " Vim起動時にneocompleteを有効にする
  let g:neocomplete#enable_at_startup = 1
  " smartcase有効化.
  " 大文字が入力されるまで大文字小文字の区別を無視する
  let g:neocomplete#enable_smart_case = 1
  " 3文字以上の単語に対して補完を有効にする
  let g:neocomplete#min_keyword_length = 3
  " 区切り文字まで補完する
  let g:neocomplete#enable_auto_delimiter = 1
  let g:neocomplete#auto_completion_start_length = 1
  " バックスペースで補完のポップアップを閉じる
  "inoremap <expr><BS> neocomplete#smart_close_popup()."<C-h>"
  "エンターキーで補完候補の確定.スニペットの展開もエンターキーで確定・・・・・・②
  " imap <expr><CR> neosnippet#expandable() ? "<Plug>(neosnippet_expand_or_jump)" : pumvisible() ? "<C-y>" : "<CR>"
  " タブキーで補完候補の選択.スニペット内のジャンプもタブキーでジャンプ・・・・・・③
  inoremap <expr><TAB>  pumvisible() ? "\<C-n>" : "\<TAB>"
endif
