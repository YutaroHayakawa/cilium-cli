// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package defaults

import "time"

const (
	AgentContainerName      = "cilium-agent"
	AgentServiceAccountName = "cilium"
	AgentClusterRoleName    = "cilium"
	AgentSecretsRoleName    = "cilium-secrets"
	AgentDaemonSetName      = "cilium"
	AgentPodSelector        = "k8s-app=cilium"
	AgentResourceQuota      = "cilium-resource-quota"
	AgentImage              = "quay.io/cilium/cilium"

	CASecretName     = "cilium-ca"
	CASecretKeyName  = "ca.key"
	CASecretCertName = "ca.crt"

	EncryptionSecretName = "cilium-ipsec-keys"
	AKSSecretName        = "cilium-azure"

	NodeInitDaemonSetName = "cilium-node-init"

	OperatorServiceAccountName = "cilium-operator"
	OperatorClusterRoleName    = "cilium-operator"
	OperatorSecretsRoleName    = "cilium-operator-secrets"
	OperatorContainerName      = "cilium-operator"
	OperatorDeploymentName     = "cilium-operator"
	OperatorResourceQuota      = "cilium-operator-resource-quota"
	OperatorImage              = "quay.io/cilium/operator-generic"
	OperatorImageAWS           = "quay.io/cilium/operator-aws"
	OperatorImageAzure         = "quay.io/cilium/operator-azure"

	HubbleServerSecretName = "hubble-server-certs"

	RelayContainerName       = "hubble-relay"
	RelayDeploymentName      = "hubble-relay"
	RelayClusterRoleName     = "hubble-relay"
	RelayServiceAccountName  = "hubble-relay"
	RelayConfigMapName       = "hubble-relay-config"
	RelayImage               = "quay.io/cilium/hubble-relay"
	RelayServerSecretName    = "hubble-relay-server-certs"
	RelayClientSecretName    = "hubble-relay-client-certs"
	HubbleUIClientSecretName = "hubble-ui-client-certs"

	HubbleUIClusterRoleName    = "hubble-ui"
	HubbleUIServiceAccountName = "hubble-ui"
	HubbleUIDeploymentName     = "hubble-ui"
	HubbleUIImage              = "quay.io/cilium/hubble-ui"
	HubbleUIBackendImage       = "quay.io/cilium/hubble-ui-backend"

	ClusterMeshDeploymentName             = "clustermesh-apiserver"
	ClusterMeshContainerName              = "apiserver"
	ClusterMeshServiceAccountName         = "clustermesh-apiserver"
	ClusterMeshClusterRoleName            = "clustermesh-apiserver"
	ClusterMeshApiserverImage             = "quay.io/cilium/clustermesh-apiserver"
	ClusterMeshServiceName                = "clustermesh-apiserver"
	ClusterMeshSecretName                 = "cilium-clustermesh" // Secret which contains the clustermesh configuration
	ClusterMeshServerSecretName           = "clustermesh-apiserver-server-cert"
	ClusterMeshAdminSecretName            = "clustermesh-apiserver-admin-cert"
	ClusterMeshClientSecretName           = "clustermesh-apiserver-client-cert"
	ClusterMeshExternalWorkloadSecretName = "clustermesh-apiserver-external-workload-cert"

	ConnectivityCheckNamespace = "cilium-test"

	ConnectivityCheckAlpineCurlImage = "quay.io/cilium/alpine-curl:v1.5.0@sha256:7b286939730d8af1149ef88dba15739d8330bb83d7d9853a23e5ab4043e2d33c"
	ConnectivityPerformanceImage     = "quay.io/cilium/network-perf:bf58fb8bc57c4933dfa6e2a9581d3925c0a0571e@sha256:9bef508b2dcaeb3e288a496b8d3f065e8636a4937ba3aebcb1732afffaccea34"
	ConnectivityCheckJSONMockImage   = "quay.io/cilium/json-mock:v1.3.2@sha256:bc6c46c74efadb135bc996c2467cece6989302371ef4e3f068361460abaf39be"
	ConnectivityDNSTestServerImage   = "docker.io/coredns/coredns:1.9.4@sha256:b82e294de6be763f73ae71266c8f5466e7e03c69f3a1de96efd570284d35bb18"

	ConfigMapName = "cilium-config"
	Version       = "v1.12.2"

	StatusWaitDuration = 5 * time.Minute

	WaitRetryInterval   = 2 * time.Second
	WaitWarningInterval = 10 * time.Second

	FlowWaitTimeout   = 10 * time.Second
	FlowRetryInterval = 500 * time.Millisecond

	PolicyWaitTimeout = 15 * time.Second

	IngressClassName        = "cilium"
	IngressControllerName   = "cilium.io/ingress-controller"
	IngressSecretsNamespace = "cilium-secrets"

	HelmValuesSecretName          = "cilium-cli-helm-values"
	HelmValuesSecretKeyName       = "io.cilium.cilium-cli"
	HelmChartVersionSecretKeyName = "io.cilium.chart-version"
)

var (
	// ClusterMeshDeploymentLabels are the labels set on the clustermesh API server by default.
	ClusterMeshDeploymentLabels = map[string]string{
		"k8s-app": "clustermesh-apiserver",
	}

	// CiliumPodSelector is the pod selector to be used for the Cilium agents.
	CiliumPodSelector = "k8s-app=cilium"
)

// All hubble values from `cilium-config` configmap:
// https://github.com/cilium/cilium/blob/d9a04be9d714e5f5544cbca7ef8db7a151bfce96/install/kubernetes/cilium/templates/cilium-configmap.yaml#L709-L750
// this list is used to cherry-pick only hubble related values for configmap patch
// when running in unknown install state (i.e. when `cilium-cli-helm-values` doesn't exist)
var HubbleKeys = []string{
	"enable-hubble",
	"hubble-disable-tls",
	"hubble-event-buffer-capacity",
	"hubble-event-queue-size",
	"hubble-flow-buffer-size",
	"hubble-listen-address",
	"hubble-metrics",
	"hubble-metrics-server",
	"hubble-socket-path",
	"hubble-tls-cert-file",
	"hubble-tls-client-ca-files",
	"hubble-tls-key-file",
}
