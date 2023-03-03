package kube

import (
	"context"
	"fmt"
	"testing"
)

func TestListPodsOnNode(t *testing.T) {
	cs, err := CreateClientsetFromLocal()
	if err != nil {
		t.Error(err)
		return
	}
	kubeHelper := NewKubeHelper(cs)
	ctx := context.TODO()
	nodes, err1 := kubeHelper.ListNodesWithLabel(ctx, "")
	if err1 != nil {
		t.Error(err1)
		return
	}
	fmt.Println("Number of node:", len(nodes))
	fmt.Println("Pods on the node:", nodes[len(nodes)-1].Name)
	pods, err2 := kubeHelper.ListPodsOnNode(ctx, nodes[len(nodes)-1].Name)
	if err2 != nil {
		t.Error(err2)
		return
	}
	for _, pod := range pods {
		fmt.Println(pod.Name)
	}
}

func TestListOD_Node(t *testing.T) {
	cs, err := CreateClientsetFromLocal()
	if err != nil {
		t.Error(err)
		return
	}
	kubeHelper := NewKubeHelper(cs)
	ctx := context.TODO()
	nodes, err1 := kubeHelper.ListNodesWithLabel(ctx, "cloud.google.com/gke-spot=true")
	if err1 != nil {
		t.Error(err1)
		return
	}
	fmt.Println("Number of node:", len(nodes))
}
