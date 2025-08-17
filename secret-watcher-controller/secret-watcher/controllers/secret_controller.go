package controllers

import (
    "context"
    "crypto/sha256"
    "encoding/hex"
    "fmt"

    corev1 "k8s.io/api/core/v1"
    "k8s.io/apimachinery/pkg/labels"
    "k8s.io/apimachinery/pkg/runtime"
    "sigs.k8s.io/controller-runtime/pkg/client"
    ctrl "sigs.k8s.io/controller-runtime"
    "sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var Scheme = runtime.NewScheme()

func init() {
    _ = corev1.AddToScheme(Scheme)
}

type SecretReconciler struct {
    client.Client
    TargetSecretName       string
    TargetSecretNamespace  string
    TargetPodLabelSelector map[string]string
    previousHash           string
}

func (r *SecretReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
    if req.Name != r.TargetSecretName || req.Namespace != r.TargetSecretNamespace {
        return reconcile.Result{}, nil
    }

    var secret corev1.Secret
    if err := r.Get(ctx, req.NamespacedName, &secret); err != nil {
        return reconcile.Result{}, client.IgnoreNotFound(err)
    }

    cert := secret.Data["tls.crt"]
    if cert == nil {
        return reconcile.Result{}, nil
    }

    hash := sha256.Sum256(cert)
    hashStr := hex.EncodeToString(hash[:])

    if hashStr == r.previousHash {
        return reconcile.Result{}, nil
    }
    r.previousHash = hashStr

    var pods corev1.PodList
    labelSelector := labels.SelectorFromSet(r.TargetPodLabelSelector)
    if err := r.List(ctx, &pods,
        client.InNamespace(r.TargetSecretNamespace),
        client.MatchingLabelsSelector{Selector: labelSelector}); err != nil {
        return reconcile.Result{}, err
    }

    for _, pod := range pods.Items {
        fmt.Printf("Deleting pod %s/%s due to cert change\n", pod.Namespace, pod.Name)
        if err := r.Delete(ctx, &pod); err != nil {
            return reconcile.Result{}, err
        }
    }

    return reconcile.Result{}, nil
}

func (r *SecretReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).
        For(&corev1.Secret{}).
        Complete(r)
}
