| Metric name | Metric Type | Description | Unit (where applicable) | Labels/tags | Status |
| ----------- | ----------- | ----------- | ----------------------- | ----------- | ------ |
|kube_resourcequota|Gauge|Information about resource quota.||`resourcequota`=&lt;quota-name&gt; <br> `namespace`=&lt;namespace&gt; <br> `resource`=&lt;ResourceName&gt; <br> `type`=&lt;quota-type&gt;|STABLE|
|kube_resourcequota_created|Gauge|Unix creation timestamp|seconds|`resourcequota`=&lt;quota-name&gt; <br> `namespace`=&lt;namespace&gt;|STABLE|