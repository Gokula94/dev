/*
Copyright 2024 Gokul.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"fmt"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	apiv1alpha1 "dev/api/v1alpha1"
)

var logger = log.Log.WithName("controller_dev")

// DevReconciler reconciles a Dev object
type DevReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=api.gokula.dev,resources=devs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=api.gokula.dev,resources=devs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=api.gokula.dev,resources=devs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Dev object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.3/pkg/reconcile
func (r *DevReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logger.WithValues("Request.Namespace", req.Namespace, "Request.Name", req.Name)

	log.Info("Reconcile called by gokul")

	fmt.Println("Reconcile is called for testing")
	//if req.Namespace == "default" {
	//	log.Info("Pod created\n")
	//} else {
	//	log.Info("Pod deleted\n")
	//}
	//return ctrl.Result{}, nil}

	dev := &apiv1alpha1.Dev{}

	err := r.Get(ctx, req.NamespacedName, dev)
	if err != nil {
		if apierrors.IsNotFound(err) {
			log.Info("dev resource is no found")
			return ctrl.Result{}, nil
			// else {
			// 	log.Info("dev resource is up")
			// }
		}
		log.Error(err, "Failed")
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, err
}

// SetupWithManager sets up the controller with the Manager.
func (r *DevReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apiv1alpha1.Dev{}).
		Complete(r)
}
