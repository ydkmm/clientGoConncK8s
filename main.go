package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"time"

	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var kubeconfig *string
	// 使用 home 路径
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "testConfig"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	// 使用绝对路径
	//kubeconfig = flag.String("kubeconfig", "C:\\Users\\23042\\.kube\\testConfig", "absolute path to the kubeconfig file")

	// 使用项目的工作路径，打包镜像可以使用
	//currentDir, err := os.Getwd()
	//kubeconfig = flag.String("kubeconfig", filepath.Join(currentDir, "/testConfig"), "absolute path to the kubeconfig file")
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
		for _, pod := range pods.Items {
			fmt.Println(pod.Name)
		}
		time.Sleep(10 * time.Second)
	}
}
