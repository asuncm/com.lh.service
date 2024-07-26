package pgx

import (
	"com.lh.service/tools"
	"context"
	"fmt"
	"github.com/yugabyte/pgx/v5"
	"github.com/yugabyte/pgx/v5/pgxpool"
	"gopkg.in/yaml.v3"
	"os"
	"reflect"
	"strings"
)

type Option struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int16  `json:"port"`
}

type Config struct {
	Name string `json:"name"`
	DB   string `json:"db"`
	Key  string `json:"key"`
	Type string `json:"type"`
	Role string `json:"role"`
	Rule string `json:"rule"`
	User string `json:"user"`
}

type dbConf = map[string]Option

var options dbConf

func getConfig() (dbConf, error) {
	platform := tools.Platform("")
	filename := tools.GetPath("LHPATH", fmt.Sprintf("%s%s%s", "com.lh.service/config/database.", platform.Env, ".yaml"))
	config, err := os.ReadFile(filename)
	opts := dbConf{}
	if err != nil {
		return opts, err
	}
	err = yaml.Unmarshal(config, &opts)
	if err != nil {
		return opts, err
	}
	return opts, err
}

func getUri(key string) string {
	conf := reflect.ValueOf(options)
	opts := conf.FieldByName(key)
	user := opts.FieldByName("Name").String()
	pw := opts.FieldByName("Password").String()
	host := opts.FieldByName("Host").String()
	port := opts.FieldByName("Port").Int()
	return fmt.Sprintf("%s:%s@%s:%d", user, pw, host, port)
}

func InitConfig() {
	var err error
	options, err = getConfig()
	if err != nil {
		options = dbConf{}
	}
}

func openDB(opts Config) (*pgx.Conn, error) {
	uri := getUri(opts.Name)
	url := fmt.Sprintf("postgres://%s/%s?load_balance=true", uri, opts.DB)
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func poolDB(opts Config) (*pgxpool.Pool, error) {
	uri := getUri(opts.Name)
	url := fmt.Sprintf("postgres://%s/%s?load_balance=true", uri, opts.DB)
	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func InitCreate(opts Config) bool {
	pool, err := poolDB(opts)
	if err != nil {
		return false
	}
	_, err = pool.Exec(context.Background(), fmt.Sprintf("%s %s %s;", "CREATE", opts.Type, opts.Key))
	if err != nil {
		return false
	}
	return true
}

func InitDrop(opts Config) bool {
	pool, err := poolDB(opts)
	if err != nil {
		return false
	}
	_, err = pool.Exec(context.Background(), fmt.Sprintf("%s %s %s;", "CREATE", opts.Type, opts.Key))
	if err != nil {
		return false
	}
	return true
}

func InitAlter(opts Config) bool {
	pool, err := poolDB(opts)
	if err != nil {
		return false
	}
	_, err = pool.Exec(context.Background(), fmt.Sprintf("%s %s %s;", "ALTER", opts.Type, opts.Key))
	if err != nil {
		return false
	}
	return true
}

func InitGrant(opts Config) bool {
	pool, err := poolDB(opts)
	if err != nil {
		return false
	}
	_, err = pool.Exec(context.Background(), fmt.Sprintf("%s %s %s;", "GRANT", opts.Role, "TO", opts.User))
	if err != nil {
		return false
	}
	return true
}

func InitRevoke(opts Config) bool {
	pool, err := poolDB(opts)
	if err != nil {
		return false
	}
	_, err = pool.Exec(context.Background(), fmt.Sprintf("%s %s %s;", "REVOKE", opts.Role, "FROM", opts.User))
	if err != nil {
		return false
	}
	return true
}

func Ping(opts Config) bool {
	pool, err := poolDB(opts)
	if err != nil {
		return false
	}
	err = pool.Ping(context.Background())
	if err != nil {
		return false
	}
	return true
}
func FormatKeys(values []string) string {
	return strings.Join(values, ", ")
}
