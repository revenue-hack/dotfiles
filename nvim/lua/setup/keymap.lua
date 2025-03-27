local map = vim.keymap.set
local opts = { silent = true, noremap = true }

-- 括弧の自動挿入（インサートモード）
map("i", "{", "{}<Left>", { noremap = true })
map("i", "[", "[]<Left>", { noremap = true })
map("i", "(", "()<Left>", { noremap = true })

-- NERDTree toggle
map("n", "<C-e>", ":NERDTreeToggle<CR>", opts)

-- ウィンドウ移動（左右上下）
map("n", "<C-h>", "<C-w>h", opts)
map("n", "<C-j>", "<C-w>j", opts)
map("n", "<C-k>", "<C-w>k", opts)
map("n", "<C-l>", "<C-w>l", opts)

-- ファイルを再読み込み
map("n", "<S-y>", ":e!<CR>", opts)

-- ターミナル起動
map("n", "<C-t>", ":terminal<CR>", opts)

-- 括弧内選択（カッコ内全選択）
map("n", "<C-u>", "vi)", opts)

