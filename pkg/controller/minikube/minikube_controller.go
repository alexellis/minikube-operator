package minikube

import (
	"context"
	"fmt"

	alexellisv2alpha1 "github.com/alexellis/minikube-operator/pkg/apis/alexellis/v2alpha1"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_minikube")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new Minikube Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileMinikube{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("minikube-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource Minikube
	err = c.Watch(&source.Kind{Type: &alexellisv2alpha1.Minikube{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner Minikube
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &alexellisv2alpha1.Minikube{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileMinikube{}

// ReconcileMinikube reconciles a Minikube object
type ReconcileMinikube struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a Minikube object and makes changes based on the state read
// and what is in the Minikube.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileMinikube) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling Minikube")

	// Fetch the Minikube instance
	instance := &alexellisv2alpha1.Minikube{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Define a new Pod object
	pod := newMinikubePod(instance)

	// Set Minikube instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, pod, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this Pod already exists
	found := &corev1.Pod{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: pod.Name, Namespace: pod.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Pod", "Pod.Namespace", pod.Namespace, "Pod.Name", pod.Name)
		err = r.client.Create(context.TODO(), pod)
		if err != nil {
			return reconcile.Result{}, err
		}

		// Pod created successfully - don't requeue
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	// Pod already exists - don't requeue
	reqLogger.Info("Skip reconcile: Pod already exists", "Pod.Namespace", found.Namespace, "Pod.Name", found.Name)
	return reconcile.Result{}, nil
}

// // newPodForCR returns a busybox pod with the same name/namespace as the cr
// func newPodForCR(cr *alexellisv2alpha1.Minikube) *corev1.Pod {
// 	labels := map[string]string{
// 		"app": cr.Name,
// 	}
// 	return &corev1.Pod{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name:      cr.Name + "-pod",
// 			Namespace: cr.Namespace,
// 			Labels:    labels,
// 		},
// 		Spec: corev1.PodSpec{
// 			Containers: []corev1.Container{
// 				{
// 					Name:    "busybox",
// 					Image:   "busybox",
// 					Command: []string{"sleep", "3600"},
// 				},
// 			},
// 		},
// 	}
// }

func newMinikubePod(cr *alexellisv2alpha1.Minikube) *corev1.Pod {
	labels := map[string]string{
		"app": cr.ObjectMeta.Name + "-minikube",
	}
	privileged := true
	propagate := corev1.MountPropagationHostToContainer
	// user := int64(0)
	pathTypeHostDir := corev1.HostPathDirectory
	return &corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.ObjectMeta.Name + "-minikube",
			Namespace: cr.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(cr, schema.GroupVersionKind{
					Group:   alexellisv2alpha1.SchemeGroupVersion.Group,
					Version: alexellisv2alpha1.SchemeGroupVersion.Version,
					Kind:    "Minikube",
				}),
			},
			Labels: labels,
		},
		Spec: corev1.PodSpec{
			Volumes: []corev1.Volume{
				{
					Name: "kvm",
					VolumeSource: corev1.VolumeSource{
						HostPath: &corev1.HostPathVolumeSource{
							Path: "/dev/kvm",
						},
					},
				},
				{
					Name: "virsh-lib",
					VolumeSource: corev1.VolumeSource{
						HostPath: &corev1.HostPathVolumeSource{
							Path: "/var/lib/libvirt",
							Type: &pathTypeHostDir,
						},
					},
				},
				{
					Name: "virsh-run",
					VolumeSource: corev1.VolumeSource{
						HostPath: &corev1.HostPathVolumeSource{
							Path: "/var/run/libvirt",
							Type: &pathTypeHostDir,
						},
					},
				},
				{
					Name: "var-mkaas",
					VolumeSource: corev1.VolumeSource{
						HostPath: &corev1.HostPathVolumeSource{
							Path: "/var/mkaas",
							Type: &pathTypeHostDir,
						},
					},
				},
			},
			HostNetwork: true,
			Containers: []corev1.Container{
				{
					Name:  "libvirt",
					Image: "alexellis2/libvirt-xenial-minikube:0.4",
					SecurityContext: &corev1.SecurityContext{
						Privileged: &privileged,
					},
					Env: []corev1.EnvVar{
						{
							Name:  "CLUSTER_CPUS",
							Value: fmt.Sprintf("%d", cr.Spec.CPUCount),
						},
						{
							Name:  "CLUSTER_MEMORY",
							Value: fmt.Sprintf("%d", cr.Spec.MemoryMB),
						},
						{
							Name:  "CLUSTER_NAME",
							Value: cr.Spec.ClusterName,
						},
					},
					Ports: []corev1.ContainerPort{},
					VolumeMounts: []corev1.VolumeMount{
						{
							MountPath:        "/var/run/libvirt",
							Name:             "virsh-run",
							MountPropagation: &propagate,
						},
						{
							MountPath:        "/var/lib/libvirt",
							Name:             "virsh-lib",
							MountPropagation: &propagate,
						},
						{
							MountPath:        "/var/mkaas",
							Name:             "var-mkaas",
							MountPropagation: &propagate,
						},
					},
				},
			},
		},
	}
}
