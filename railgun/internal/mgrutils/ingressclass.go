package mgrutils

import "github.com/kong/kubernetes-ingress-controller/pkg/annotations"

// IngressClass is the current "kubernetes.io/ingress.class" configured for controllers
var IngressClass = annotations.DefaultIngressClass
