ARG GO_VERSION
ARG LINTER_VERSION

FROM golang:${GO_VERSION}-alpine

RUN apk --update-cache add --no-cache git make gcc musl-dev curl bash-completion

# wireを利用する場合はアンコメントしてください
# RUN go install github.com/google/wire/cmd/wire@vlatest
# airを利用する場合はアンコメントしてください
# RUN go install github.com/cosmtrek/air@latest

# golangci-lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin ${LINT_VERSION}

# タイムゾーンをJSTに設定
RUN apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
    && apk del tzdata

# bashrcに設定追加
RUN echo 'source /etc/profile.d/bash_completion.sh' >> ~/.bashrc

# KGMを使うための設定
ENV GOPRIVATE gitlab.kaonavi.jp/ae

# その他の環境変数
ENV HISTCONTROL ignoredups

EXPOSE 8080
