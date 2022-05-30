package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
	"runtime"
)

func init() {
	env := os.Getenv("env")
	//if env == "" {
	//	env = "dev"
	//}

	viper.SetConfigName("app-" + env)
	viper.SetConfigType("toml")
	viper.AddConfigPath("conf")

	if env == "dev"{
		//获取调用栈路径
		currentPath := GetCurrentPath()

		//获取项目根目录的绝对路径
		if len(currentPath) >= 11 {
			currentPath = currentPath[:len(currentPath) - 11]
			viper.AddConfigPath(currentPath + "/conf")
		}

		err := viper.ReadInConfig()

		if err != nil {
			fmt.Printf("read config failed: %+v", err)
		}

		return
	}
}

type Config struct {
	App             App             `toml:"app"`
	Log             Log             `toml:"log"`
	Whitelist       Whitelist       `toml:"whitelist"`
	Sync            Sync            `toml:"sync"`
	Db              Db              `toml:"db"`
	Cache           Cache           `toml:"cache"`
	Redis           Redis           `toml:"redis"`
	Kafka           Kafka           `toml:"kafka"`
	Task            Task            `toml:"task"`
	Google          Google          `toml:"google"`
	Facebook        Facebook        `toml:"facebook"`
	RPC             RPC             `toml:"rpc"`
	Grpc            Grpc            `toml:"grpc"`
	HTTP            HTTP            `toml:"http"`
	Nats            Nats            `toml:"nats"`
	ClickhouseRead  ClickhouseRead  `toml:"clickhouse-read"`
	ClickhouseWrite ClickhouseWrite `toml:"clickhouse-write"`
	DbType          DbType          `toml:"db-type"`
}

type App struct {
	Name               string `toml:"name"`
	Version            string `toml:"version"`
	CampaignThreadSize int    `toml:"campaign_thread_size"`
	HistoryThreadSize  int    `toml:"history_thread_size"`
	AdSize             int    `toml:"ad_size"`
	UpdateSize         int    `toml:"update_size"`
	InsertSize         int    `toml:"insert_size"`
	TopicSize          int    `toml:"topic_size"`
	DumperPath         string `toml:"dumper_path"`
	ZipPath            string `toml:"zip_path"`
	SyncPath           string `toml:"sync_path"`
	Debug              int    `toml:"debug"`
	UpdateTable        string `toml:"update_table"`
	UpdateType         string `toml:"update_type"`
}
type Log struct {
	Path  string `toml:"path"`
	Debug int    `toml:"debug"`
}
type Whitelist struct {
	ProfileID int64 `toml:"profileId"`
}
type Sync struct {
	BatchSize   int    `toml:"batchSize"`
	MinSize     int    `toml:"minSize"`
	ResourceID  int    `toml:"resourceId"`
	Op          string `toml:"op"`
	WaitingTime int    `toml:"waiting_time"`
	Timeout     int    `toml:"timeout"`
}
type Db struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Database string `toml:"database"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}
type Cache struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Password string `toml:"password"`
	Database int    `toml:"database"`
}
type Redis struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Password string `toml:"password"`
	Database int    `toml:"database"`
}
type Kafka struct {
	Type            string `toml:"type"`
	Broker          string `toml:"broker"`
	SaslEnable      int    `toml:"sasl_enable"`
	MaxOpenRequests int    `toml:"max_open_requests"`
}
type Task struct {
	ReportURL  string `toml:"report_url"`
	RequestURL string `toml:"request_url"`
}
type Google struct {
	Path            string `toml:"path"`
	DevToken        string `toml:"devToken"`
	LoginCustomerID int    `toml:"loginCustomerID"`
	ClientID        string `toml:"clientId"`
	ClientSecret    string `toml:"clientSecret"`
	RefreshToken    string `toml:"refreshToken"`
}
type Facebook struct {
	Path         string `toml:"path"`
	ClientID     int64  `toml:"clientId"`
	ClientSecret string `toml:"clientSecret"`
	RefreshToken string `toml:"refreshToken"`
}
type RPC struct {
	Network string `toml:"network"`
	Address string `toml:"address"`
	Host    string `toml:"host"`
	Port    int    `toml:"port"`
	Timeout int    `toml:"timeout"`
}
type Grpc struct {
	Port string `toml:"port"`
}
type HTTP struct {
	Port string `toml:"port"`
}
type Nats struct {
	URL string `toml:"url"`
}
type ClickhouseRead struct {
	Dsn string `toml:"dsn"`
}
type ClickhouseWrite struct {
	Dsn string `toml:"dsn"`
}
type DbType struct {
	Value string `toml:"value"`
}

type ServiceOption struct {
	ClientType              string `json:"client_type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	PrivateKeyPath          string `json:"private_key_path"`
	ClientEmail             string `json:"client_email"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x_509_cert_url"`
	ClientX509CertURL       string `json:"client_x_509_cert_url"`
}

func Get(name string) map[string]string {
	return viper.GetStringMapString(name)
}

func GetCurrentPath() string {
	//获取程序运行时的文件名
	_, filename, _, _ := runtime.Caller(1)
	//获取具体路径
	return path.Dir(filename)
}

func ReadConfig(path string, fileType string, name string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(name)
	v.SetConfigType(fileType)
	v.AddConfigPath(path)

	currentPath := GetCurrentPath()
	if len(currentPath) >= 11 {
		currentPath = currentPath[:len(currentPath)-11]
		v.AddConfigPath(currentPath + "/conf")
	}

	err := v.ReadInConfig()

	if err != nil{
		fmt.Printf("read config failed: %+v\n", err)
		return nil, err
	}

	return v, nil
}

func GetString(name string) string {
	return viper.GetString(name)
}

func GetStringOrDefault(name string, defaultValue string) string {
	value := viper.GetString(name)
	if value == "" {
		value = defaultValue
	}

	return value
}

func GetInt(name string) int {
	return viper.GetInt(name)
}

func GetIntOrDefault(name string, defaultValue int) int {
	value := viper.GetInt(name)
	if value == 0 {
		value = defaultValue
	}

	return value
}
