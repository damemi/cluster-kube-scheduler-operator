module github.com/openshift/cluster-kube-scheduler-operator

go 1.13

require (
	github.com/blang/semver v3.5.0+incompatible
	github.com/ghodss/yaml v1.0.0
	github.com/go-bindata/go-bindata v3.1.2+incompatible
	github.com/openshift/api v0.0.0-20200609191024-dca637550e8c
	github.com/openshift/build-machinery-go v0.0.0-20200512074546-3744767c4131
	github.com/openshift/client-go v0.0.0-20200521150516-05eb9880269c
	github.com/openshift/library-go v0.0.0-20200527213645-a9b77f5402e3
	github.com/prometheus/client_golang v1.1.0
	github.com/prometheus/common v0.6.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	k8s.io/api v0.18.3
	k8s.io/apimachinery v0.18.3
	k8s.io/apiserver v0.18.3
	k8s.io/client-go v0.18.3
	k8s.io/component-base v0.18.3
	k8s.io/klog v1.0.0
)

replace github.com/jteeuwen/go-bindata => github.com/jteeuwen/go-bindata v3.0.8-0.20151023091102-a0ff2567cfb7+incompatible

replace github.com/openshift/api => github.com/damemi/api v0.0.0-20200629190638-f973ae2834fe
