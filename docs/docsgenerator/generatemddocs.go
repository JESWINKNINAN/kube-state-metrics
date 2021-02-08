package main

import "k8s.io/kube-state-metrics/v2/pkg/docscollector"

func main() {

	docscollector.DocsCreate("lease-metrics")
	docscollector.DocsCreate("resourcequota-metrics")

}
