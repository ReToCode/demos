# A simple Knative Service

```bash
kubectl apply -f - <<EOF
apiVersion: serving.knative.dev/v1
kind: Service
metadata:
  name: helloworld-go
  namespace: default
spec:
  template:
    spec:
      containers:
        - image: ghcr.io/knative/helloworld-go:latest
EOF
```

Creates all the necessary Kubernetes resources:

<details>
  <summary>Deployment</summary>

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: helloworld-go-00001
    serving.knative.dev/configuration: helloworld-go
    serving.knative.dev/configurationGeneration: "1"
    serving.knative.dev/configurationUID: f92c2dca-758a-4885-9efc-d7cf9b6ae319
    serving.knative.dev/revision: helloworld-go-00001
    serving.knative.dev/revisionUID: 08f0a478-8d37-4698-962a-003c95aff80a
    serving.knative.dev/service: helloworld-go
    serving.knative.dev/serviceUID: d2dc79de-7c66-4000-9f86-e11a1e719c5e
  name: helloworld-go-00001-deployment
  namespace: default
  ownerReferences:
    - apiVersion: serving.knative.dev/v1
      blockOwnerDeletion: true
      controller: true
      kind: Revision
      name: helloworld-go-00001
      uid: 08f0a478-8d37-4698-962a-003c95aff80a
spec:
  progressDeadlineSeconds: 600
  replicas: 0
  revisionHistoryLimit: 10
  selector:
    matchLabels:
      serving.knative.dev/revisionUID: 08f0a478-8d37-4698-962a-003c95aff80a
  strategy:
    rollingUpdate:
      maxSurge: 25%
      maxUnavailable: 0
    type: RollingUpdate
  template:
    metadata:
      labels:
        app: helloworld-go-00001
        serving.knative.dev/configuration: helloworld-go
        serving.knative.dev/configurationGeneration: "1"
        serving.knative.dev/configurationUID: f92c2dca-758a-4885-9efc-d7cf9b6ae319
        serving.knative.dev/revision: helloworld-go-00001
        serving.knative.dev/revisionUID: 08f0a478-8d37-4698-962a-003c95aff80a
        serving.knative.dev/service: helloworld-go
        serving.knative.dev/serviceUID: d2dc79de-7c66-4000-9f86-e11a1e719c5e
    spec:
      containers:
        - env:
            - name: TARGET
              value: Go Sample v1
            - name: PORT
              value: "8080"
            - name: K_REVISION
              value: helloworld-go-00001
            - name: K_CONFIGURATION
              value: helloworld-go
            - name: K_SERVICE
              value: helloworld-go
          image: ghcr.io/knative/helloworld-go@sha256:530609c992180374eea5e5cef1b4beab172c085bcaf77193103f1e8eb17ef999
          imagePullPolicy: IfNotPresent
          lifecycle:
            preStop:
              httpGet:
                path: /wait-for-drain
                port: 8022
                scheme: HTTP
          name: user-container
          ports:
            - containerPort: 8080
              name: user-port
              protocol: TCP
          resources: {}
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: FallbackToLogsOnError
        - env:
            - name: SERVING_NAMESPACE
              value: default
            - name: SERVING_SERVICE
              value: helloworld-go
            - name: SERVING_CONFIGURATION
              value: helloworld-go
            - name: SERVING_REVISION
              value: helloworld-go-00001
            - name: QUEUE_SERVING_PORT
              value: "8012"
            - name: QUEUE_SERVING_TLS_PORT
              value: "8112"
            - name: CONTAINER_CONCURRENCY
              value: "0"
            - name: REVISION_TIMEOUT_SECONDS
              value: "300"
            - name: REVISION_RESPONSE_START_TIMEOUT_SECONDS
              value: "0"
            - name: REVISION_IDLE_TIMEOUT_SECONDS
              value: "0"
            - name: SERVING_POD
              valueFrom:
                fieldRef:
                  apiVersion: v1
                  fieldPath: metadata.name
            - name: SERVING_POD_IP
              valueFrom:
                fieldRef:
                  apiVersion: v1
                  fieldPath: status.podIP
            - name: SERVING_LOGGING_CONFIG
            - name: SERVING_LOGGING_LEVEL
            - name: SERVING_REQUEST_LOG_TEMPLATE
              value: '{"httpRequest": {"requestMethod": "{{.Request.Method}}", "requestUrl":
            "{{js .Request.RequestURI}}", "requestSize": "{{.Request.ContentLength}}",
            "status": {{.Response.Code}}, "responseSize": "{{.Response.Size}}", "userAgent":
            "{{js .Request.UserAgent}}", "remoteIp": "{{js .Request.RemoteAddr}}",
            "serverIp": "{{.Revision.PodIP}}", "referer": "{{js .Request.Referer}}",
            "latency": "{{.Response.Latency}}s", "protocol": "{{.Request.Proto}}"},
            "traceId": "{{index .Request.Header "X-B3-Traceid"}}"}'
            - name: SERVING_ENABLE_REQUEST_LOG
              value: "false"
            - name: SERVING_REQUEST_METRICS_BACKEND
              value: prometheus
            - name: TRACING_CONFIG_BACKEND
              value: none
            - name: TRACING_CONFIG_ZIPKIN_ENDPOINT
            - name: TRACING_CONFIG_DEBUG
              value: "false"
            - name: TRACING_CONFIG_SAMPLE_RATE
              value: "0.1"
            - name: USER_PORT
              value: "8080"
            - name: SYSTEM_NAMESPACE
              value: knative-serving
            - name: METRICS_DOMAIN
              value: knative.dev/internal/serving
            - name: SERVING_READINESS_PROBE
              value: '{"tcpSocket":{"port":8080,"host":"127.0.0.1"},"successThreshold":1}'
            - name: ENABLE_PROFILING
              value: "false"
            - name: SERVING_ENABLE_PROBE_REQUEST_LOG
              value: "false"
            - name: METRICS_COLLECTOR_ADDRESS
            - name: HOST_IP
              valueFrom:
                fieldRef:
                  apiVersion: v1
                  fieldPath: status.hostIP
            - name: ENABLE_HTTP2_AUTO_DETECTION
              value: "false"
            - name: ROOT_CA
          image: gcr.io/knative-releases/knative.dev/serving/cmd/queue@sha256:65c427aaab3be9cea1afea32cdef26d5855c69403077d2dc3439f75c26a1e83f
          imagePullPolicy: IfNotPresent
          name: queue-proxy
          ports:
            - containerPort: 8022
              name: http-queueadm
              protocol: TCP
            - containerPort: 9090
              name: http-autometric
              protocol: TCP
            - containerPort: 9091
              name: http-usermetric
              protocol: TCP
            - containerPort: 8012
              name: queue-port
              protocol: TCP
            - containerPort: 8112
              name: https-port
              protocol: TCP
          readinessProbe:
            failureThreshold: 3
            httpGet:
              httpHeaders:
                - name: K-Network-Probe
                  value: queue
              path: /
              port: 8012
              scheme: HTTP
            periodSeconds: 10
            successThreshold: 1
            timeoutSeconds: 1
          resources:
            requests:
              cpu: 25m
          securityContext:
            allowPrivilegeEscalation: false
            capabilities:
              drop:
                - ALL
            readOnlyRootFilesystem: true
            runAsNonRoot: true
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
      dnsPolicy: ClusterFirst
      enableServiceLinks: false
      restartPolicy: Always
      schedulerName: default-scheduler
      securityContext: {}
      terminationGracePeriodSeconds: 300
```
</details>

<details>
  <summary>Service</summary>

```yaml
apiVersion: v1
kind: Service
metadata:
  annotations:
    serving.knative.dev/creator: kubernetes-admin
    serving.knative.dev/lastModifier: kubernetes-admin
  labels:
    serving.knative.dev/route: helloworld-go
    serving.knative.dev/service: helloworld-go
  name: helloworld-go
  namespace: default
  ownerReferences:
    - apiVersion: serving.knative.dev/v1
      blockOwnerDeletion: true
      controller: true
      kind: Route
      name: helloworld-go
      uid: 7fefbb19-052a-4b70-9e68-65ff8822ef16
spec:
  externalName: kourier-internal.kourier-system.svc.cluster.local
  ports:
    - name: http2
      port: 80
      protocol: TCP
      targetPort: 80
  sessionAffinity: None
  type: ExternalName
```
</details>

And also a further Knative internal resources

```bash
kubectl tree ksvc/helloworld-go -n default
```

```text
W0517 07:57:46.483583   78466 warnings.go:70] metallb.io v1beta1 AddressPool is deprecated, consider using IPAddressPool
NAMESPACE  NAME                                                          READY    REASON             AGE
default    Service/helloworld-go                                         True                        20m
default    ├─Configuration/helloworld-go                                 True                        20m
default    │ └─Revision/helloworld-go-00001                              True                        20m
default    │   ├─Deployment/helloworld-go-00001-deployment               -                           20m
default    │   │ └─ReplicaSet/helloworld-go-00001-deployment-557b74997f  -                           20m
default    │   ├─Image/helloworld-go-00001-cache-user-container          -                           20m
default    │   └─PodAutoscaler/helloworld-go-00001                       False    NoTraffic          20m
default    │     ├─Metric/helloworld-go-00001                            True                        20m
default    │     └─ServerlessService/helloworld-go-00001                 Unknown  NoHealthyBackends  20m
default    │       ├─Endpoints/helloworld-go-00001                       -                           20m
default    │       │ └─EndpointSlice/helloworld-go-00001-h7zd4           -                           20m
default    │       ├─Service/helloworld-go-00001                         -                           20m
default    │       └─Service/helloworld-go-00001-private                 -                           20m
default    │         └─EndpointSlice/helloworld-go-00001-private-xvlkd   -                           20m
default    └─Route/helloworld-go                                         True                        20m
default      ├─Ingress/helloworld-go                                     True                        20m
default      └─Service/helloworld-go                                     -                           20m
```
