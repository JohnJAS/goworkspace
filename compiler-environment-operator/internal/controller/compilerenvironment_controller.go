/*
Copyright 2025.

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
	"time"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	compilerv1 "example.com/compiler-environment-operator/api/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CompilerEnvironmentReconciler reconciles a CompilerEnvironment object
type CompilerEnvironmentReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=compiler.example.com,resources=compilerenvironments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=compiler.example.com,resources=compilerenvironments/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=compiler.example.com,resources=compilerenvironments/finalizers,verbs=update
// +kubebuilder:rbac:groups="",resources=pods,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups="",resources=persistentvolumeclaims,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the CompilerEnvironment object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.21.0/pkg/reconcile
func (r *CompilerEnvironmentReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logf.FromContext(ctx, "compilerenvironment", req.NamespacedName)
	// 1. 获取 CompilerEnvironment 实例
	var ce compilerv1.CompilerEnvironment
	if err := r.Get(ctx, req.NamespacedName, &ce); err != nil {
		// 如果资源不存在，则忽略（可能是被删除了）
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// 2. 检查资源是否正在被删除
	if !ce.ObjectMeta.DeletionTimestamp.IsZero() {
		// 资源正在被删除，执行清理逻辑（如果有Finalizer）
		return ctrl.Result{}, nil
	}

	// 3. 处理 PVC
	pvcName := getPVCName(&ce)
	var pvc corev1.PersistentVolumeClaim
	if err := r.Get(ctx, client.ObjectKey{Namespace: req.Namespace, Name: pvcName}, &pvc); err != nil {
		// PVC 不存在，创建它
		newPvc := r.constructPVCForCompilerEnvironment(&ce, pvcName)
		if err := r.Create(ctx, newPvc); err != nil {
			log.Error(err, "Failed to create PVC")
			return ctrl.Result{}, err
		}
		log.Info("PVC created successfully", "name", pvcName)
		// 等待下一次 Reconcile，等待 PVC 绑定
		return ctrl.Result{Requeue: true}, nil
	}

	// 4. 处理 Pod
	podName := getPodName(&ce)
	var pod corev1.Pod
	if err := r.Get(ctx, client.ObjectKey{Namespace: req.Namespace, Name: podName}, &pod); err != nil {
		// Pod 不存在，创建它
		newPod := r.constructPodForCompilerEnvironment(&ce, podName, pvcName)
		// 设置 ControllerReference，让 Pod 被 CE 资源管理
		if err := ctrl.SetControllerReference(&ce, newPod, r.Scheme); err != nil {
			return ctrl.Result{}, err
		}
		if err := r.Create(ctx, newPod); err != nil {
			log.Error(err, "Failed to create Pod")
			return ctrl.Result{}, err
		}
		log.Info("Pod created successfully", "name", podName)
		// 更新状态
		ce.Status.PodName = podName
		ce.Status.PVCName = pvcName
		ce.Status.Phase = "Creating"
		if err := r.Status().Update(ctx, &ce); err != nil {
			log.Error(err, "Failed to update status")
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, nil
	}

	// 5. 更新状态：根据 Pod 的实际状态更新 CE 的状态
	currentPhase := string(pod.Status.Phase)
	if ce.Status.Phase != currentPhase {
		ce.Status.Phase = currentPhase
		if err := r.Status().Update(ctx, &ce); err != nil {
			log.Error(err, "Failed to update status phase")
			return ctrl.Result{}, err
		}
	}

	// 6. 如果 Pod 运行完成（Succeeded/Failed），可以选择在一段时间后自动清理整个 CE 资源
	if pod.Status.Phase == corev1.PodSucceeded || pod.Status.Phase == corev1.PodFailed {
		// 例如：5分钟后自动清理
		deletionTimestamp := ce.ObjectMeta.CreationTimestamp.Add(5 * time.Minute)
		if time.Now().After(deletionTimestamp) {
			log.Info("Pod has completed, deleting CompilerEnvironment resource")
			if err := r.Delete(ctx, &ce); err != nil {
				log.Error(err, "Failed to delete CompilerEnvironment")
				return ctrl.Result{}, err
			}
			// Kubernetes 的垃圾收集器会因为 OwnerReference 自动删除 Pod 和 PVC
			return ctrl.Result{}, nil
		}
		// 还没到清理时间，等一会儿再检查
		return ctrl.Result{RequeueAfter: time.Until(deletionTimestamp)}, nil
	}

	// Pod 还在运行，无需额外操作
	return ctrl.Result{}, nil
}

// Helper function to construct PVC
func (r *CompilerEnvironmentReconciler) constructPVCForCompilerEnvironment(ce *compilerv1.CompilerEnvironment, name string) *corev1.PersistentVolumeClaim {
	storageClassName := "standard" // 使用你的集群中存在的 StorageClass 名称
	return &corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: ce.Namespace,
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes: []corev1.PersistentVolumeAccessMode{corev1.ReadWriteOnce},
			Resources: corev1.VolumeResourceRequirements{
				Requests: corev1.ResourceList{
					corev1.ResourceStorage: resource.MustParse(ce.Spec.StorageRequest),
				},
			},
			StorageClassName: &storageClassName,
		},
	}
}

// Helper function to construct Pod
func (r *CompilerEnvironmentReconciler) constructPodForCompilerEnvironment(ce *compilerv1.CompilerEnvironment, podName, pvcName string) *corev1.Pod {
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podName,
			Namespace: ce.Namespace,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:      "compiler",
					Image:     ce.Spec.Image,
					Command:   ce.Spec.Command,
					Args:      ce.Spec.Args,
					Resources: ce.Spec.Resources,
					Env:       ce.Spec.Env,
					VolumeMounts: []corev1.VolumeMount{
						{
							Name:      "workspace",
							MountPath: "/workspace", // 容器内代码挂载路径
						},
					},
				},
			},
			RestartPolicy: corev1.RestartPolicyNever, // 任务型，不重启
			Volumes: []corev1.Volume{
				{
					Name: "workspace",
					VolumeSource: corev1.VolumeSource{
						PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
							ClaimName: pvcName,
						},
					},
				},
			},
		},
	}
}

// Helper functions to generate names
func getPVCName(ce *compilerv1.CompilerEnvironment) string {
	return ce.Name + "-pvc"
}
func getPodName(ce *compilerv1.CompilerEnvironment) string {
	return ce.Name + "-pod"
}

// SetupWithManager sets up the controller with the Manager.
func (r *CompilerEnvironmentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&compilerv1.CompilerEnvironment{}).
		Owns(&corev1.Pod{}).                   // 监视管理的 Pod 变化并触发 Reconcile
		Owns(&corev1.PersistentVolumeClaim{}). // 监视管理的 PVC 变化
		Complete(r)
}
