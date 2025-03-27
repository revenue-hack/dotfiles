-- ~/.config/nvim/lua/config/lsp.lua

local lspconfig = require("lspconfig")

-- PHP (intelephense)
lspconfig.intelephense.setup({})

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
