package cluster

type Cluster struct {
	tableName              struct{}          `sql:"cluster" pg:",discard_unknown_columns"`
	Id                     int               `sql:"id,pk"`
	ClusterName            string            `sql:"cluster_name"`
	Description            string            `sql:"description"`
	ServerUrl              string            `sql:"server_url"`
	PrometheusEndpoint     string            `sql:"prometheus_endpoint"`
	Active                 bool              `sql:"active,notnull"`
	Config                 map[string]string `sql:"config"`
}
