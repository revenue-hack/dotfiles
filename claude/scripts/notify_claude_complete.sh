#!/bin/bash

# Claude Code完了通知スクリプト
# 使い方: ./notify_claude_complete.sh "完了メッセージ"

MESSAGE=${1:-"Claude Codeの処理が完了しました"}

osascript -e "display alert \"$MESSAGE\""

# 成功時のメッセージ
echo "アラートを表示しました: $MESSAGE"