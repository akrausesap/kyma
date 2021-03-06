// Code generated by informer-gen. DO NOT EDIT.

package v1alpha2

import (
	internalinterfaces "github.com/kyma-project/kyma/components/application-registry/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Handlers returns a HandlerInformer.
	Handlers() HandlerInformer
	// Instances returns a InstanceInformer.
	Instances() InstanceInformer
	// Rules returns a RuleInformer.
	Rules() RuleInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Handlers returns a HandlerInformer.
func (v *version) Handlers() HandlerInformer {
	return &handlerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Instances returns a InstanceInformer.
func (v *version) Instances() InstanceInformer {
	return &instanceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Rules returns a RuleInformer.
func (v *version) Rules() RuleInformer {
	return &ruleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
