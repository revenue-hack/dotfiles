return {
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
              "ts_ls",
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
        "yetone/avante.nvim",
        event = "VeryLazy",
        version = false, -- Never set this value to "*"! Never!
        opts = {
          -- add any opts here
          -- for example
          provider = "openai",
          openai = {
            endpoint = "https://api.openai.com/v1",
            model = "gpt-4o", -- your desired model (or use gpt-4o, etc.)
            timeout = 30000, -- Timeout in milliseconds, increase this for reasoning models
            temperature = 0,
            max_completion_tokens = 8192, -- Increase this to include reasoning tokens (for reasoning models)
            --reasoning_effort = "medium", -- low|medium|high, only used for reasoning models
          },
          behaviour = {
            auto_apply_diff_after_generation = true
          },
        },
        -- if you want to build from source then do `make BUILD_FROM_SOURCE=true`
        build = "make",
        -- build = "powershell -ExecutionPolicy Bypass -File Build.ps1 -BuildFromSource false" -- for windows
        dependencies = {
          "nvim-treesitter/nvim-treesitter",
          "stevearc/dressing.nvim",
          "nvim-lua/plenary.nvim",
          "MunifTanjim/nui.nvim",
          --- The below dependencies are optional,
          "echasnovski/mini.pick", -- for file_selector provider mini.pick
          "nvim-telescope/telescope.nvim", -- for file_selector provider telescope
          "hrsh7th/nvim-cmp", -- autocompletion for avante commands and mentions
          "ibhagwan/fzf-lua", -- for file_selector provider fzf
          "nvim-tree/nvim-web-devicons", -- or echasnovski/mini.icons
          "zbirenbaum/copilot.lua", -- for providers='copilot'
          {
            -- support for image pasting
            "HakonHarnes/img-clip.nvim",
            event = "VeryLazy",
            opts = {
              -- recommended settings
              default = {
                embed_image_as_base64 = false,
                prompt_for_file_name = false,
                drag_and_drop = {
                  insert_mode = true,
                },
                -- required for Windows users
                use_absolute_path = true,
              },
            },
          },
          {
            -- Make sure to set this up properly if you have lazy=true
            'MeanderingProgrammer/render-markdown.nvim',
            opts = {
              file_types = { "markdown", "Avante" },
            },
            ft = { "markdown", "Avante" },
          },
        },
      },
    }
