package main

import (
	"fmt"
	"os"

	"k8s.io/client-go/tools/clientcmd"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/ONE7live/kompass/pkg/apis/v1alpha1"
	"github.com/ONE7live/kompass/pkg/controller"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "C:\\Users\\ONE7\\.kube\\config-52")
	if err != nil {
		panic(err)
	}

	mgr, err := ctrl.NewManager(config, ctrl.Options{})
	if err != nil {
		fmt.Errorf("unable to start manager, err: %v", err)
		os.Exit(1)
	}

	if err = v1alpha1.AddToScheme(mgr.GetScheme()); err != nil {
		fmt.Errorf("unable to add scheme: %v", err)
		os.Exit(1)
	}

	if err = (&controller.LeafPodConvertReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		fmt.Errorf("unable to create controller LeafPodConvert, err: %v", err)
		os.Exit(1)
	}

	if err = mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		fmt.Errorf("problem running manager, err: %v", err)
		os.Exit(1)
	}
}
