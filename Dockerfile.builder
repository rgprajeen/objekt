FROM rockylinux:8

ARG GO_VERSION

LABEL GO_VERSION="${GO_VERSION}"

ENV PATH="${PATH}:/usr/local/go/bin" \
  GO_VERSION="${GO_VERSION}"

RUN dnf install -y git gzip wget && \
  dnf clean all && \
  rm -rf /var/lib/yum && \
  wget --progress=dot:mega -O go.tar.gz \
    "https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz" && \
  rm -rf /usr/local/go && tar -C /usr/local -xzf go.tar.gz && \
  rm go.tar.gz && \
  adduser -u 1000 -d /buildman buildman

USER buildman

WORKDIR /buildman

RUN git config --global user.name 'Objekt Buildman' && \
  git config --global user.email '<>'

CMD ["go", "version"]
