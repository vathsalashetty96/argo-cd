FROM golang@sha256:41cc61ae20fef3cbf18c315a82886100de9539da3ea05cb625d74100fd609e0b
RUN go version
RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.40.1

RUN golangci-lint --version
RUN export GOPATH=$HOME/go
RUN export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

ADD hack/install.sh .
ADD hack/installers installers
ADD hack/tool-versions.sh .
ADD hack hack

RUN go mod vendor
#RUN ./hack/update-codegen.sh
RUN ./hack/update-manifests.sh

