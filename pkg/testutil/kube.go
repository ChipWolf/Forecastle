package testutil

import (
	routev1 "github.com/openshift/api/route/v1"
	v1alpha1 "github.com/stakater/Forecastle/pkg/apis/forecastle/v1alpha1"
	"k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func CreateIngress(name string) *v1beta1.Ingress {
	return &v1beta1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
}

func CreateRoute(name string) *routev1.Route {
	return &routev1.Route{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
}

func CreateIngressWithNamespace(name string, namespace string) *v1beta1.Ingress {
	ingress := CreateIngress(name)
	ingress.ObjectMeta.Namespace = namespace

	return ingress
}

func CreateRouteWithHost(name string, url string) *routev1.Route {
	route := CreateRoute(name)
	route.Spec.Host = url

	return route
}

func CreateIngressWithHost(name string, url string) *v1beta1.Ingress {
	ingress := CreateIngress(name)
	ingress.Spec.Rules = []v1beta1.IngressRule{
		{
			Host: url,
		},
	}

	return ingress
}

func AddAnnotationToIngress(ingress *v1beta1.Ingress, annotationKey string, annotationValue string) *v1beta1.Ingress {
	if ingress.Annotations == nil {
		ingress.Annotations = make(map[string]string)
	}

	ingress.Annotations[annotationKey] = annotationValue

	return ingress
}

func CreateIngressWithHostAndSubPath(name string, url string, subpath string, port string) *v1beta1.Ingress {
	ingress := CreateIngressWithHost(name, url)
	ingress.Spec.Rules[0].HTTP = &v1beta1.HTTPIngressRuleValue{
		Paths: []v1beta1.HTTPIngressPath{
			{
				Backend: v1beta1.IngressBackend{
					ServicePort: intstr.FromString(port),
				},
				Path: subpath,
			},
		},
	}

	return ingress
}

func CreateIngressWithTLSHost(name string, tlsurl string) *v1beta1.Ingress {
	ingress := CreateIngress(name)
	ingress.Spec.TLS = []v1beta1.IngressTLS{
		{
			Hosts: []string{
				tlsurl,
			},
		},
	}

	return ingress
}

func CreateIngressWithHostAndTLSHost(name string, host string, tlsurl string) *v1beta1.Ingress {
	ingress := CreateIngressWithHost(name, host)
	ingress.Spec.TLS = []v1beta1.IngressTLS{
		{
			Hosts: []string{
				tlsurl,
			},
		},
	}

	return ingress
}

func CreateForecastleApp(name string, url string, group string, icon string) *v1alpha1.ForecastleApp {
	return &v1alpha1.ForecastleApp{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: v1alpha1.ForecastleAppSpec{
			Name:  name,
			URL:   url,
			Icon:  icon,
			Group: group,
		},
	}
}

func CreateForecastleAppWithURLFromRoute(name string, group string, icon string, routeName string) *v1alpha1.ForecastleApp {
	forecastleApp := CreateForecastleApp(name, "", group, icon)
	forecastleApp.Spec.URLFrom = &v1alpha1.URLSource{
		RouteRef: &v1alpha1.RouteURLSource{
			LocalObjectReference: v1alpha1.LocalObjectReference{
				Name: routeName,
			},
		},
	}

	return forecastleApp
}

func CreateForecastleAppWithURLFromIngress(name string, group string, icon string, ingressName string) *v1alpha1.ForecastleApp {
	forecastleApp := CreateForecastleApp(name, "", group, icon)
	forecastleApp.Spec.URLFrom = &v1alpha1.URLSource{
		IngressRef: &v1alpha1.IngressURLSource{
			LocalObjectReference: v1alpha1.LocalObjectReference{
				Name: ingressName,
			},
		},
	}

	return forecastleApp
}
