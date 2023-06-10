package predict

import (
	"log"
	"os"
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
)

const configPath = ".fstorage"

type EncodeType uint8

const (
	None EncodeType = iota
	Brotli
	Bzip2
	Gzip
	Lzma
	Zstd
	Snappy
	Lz4
)

type AlgorithmConfig struct {
	Brotli struct {
		Metric float32 `yaml:"metric"`
		Level  int     `yaml:"level"`
	} `yaml:"brotli"`
	bzip2 struct {
		Metric float32 `yaml:"metric"`
		Level  int     `yaml:"level"`
	} `yaml:"bzip2"`
	gzip struct {
		Metric float32 `yaml:"metric"`
		Level  int     `yaml:"level"`
	} `yaml:"gzip"`
	lzma struct {
		Metric float32 `yaml:"metric"`
		Level  int     `yaml:"level"`
	} `yaml:"lzma"`
	zstd struct {
		Metric float32 `yaml:"metric"`
		Level  int     `yaml:"level"`
	} `yaml:"zstd"`
	snappy struct {
		Metric float32 `yaml:"metric"`
		Level  int     `yaml:"level"`
	} `yaml:"snappy"`
	lz4 struct {
		Metric float32 `yaml:"metric"`
		Level  int     `yaml:"level"`
	} `yaml:"lz4"`
}

func ReadConfig() (cfg AlgorithmConfig) {
	cfg = AlgorithmConfig{}
	err := cleanenv.ReadConfig(filepath.Join(configPath, "config.yml"), &cfg)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return cfg
}
