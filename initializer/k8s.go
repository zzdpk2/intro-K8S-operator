package initializer

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"kubeproj.com/global"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

func KubernetesClient() {
	// var kubeconfig *string
	// if home := homedir.HomeDir(); home != "" {
	// 	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	// } else {
	// 	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	// }
	// flag.Parse()
	kubeconfig := "./.kube/config"

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	global.KubeConfigSet = clientset
	// for {
	// 	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	//
	// 	// Examples for error handling:
	// 	// - Use helper functions like e.g. errors.IsNotFound()
	// 	// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
	// 	namespace := "default"
	// 	pod := "example-xxxxx"
	// 	_, err = clientset.CoreV1().Pods(namespace).Get(context.TODO(), pod, metav1.GetOptions{})
	// 	if errors.IsNotFound(err) {
	// 		fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
	// 	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
	// 		fmt.Printf("Error getting pod %s in namespace %s: %v\n",
	// 			pod, namespace, statusError.ErrStatus.Message)
	// 	} else if err != nil {
	// 		panic(err.Error())
	// 	} else {
	// 		fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
	// 	}
	//
	// 	time.Sleep(10 * time.Second)
}
