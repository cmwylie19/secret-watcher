# Deploy

The environment to build, one of: dev, stage, prod
Default to dev if not set.
- **dev** is compiled for amd64 architecture (Kind)
- **stage** is compiled for arm64 architecture (Rasberry Pi)
- **prod** prod is compiled for amd64 architecture (OpenShift)


```bash
make deploy/<environment> |  kubectl apply -f -
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


```bash
make deploy/<environment> | kubectl delete -f -
```