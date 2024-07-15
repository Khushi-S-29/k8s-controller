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

package controllers

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	mathsv1 "gitlab.com/avengehers/k8s-controller/api/v1"
)

// ArithmeticReconciler reconciles a Arithmetic object
type ArithmeticReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=maths.stream.com,resources=arithmetics,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=maths.stream.com,resources=arithmetics/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=maths.stream.com,resources=arithmetics/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Arithmetic object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.1/pkg/reconcile

func (r *ArithmeticReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

	log := log.FromContext(ctx)

	var problem mathsv1.Arithmetic
	if err := r.Get(ctx, req.NamespacedName, &problem); err != nil {
		log.Error(err, "could not get the Arithmetic object")
	}

	log.Info(fmt.Sprintf("Reconciling for %s", req.NamespacedName))
	log.Info(fmt.Sprintf("Expression: %s", problem.Spec.Expression))

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ArithmeticReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&mathsv1.Arithmetic{}).
		Complete(r)
}
