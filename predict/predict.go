package predict

import (
	"github.com/meyakovenkoj/go-compress/typing"
)

func findMaxIndex(stat typing.FileStats) int {
	max := stat[0]
	index := 0
	for ind, value := range stat {
		if value > max {
			max = value
			index = ind
		}
	}
	return index
}

func Predict(stat typing.FileStats) (EncodeType, AlgorithmConfig) {
	config := ReadConfig()

	mostType := findMaxIndex(stat)

	selectedConfig := config[mostType]

	maxVal := float32(0.0)
	bestAlg := None
	if selectedConfig.bzip2.Metric > maxVal {
		maxVal = selectedConfig.bzip2.Metric
		bestAlg = Bzip2
	}
	if selectedConfig.gzip.Metric > maxVal {
		maxVal = selectedConfig.gzip.Metric
		bestAlg = Gzip
	}
	if selectedConfig.lz4.Metric > maxVal {
		maxVal = selectedConfig.lz4.Metric
		bestAlg = Lz4
	}
	if selectedConfig.lzma.Metric > maxVal {
		maxVal = selectedConfig.lzma.Metric
		bestAlg = Lzma
	}
	if selectedConfig.zstd.Metric > maxVal {
		maxVal = selectedConfig.zstd.Metric
		bestAlg = Zstd
	}
	if selectedConfig.snappy.Metric > maxVal {
		maxVal = selectedConfig.snappy.Metric
		bestAlg = Snappy
	}
	if selectedConfig.Brotli.Metric > maxVal {
		maxVal = selectedConfig.Brotli.Metric
		bestAlg = Brotli
	}
	return bestAlg, selectedConfig
}
