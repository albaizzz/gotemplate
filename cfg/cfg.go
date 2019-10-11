package cfg

//Configuration structure
type Configuration struct {
	Common         *Common         `yaml:"common"`
	DB             *DB             `yaml:"db"`
	Breaker        *Breaker        `yaml:"breaker"`
	AWS            *AWS            `yaml:"aws"`
	Country        *Country        `yaml:"country"`
	APM            *APM            `yaml:"apm"`
	Monitor        *Monitor        `yaml:"monitor"`
	HTTP           *HTTP           `yaml:"http"`
	Kafka          *Kafka          `yaml:"kafka"`
	PartnerWallet  *PartnerWallet  `yaml:"partner_wallet"`
	CircuitBreaker *CircuitBreaker `yaml:"circuit_breaker"`
}

type DB struct {
	Maria *Maria       `yaml:"maria"`
	Redis *RedisCluter `yaml:"redis_cluster"`
}

//RedisCluter configuration
type RedisCluter struct {
	Host                  string `yaml:"host"`
	RetryCount            int    `yaml:"retry_count"`
	RetryDuration         int    `yaml:"retry_duration"`
	MaxActive             int    `yaml:"max_active"`
	MaxIdle               int    `yaml:"max_idle"`
	IdleTimeout           int    `yaml:"idle_timeout"`
	DialConnectionTimeout int    `yaml:"dial_connection_timeout"`
}

type Maria struct {
	MasterDB *Database `yaml:"master_db"`
	SlaveDB  *Database `yaml:"slave_db"`
}

//Common structure
type Common struct {
	AppName       string `yaml:"app_name"`
	APIKey        string `yaml:"api_key"`
	IsDebug       bool   `yaml:"is_debug"`
	UseNewRelic   bool   `yaml:"use_new_relic"`
	UseDatadog    bool   `yaml:"use_datadog"`
	UseSentry     bool   `yaml:"use_sentry"`
	UseSentryDsn  bool   `yaml:"use_sentry_dsn"`
	IsMaintenance bool   `yaml:"is_maintenance"`
	Timezone      string `yaml:"timezone"`
	Lang          string `yaml:"lang"`
	Environment   string `yaml:"environments"`
}

//Database configuration structure
type Database struct {
	Name string `yaml:"name"`
	User string `yaml:"user"`
	Host string `yaml:"host`
	Port int    `yaml:"port`
	Pass string `yaml:"pass"`
}

//PartnerWallet structure
type PartnerWallet struct {
	ClientID    string `yaml:"client_id"`
	SecretKey   string `yaml:"secret_key"`
	Host        string `yaml:"host"`
	LastMigrate string `yaml:"last_migrate"`
}

//Breaker structure
type Breaker struct {
	MaxFailure   int `yaml:"max_failure`
	ResetTimeout int `yaml:"reset_timeout"`
}

//HTTP generic config
type HTTP struct {
	Timeout int
}

//AWS structure
type AWS struct {
	Region string `yaml:"region"`
}

//Country structure
type Country struct {
	Name     string `yaml:"name"`
	Currency string `yaml:"currency"`
}

//APM structure
type APM struct {
	Address   string `yaml:"address"`
	Port      int    `yaml:"port"`
	IsEnabled bool   `yaml:"is_enabled"`
}

//Monitor as monitoring structure
type Monitor struct {
	Address   string `yaml:"address"`
	Port      int    `yaml:"port"`
	IsEnabled string `yaml:"is_enabled"`
}

//Kafka as kafka config
type Kafka struct {
	Brokers           []string `yaml:"brokers"`
	GroupID           string   `yaml:"group_id"`
	Protocol          string   `yaml:"protocol"`
	UseRetry          bool     `yaml:"use_retry"`
	MaxInFlight       int      `yaml:"max_in_flight"`
	MaxBytes          int      `yaml:"max_bytes"`
	MinBytes          int      `yaml:"min_bytes"`
	MaxRetryAttempt   int      `yaml:"max_retry_attempt"`
	DialTimeoutSecond int      `yaml:"dial_timeout_second"`
	SendTimeoutSecond int      `yaml:"send_timeout_second"`
	DefaultPartition  int      `yaml:"default_partition"`
}

//CircuitBreaker configuration
type CircuitBreaker struct {
	MaxConcurrency           int `yaml:"max_concurrency"`
	ErrorPercentageThreshold int `yaml:"error_threshold"`
	TimeoutSecond            int `yaml:"timeout_second"`
}
