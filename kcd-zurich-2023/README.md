# KCD Zurich 2023

* Slides - will be available after the event.


## Setup

You need:
* [Ko](https://github.com/ko-build/ko) installed and registry configured
* A Kubernetes cluster
* Knative Serving with Kourier installed [script](https://github.com/ReToCode/local-kind-setup/blob/main/install_serving_kourier.sh)
* Knative Eventing installed [script](https://github.com/ReToCode/local-kind-setup/blob/main/install_eventing_kafka.sh)

```bash
./create_cluster.sh
./install_serving_kourier.sh
./install_eventing_kafka.sh
```

### Tweak Serving configuration

> 📝 Knative provides a lot of default values. For the demo we make them a bit more snappy

```bash
kubectl patch cm config-autoscaler -n knative-serving -p '{"data": {"stable-window": "10s"}}'
kubectl patch cm config-autoscaler -n knative-serving -p '{"data": {"allow-zero-initial-scale": "true"}}'
kubectl patch cm config-autoscaler -n knative-serving -p '{"data": {"scale-to-zero-grace-period": "1s"}}'
kubectl patch cm config-autoscaler -n knative-serving -p '{"data": {"container-concurrency-target-percentage": "0.7"}}'
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
