
# Image URL to use all building/pushing image targets
IMG ?= metadata-controller:latest

default: manager

all: manifests test manager

# Run tests
test: generate fmt vet manifests
	go test ./src/... ./cmd/... -coverprofile cover.out

# Build manager binary
manager: generate fmt
	go build -o bin/manager github.com/alphagov/verify-metadata-controller/cmd/manager

# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate fmt vet
	go run ./cmd/manager/main.go

# Install CRDs into a cluster
install: manifests
	kubectl apply -f config/crds

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
deploy: manifests
	kubectl apply -f config/crds
	kustomize build config/default | kubectl apply -f -

# Generate manifests e.g. CRD, RBAC etc.
manifests:
	go run vendor/sigs.k8s.io/controller-tools/cmd/controller-gen/main.go all
	cp config/crds/*.yaml chart/templates/
	echo "the helm chart does not get automatically updated you may need to tweak it if new values generated"

# Run go fmt against code
fmt:
	go fmt ./src/... ./cmd/...

# Run go vet against code
vet:
	go vet ./src/... ./cmd/...

# Generate code
generate:
	go generate ./src/... ./cmd/...
	counterfeiter src/hsm/ Client

# Build the docker image
docker-build:
	docker build . -t ${IMG}
	@echo "updating kustomize image patch file for manager resource"
	sed -i'' -e 's@image: .*@image: '"${IMG}"'@' ./config/default/manager_image_patch.yaml

# Push the docker image
docker-push:
	docker push ${IMG}


generator:
	mkdir -p build
	test -d build/verify-metadata || git clone git@github.com:alphagov/verify-metadata.git build/verify-metadata
	test -d build/verify-metadata-generator || git clone git@github.com:alphagov/verify-metadata-generator.git build/verify-metadata-generator
	docker build -f Dockerfile.generator .
