# Loudspeaker Operator

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Docker](https://img.shields.io/docker/v/masanetes/loudspeaker/v0.0.3?color=blue&logo=docker)](https://hub.docker.com/repository/docker/masanetes/loudspeaker)
[![Go Reference](https://pkg.go.dev/badge/github.com/masanetes/loudspeaker.svg)](https://pkg.go.dev/github.com/masanetes/loudspeaker)
[![Test](https://github.com/masanetes/loudspeaker/actions/workflows/test.yaml/badge.svg)](https://github.com/masanetes/loudspeaker/actions/workflows/test.yaml)
[![report](https://goreportcard.com/badge/github.com/masanetes/loudspeaker)](https://goreportcard.com/report/github.com/masanetes/loudspeaker)
[![codecov](https://codecov.io/gh/masanetes/loudspeaker/branch/master/graph/badge.svg?token=9HT5CC8XDK)](https://codecov.io/gh/masanetes/loudspeaker)

Loudspeaker uses the operator pattern to manage the runtime that delivers events from KubeAPI to listeners.

It can be configured which namespace events to observe, which events to discard, etc.
These settings are managed by configmap, and the runtime can change its observation target in real time by overwriting its own settings when it detects a configuration change.

[See runtime doc for details](https://github.com/masanetes/loudspeaker-runtime)


```mermaid
flowchart LR
subgraph 4_[Kubernetes Cluster]
  A[KubeAPI] -->|Events| B(loudspeaker)
end  
B -->|Events| C[Listener1]
B -->|Events| D[Listener2]
B -->|Events| E[Listener3]
```

# 🚀 Quick Start

## Install Loudspeaker Operator

```
$ kubectl apply -f https://raw.githubusercontent.com/masanetes/loudspeaker/master/install/install.yaml
```


## Preparation of runtime setting

Create serviceaccount and clusterrolebindings for runtime to observe events. 
This serviceaccount is associated with the serviceaccountname of the Pod by the operator.

```
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: ServiceAccount
metadata:
  name: loudspeaker-runtime
  namespace: default
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: loudspeaker-runtime
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: view
subjects:
  - kind: ServiceAccount
    name: loudspeaker-runtime
    namespace: default
EOF
```

## Preparation of confidential listener information

Confidential information is needed to post the event to the destination, so it is managed in secrets. CRD also needs to set a type for each listener, since the format of the setting varies depending on the destination.

```
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: sentry-secrets
type: Opaque
stringData:
  credentilas.yaml: |
    dsn: "sample"
EOF
```

## Sample custom resource

https://github.com/masanetes/loudspeaker/blob/master/config/samples/loudspeaker_v1alpha1_loudspeaker.yaml

```yaml
apiVersion: loudspeaker.masanetes.github.io/v1alpha1
kind: Loudspeaker
metadata:
  name: loudspeaker-sample
  namespace: default  
spec:
  image: masanetes/loudspeaker-runtime:latest
  serviceAccountName: loudspeaker-runtime  
  listeners:
    - name: foo
      type: sentry
      credentials: sentry-secrets
      subscribes:
        - namespace: "" # all namespaces
          ignore: ["BackoffLimitExceeded"]
        - namespace: "default"
          ignore: ["Unhealthy"]
    
    - name: bar
      type: sentry
      credentials: sentry-secrets
      subscribes:
        - namespace: "" # all namespaces
          ignore: ["BackoffLimitExceeded"]
        - namespace: "default"
          ignore: ["Unhealthy"]
    
    - name: baz
      type: sentry
      credentials: sentry-secrets
      subscribes:
        - namespace: "" # all namespaces
          ignore: ["BackoffLimitExceeded"]
        - namespace: "default"
          ignore: ["Unhealthy"]
```
