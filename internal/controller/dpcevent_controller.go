// internal/controller/dpcevent_controller.go
package controller

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	correlationv1alpha1 "github.com/Joshua-Pok/event-correlation-controller/api/v1alpha1"
)

type DpcEventReconciler struct {
	client.Client
}

//when reconcile is called, we receive a request with just a name, we need to figure out the current Status

/*
1) Fetch the resource- ALWAYS

2) Check resource state

3) compare current state with desired state, and take actions to close the gap
*/
func (r *DpcEventReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	//Fetch the event
	var event correlationv1alpha1.DpcEvent
	if err := r.Get(ctx, req.NamespacedName, &event); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	logger.Info("reconciling event", "name", event.Name, "severity", event.Spec.Severity)

	//If already correlated, skip event
	if event.Status.RoutedTo

	return ctrl.Result{}, nil
}

func (r *DpcEventReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&correlationv1alpha1.DpcEvent{}).
		Complete(r)
}
