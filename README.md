# minikube-operator

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

Install the pre-reqs from mkaas.

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
operator-sdk build alexellis2/minikube-operator

docker push alexellis2/minikube-operator:0.1.1
```
