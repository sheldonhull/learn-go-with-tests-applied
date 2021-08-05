# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.187.0/containers/codespaces-linux/.devcontainer/base.Dockerfile

FROM mcr.microsoft.com/vscode/devcontainers/universal:1-focal

# ** [Optional] Uncomment this section to install additional packages. **
# USER root
#
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>
#
# USER codespace
RUN curl -sf https://gobinaries.com/chriswalz/bit | sh; \
    bit complete;

RUN go install github.com/owenrumney/squealer/cmd/squealer@latest \
    go install github.com/evilmartians/lefthook@latest \
    GO111MODULE='off' go get github.com/daixiang0/gci \
    GO111MODULE='off'  go get golang.org/x/tools/cmd/goimports \
    GO111MODULE=on go get mvdan.cc/gofumpt \
    go install github.com/gechr/yamlfmt@latest \
    GO111MODULE=on go get github.com/zricethezav/gitleaks/v7
