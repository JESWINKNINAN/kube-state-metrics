| Metric name | Metric Type | Description | Unit (where applicable) | Labels/tags | Status |
| ----------- | ----------- | ----------- | ----------------------- | ----------- | ------ |
|kube_lease_owner|Gauge|Information about the Lease's owner||`lease`=&lt;lease-name&gt; <br> `owner_kind`=&lt;onwer kind&gt; <br> `owner_name`=&lt;owner name&gt;|EXPERIMENTAL|
|kube_lease_renew_time |Gauge|Kube lease renew time.|seconds|`lease`=&lt;lease-name&gt;|EXPERIMENTAL|