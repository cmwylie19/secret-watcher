# Secret Watcher

_The secret watcher is a loop that watches secrets._

- [Build](#build)
- [Usage](#usage)
- [Reconcile Loop](#reconcile-loop)



## Usage

Deploy `secret-watcher` manifests

```bash
kubectl apply -f k8s
```

Wait for the secret-watcher to be ready

```bash
kubectl wait --for=condition=Ready pod -l app=secret-watcher -n secret-watcher
```


Curl secrets from all namespaces

```bash
curl $(kubectl get route secret-watcher -n secret-watcher --template='{{ .spec.host }}')/secrets
```

output

```bash
builder-token-brw8p
builder-token-z42rl
cert-manager-cainjector-dockercfg-64fwt
cert-manager-cainjector-token-2j8qb
cert-manager-cainjector-token-tzls7
cert-manager-dockercfg-wttq6
cert-manager-startupapicheck-dockercfg-mmk59
cert-manager-startupapicheck-token-hxdzh
cert-manager-startupapicheck-token-ktdwz
cert-manager-token-jq6ck
cert-manager-token-w7xs2
cert-manager-webhook-ca
cert-manager-webhook-dockercfg-s22v6
cert-manager-webhook-token-42t7j
cert-manager-webhook-token-dh7k2
default-dockercfg-zz77r
default-token-8rnjt
default-token-kxrwh
deployer-dockercfg-8hccg
```

Curl secrets from a given namespace

```bash
curl $(kubectl get route secret-watcher -n secret-watcher --template='{{ .spec.host }}')/secrets\?namespace\=default
```

output

```bash
builder-dockercfg-tq8bq
builder-token-2zzzs
builder-token-44s7t
default-dockercfg-fkf7c
default-token-b8l64
default-token-dszf2
deployer-dockercfg-hbrgv
deployer-token-gkkpg
deployer-token-rnj86
```

## Build 

_This part is only necessary if you are developing._  

Build application

```bash
make image/secret-watcher
```

Build container image

```bash
make docker-image
```

Push container image

```bash
make push-image
```

Fast Build and Deploy

```bash
make image/secret-watcher;make docker-image;make push-image
```


## Reconcile Loop

A reconcile loop for secrets can be implemented with an infinity for loop.  

You can search for a specific secret in a specific namespace

```go
clientset.CoreV1().Secrets("namespace-name").Get(context.TODO(), "secret-name", metav1.GetOptions{})
```

Or, you could search for all secrets

```go
// creates the in-cluster config
config, err := rest.InClusterConfig()
if err != nil {
    panic(err.Error())
}
// creates the clientset
clientset, err := kubernetes.NewForConfig(config)
if err != nil {
    panic(err.Error())
}
for {
    // get secrets in all the namespaces by omitting namespace
    // Or specify namespace to get secrets in particular namespace
    secrets, err := clientset.CoreV1().Secrets("").List(context.TODO(), metav1.ListOptions{})
    if err != nil {
        panic(err.Error())
    }
    fmt.Printf("There are %d secrets in the cluster\n", len(secrets.Items))

    // Examples for error handling:
    // - Use helper functions e.g. errors.IsNotFound()
    // - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
    _, err = clientset.CoreV1().Secrets("default").Get(context.TODO(), "example-xxxxx", metav1.GetOptions{})
    if errors.IsNotFound(err) {
        fmt.Printf("Secret example-xxxxx not found in default namespace\n")
    } else if statusError, isStatus := err.(*errors.StatusError); isStatus {
        fmt.Printf("Error getting secret %v\n", statusError.ErrStatus.Message)
    } else if err != nil {
        panic(err.Error())
    } else {
        fmt.Printf("Found example-xxxxx secret in default namespace\n")
    }

    time.Sleep(10 * time.Second) //configurable
}
```