package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := os.Getenv("KUBECONFIG")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// 创建共享 informer factory，指定 resync 周期为 10 分钟
	factory := informers.NewSharedInformerFactory(clientset, 10*time.Minute)
	cmInformer := factory.Core().V1().ConfigMaps().Informer()

	// 注册事件处理回调函数
	cmInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			cm := obj.(*corev1.ConfigMap)
			fmt.Printf("[ADD] %s/%s\n", cm.Namespace, cm.Name)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oldCM := oldObj.(*corev1.ConfigMap)
			newCM := newObj.(*corev1.ConfigMap)
			fmt.Printf("[UPDATE] %s/%s: %d -> %d \n",
				newCM.Namespace, newCM.Name,
				len(oldCM.Data), len(oldCM.Data))
		},
		DeleteFunc: func(obj interface{}) {
			cm := obj.(*corev1.ConfigMap)
			fmt.Printf("[DELETE] %s/%s\n", cm.Namespace, cm.Name)
		},
	})

	stopCh := make(chan struct{})
	defer close(stopCh)
	// 捕捉 panic 并优雅退出
	defer runtime.HandleCrash()

	// 启动所有 informers（包括 podsInformer）
	factory.Start(stopCh)
	// 等待缓存同步完成后再处理事件
	factory.WaitForCacheSync(stopCh)

	fmt.Println("Informer started. Listening for Pod events...")

	// 等待 SIGINT/SIGTERM 信号优雅退出
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	<-sigCh
	fmt.Println("Shutting down...")
}
