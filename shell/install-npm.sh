#!/bin/sh

if type "npm" > /dev/null 2>&1 ; then
    npm install -g typescript typescript-language-server
    # PHPのLSP
    npm -g install intelephense
    # VueのLSP
    npm i -g vls
    npm i -g eslint
    npm i -g eslint-loader
    npm i -g eslint-plugin-vue
    npm install -g dockerfile-language-server-nodejs
    npm i -g bash-language-server
    npm install -g vim-language-server
    npm i -g sql-language-server
    npm i -g diagnostic-languageserver
    npm i -g graphql
    npm i -g graphql-language-service-cli
    npm install -g @anthropic-ai/claude-code
    npm install -g @google/gemini-cli
    npm i -g reviewit
    npm i -g macos-notify-mcp

fi
