FROM mcr.microsoft.com/vscode/devcontainers/universal:linux

# ** [Optional] Uncomment this section to install additional packages. **
USER root

ENV PACKAGE_LIST build-essential \
    pv \
    curl \
    wget \
    unzip
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -yqq install --no-install-recommends ${PACKAGE_LIST} \
    && apt-get clean -y && rm -rf /var/lib/apt/lists/* /tmp/library-scripts

# USER codespace
# ARG USERNAME=$USERNAME

# CACHED EXTENSIONS FOR FASTER STARTUP
USER root
RUN mkdir -p /home/codespace/.vscode-server/extensions \
    /home/codespace/.vscode-server/extensions && \
    chown -R codespace \
        /home/codespace/.vscode-server \
        /home/codespace/.vscode-server

COPY library-scripts/*.sh library-scripts/*.env /tmp/library-scripts/
RUN if [ "$INSTALL_AZURE_CLI" = "true" ]; then bash /tmp/library-scripts/azcli-debian.sh; fi \
    && apt-get clean -y && rm -rf /var/lib/apt/lists/* /tmp/library-scripts

COPY library-scripts/*.sh library-scripts/*.env /tmp/library-scripts/
RUN if [ "$INSTALL_AWS_CLI" = "true" ]; then bash /tmp/library-scripts/ubuntu-aws-cli.sh; fi \
    && apt-get clean -y && rm -rf /var/lib/apt/lists/* /tmp/library-scripts

COPY library-scripts/*.sh library-scripts/*.env /tmp/library-scripts/
RUN if [ "$INSTALL_GITVERSION" = "true" ]; then bash /tmp/library-scripts/dotnet-gitversion.sh; fi \
    && apt-get clean -y && rm -rf /var/lib/apt/lists/* /tmp/library-scripts
COPY library-scripts/*.sh library-scripts/*.env /tmp/library-scripts/
RUN echo "✅ install starship prompt" \
    && bash /tmp/library-scripts/install-starship.sh \
    && apt-get clean -y && rm -rf /var/lib/apt/lists/* /tmp/library-scripts

RUN echo "✅ install chezmoi for dotfiles syncing" && /bin/sh -c "$(wget -qO- git.io/chezmoi)" && mv ./bin/chezmoi /usr/local/bin/chezmoi && rm -rf ./bin

# RUN curl --silent -fsSL https://starship.rs/install.sh | sh -- --bin-dir /usr/local/bin --force && echo "completed setup of starship.rs"

USER codespace



RUN echo "✅ installing task runners" \
  && go version \
  && go install github.com/magefile/mage@latest && go install github.com/iwittkau/mage-select@latest \
  && go install github.com/go-task/task/v3/cmd/task@latest

RUN pwsh -nologo -c 'New-Item -Path ($Profile | Split-Path -Parent) -ItemType Directory' \
    && pwsh -nologo -c 'Set-Content $PROFILE ''iex $(&starship init powershell --print-full-init | Out-String)'' -Force'

RUN go version && go install github.com/owenrumney/squealer/cmd/squealer@latest \
    && go install github.com/evilmartians/lefthook@latest \
    && GO111MODULE='off' go get github.com/daixiang0/gci \
    && GO111MODULE='off' go get golang.org/x/tools/cmd/goimports \
    && GO111MODULE=on go get mvdan.cc/gofumpt \
    && go install github.com/gechr/yamlfmt@latest \
    && GO111MODULE=on go get github.com/zricethezav/gitleaks/v7

# Set multiple labels at once, using line-continuation characters to break long lines
LABEL vendor=misc \
      type=codespaces \
      is-beta=true \
      user=codespace \
      is-production=false \
      version="0.0.1-beta" \
      release-date="2021-06-01" \
      description="development image with tooling for vscode codespaces" \
      maintainer="sheldonhull" \
      name=codespaces-general
