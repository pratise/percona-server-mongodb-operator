package exporter

import (
	"fmt"
	api "github.com/percona/percona-server-mongodb-operator/pkg/apis/psmdb/v1"
	"github.com/percona/percona-server-mongodb-operator/pkg/psmdb"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
)

const (
	agentContainerName           = "exporter-agent"
	exporterUser                 = "MONGODB_CLUSTER_MONITOR_USER"
	exporterPassword             = "MONGODB_CLUSTER_MONITOR_PASSWORD"
	agentHealthStatusFilePathEnv = "AGENT_STATUS_FILEPATH"
	awsAccessKeySecretKey        = "AWS_ACCESS_KEY_ID"
	awsSecretAccessKeySecretKey  = "AWS_SECRET_ACCESS_KEY"
)

// AgentContainer creates the container object for a backup agent
func AgentContainer(cr *api.PerconaServerMongoDB, exporterSec corev1.Secret) (corev1.Container, error) {
	res, err := psmdb.CreateResources(cr.Spec.Exporter.Resources)
	if err != nil {
		return corev1.Container{}, errors.Wrap(err, "create resources")
	}
	user := exporterSec.Data[exporterUser]
	password := exporterSec.Data[exporterPassword]

	c := corev1.Container{
		Name:            agentContainerName,
		Image:           cr.Spec.Exporter.Image,
		ImagePullPolicy: corev1.PullAlways,
		Command: []string{
			"/bin/sh",
			"-c",
			fmt.Sprintf("/mongodb_exporter --mongodb.uri=mongodb://%s:%s@127.0.0.1/admin", user, password),
		},
		Env: []corev1.EnvVar{
			{
				Name:  agentHealthStatusFilePathEnv,
				Value: "/healthstatus/agent-health-status.json",
			},
		},
		Ports: []corev1.ContainerPort{
			{

				ContainerPort: 9216,
				Protocol:      corev1.ProtocolTCP,
			},
		},
		TerminationMessagePath:   "/dev/termination-log",
		TerminationMessagePolicy: corev1.TerminationMessageReadFile,
		SecurityContext:          cr.Spec.Backup.ContainerSecurityContext,
		Resources:                res,
	}

	return c, nil
}
