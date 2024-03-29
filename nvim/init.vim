let s:dein_dir = expand('~/.cache/dein')
let s:dein_repo_dir = s:dein_dir . '/repos/github.com/Shougo/dein.vim'
" when dein nothing
if &runtimepath !~# '/dein.vim'
  if !isdirectory(s:dein_repo_dir)
    execute '!git clone https://github.com/Shougo/dein.vim' s:dein_repo_dir
  endif
  execute 'set runtimepath^=' . fnamemodify(s:dein_repo_dir, ':p')
endif

augroup vimrc
  autocmd!
augroup end
" dein install
if dein#load_state(s:dein_dir)
  call dein#begin(s:dein_dir)
    let s:rc_dir = expand('~/dotfiles/nvim')
    let s:toml = s:rc_dir . '/dein.toml'
    let s:lazy_toml = s:rc_dir . '/dein_lazy.toml'
    call dein#load_toml(s:toml, {'lazy': 0})
    call dein#load_toml(s:lazy_toml, {'lazy': 1})
  call dein#end()
  call dein#save_state()

  let g:dein#install_github_api_token = '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'
endif

" still installed
if dein#check_install()
  call dein#install()
endif

set runtimepath+=~/dotfiles/nvim

"fileタイプの検索を有効にする
" filetype plugin on
runtime! setup/*.vim

