[![Go Reference](https://pkg.go.dev/badge/github.com/ntnguyencse/L-KaaS.svg)](https://pkg.go.dev/github.com/ntnguyencse/L-KaaS)
# Logical Kubernetes as a Service

* L-KaaS is a project focused on providing declarative APIs and tooling to simplify, abstract, be easy to use for users who don’t have deep technical knowledge of infrastructure and shield them from low-level concepts and technologies.


## Goals of L-KaaS

* Powerful abstraction implemented on top of existing project Cluster API

* Provide a simple, automatic, and easy-to-manage lifecycle of multi-Kubernetes clusters using declarative methods.

* Based on GitOps supports, taking advantage of git for management from Day 0 through Day 2.

* Reuse and integrate existing ecosystems (Cluster API,..) rather than duplicating their functionality.

* Simplifying and uniform automation all the way to onboarding, the complexity of provisioning and managing a multi-provider, multi-site deployment of underlying cloud infrastructure or distributed cloud, getting rid of all complex configurations

## Documentation
This is a very early development; right now it is just a starter page with links to other resources. In time we develop comprehensive documentation.
For concepts, glossary, designs, see [Docs](/docs)

## Getting Started
You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

### Running on the cluster
1. Install Instances of Custom Resources:

```sh
kubectl apply -f config/crd/bases/
kubectl apply -f config/samples/profiles/
kubectl apply -f config/samples/clustercatalogs/
```
2. Install Cluster API Controllers
```sh
./hack/capi-install.sh
```
3. Run Controllers

```sh
make run
```
4. Create Logical Cluster
```sh
kubectl apply -f config/samples/intent_v1_logicalcluster.yaml
```
### Uninstall CRDs
To delete the CRDs from the cluster:

```sh
make uninstall
```

### How it works
This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/).

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/),
which provide a reconcile function responsible for synchronizing resources until the desired state is reached on the cluster.

### Test It Out
1. Install the CRDs into the cluster:

```sh
make install
```

2. Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

```sh
make run
```

**NOTE:** You can also run this in one step by running: `make install run`

### Modifying the API definitions
If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2023 Nguyen Thanh Nguyen.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

