# Deploy

The environment to build, one of: dev, stage, prod
Default to dev if not set.
- **dev** is compiled for amd64 architecture (Kind)
- **stage** is compiled for arm64 architecture (Raspberry Pi)
- **prod** prod is compiled for amd64 architecture (OpenShift)


```bash
make deploy/<environment> |  kubectl apply -f -
```

output

```bash
namespace/secret-watcher created
serviceaccount/secret-watcher created
clusterrole.rbac.authorization.k8s.io/secret-view created
clusterrolebinding.rbac.authorization.k8s.io/secret-view-binding created
service/secret-watcher created
deployment.apps/secret-watcher created
```

Wait for the secret-watcher to be ready

```bash
kubectl wait --for=condition=Ready pod -l app=secret-watcher -n secret-watcher
```

## Uninstall

```bash
kubectl delete deploy,svc,sa,po -n secret-watcher --all --force --grace-period=0

kubectl delete clusterrolebinding secret-view-binding

kubectl delete clusterrole secret-view

kubectl delete ns secret-watcher
```

or (Risky to delete namespace before other objects)

```bash
make deploy/<environment> | kubectl delete -f -
```