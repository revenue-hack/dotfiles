-- Python / Ruby インタープリタ設定
vim.g.python_host_prog = os.getenv("PYENV_ROOT") .. "/shims/python2"
vim.g.python3_host_prog = os.getenv("PYENV_ROOT") .. "/shims/python3"
vim.g.ruby_host_prog = os.getenv("RBENV_ROOT") .. "/versions/3.0.0/bin/ruby"

-- カラースキーム
vim.cmd("colorscheme molokai")
vim.opt.background = "dark"
vim.g.molokai_original = 1
vim.g.rehash256 = 1

-- 表示・見た目
vim.opt.number = true
vim.opt.ambiwidth = "double"
vim.opt.list = true
vim.opt.listchars = {
  tab = "»-",
  trail = "-",
  eol = "↲",
  extends = "»",
  precedes = "«",
  nbsp = "%",
}
vim.cmd([[highlight NonText guibg=NONE guifg=DarkGreen]])
-- vim.cmd([[highlight SpecialKey guibg=NONE guifg=Gray40]])

-- 編集系
vim.opt.expandtab = true
vim.opt.autoindent = true
vim.opt.smartindent = true
vim.opt.encoding = "utf-8"
vim.opt.fileencodings = { "utf-8" }
vim.opt.backspace = { "indent", "eol", "start" }
vim.opt.swapfile = false
vim.opt.modifiable = true
vim.opt.write = true

-- ステータス・情報表示
vim.opt.ruler = true
vim.opt.cmdheight = 2
vim.opt.laststatus = 2
vim.opt.title = true
vim.opt.showcmd = true

-- 検索
vim.opt.incsearch = true
vim.opt.ignorecase = true
vim.opt.smartcase = true
vim.opt.hlsearch = true

-- 入力補助
vim.opt.wildmenu = true
vim.opt.history = 5000

-- クリップボード
vim.opt.clipboard = "unnamedplus"

-- NERDTree 設定
vim.g.NERDTreeShowHidden = 1

