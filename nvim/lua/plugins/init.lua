return {
  {
    "nvim-lua/plenary.nvim",
  },
  {
    "greggh/claude-code.nvim",
    dependencies = {
      "nvim-lua/plenary.nvim", -- Required for git operations
    },
    config = function()
      require("claude-code").setup({
        -- Terminal window settings
        window = {
          --split_ratio = 0.3,      -- Percentage of screen for the terminal window (height for horizontal, width for vertical splits)
          vertical_ratio = 0.3,      -- Percentage of screen for the terminal window (height for horizontal, width for vertical splits)
          position = "rightbelow vsplit",  -- Position of the window: "botright", "topleft", "vertical", "rightbelow vsplit", etc.
          --enter_insert = true,    -- Whether to enter insert mode when opening Claude Code
          --hide_numbers = true,    -- Hide line numbers in the terminal window
          --hide_signcolumn = true, -- Hide the sign column in the terminal window
        },
        -- File refresh settings
        refresh = {
          enable = true,           -- Enable file change detection
          updatetime = 100,        -- updatetime when Claude Code is active (milliseconds)
          timer_interval = 1000,   -- How often to check for file changes (milliseconds)
          show_notifications = true, -- Show notification when files are reloaded
        },
        -- Git project settings
        git = {
          use_git_root = true,     -- Set CWD to git root when opening Claude Code (if in git project)
        },
        -- Shell-specific settings
        shell = {
          separator = '&&',        -- Command separator used in shell commands
          pushd_cmd = 'pushd',     -- Command to push directory onto stack (e.g., 'pushd' for bash/zsh, 'enter' for nushell)
          popd_cmd = 'popd',       -- Command to pop directory from stack (e.g., 'popd' for bash/zsh, 'exit' for nushell)
        },
        -- Command settings
        command = "claude",        -- Command used to launch Claude Code
        -- Command variants
        command_variants = {
          -- Conversation management
          continue = "--continue", -- Resume the most recent conversation
          --resume = "--resume",     -- Display an interactive conversation picker

          -- Output options
          --verbose = "--verbose",   -- Enable verbose logging with full turn-by-turn output
        },
        -- Keymaps
        keymaps = {
          toggle = {
            normal = "<C-,>",       -- Normal mode keymap for toggling Claude Code, false to disable
            terminal = "<C-,>",     -- Terminal mode keymap for toggling Claude Code, false to disable
            variants = {
              continue = "<leader>cC", -- Normal mode keymap for Claude Code with continue flag
              verbose = "<leader>cV",  -- Normal mode keymap for Claude Code with verbose flag
            },
          },
          window_navigation = true, -- Enable window navigation keymaps (<C-h/j/k/l>)
          scrolling = true,         -- Enable scrolling keymaps (<C-f/b>) for page up/down
        }
      })
      vim.keymap.set('n', '<leader>cc', '<cmd>ClaudeCode<CR>', { desc = 'Toggle Claude Code' })
    end
  },
  {
    "catppuccin/nvim",
    name = "catppuccin",
    priority = 1000,
    lazy = false,
    config = function()
      require("catppuccin").setup({
        flavour = "mocha", -- latte, frappe, macchiato, mocha
        transparent_background = false,
        term_colors = true,
      })
      vim.cmd.colorscheme("catppuccin")
    end,
  },
  {
    "stevearc/conform.nvim",
    config = function()
      require("conform").setup({
        formatters_by_ft = {
          python = { "ruff_format" },
        },
      })

      vim.api.nvim_create_autocmd("BufWritePre", {
        pattern = "*.py",
        callback = function(args)
          require("conform").format({
            bufnr = args.buf,
            async = false,
            lsp_fallback = true,
          })
        end,
      })
    end,
  },
  {
    "nvim-neo-tree/neo-tree.nvim",
    branch = "v3.x",  -- 安定版
    dependencies = {
      "nvim-lua/plenary.nvim",
      "nvim-tree/nvim-web-devicons", -- アイコン表示
      "MunifTanjim/nui.nvim",
    },
    config = function()
      require("neo-tree").setup({
        window = {
          width = 30,
          mappings = {
            ["<space>"] = "none",
          },
        },
        filesystem = {
          filtered_items = {
            visible = true,
            hide_dotfiles = false,
            hide_gitignored = false,
          },
        },
      })

      -- おすすめキーマップ（必要に応じて）
      vim.keymap.set("n", "<C-e>", ":Neotree toggle<CR>", { desc = "Toggle NeoTree", silent = true })
    end,
  },

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
      -- capabilities（cmpとの連携用。無くても最低限動く）
      local capabilities = require("cmp_nvim_lsp").default_capabilities()

      -- LSPがアタッチされたときのキーマップ
      local on_attach = function(client, bufnr)
        local opts = { buffer = bufnr, silent = true }

        vim.keymap.set("n", "<Tab>]", vim.lsp.buf.definition, opts)
        vim.keymap.set("n", "<Tab>[", vim.lsp.buf.declaration, opts)
        vim.keymap.set("n", "<Tab>i", vim.lsp.buf.implementation, opts)
        vim.keymap.set("n", "<C-r>", vim.lsp.buf.references, opts)
        vim.keymap.set("n", "K",  vim.lsp.buf.hover, opts)
        vim.keymap.set("n", "<leader>d", vim.diagnostic.open_float)
        vim.keymap.set("n", "cn", vim.lsp.buf.rename, { desc = "rename" })


        if client.server_capabilities.documentFormattingProvider then
          vim.api.nvim_create_autocmd("BufWritePre", {
            buffer = bufnr,
            callback = function()
              vim.lsp.buf.format({ bufnr = bufnr })
            end,
          })
        end
      end

      -- LSPサーバーの設定をここで行う
      local lspconfig = require("lspconfig")

      -- masonの設定からsolargraphを削除し、手動で設定
      lspconfig.solargraph.setup{
        cmd = { "solargraph", "stdio" },
        filetypes = { "ruby" },
        root_dir = require'lspconfig'.util.root_pattern('.git', 'Gemfile', '.ruby-version'),
      }

      lspconfig.intelephense.setup({
        on_attach = on_attach,
        capabilities = capabilities,
        settings = {
          intelephense = {
            environment = {
              includePaths = { "./vendor" },
            },
          },
        },
      })

      require("lspconfig").pyright.setup({
        on_attach = on_attach,
        capabilities = capabilities,
        settings = {
          python = {
            analysis = {
              autoSearchPaths = true,
              useLibraryCodeForTypes = true,
              diagnosticMode = "openFilesOnly", -- or "workspace"
            },
          },
        },
      })

      -- Go (gopls)
      lspconfig.gopls.setup({
        on_attach = on_attach,
        capabilities = capabilities,
        settings = {
          gopls = {
            gofumpt = false,

            analyses = {
              unusedparams = true,
            },
            staticcheck = true,
            codelenses = {
              generate = true,
              gc_details = true,
            },
          },
        },
      })


      -- TypeScript / JavaScript (ts_ls)
      lspconfig.ts_ls.setup({
        on_attach = on_attach,
        capabilities = capabilities,
      })

      -- Ruby (solargraph)
      lspconfig.solargraph.setup({
        on_attach = on_attach,
        capabilities = capabilities,
      })
      -- Shell (bash)
      lspconfig.bashls.setup({
        on_attach = on_attach,
        capabilities = capabilities,
      })

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
      lspconfig.vimls.setup({
        on_attach = on_attach,
        capabilities = capabilities,
      })

      -- Docker
      lspconfig.dockerls.setup({
        on_attach = on_attach,
        capabilities = capabilities,
      })

      -- Terraform
      lspconfig.terraformls.setup({
        on_attach = on_attach,
        capabilities = capabilities,
      })

      -- SQL
      lspconfig.sqlls.setup({
        on_attach = on_attach,
        capabilities = capabilities,
      })
    end,
  },
  { "plasticboy/vim-markdown" },

  -- previm + markdown設定
  {
    "previm/previm",
    ft = { "markdown" },
    config = function()
      vim.cmd [[
      nnoremap <silent> <TAB>o :PrevimOpen<CR>
      let g:vim_markdown_folding_disabled = 1
      let g:previm_open_cmd = 'open -a "Google Chrome"'
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
        "williamboman/mason-lspconfig.nvim",
        dependencies = { "mason.nvim", "neovim/nvim-lspconfig" },
        config = function()
          require("mason-lspconfig").setup({
            ensure_installed = {
              "lua_ls",
              "ruff",
              "pyright",
              "ts_ls",
              "gopls",
              "intelephense",
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
        "nvim-treesitter/nvim-treesitter",
        build = ":TSUpdate",
        event = { "BufReadPost", "BufNewFile" },
        config = function()
          require("nvim-treesitter.configs").setup({
            ensure_installed = {
              "go",
              "php",
              "python",
              "typescript",
              "javascript",
              "ruby",
              "lua",
              "bash",
              "vim",
              "json",
              "yaml",
              "markdown",
              "html",
              "css",
            },
            highlight = {
              enable = true,
              additional_vim_regex_highlighting = false,
            },
            indent = {
              enable = true,
            },
            incremental_selection = {
              enable = true,
              keymaps = {
                init_selection = "gnn",        -- 開始
                node_incremental = "grn",      -- ノードを広げる
                scope_incremental = "grc",     -- スコープを広げる
                node_decremental = "grm",      -- ノードを戻す
              },
            },
            textobjects = {
              select = {
                enable = true,
                lookahead = true,
                keymaps = {
                  ["af"] = "@function.outer",  -- 関数全体
                  ["if"] = "@function.inner",  -- 関数内側
                  ["ac"] = "@class.outer",
                  ["ic"] = "@class.inner",
                },
              },
            },
          })
        end,
      },
      --{
      --  "yetone/avante.nvim",
      --  event = "VeryLazy",
      --  version = false, -- Never set this value to "*"! Never!
      --  opts = {
      --    -- add any opts here
      --    -- for example
      --    provider = "openai",
      --    openai = {
      --      endpoint = "https://api.openai.com/v1",
      --      model = "gpt-4.1", -- your desired model (or use gpt-4o, etc.)
      --      timeout = 30000, -- Timeout in milliseconds, increase this for reasoning models
      --      temperature = 0,
      --      max_tokens = 8192, -- Increase this to include reasoning tokens (for reasoning models)
      --      --reasoning_effort = "medium", -- low|medium|high, only used for reasoning models
      --    },
      --    behaviour = {
      --      auto_apply_diff_after_generation = true
      --    },
      --  },
      --  -- if you want to build from source then do `make BUILD_FROM_SOURCE=true`
      --  build = "make",
      --  -- build = "powershell -ExecutionPolicy Bypass -File Build.ps1 -BuildFromSource false" -- for windows
      --  dependencies = {
      --    "nvim-treesitter/nvim-treesitter",
      --    {
      --      "stevearc/dressing.nvim",
      --      event = "VeryLazy",
      --      opts = {
      --        input = {
      --          -- safe fallback width
      --          win_options = {
      --            list = false, -- dressing用のpopupウィンドウでもlist無効
      --          },
      --        },
      --      },
      --    },
      --    "nvim-lua/plenary.nvim",
      --    "MunifTanjim/nui.nvim",
      --    --- The below dependencies are optional,
      --    "echasnovski/mini.pick", -- for file_selector provider mini.pick
      --    "nvim-telescope/telescope.nvim", -- for file_selector provider telescope
      --    "hrsh7th/nvim-cmp", -- autocompletion for avante commands and mentions
      --    "ibhagwan/fzf-lua", -- for file_selector provider fzf
      --    "nvim-tree/nvim-web-devicons", -- or echasnovski/mini.icons
      --    "zbirenbaum/copilot.lua", -- for providers='copilot'
      --    {
      --      -- support for image pasting
      --      "HakonHarnes/img-clip.nvim",
      --      event = "VeryLazy",
      --      opts = {
      --        -- recommended settings
      --        default = {
      --          embed_image_as_base64 = false,
      --          prompt_for_file_name = false,
      --          drag_and_drop = {
      --            insert_mode = true,
      --          },
      --          -- required for Windows users
      --          use_absolute_path = true,
      --        },
      --      },
      --    },
      --    {
      --      "iamcco/markdown-preview.nvim",
      --      ft = { "markdown" },
      --      build = "cd app && yarn install",
      --      config = function()
      --        vim.g.mkdp_auto_start = 1
      --        vim.g.mkdp_open_to_the_world = 0
      --      end,
      --    },
      --},
}
