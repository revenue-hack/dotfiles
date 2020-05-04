#!/bin/sh

if type "npm" > /dev/null 2>&1 ; then
    # TypescriptのLSP
    npm install -g typescript typescript-language-server
    # PHPのLSP
    npm -g install intelephense
    # VueのLSP
    npm i -g vls
    npm i -g eslint
    npm i -g eslint-loader
    npm i -g eslint-plugin-vue
    # Dockerfile LSP
    npm install -g dockerfile-language-server-nodejs
    # Bash LSP
    npm i -g bash-language-server
    # Vim LSP
    npm install -g vim-language-server
    # SQL LSP
    npm i -g sql-language-server
fi
