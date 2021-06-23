FROM golang@sha256:41cc61ae20fef3cbf18c315a82886100de9539da3ea05cb625d74100fd609e0b
RUN go version
RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.40.1

RUN golangci-lint --version
RUN export GOPATH=$HOME/go
RUN export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin


WORKDIR /go/src/github.com/vathsalashetty25/argo-cd

COPY go.mod go.mod
COPY go.sum go.sum


ADD hack/install.sh .
ADD hack/installers installers
ADD hack/tool-versions.sh .
ADD hack hack
#RUN go clean -modcache
#RUN ./hack/update-manifests.sh
#RUN go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.3.0
#ENV PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
#RUN go mod tidy
#RUN go mod vendor
#RUN ./hack/update-openapi.sh
#RUN ./hack/generate-proto.sh
#RUN ./hack/update-codegen.sh
RUN go mod vendor
#RUN ./hack/update-codegen.sh
RUN ./hack/update-manifests.sh
#RUN ./hack/generate-proto.sh
#RUN ./hack/update-openapi.sh
