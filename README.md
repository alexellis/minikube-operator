# minikube-operator

## Status

This is the port of my [mkaas](https://github.com/alexellis/mkaas) project, which used an early version of the operator-framework. CoreOS made so many changes that the code stopped working all-together, so this repo contains the updated versions.

Goals
- [x] Port over CRD / operator
- [x] Update to Kubernetes v1.13
- [ ] Port over Nginx DaemonSet for accessing the clusters remotely

## Preview

See [my Tweet](https://twitter.com/alexellisuk/status/1101185378277564417) with a Kubernetes cluster booting up on one of my Kubernetes nodes.

![](https://pbs.twimg.com/media/D0gyGAAWwAAiuNE.jpg)

## Example

```yaml
apiVersion: alexellis.io/v2alpha1
kind: Minikube
metadata:
  name: "inlets"
  namespace: clusters
spec:
  clusterName: "inlets"
  cpuCount: 2
  memoryMB: 2048
```

## Try it out

Install the pre-reqs from [mkaas](https://github.com/alexellis/mkaas).

```
git clone git@github.com:alexellis/minikube-operator.git
cd minikube-operator

kubectl create ns clusters
kubectl apply -f ./deploy
kubectl apply -f ./deploy/crds/
```

View logs:

```
# See CR get picked up and created
kubectl logs deploy/minikube-operator -n clusters -f

# See logs of CR's boot-up script where it creates the minikube cluster with KVM
# Wait until you see text about creating the bundle.
kubectl logs pod/inlets-minikube -n clusters -f
```

Minikube data will be available as a bundle in `/var/mkaas` on the host.

You can un-tar it to /root/ and run `kubectl get node`

## Hacking

Build:

```bash
operator-sdk build alexellis2/minikube-operator:0.2.0
docker push alexellis2/minikube-operator:0.2.0
```
## License

MIT
