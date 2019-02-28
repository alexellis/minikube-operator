package controller

import (
	"github.com/alexellis/minikube-operator/pkg/controller/minikube"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, minikube.Add)
}
