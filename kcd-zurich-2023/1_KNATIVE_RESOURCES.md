# A simple Knative Service

```bash
kubectl apply -f - <<EOF
apiVersion: serving.knative.dev/v1
kind: Service
metadata:
  name: horse-jaskier
  namespace: default
spec:
  template:
    spec:
      containers:
        - image: quay.io/rlehmann/horse-jaskier
EOF
```

Creates all the necessary Kubernetes resources:

<details>
  <summary>Deployment</summary>

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  annotations:
    deployment.kubernetes.io/revision: "1"
    serving.knative.dev/creator: kubernetes-admin
  creationTimestamp: "2023-06-06T11:32:15Z"
  generation: 1
  labels:
    app: horse-jaskier-00001
    serving.knative.dev/configuration: horse-jaskier
    serving.knative.dev/configurationGeneration: "1"
    serving.knative.dev/configurationUID: 2e0fb544-af0c-42a8-b92b-fdd2f0bf48fa
    serving.knative.dev/revision: horse-jaskier-00001
    serving.knative.dev/revisionUID: b0c332f0-7dbe-4cf8-9d56-829afaa15019
    serving.knative.dev/service: horse-jaskier
    serving.knative.dev/serviceUID: 82240fdb-1349-4ac0-ac76-592d583c9c94
  name: horse-jaskier-00001-deployment
  namespace: default
  ownerReferences:
    - apiVersion: serving.knative.dev/v1
      blockOwnerDeletion: true
      controller: true
      kind: Revision
      name: horse-jaskier-00001
      uid: b0c332f0-7dbe-4cf8-9d56-829afaa15019
  resourceVersion: "338594"
  uid: d7339d98-304a-41d2-b329-78b840e255c0
spec:
  progressDeadlineSeconds: 600
  replicas: 1
  revisionHistoryLimit: 10
  selector:
    matchLabels:
      serving.knative.dev/revisionUID: b0c332f0-7dbe-4cf8-9d56-829afaa15019
  strategy:
    rollingUpdate:
      maxSurge: 25%
      maxUnavailable: 0
    type: RollingUpdate
  template:
    metadata:
      annotations:
        serving.knative.dev/creator: kubernetes-admin
      creationTimestamp: null
      labels:
        app: horse-jaskier-00001
        serving.knative.dev/configuration: horse-jaskier
        serving.knative.dev/configurationGeneration: "1"
        serving.knative.dev/configurationUID: 2e0fb544-af0c-42a8-b92b-fdd2f0bf48fa
        serving.knative.dev/revision: horse-jaskier-00001
        serving.knative.dev/revisionUID: b0c332f0-7dbe-4cf8-9d56-829afaa15019
        serving.knative.dev/service: horse-jaskier
        serving.knative.dev/serviceUID: 82240fdb-1349-4ac0-ac76-592d583c9c94
    spec:
      containers:
        - env:
            - name: PORT
              value: "8080"
            - name: K_REVISION
              value: horse-jaskier-00001
            - name: K_CONFIGURATION
              value: horse-jaskier
            - name: K_SERVICE
              value: horse-jaskier
          image: quay.io/rlehmann/horse-jaskier@sha256:bb54bec6ef0d2f3ada70f7759435e8f3e76dc3076f0599a3c0832e9ffe5721e0
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
              value: horse-jaskier
            - name: SERVING_CONFIGURATION
              value: horse-jaskier
            - name: SERVING_REVISION
              value: horse-jaskier-00001
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
status:
  availableReplicas: 1
  conditions:
    - lastTransitionTime: "2023-06-06T11:32:21Z"
      lastUpdateTime: "2023-06-06T11:32:21Z"
      message: Deployment has minimum availability.
      reason: MinimumReplicasAvailable
      status: "True"
      type: Available
    - lastTransitionTime: "2023-06-06T11:32:15Z"
      lastUpdateTime: "2023-06-06T11:32:21Z"
      message: ReplicaSet "horse-jaskier-00001-deployment-c5ddb99d6" has successfully
        progressed.
      reason: NewReplicaSetAvailable
      status: "True"
      type: Progressing
  observedGeneration: 1
  readyReplicas: 1
  replicas: 1
  updatedReplicas: 1
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
  creationTimestamp: "2023-06-06T11:32:21Z"
  labels:
    serving.knative.dev/route: horse-jaskier
    serving.knative.dev/service: horse-jaskier
  name: horse-jaskier
  namespace: default
  ownerReferences:
    - apiVersion: serving.knative.dev/v1
      blockOwnerDeletion: true
      controller: true
      kind: Route
      name: horse-jaskier
      uid: 36259b4f-491e-4de8-a382-a84a741d607d
  resourceVersion: "338626"
  uid: 48186bfd-4035-4a03-8a36-a646ee29fc45
spec:
  externalName: kourier-internal.kourier-system.svc.cluster.local
  ports:
    - name: http2
      port: 80
      protocol: TCP
      targetPort: 80
  sessionAffinity: None
  type: ExternalName
status:
  loadBalancer: {}
```
</details>

And also a further Knative internal resources

```bash
kubectl tree ksvc/horse-jaskier -n default
```

```text
NAMESPACE  NAME                                                          READY    REASON             AGE
default    Service/horse-jaskier                                         True                        77s
default    ├─Configuration/horse-jaskier                                 True                        77s
default    │ └─Revision/horse-jaskier-00001                              True                        77s
default    │   ├─Deployment/horse-jaskier-00001-deployment               -                           77s
default    │   │ └─ReplicaSet/horse-jaskier-00001-deployment-c5ddb99d6   -                           77s
default    │   │   └─Pod/horse-jaskier-00001-deployment-c5ddb99d6-glw9d  True                        77s
default    │   ├─Image/horse-jaskier-00001-cache-user-container          -                           77s
default    │   └─PodAutoscaler/horse-jaskier-00001                       False    NoTraffic          77s
default    │     ├─Metric/horse-jaskier-00001                            True                        77s
default    │     └─ServerlessService/horse-jaskier-00001                 Unknown  NoHealthyBackends  77s
default    │       ├─Endpoints/horse-jaskier-00001                       -                           77s
default    │       │ └─EndpointSlice/horse-jaskier-00001-l2f8w           -                           77s
default    │       ├─Service/horse-jaskier-00001                         -                           77s
default    │       └─Service/horse-jaskier-00001-private                 -                           77s
default    │         └─EndpointSlice/horse-jaskier-00001-private-bxsz9   -                           77s
default    └─Route/horse-jaskier                                         True                        77s
default      ├─Ingress/horse-jaskier                                     True                        71s
default      └─Service/horse-jaskier                                     -                           71s
```

```bash
kubectl get -n default ksvc
```

```text
NAME            URL                                                 LATESTCREATED         LATESTREADY           READY   REASON
horse-jaskier   https://horse-jaskier-default.apps.rlehmann-ocp-4-12.serverless.devcluster.openshift.com   horse-jaskier-00001   horse-jaskier-00001   True
```