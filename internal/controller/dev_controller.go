/*
Copyright 2024 Gokul.

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
	"fmt"

	apiv1alpha1 "dev/api/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

//var logger = log.Log.WithName("controller_dev")

// DevReconciler reconciles a Dev object
type DevReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

type controller struct {
	client kubernetes.Interface
}

//+kubebuilder:rbac:groups=api.gokula.dev,resources=devs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=api.gokula.dev,resources=devs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=api.gokula.dev,resources=devs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Dev object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.3/pkg/reconcile
func (r *DevReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	//log := logger.WithValues("Request.Namespace", req.Namespace, "Request.Name", req.Name)
	logger := log.FromContext(ctx)

	logger.Info("Reconcile called by gokul")

	fmt.Println("Reconcile is called for testing")
	logger.Info(fmt.Sprintf("Pod created is %v", req.NamespacedName))

	//dev := &apiv1alpha1.Dev{}

	//err := r.Get(ctx, req.NamespacedName, dev)

	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})

	config, err := kubeconfig.ClientConfig()
	if err != nil {
		panic(err)
	}
	clientset := kubernetes.NewForConfigOrDie(config)
	nodeList, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, n := range nodeList.Items {
		fmt.Println(n.Name)
	}

	// var kubeconfig *string
	// if home := homedir.HomeDir(); home != "" {
	// 	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	// } else {
	// 	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	// }
	// flag.Parse()

	// // use the current context in kubeconfig
	// config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	// if err != nil {
	// 	_ = fmt.Errorf("failed to config %s", err.Error())
	// 	config, err = rest.InClusterConfig()
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// }
	// // create the clientset
	// clientset, err := kubernetes.NewForConfig(config)

	// stopper := make(chan struct{})
	// defer close(stopper)
	// var ctr controller
	// //config, err := clientcmd.BuildConfigFromFlags("", ctr)

	// kc.err := utils.GetClient(config)
	// if err != nil {

	// }
	// factory := informers.NewSharedInformerFactory(ctr.client, 0)

	// informer := factory.Core().V1().Pods().Informer()

	// informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
	// 	AddFunc: func(obj interface{}) {
	// 		fmt.Println("add event")
	// 	},
	// 	UpdateFunc: func(obj1, obj2 interface{}) {
	// 		fmt.Println("update event")
	// 	},
	// 	DeleteFunc: func(obj interface{}) {
	// 		fmt.Println("delete event")
	// 	},
	// })

	// go informer.Run(stopper)
	// <-stopper

	return ctrl.Result{}, err
}

// SetupWithManager sets up the controller with the Manager.
func (r *DevReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apiv1alpha1.Dev{}).
		Complete(r)
}
