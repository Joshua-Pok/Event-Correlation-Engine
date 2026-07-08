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

	if event.Status.Correlated {
		logger.Info("Event already correlated, skipping")
		return ctrl.Result{}, nil
	}

	group, found, err := r.findOpenGroup(ctx, event.Spec.Node, event.Namespace)
	if err != nil {
		logger.Error(err, "Failed to open group")
		return ctrl.Result{}, err
	}

	//create if does not exist if not update
	if !found {
		group, err = r.Create(ctx, &event)
		if err != nil {
			logger.Error(err, "Failed to create group")
			return ctrl.Result{}, err
		}
		logger.Info("Created a new DPCEVent: ", "group", group.Name)
	} else {
		err = r.addEventToGroup(ctx, &grp, &event)
		if err != nil {
			logger.Error(err, "Failed to add event to group")
			return ctrl.Result{}, err
		}
	}

	//cleanup --> mark event as correlated

	event.Status.Correlated = true
	event.Status.GroupRef = group.Name
	event.Status.ProcessedAt = metav1.Now()

	logger.Info("Successfully correlated event")

	return ctrl.Result{}, nil
}

func (r *DpcEventReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&correlationv1alpha1.DpcEvent{}).
		Complete(r)
}
