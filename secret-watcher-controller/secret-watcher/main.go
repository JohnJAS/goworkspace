package main

import (
	"flag"
	"os"

	"secretwatcher/controllers"

	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
)

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "enable-leader-election", false, "Enable leader election for controller manager.")
	flag.Parse()

	ctrl.SetLogger(zap.New(zap.UseDevMode(true)))

	cfg, err := rest.InClusterConfig()
	if err != nil {
		os.Exit(1)
	}

	scheme := controllers.Scheme

	metrics := metricsserver.Options{
		BindAddress: metricsAddr,
	}

	mgr, err := ctrl.NewManager(cfg, ctrl.Options{
		Scheme:           scheme,
		Metrics:          metrics,
		LeaderElection:   enableLeaderElection,
		LeaderElectionID: "secret-watcher-controller",
	})
	if err != nil {
		os.Exit(1)
	}

	r := &controllers.SecretReconciler{
		Client:                 mgr.GetClient(),
		TargetSecretName:       "my-tls-secret",
		TargetSecretNamespace:  "default",
		TargetPodLabelSelector: map[string]string{"app": "myapp"},
	}

	if err = r.SetupWithManager(mgr); err != nil {
		os.Exit(1)
	}

	if err = mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		os.Exit(1)
	}
}
