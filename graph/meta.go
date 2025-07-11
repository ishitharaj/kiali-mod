package graph

// MetadataKey is a mnemonic type name for string
type MetadataKey string

// Metadata is a map for storing node and edge metadata values reported by the vendors
type Metadata map[MetadataKey]interface{}

// NewMetadata returns an empty Metadata map
func NewMetadata() Metadata {
	return make(map[MetadataKey]interface{})
}

// Metadata keys to be used instead of literal strings
const (
	Aggregate             MetadataKey = "aggregate" // the prom attribute used for aggregation
	AggregateValue        MetadataKey = "aggregateValue"
	DestPrincipal         MetadataKey = "destPrincipal"
	DestServices          MetadataKey = "destServices"
	HealthData            MetadataKey = "healthData"
	HealthDataApp         MetadataKey = "healthDataApp" // for storing app health on versioned app nodes
	HasCB                 MetadataKey = "hasCB"
	HasFaultInjection     MetadataKey = "hasFaultInjection"
	HasHealthConfig       MetadataKey = "hasHealthConfig"
	HasIngressWaypoint    MetadataKey = "hasIngressWaypoint"
	HasMirroring          MetadataKey = "hasMirroring"
	HasTCPTrafficShifting MetadataKey = "hasTCPTrafficShifting"
	HasTrafficShifting    MetadataKey = "hasTrafficShifting"
	HasRequestRouting     MetadataKey = "hasRequestRouting"
	HasRequestTimeout     MetadataKey = "hasRequestTimeout"
	HasVS                 MetadataKey = "hasVS"
	HasWorkloadEntry      MetadataKey = "hasWorkloadEntry"
	IsAmbient             MetadataKey = "isAmbient"
	IsDead                MetadataKey = "isDead"
	IsEgressCluster       MetadataKey = "isEgressCluster"  // PassthroughCluster or BlackHoleCluster
	IsEgressGateway       MetadataKey = "isEgressGateway"  // Identifies a node that is an Istio egress gateway
	IsExtension           MetadataKey = "isExtension"      // Identifies a node from an configured extension
	IsGatewayAPI          MetadataKey = "isGatewayAPI"     // Identifies a node that is a Gateway API gateway (ingress)
	IsIngressGateway      MetadataKey = "isIngressGateway" // Identifies a node that is an Istio ingress gateway
	IsIdle                MetadataKey = "isIdle"
	IsInaccessible        MetadataKey = "isInaccessible"
	IsInjected            MetadataKey = "isInjected"      // Identifies an injected service node (server-side use only)
	IsK8sGatewayAPI       MetadataKey = "isK8sGatewayAPI" // true when config is autogenerated from K8s API Gateway
	IsMTLS                MetadataKey = "isMTLS"
	IsOutOfMesh           MetadataKey = "isOutOfMesh"
	IsOutside             MetadataKey = "isOutside"
	IsRoot                MetadataKey = "isRoot"
	IsServiceEntry        MetadataKey = "isServiceEntry"
	IsWaypoint            MetadataKey = "isWaypoint"
	Labels                MetadataKey = "labels"
	ProtocolKey           MetadataKey = "protocol"
	ResponseTime          MetadataKey = "responseTime"
	SourcePrincipal       MetadataKey = "sourcePrincipal"
	Throughput            MetadataKey = "throughput"
	Waypoint              MetadataKey = "waypoint" // Information for edges to or from a waypoint
)

// DestServicesMetadata key=Service.Key()
type DestServicesMetadata map[string]ServiceName

// NewDestServicesMetadata returns an empty DestServicesMetadata map
func NewDestServicesMetadata() DestServicesMetadata {
	return make(map[string]ServiceName)
}

// Add adds or replaces a destService
func (dsm DestServicesMetadata) Add(key string, service ServiceName) DestServicesMetadata {
	dsm[key] = service
	return dsm
}

type GatewaysMetadata map[string][]string
type LabelsMetadata map[string]string
type VirtualServicesMetadata map[string][]string
