vim.api.nvim_create_augroup("MyFiletypes", { clear = true })

-- ファイルタイプで基本インデントを設定
local indent_2 = {
  "html", "javascript", "sql", "twig", "xml", "yaml", "zsh", "vim", "css", "scss", "java", "sass", "jade",
  "vcl", "swift", "ruby", "json", "toml", "proto", "tf", "json5", "graphql", "vue", "typescript", "markdown", "sh"
}
for _, ft in ipairs(indent_2) do
  vim.api.nvim_create_autocmd("FileType", {
    group = "MyFiletypes",
    pattern = ft,
    command = "setlocal sw=2 sts=2 ts=2 et",
  })
end

vim.api.nvim_create_autocmd("FileType", {
  group = "MyFiletypes",
  pattern = { "php", "python" },
  command = "setlocal sw=4 sts=4 ts=4 et",
})

-- 特殊なファイルタイプ定義
vim.api.nvim_create_autocmd({ "BufNewFile", "BufRead" }, {
  group = "MyFiletypes",
  pattern = "*.svelte",
  command = "set filetype=svelte",
})

vim.api.nvim_create_autocmd({ "BufNewFile", "BufRead" }, {
  group = "MyFiletypes",
  pattern = { "*.tsx", "*.jsx" },
  command = "set filetype=typescriptreact | setlocal sw=2 sts=2 ts=2 et",
})

vim.api.nvim_create_autocmd({ "BufNewFile", "BufRead" }, {
  group = "MyFiletypes",
  pattern = "*.graphql",
  command = "set filetype=graphql",
})

vim.api.nvim_create_autocmd({ "BufNewFile", "BufRead" }, {
  group = "MyFiletypes",
  pattern = "*.pu",
  command = "setlocal sw=2 sts=2 ts=2 et",
})

-- markdownファイルのファイルタイプ設定（念のため）
vim.api.nvim_create_autocmd({ "BufRead", "BufNewFile" }, {
  group = "MyFiletypes",
  pattern = "*.md",
  command = "set filetype=markdown",
})

-- luaファイルのときにスペース2に設定
vim.api.nvim_create_autocmd("FileType", {
  pattern = "lua",
  group = vim.api.nvim_create_augroup("LuaIndent", { clear = true }),
  command = "setlocal shiftwidth=2 softtabstop=2 tabstop=2 expandtab",
})


-- goimports on save
vim.api.nvim_create_autocmd("BufWritePre", {
  pattern = "*.go",
  callback = function()
    local params = vim.lsp.util.make_range_params()
    params.context = {only = {"source.organizeImports"}}
    -- buf_request_sync defaults to a 1000ms timeout. Depending on your
    -- machine and codebase, you may want longer. Add an additional
    -- argument after params if you find that you have to write the file
    -- twice for changes to be saved.
    -- E.g., vim.lsp.buf_request_sync(0, "textDocument/codeAction", params, 3000)
    local result = vim.lsp.buf_request_sync(0, "textDocument/codeAction", params)
    for cid, res in pairs(result or {}) do
      for _, r in pairs(res.result or {}) do
        if r.edit then
          local enc = (vim.lsp.get_client_by_id(cid) or {}).offset_encoding or "utf-16"
          vim.lsp.util.apply_workspace_edit(r.edit, enc)
        end
      end
    end
    vim.lsp.buf.format({async = false})
  end
})
