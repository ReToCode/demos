# KCD Zurich 2023

* [Slides](./slides.pdf)


## Setup

You need:
* [Ko](https://github.com/ko-build/ko) installed and registry configured
* A Kubernetes cluster
* Knative Serving with Kourier installed [script](https://github.com/ReToCode/local-kind-setup)
* Knative Eventing installed [script](https://github.com/ReToCode/local-kind-setup)

```bash
./create_cluster.sh
./install_serving_kourier.sh
./install_eventing_kafka.sh
```

### Tweak Serving configuration

> 📝 Knative provides a lot of default values. For the demo we make them a bit more snappy

```bash
oc -n knative-serving patch knativeserving/knative-serving --type=merge --patch='{"spec": {"config": { "autoscaler": {"stable-window": "10s"}}}}'
oc -n knative-serving patch knativeserving/knative-serving --type=merge --patch='{"spec": {"config": { "autoscaler": {"allow-zero-initial-scale": "true"}}}}'
oc -n knative-serving patch knativeserving/knative-serving --type=merge --patch='{"spec": {"config": { "autoscaler": {"scale-to-zero-grace-period": "1s"}}}}'
oc -n knative-serving patch knativeserving/knative-serving --type=merge --patch='{"spec": {"config": { "autoscaler": {"container-concurrency-target-percentage": "0.7"}}}}'
```

### Create the demo resources

```bash
# Kafka resources
kubectl get brokers -A

# Applications
kubectl apply -f oger/oger.yaml
```

### Build and push the applications

```bash
KO_DOCKER_REPO=quay.io/rlehmann ko build --platform=linux/arm64,linux/amd64 --sbom=none -B oger
KO_DOCKER_REPO=quay.io/rlehmann ko build --platform=linux/arm64,linux/amd64 --sbom=none -B horse-jaskier
```
