# Knative Autoscaling

## Tweak Knative settings for demo purposes

> 📝 Knative provides a lot of default values. For the demo we make them a bit more snappy

```bash
kubectl patch cm config-autoscaler -n knative-serving -p '{"data": {"stable-window": "10s"}}'
kubectl patch cm config-autoscaler -n knative-serving -p '{"data": {"allow-zero-initial-scale": "true"}}'
kubectl patch cm config-autoscaler -n knative-serving -p '{"data": {"scale-to-zero-grace-period": "1s"}}'
kubectl patch cm config-autoscaler -n knative-serving -p '{"data": {"container-concurrency-target-percentage": "0.7"}}'
```

## Watching pods and Knative internals

```bash
watch -n 1 kubectl get pod -n default
```
```bash
watch -n 1 kubectl get podautoscaler -n default
```

## Scale to zero

```bash
kubectl apply -f - <<EOF
apiVersion: serving.knative.dev/v1
kind: Service
metadata:
  name: scale-to-zero
  namespace: default
spec:
  template:
    spec:
      containers:
        - image: ghcr.io/knative/helloworld-go:latest
          env:
            - name: TARGET
              value: "Siemens folks 👋"        
EOF
```

```bash
curl http://scale-to-zero.default.10.89.0.200.sslip.io
```

## Scale from zero

> 📝 Per default a service will always be scaled to 1 at creation to verify that it works properly

```bash
kubectl apply -f - <<EOF
apiVersion: serving.knative.dev/v1
kind: Service
metadata:
  name: scale-to-zero
  namespace: default
spec:
  template:
    spec:
      containers:
        - image: ghcr.io/knative/helloworld-go:latest
          env:
            - name: TARGET
              value: "Siemens folks 👋"        
EOF
```

```bash
kubectl apply -f - <<EOF
apiVersion: serving.knative.dev/v1
kind: Service
metadata:
  name: scale-from-zero
  namespace: default
spec:
  template:
    metadata:
      annotations:
        autoscaling.knative.dev/initial-scale: "0"  
    spec:
      containers:
        - image: ghcr.io/knative/helloworld-go:latest
          env:
            - name: TARGET
              value: "Siemens folks 👋"        
EOF
```

```bash
curl http://scale-from-zero.default.10.89.0.200.sslip.io
```

## Scaling based on targets

> 📝 Knative supports multiple [metric targets](https://knative.dev/docs/serving/autoscaling/autoscaling-metrics/), this example uses concurrent requests

```bash
kubectl apply -f - <<EOF
apiVersion: serving.knative.dev/v1
kind: Service
metadata:
  name: scale-target
  namespace: default
spec:
  template:
    metadata:
      annotations:
        autoscaling.knative.dev/metric: "concurrency"
        autoscaling.knative.dev/target: "1"
        autoscaling.knative.dev/initial-scale: "0"
    spec:
      containers:
        - image: ghcr.io/knative/helloworld-go:latest
          env:
            - name: TARGET
              value: "Siemens folks 👋"        
EOF
```

### 🙋‍♂️🙋‍♀ How many pods are we going to see?️

```bash
curl http://scale-target.default.10.89.0.200.sslip.io
```

```bash
kubectl patch cm config-autoscaler -n knative-serving -p '{"data": {"container-concurrency-target-percentage": "1"}}'
```

### 🙋‍♂️🙋‍♀ And now?

```bash
curl http://scale-target.default.10.89.0.200.sslip.io
```

### 🙋‍♂️🙋‍♀ And now?

```bash
# Send 20 concurrent requests for 60 seconds
hey -z 60s -c 20 http://scale-target.default.10.89.0.200.sslip.io
```


## Knative can do even more

> 📝 You can find more details on the autoscaler [here](https://knative.dev/docs/serving/autoscaling/)
