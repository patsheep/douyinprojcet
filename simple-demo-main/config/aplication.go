package config

// GinConfig 定义 Gin 配置文件的结构体
type GinConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

// MySQLConfig 定义 mysql 配置文件结构体
type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	DBname       string `mapstructure:"db_name"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

// RedisConfig 定义 redis 配置文件结构体
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}


// OssConfig 定义oss结构体
type OssConfig struct {
	Key string `mapstructure:"key"`
	Secret string `mapstructure:"secret"`
	Endpoint string `mapstructure:"endpoint"`
	Bucket string `mapstructure:"bucket"`
}
// SnowflakeConfig 定义雪花算法接口结构体
type SnowFlakeConfig struct {
	MechineId int64 `mapstructure:"mechineid"`
}

// System 定义项目配置文件结构体
type SystemConf struct {
	GinConfig   *GinConfig   `mapstructure:"gin"`
	MySQLConfig *MySQLConfig `mapstructure:"mysql"`
	RedisConfig *RedisConfig `mapstructure:"redis"`
	OssConfig	*OssConfig   `mapstructure:"oss"`
	SnowFlakeConfig *SnowFlakeConfig `mapstructure:"snowflake"`
}

