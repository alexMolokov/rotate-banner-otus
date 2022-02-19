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
	ExchangeName string `config:"queue-exchangeName"`
	ExchangeType string `config:"queue-exchangeType"`
}
