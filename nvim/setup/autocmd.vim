if has("autocmd")
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
  autocmd FileType vcl setlocal sw=2 sts=2 ts=2 et
  autocmd FileType swift setlocal sw=2 sts=2 ts=2 et
  autocmd FileType ruby setlocal sw=2 sts=2 ts=2 et
  autocmd FileType json setlocal sw=2 sts=2 ts=2 et
  autocmd FileType toml setlocal sw=2 sts=2 ts=2 et
  autocmd FileType vue setlocal sw=2 sts=2 ts=2 et
  autocmd FileType python setlocal sw=4 sts=4 ts=4 et
  syntax enable
endif

