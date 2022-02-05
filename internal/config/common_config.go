package config

type LoggerConf struct {
	Level    string `config:"logger-level"`
	Encoding string `config:"logger-encoding"`
	Output   string `config:"logger-output"`
}

type DBConf struct {
	Driver            string `config:"postgres-driver"`
	Host              string `config:"postgres-host"`
	Port              int    `config:"postgres-port"`
	Name              string `config:"postgres-db"`
	User              string `config:"postgres-user"`
	Password          string `config:"postgres-password"`
	MaxConnectionPool int    `config:"postgres-maxConnectionPool"`
	SslMode           string `config:"postgres-sslMode"`
}

type GRPCConf struct {
	Host string `config:"grpc-host"`
	Port int    `config:"grpc-port"`
}

type QueueConf struct {
	// параметры подключения
	Host     string `config:"queue-host"`
	Port     int    `config:"queue-port"`
	User     string `config:"queue-user"`
	Password string `config:"queue-password"`

	// параметры очереди
	Name         string `config:"queue-name"`
	ExchangeName string `config:"queue-exchangeName"`
	ExchangeType string `config:"queue-exchangeType"`
	// When setting up the channel after a channel exception in seconds
	ReInitDelay int `config:"queue-reInitDelay"`
	// When reconnecting to the server after connection failure in seconds
	ReconnectDelay int `config:"queue-reconnectDelay"`
	// resending messages the server didn't confirm in seconds
	ResendDelay int `config:"queue-resendDelay"`
}
