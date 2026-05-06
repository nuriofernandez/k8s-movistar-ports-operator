package controller

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/nuriofernandez/movistarapi/hgu"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	movistarv1alpha1 "github.com/nuriofernandez/k8s-movistar-ports-operator/api/v1alpha1"
	"github.com/nuriofernandez/movistarapi"
)

type MovistarPortReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r *MovistarPortReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	var mp movistarv1alpha1.MovistarPort
	if err := r.Get(ctx, req.NamespacedName, &mp); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	logger.Info("reconciling MovistarPort", "name", req.Name, "port", mp.Spec.ExternalPort, "protocol", mp.Spec.Protocol)

	// Attempt to connect to router
	hguRouter, err := movistarapi.HGULogin(os.Getenv("MOVISTAR_ROUTER_PASS"))
	if err != nil {
		logger.Error(err, "reconciling MovistarPort failed. Unable to connect to router", "name", req.Name, "port", mp.Spec.ExternalPort, "protocol", mp.Spec.Protocol)
		return ctrl.Result{RequeueAfter: time.Second * 30}, nil
	}

	// Connected! Verify if port is already open
	ports, err := hguRouter.OpenPorts()
	if err != nil {
		logger.Error(err, "reconciling MovistarPort failed. Unable to verify router open ports", "name", req.Name, "port", mp.Spec.ExternalPort, "protocol", mp.Spec.Protocol)
		return ctrl.Result{RequeueAfter: time.Second * 30}, nil
	}

	// If already open, no action needed
	for _, existingPort := range ports {
		if existingPort.ExternalPortStart == int(mp.Spec.ExternalPort) {
			logger.Info("port is already open, no action required")
			return ctrl.Result{}, nil
		}
	}

	// Not open! open it!
	err = hguRouter.OpenPort(hgu.OpenPort{
		Name:              "k8s-m-" + strconv.Itoa(int(mp.Spec.ExternalPort)),
		Protocol:          hgu.Protocol(mp.Spec.Protocol),
		Address:           mp.Spec.Host,
		ExternalPortStart: int(mp.Spec.ExternalPort),
		ExternalPortEnd:   0,
		InternalPortStart: int(mp.Spec.InternalPort),
		Enabled:           true,
		Interface:         "ppp0.1",
	})
	if err != nil {
		logger.Error(err, "reconciling MovistarPort failed. Unable to verify router open ports", "name", req.Name, "port", mp.Spec.ExternalPort, "protocol", mp.Spec.Protocol)
		return ctrl.Result{RequeueAfter: time.Second * 30}, nil
	}

	logger.Info("successfully opened port!", "name", req.Name, "port", mp.Spec.ExternalPort, "protocol", mp.Spec.Protocol)
	return ctrl.Result{}, nil
}

func (r *MovistarPortReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&movistarv1alpha1.MovistarPort{}).
		Complete(r)
}
