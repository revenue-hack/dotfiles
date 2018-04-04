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

"nerdtreeのalias
nnoremap <silent><C-e> :NERDTreeToggle<CR>
"ファイル移動のalias
noremap <C-l> <C-w>l
noremap <C-j> <C-w>j
noremap <C-h> <C-w>h
noremap <C-k> <C-w>k

