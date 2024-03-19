package controller

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"time"

	"github.com/ONE7live/kompass/pkg/apis/v1alpha1"
)

type LeafPodConvertReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

var predicatesFunc = predicate.Funcs{
	CreateFunc: func(createEvent event.CreateEvent) bool {
		fmt.Println("Create: ", createEvent.Object)
		return true
	},
	UpdateFunc: func(updateEvent event.UpdateEvent) bool {
		fmt.Println("Update: ", updateEvent.ObjectNew)
		return true
	},
	DeleteFunc: func(deleteEvent event.DeleteEvent) bool {
		fmt.Println("Delete: ", deleteEvent.Object)
		return true
	},
	GenericFunc: func(genericEvent event.GenericEvent) bool {
		fmt.Println("Generic: ", genericEvent.Object)
		return false
	},
}

func (r *LeafPodConvertReconciler) Reconcile(ctx context.Context, req ctrl.Request) (reconcile.Result, error) {
	klog.V(4).Infof("============ %s starts to reconcile %s ============", "LeafPodConvert-Controller", req.NamespacedName)

	leafPodConvertPolicyList := &v1alpha1.LeafPodConvertPolicyList{}
	if err := r.Client.List(ctx, leafPodConvertPolicyList); err != nil {
		klog.Errorf("failed to list clusters, error: %v", err)
		return reconcile.Result{RequeueAfter: 10 * time.Second}, nil
	}

	klog.Infof("leaf pod convert obj: %v", req.NamespacedName)

	return ctrl.Result{}, nil
}

func (r *LeafPodConvertReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.LeafPodConvertPolicy{}, builder.WithPredicates(predicatesFunc)).
		Complete(r)
}
