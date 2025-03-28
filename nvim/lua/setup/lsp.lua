-- ~/.config/nvim/lua/config/lsp.lua
--
-- capabilities（cmpとの連携用。無くても最低限動く）
local capabilities = require("cmp_nvim_lsp").default_capabilities()

-- LSPがアタッチされたときのキーマップ
local on_attach = function(_, bufnr)
  local opts = { buffer = bufnr, silent = true }

  vim.keymap.set("n", "gd", vim.lsp.buf.definition, opts)
  vim.keymap.set("n", "gD", vim.lsp.buf.declaration, opts)
  vim.keymap.set("n", "gi", vim.lsp.buf.implementation, opts)
  vim.keymap.set("n", "gr", vim.lsp.buf.references, opts)
  vim.keymap.set("n", "K",  vim.lsp.buf.hover, opts)
end

local lspconfig = require("lspconfig")

-- PHP (intelephense)
lspconfig.intelephense.setup({
  on_attach = on_attach,
  capabilities = capabilities,
})

-- Go (gopls)
lspconfig.gopls.setup({})

-- TypeScript / JavaScript (tsserver)
lspconfig.tsserver.setup({})

-- Ruby (solargraph)
lspconfig.solargraph.setup({})
-- Shell (bash)
lspconfig.bashls.setup({})

-- Lua (Neovim Lua 用設定推奨)
lspconfig.lua_ls.setup({
  settings = {
    Lua = {
      runtime = {
        version = "LuaJIT",
      },
      diagnostics = {
        globals = { "vim" }, -- `vim` グローバルを未定義と警告しないように
      },
      workspace = {
        library = vim.api.nvim_get_runtime_file("", true),
        checkThirdParty = false,
      },
      telemetry = {
        enable = false,
      },
    },
  },
})

-- Vim script
lspconfig.vimls.setup({})

-- Docker
lspconfig.dockerls.setup({})

-- Terraform
lspconfig.terraformls.setup({})

-- SQL
lspconfig.sqlls.setup({})
