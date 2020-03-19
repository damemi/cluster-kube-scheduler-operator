package configmetrics

import (
	"github.com/blang/semver"
	"github.com/prometheus/client_golang/prometheus"
	k8smetrics "k8s.io/component-base/metrics"
	"k8s.io/component-base/metrics/legacyregistry"

	configinformers "github.com/openshift/client-go/config/informers/externalversions"
	configlisters "github.com/openshift/client-go/config/listers/config/v1"
)

func Register(configInformer configinformers.SharedInformerFactory) {
	legacyregistry.MustRegister(&configMetrics{
		configLister: configInformer.Config().V1().Schedulers().Lister(),
		config: k8smetrics.NewGauge(&k8smetrics.GaugeOpts{
			Name: "cluster_master_schedulable",
			Help: "Reports whether the cluster master nodes are schedulable.",
		}),
	})
}

// configMetrics implements metrics gathering for this component.
type configMetrics struct {
	configLister configlisters.SchedulerLister
	config       *k8smetrics.Gauge
}

func (m *configMetrics) ClearState() {}

func (m *configMetrics) Create(version *semver.Version) bool {
	return true
}

// Describe reports the metadata for metrics to the prometheus collector.
func (m *configMetrics) Describe(ch chan<- *prometheus.Desc) {
	m.config.Describe(ch)
}

// Collect calculates metrics from the cached config and reports them to the prometheus collector.
func (m *configMetrics) Collect(ch chan<- prometheus.Metric) {
	if config, err := m.configLister.Get("cluster"); err == nil {
		g := m.config
		if config.Spec.MastersSchedulable {
			g.Set(1)
		} else {
			g.Set(0)
		}
		g.Collect(ch)
	}
}

func (m *configMetrics) FQName() string {
	return m.FQName()
}
