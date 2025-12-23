package messaging

// ClusterConfig - message broker cluster config object
type ClusterConfig struct {
	// Host - message cluster host
	Host string `env:"MESSAGE_CLUSTER_HOST, default=localhost"`
	// User - message cluster user
	User string `env:"MESSAGE_CLUSTER_USER, default=user"`
	// Password - message cluster password
	Password string `env:"MESSAGE_CLUSTER_PASSWORD, default=password"`
}

// ConsumerConfig - message consumer config object
type ConsumerConfig struct {
	ClusterConfig
	// Topics - list of topic names to consume from
	Topics []string `env:"MESSAGE_CONSUMER_TOPICS, required"`

	// MaxPendingMsgs - message consumer flow control for max pending msgs for all consumers
	MaxPendingMsgs int `env:"MESSAGE_CONSUMER_MAX_PENDING_MESSAGES, default=0"`
}

// PublisherConfig - message publisher config object
type PublisherConfig struct {
	ClusterConfig
}
