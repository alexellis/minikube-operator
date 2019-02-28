FROM ubuntu:xenial
MAINTAINER alex@openfaas.com

RUN apt update \
    && DEBIAN_FRONTEND=noninteractive DEBCONF_NONINTERACTIVE_SEEN=true  apt install -qy libvirt-bin curl ca-certificates --no-install-recommends

RUN curl -SLO https://github.com/kubernetes/minikube/releases/download/v0.30.0/docker-machine-driver-kvm2 \
 && curl -SLO https://github.com/kubernetes/minikube/releases/download/v0.34.1/minikube-linux-amd64 \
 && chmod +x docker-machine-driver-kvm2 \
 && chmod +x minikube-linux-amd64 \
 && mv docker-machine-driver-kvm2 /usr/local/bin \
 && mv minikube-linux-amd64 /usr/local/bin/minikube \
 && curl -Lo kubectl https://storage.googleapis.com/kubernetes-release/release/v1.13.0/bin/linux/amd64/kubectl && chmod +x kubectl && mv kubectl /usr/local/bin/

COPY entry.sh	.
RUN chmod +x ./entry.sh

CMD ["/bin/bash", "./entry.sh"]
