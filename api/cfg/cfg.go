package cfg

import (
	"errors"
	"os"
)

type Env struct {
	ServiceKeyFile string
	PortStr        string
}

func LoadEnv() (*Env, error) {
	serviceKeyFile := os.Getenv("SERVICE_KEY_FILE")
	if serviceKeyFile == "" {
		return nil, errors.New("missing service key file")
	}

	apiPortStr := os.Getenv("API_PORT")
	if apiPortStr == "" {
		return nil, errors.New("missing api port")
	}

	e := &Env{
		ServiceKeyFile: serviceKeyFile,
		PortStr:        apiPortStr,
	}
	return e, nil
}
