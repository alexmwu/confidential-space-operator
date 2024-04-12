/*
Copyright 2024.

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

	"google.golang.org/api/googleapi"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	teev1alpha1 "confidential-space-operator/api/v1alpha1"
	"confidential-space-operator/pkg/confidentialspace"
)

// ConfidentialSpaceReconciler reconciles a ConfidentialSpace object
type ConfidentialSpaceReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	csClient confidentialspace.Client
}

//+kubebuilder:rbac:groups=tee.cloud.google.com,resources=confidentialspaces,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=tee.cloud.google.com,resources=confidentialspaces/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=tee.cloud.google.com,resources=confidentialspaces/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ConfidentialSpace object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.2/pkg/reconcile
func (r *ConfidentialSpaceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logc := log.FromContext(ctx)

	instances, err := r.csClient.List(ctx)
	if err != nil {
		logc.Error(err, "unable to fetch instance")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		if e, ok := err.(*googleapi.Error); ok {
			if e.Code == 404 {
				return ctrl.Result{}, nil
			}
		}
		return ctrl.Result{}, err
	}
	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ConfidentialSpaceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&teev1alpha1.ConfidentialSpace{}).
		Complete(r)
}
