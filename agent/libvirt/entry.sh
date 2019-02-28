#!/bin/bash

export MINIKUBE_HOME=/var/mkaas/.minikube/

exit_script() {
    echo "Stopping cluster $CLUSTER_NAME"
    minikube delete --profile $CLUSTER_NAME

    trap - SIGINT SIGTERM # clear the trap
}

trap exit_script SIGINT SIGTERM

minikube start --bootstrapper=kubeadm --vm-driver=kvm2 \
    --memory $CLUSTER_MEMORY \
    --cpus $CLUSTER_CPUS \
    --profile $CLUSTER_NAME

cd /var/mkaas/

sed -ie 's#/var/mkaas/#../#g' /root/.kube/config
mkdir -p .kube
cp /root/.kube/config ./.kube/

tar -czf /var/mkaas/bundle.tgz .minikube/*.crt .minikube/*.key .minikube/*.pem .kube/config

echo "/var/mkaas/bundle.tgz written."

while [ true ] ;
do
   sleep 5
done
