package yugabyte

import (
	"com.lh.service/tools"
	"context"
	"fmt"
	"github.com/yugabyte/pgx/v5"
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
	Allkic   Option `json:"allkic"`
	Yugabyte Option `json:"yugabyte"`
}

func getConfig() (Config, error) {
	platform := tools.Platform("")
	filename := tools.GetPath("LHPATH", fmt.Sprintf("%s%s%s", "com.lh.service/config/database.", platform.Env, ".yaml"))
	config, err := os.ReadFile(filename)
	opts := Config{}
	if err != nil {
		return opts, err
	}
	err = yaml.Unmarshal(config, &opts)
	if err != nil {
		return opts, err
	}
	return opts, err
}

func OpenDB(key string, value string) (*pgx.Conn, error) {
	config, err := getConfig()
	if err != nil {
		return nil, err
	}
	conf := reflect.ValueOf(config)
	opts := conf.FieldByName(key)
	user := opts.FieldByName("Name").String()
	pw := opts.FieldByName("Password").String()
	host := opts.FieldByName("Host").String()
	port := opts.FieldByName("Port").Int()
	yUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, pw, host, port, value)
	url := fmt.Sprintf("%s?load_balance=true", yUrl)
	conn, err := pgx.Connect(context.Background(), url)
	defer conn.Close(context.Background())
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func PoolDB() {

}

func FormatKeys(values []string) string {
	return strings.Join(values, ", ")
}
