/*
Copyright 2023 Nguyen Thanh Nguyen.

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

	"github.com/go-logr/logr"
	intentv1 "github.com/ntnguyencse/intent-kaas/api/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// BlueprintReconciler reconciles a Blueprint object
type BlueprintReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	l      logr.Logger
	s      *json.Serializer
}

//+kubebuilder:rbac:groups=intent.automation.dcn.ssu.ac.kr,resources=blueprints,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=intent.automation.dcn.ssu.ac.kr,resources=blueprints/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=intent.automation.dcn.ssu.ac.kr,resources=blueprints/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Blueprint object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *BlueprintReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.l = log.FromContext(ctx)
	pd, err := r.startRequest(ctx, req)
	if err != nil {
		if client.IgnoreNotFound(err) != nil {
			r.l.Error(err, "cannot get pd")
			return ctrl.Result{}, err
		}
		r.l.Info("cannot get resource, probably deleted", "error", err.Error())
		r.l.Error(err, "Error when start Request Reconcile function", "error")
		return ctrl.Result{}, nil
	}
	// TODO(user): your logic here
	r.l.Info("Print Object", pd.Name)
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *BlueprintReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&intentv1.Blueprint{}).
		Complete(r)
}

func (r *BlueprintReconciler) startRequest(ctx context.Context, req ctrl.Request) (*intentv1.Blueprint, error) {
	r.l = log.FromContext(ctx)
	r.l.Info("Start Request: Get Blueprint", "req", req)
	var blueprints intentv1.Blueprint
	if err := r.Get(ctx, req.NamespacedName, &blueprints); err != nil {
		r.l.Error(err, "Unable to fetch Blueprint List from Kubernetes API Server")
		return nil, err
	}
	return &blueprints, nil
}
