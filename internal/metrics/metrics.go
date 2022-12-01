// Package metrics author: weiqiang; date: 2022-12
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "doumeng"
)

type Exporter struct {
	NodeState *prometheus.GaugeVec
	UserCount *prometheus.GaugeVec
}

func NewExporter() *Exporter {
	return &Exporter{
		NodeState: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "node_state",
			Help:      "节点服务状态，0 关闭状态，1 服务状态",
		}, []string{"node_ip", "node_host", "node_area_code", "node_area_name"}),
		UserCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "user_count",
			Help:      "用户数量统计",
		}, []string{}),
	}
}
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	e.NodeState.Describe(ch)
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	e.NodeState.Reset()
	e.UserCount.Reset()

	// 定义值
	e.NodeState.WithLabelValues("node.NodeIp", "node.NodeHost", "node.NodeAreaCode", "node.NodeAreaName").Set(0)
	e.UserCount.WithLabelValues().Set(0)

	e.NodeState.Collect(ch)
	e.UserCount.Collect(ch)
}
