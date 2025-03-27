return {
  { "scrooloose/nerdtree" },

  { "tomasr/molokai" },

  { "GutenYe/json5.vim" },

  {
    "sebdah/vim-delve",
    ft = "go",
    config = function()
      vim.cmd [[
      autocmd FileType go nmap dbp :DlvAddBreakpoint<CR>
      autocmd FileType go nmap dtp :DlvAddTracepoint<CR>
      autocmd FileType go nmap dt :DlvTest<CR>
      autocmd FileType go nmap dc :DlvClearAll<CR>
      ]]
    end,
  },

  {
    "junegunn/fzf",
    build = "./install --all",
  },

  {
    "junegunn/fzf.vim",
    config = function()
      vim.cmd [[
      let g:fzf_action = { 'ctrl-s': 'split' }
      nnoremap <silent><C-p> :Files<CR>
      nnoremap <C-q> :Buffers<CR>
      nnoremap <C-g> :Rg<CR>
      let $FZF_DEFAULT_COMMAND="rg --files --hidden --glob '!.git/**' --glob '!*/node_modules/*'"
      let g:fzf_buffers_jump = 1
      command! -bang -nargs=? -complete=dir Files
      \ call fzf#vim#files(<q-args>, fzf#vim#with_preview(), <bang>0)
      ]]
    end,
  },
  {
    "hrsh7th/nvim-cmp",
    dependencies = {
      "hrsh7th/cmp-nvim-lsp",
      "hrsh7th/cmp-buffer",
      "hrsh7th/cmp-path",
      "L3MON4D3/LuaSnip",
    },
    config = function()
      local cmp = require("cmp")
      cmp.setup({
        snippet = {
          expand = function(args)
            require("luasnip").lsp_expand(args.body)
          end,
        },
        mapping = cmp.mapping.preset.insert({
          ["<C-Space>"] = cmp.mapping.complete(),
          ["<CR>"] = cmp.mapping.confirm({ select = true }),
        }),
        sources = cmp.config.sources({
          { name = "nvim_lsp" },
          { name = "buffer" },
          { name = "path" },
        }),
      })
    end,
  },
  {
    "neovim/nvim-lspconfig",
    config = function()
      -- LSPサーバーの設定をここで行う
      local lspconfig = require("lspconfig")

      lspconfig.intelephense.setup({
        settings = {
          intelephense = {
            environment = {
              includePaths = { "./vendor" },
            },
          },
        },
      })
    end,
  },
  { "plasticboy/vim-markdown" },

  -- previm + markdown設定
  {
    "previm/previm",
    config = function()
      vim.cmd [[
      nnoremap <silent> <TAB>o :PrevimOpen<CR>
      let g:vim_markdown_folding_disabled = 1
      let g:previm_open_cmd = 'open -a Google\\ Chrome'
      let g:previm_enable_realtime = 1
      ]]
    end,
  },

  -- open-browser.vim
  { "tyru/open-browser.vim" },

  -- Terraform
  {
    "hashivim/vim-terraform",
    ft = "terraform",
    config = function()
      vim.g.terraform_fmt_on_save = 1
    end,
  },

  -- TOML
  {
    "cespare/vim-toml",
    ft = "toml",
  },

  -- TypeScript
  {
    "leafgarland/typescript-vim",
    ft = { "typescriptreact", "typescript" },
  },

  -- JSX/TSX
  {
    "peitalin/vim-jsx-typescript",
    ft = { "tsx", "jsx", "typescriptreact" },
  },

  -- ALE（Linter + Fixer）
  {
    "w0rp/ale",
    ft = { "vue", "typescript", "javascript" },
    config = function()
      vim.cmd [[
      let g:ale_statusline_format = ['E%d', 'W%d', 'OK']
      nmap <silent> <C-w>j <Plug>(ale_next_wrap)
      nmap <silent> <C-w>k <Plug>(ale_previous_wrap)
      let g:ale_sign_error = '>>'
      let g:ale_sign_warning = '!!'
      let g:ale_fixers = {
        \ 'javascript': ['eslint', 'prettier'],
        \ 'vue': ['eslint', 'prettier'],
        \ 'typescript': ['eslint', 'prettier']
        \}
        let g:ale_linters = {
          \ 'javascript': ['eslint', 'prettier'],
          \ 'vue': ['eslint', 'prettier'],
          \ 'typescript': ['eslint', 'prettier']
          \}
          let g:ale_fix_on_save = 1
          let g:ale_javascript_prettier_use_local_config = 1
          ]]
        end,
      },
      {
        "williamboman/mason.nvim",
        build = ":MasonUpdate", -- optional
        config = function()
          require("mason").setup()
        end,
      },

      {
        "williamboman/mason-lspconfig.nvim",
        dependencies = { "mason.nvim", "neovim/nvim-lspconfig" },
        config = function()
          require("mason-lspconfig").setup({
            ensure_installed = {
              "lua_ls",
              "tsserver",
              "gopls",
              "intelephense",
              "solargraph",
              "bashls",
              "vimls",
              "dockerls",
              "terraformls",
              "sqlls",
            },
            automatic_installation = true,
          })
        end,
      },

      {
        "williamboman/mason.nvim",
        build = ":MasonUpdate", -- optional
        config = function()
          require("mason").setup()
        end,
      },

      {
        "williamboman/mason-lspconfig.nvim",
        dependencies = { "mason.nvim", "neovim/nvim-lspconfig" },
        config = function()
          require("mason-lspconfig").setup({
            ensure_installed = {
              "lua_ls",
              "tsserver",
              "gopls",
              "intelephense",
              "solargraph",
              "bashls",
              "vimls",
              "dockerls",
              "terraformls",
              "sqlls",
            },
            automatic_installation = true,
          })
        end,
      }


    }
