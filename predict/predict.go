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

	minVal := float32(10.0)
	bestAlg := None
	if selectedConfig.Bzip2.Metric < minVal {
		minVal = selectedConfig.Bzip2.Metric
		bestAlg = Bzip2
	}
	if selectedConfig.Gzip.Metric < minVal {
		minVal = selectedConfig.Gzip.Metric
		bestAlg = Gzip
	}
	if selectedConfig.Lz4.Metric < minVal {
		minVal = selectedConfig.Lz4.Metric
		bestAlg = Lz4
	}
	if selectedConfig.Lzma.Metric < minVal {
		minVal = selectedConfig.Lzma.Metric
		bestAlg = Lzma
	}
	if selectedConfig.Zstd.Metric < minVal {
		minVal = selectedConfig.Zstd.Metric
		bestAlg = Zstd
	}
	if selectedConfig.Snappy.Metric < minVal {
		minVal = selectedConfig.Snappy.Metric
		bestAlg = Snappy
	}
	if selectedConfig.Brotli.Metric < minVal {
		minVal = selectedConfig.Brotli.Metric
		bestAlg = Brotli
	}
	return bestAlg, selectedConfig
}
