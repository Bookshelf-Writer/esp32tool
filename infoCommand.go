package esp32tool

import (
	"fmt"
	"github.com/Bookshelf-Writer/esp32tool/esp32"
	"github.com/Bookshelf-Writer/esp32tool/lib/output"
)

type DeviceInfo struct {
	ChipType   string
	Revision   string
	Features   []string
	MacAddress string
	Partitions esp32.PartitionList
}

func InfoCommand(esp32 *esp32.ESP32ROM, logs *output.LogObj) error {
	macAddress, err := esp32.GetChipMAC()
	if err != nil {
		return fmt.Errorf("Could not retrieve MAC address: %s", err.Error())
	}

	description, err := esp32.GetChipDescription()
	if err != nil {
		return fmt.Errorf("Could not retrieve chip description: %s", err.Error())
	}

	features, err := esp32.GetFeatures()
	if err != nil {
		return fmt.Errorf("Could not retrieve chip features: %s", err.Error())
	}

	featureList := make([]string, 0)
	for feature, status := range features {
		if status {
			featureList = append(featureList, feature.String())
		}
	}

	maxSize := uint64(0)
	partitionList, err := esp32.ReadPartitionList()
	if err != nil {
		logs.Debug().Err(err).Msg("Partition")
	} else {
		for _, partition := range partitionList {
			if maxSize < uint64(partition.Size+partition.Offset) {
				maxSize = uint64(partition.Size + partition.Offset)
			}
		}
	}

	pr := logs.Info()

	pr.Str("mac", macAddress)
	pr.Int("revision", int(description.Revision))
	pr.Uint64("maxSize", maxSize)
	pr.Array("feature", output.StringArray(featureList))
	pr.Interface("partition", partitionList)

	pr.Msg(description.ChipType.String())
	return err
}
