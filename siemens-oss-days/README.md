# Open Source @ Siemens 2023

## Prerequisites

**Client**

```bash
brew install knative/client/kn
brew tap knative-sandbox/kn-plugins
brew install func

# Func registry configured
export FUNC_REGISTRY=quay.io/rlehmann
```

**Cluster**

* A Kubernetes cluster
* Knative Serving with Kourier and Knative Eventing installed [script](https://github.com/ReToCode/local-kind-setup)


## Index

* [Slides](./slides.pdf)
* [Resources](./1_RESOURCES.md)
* [Scaling](./2_SCALING.md)
* [Functions](./3_FUNCTIONS.md)
