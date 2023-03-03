package kube

import (
	"context"
	"os"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func GetDefaultKubeConfigPath() string {
	return filepath.Join(homeDir(), ".kube", "config")
}

func getRestConfig(kubecfgPath string) (*restclient.Config, error) {
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubecfgPath)
	if err != nil {
		return nil, err
	}
	return config, err
}

func CreateClientsetFromLocal() (*kubernetes.Clientset, error) {
	// use the current context in kubeconfig
	config, err := getRestConfig(GetDefaultKubeConfigPath())
	if err != nil {
		return nil, err
	}
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	return clientset, nil
}

// CreateClientsetFromPod is used when your application is running in K8S
func CreateClientsetFromPod() (*kubernetes.Clientset, error) {
	config, err := restclient.InClusterConfig()
	if err != nil {
		return nil, err
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	return clientset, nil
}

type KubeHelper struct {
	clientset *kubernetes.Clientset
}

func NewKubeHelper(clientset *kubernetes.Clientset) *KubeHelper {
	return &KubeHelper{clientset}
}

func (h *KubeHelper) ListPodsOnNode(ctx context.Context, nodeName string) ([]corev1.Pod, error) {
	pods, err := h.clientset.CoreV1().Pods("").List(ctx, metav1.ListOptions{
		FieldSelector: "spec.nodeName=" + nodeName,
	})
	if err != nil {
		return nil, err
	}
	return pods.Items, nil
}

func (h *KubeHelper) ListNodesWithLabel(ctx context.Context, labelSelector string) ([]corev1.Node, error) {
	nodes, err := h.clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{LabelSelector: labelSelector})
	if err != nil {
		return nil, err
	}
	return nodes.Items, nil
}
