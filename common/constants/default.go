package constants

import (
	"time"

	v1 "k8s.io/api/core/v1"
)

const (
	DefaultConfigDir    = "/etc/kubeedge/config/"
	DefaultCAFile       = "/etc/kubeedge/ca/rootCA.crt"
	DefaultCAKeyFile    = "/etc/kubeedge/ca/rootCA.key"
	DefaultCertFile     = "/etc/kubeedge/certs/server.crt"
	DefaultKeyFile      = "/etc/kubeedge/certs/server.key"
	DefaultEdgeCertFile = "/etc/kubeedge/certs/edge.crt"
	DefaultEdgeKeyFile  = "/etc/kubeedge/certs/edge.key"
	DefaultCADir        = "/etc/kubeedge/ca"
	DefaultCertDir      = "/etc/kubeedge/certs"

	DefaultStreamCAFile   = "/etc/kubeedge/ca/streamCA.crt"
	DefaultStreamCertFile = "/etc/kubeedge/certs/stream.crt"
	DefaultStreamKeyFile  = "/etc/kubeedge/certs/stream.key"
)

const (
	DefaultDockerAddress               = "unix:///var/run/docker.sock"
	DefaultRuntimeType                 = "docker"
	DefaultEdgedMemoryCapacity         = 7852396000
	DefaultRemoteRuntimeEndpoint       = "unix:///var/run/dockershim.sock"
	DefaultRemoteImageEndpoint         = "unix:///var/run/dockershim.sock"
	DefaultPodSandboxImage             = "kubeedge/pause:3.1"
	DefaultArmPodSandboxImage          = "kubeedge/pause-arm:3.1"
	DefaultArm64PodSandboxImage        = "kubeedge/pause-arm64:3.1"
	DefaultNodeStatusUpdateFrequency   = 10
	DefaultImagePullProgressDeadline   = 60
	DefaultRuntimeRequestTimeout       = 2
	DefaultImageGCHighThreshold        = 80
	DefaultImageGCLowThreshold         = 40
	DefaultMaximumDeadContainersPerPod = 1
	DefaultHostnameOverride            = "default-edge-node"
	DefaultRegisterNodeNamespace       = "default"
	DefaultInterfaceName               = "eth0"
	DefaultCNIConfDir                  = "/etc/cni/net.d"
	DefaultCNIBinDir                   = "/opt/cni/bin"
	DefaultCNICacheDir                 = "/var/lib/cni/cache"
	DefaultNetworkPluginMTU            = 1500
	DefaultConcurrentConsumers         = 5
	DefaultCgroupRoot                  = ""
	DefaultVolumeStatsAggPeriod        = time.Minute
	DefaultTunnelPort                  = 10004
)
const (
	DefaultPodStatusSyncInterval = 60
)

// Config
const (
	DefaultKubeContentType         = "application/vnd.kubernetes.protobuf"
	DefaultKubeConfig              = "/root/.kube/config"
	DefaultKubeNamespace           = v1.NamespaceAll
	DefaultKubeQPS                 = 100.0
	DefaultKubeBurst               = 200
	DefaultKubeUpdateNodeFrequency = 20

	DefaultUpdatePodStatusWorkers            = 1
	DefaultUpdateNodeStatusWorkers           = 1
	DefaultQueryConfigMapWorkers             = 4
	DefaultQuerySecretWorkers                = 4
	DefaultQueryServiceWorkers               = 4
	DefaultQueryEndpointsWorkers             = 4
	DefaultQueryPersistentVolumeWorkers      = 4
	DefaultQueryPersistentVolumeClaimWorkers = 4
	DefaultQueryVolumeAttachmentWorkers      = 4
	DefaultQueryNodeWorkers                  = 4
	DefaultUpdateNodeWorkers                 = 4
	DefaultDeletePodWorkers                  = 4

	DefaultUpdatePodStatusBuffer            = 1024
	DefaultUpdateNodeStatusBuffer           = 1024
	DefaultQueryConfigMapBuffer             = 1024
	DefaultQuerySecretBuffer                = 1024
	DefaultQueryServiceBuffer               = 1024
	DefaultQueryEndpointsBuffer             = 1024
	DefaultQueryPersistentVolumeBuffer      = 1024
	DefaultQueryPersistentVolumeClaimBuffer = 1024
	DefaultQueryVolumeAttachmentBuffer      = 1024
	DefaultQueryNodeBuffer                  = 1024
	DefaultUpdateNodeBuffer                 = 1024
	DefaultDeletePodBuffer                  = 1024

	DefaultETCDTimeout = 10

	DefaultEnableElection = false
	DefaultElectionTTL    = 30
	DefaultElectionPrefix = "/controller/leader"

	DefaultMessageLayer = "context"

	DefaultContextSendModuleName     = "cloudhub"
	DefaultContextReceiveModuleName  = "edgecontroller"
	DefaultContextResponseModuleName = "cloudhub"

	DefaultPodEventBuffer       = 1
	DefaultConfigMapEventBuffer = 1
	DefaultSecretEventBuffer    = 1
	DefaultServiceEventBuffer   = 1
	DefaultEndpointsEventBuffer = 1

	// Resource sep
	ResourceSep = "/"

	ResourceTypeService       = "service"
	ResourceTypeServiceList   = "servicelist"
	ResourceTypeEndpoints     = "endpoints"
	ResourceTypeEndpointsList = "endpointslist"
	ResourceTypeListener      = "listener"

	ResourceTypePersistentVolume      = "persistentvolume"
	ResourceTypePersistentVolumeClaim = "persistentvolumeclaim"
	ResourceTypeVolumeAttachment      = "volumeattachment"

	CSIResourceTypeVolume                     = "volume"
	CSIOperationTypeCreateVolume              = "createvolume"
	CSIOperationTypeDeleteVolume              = "deletevolume"
	CSIOperationTypeControllerPublishVolume   = "controllerpublishvolume"
	CSIOperationTypeControllerUnpublishVolume = "controllerunpublishvolume"
	CSISyncMsgRespTimeout                     = 1 * time.Minute

	CurrentSupportK8sVersion = "v1.17.1"
)

const (
	DefaultUpdateDeviceStatusBuffer  = 1024
	DefaultDeviceEventBuffer         = 1
	DefaultDeviceModelEventBuffer    = 1
	DefaultUpdateDeviceStatusWorkers = 1
)

const (
	// TODO put all modulename and group name together @kadisi
	DeviceTwinModuleName = "twin"
)

const (
	// ServerPort is the default port for the edgecore server on each host machine.
	// May be overridden by a flag at startup in the future.
	ServerPort = 10350
)

const (
	KubeEdge          = "kubeedge"
	KubeEdgeNameSpace = "kubeedge"
)
