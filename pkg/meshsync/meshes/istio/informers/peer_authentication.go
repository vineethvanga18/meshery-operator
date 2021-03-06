package informers

import (
	"log"

	broker "github.com/layer5io/meshery-operator/pkg/broker"
	v1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"
	"k8s.io/client-go/tools/cache"
)

func (i *Istio) PeerAuthenticationInformer() cache.SharedIndexInformer {
	// get informer
	PeerAuthenticationInformer := i.client.GetPeerAuthenticationInformer().Informer()

	// register event handlers
	PeerAuthenticationInformer.AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				PeerAuthentication := obj.(*v1beta1.PeerAuthentication)
				log.Printf("PeerAuthentication Named: %s - added", PeerAuthentication.Name)
				i.broker.Publish(Subject, broker.Message{
					Type:   "PeerAuthentication",
					Object: PeerAuthentication,
				})
			},
			UpdateFunc: func(new interface{}, old interface{}) {
				PeerAuthentication := new.(*v1beta1.PeerAuthentication)
				log.Printf("PeerAuthentication Named: %s - updated", PeerAuthentication.Name)
				i.broker.Publish(Subject, broker.Message{
					Type:   "PeerAuthentication",
					Object: PeerAuthentication,
				})
			},
			DeleteFunc: func(obj interface{}) {
				PeerAuthentication := obj.(*v1beta1.PeerAuthentication)
				log.Printf("PeerAuthentication Named: %s - deleted", PeerAuthentication.Name)
				i.broker.Publish(Subject, broker.Message{
					Type:   "PeerAuthentication",
					Object: PeerAuthentication,
				})
			},
		},
	)

	return PeerAuthenticationInformer
}
