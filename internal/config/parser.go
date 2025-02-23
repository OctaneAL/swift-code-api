package config

import (
	"fmt"

	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
)

type Parser struct {
	DataPath string
}

const (
	parserYamlKey = "parser"
)

func (c *config) DataPath() Parser {
	return c.dataPath.Do(func() interface{} {
		var cfg struct {
			DataPath *string `fig:"data_path,required"`
		}

		err := figure.
			Out(&cfg).
			From(kv.MustGetStringMap(c.getter, parserYamlKey)).
			Please()
		if err != nil {
			panic(fmt.Errorf("failed to figure out config %s: %w", parserYamlKey, err))
		}

		return Parser{
			DataPath: *cfg.DataPath,
		}
	}).(Parser)
}
