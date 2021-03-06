package informers

import (
	"log"

	broker "github.com/layer5io/meshery-operator/pkg/broker"
	informers "github.com/layer5io/meshery-operator/pkg/informers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
)

type Cluster struct {
	client *informers.Client
	broker broker.Broker
}

func New(client *informers.Client, broker broker.Broker) *Cluster {
	return &Cluster{
		client: client,
		broker: broker,
	}
}

// common resource event handler
// will get the object and log that
// and it will publish the object
func (c *Cluster) resourceEventHandlerFuncs(resourceType string) cache.ResourceEventHandlerFuncs {

	return cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			c.broker.Publish(Subject, broker.Message{
				Type:   resourceType,
				Object: obj,
			})
			object := obj.(metav1.Object)
			log.Printf("%s Named: %s - added", resourceType, object.GetName())
		},
		UpdateFunc: func(new interface{}, old interface{}) {
			c.broker.Publish(Subject, broker.Message{
				Type:   resourceType,
				Object: new,
			})
			object := new.(metav1.Object)
			log.Printf("%s Named: %s - updated", resourceType, object.GetName())
		},
		DeleteFunc: func(obj interface{}) {
			c.broker.Publish(Subject, broker.Message{
				Type:   resourceType,
				Object: obj,
			})
			object := obj.(metav1.Object)
			log.Printf("%s Named: %s - deleted", resourceType, object.GetName())
		},
	}

}

// NodeInformer - for Nodes
func (c *Cluster) NodeInformer() cache.SharedIndexInformer {
	// get informer
	NodeInformer := c.client.GetNodeInformer().Informer()

	// register event handlers
	NodeInformer.AddEventHandler(
		c.resourceEventHandlerFuncs("Node"),
	)

	return NodeInformer
}

// NamespaceInformer - for Namespaces
func (c *Cluster) NamespaceInformer() cache.SharedIndexInformer {
	// get informer
	NamespaceInformer := c.client.GetNamespaceInformer().Informer()

	// register event handlers
	NamespaceInformer.AddEventHandler(
		c.resourceEventHandlerFuncs("Namespace"),
	)

	return NamespaceInformer
}

// DeploymentInformer - for Deployments
func (c *Cluster) DeploymentInformer() cache.SharedIndexInformer {
	// get informer
	DeploymentInformer := c.client.GetDeploymentInformer().Informer()

	// register event handlers
	DeploymentInformer.AddEventHandler(
		c.resourceEventHandlerFuncs("Deployment"),
	)

	return DeploymentInformer
}

// PodInformer - for Pods
func (c *Cluster) PodInformer() cache.SharedIndexInformer {
	// get informer
	PodInformer := c.client.GetPodInformer().Informer()

	// register event handlers
	PodInformer.AddEventHandler(
		c.resourceEventHandlerFuncs("Pod"),
	)

	return PodInformer
}
