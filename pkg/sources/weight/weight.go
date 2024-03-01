package weight

import (
	"asset-locator/pkg/sources/weight/providers"
	"asset-locator/pkg/sources/weight/providers/aizhan"
	"asset-locator/pkg/sources/weight/providers/pearktrue"

	"github.com/projectdiscovery/gologger"
)

// SearchWeight search weight via aizhan, only domain can search weight
func SearchWeight(domain string) (providers.WeightLevel, error) {
	engines := loadWeightEngines()
	for _, engine := range engines {
		weight, err := engine.SearchWeight(domain)
		if err != nil {
			gologger.Debug().
				Label("Weight").
				Msgf("Engine %s get weight error", engine.Name(), err.Error())
			continue
		}
		return weight, nil
	}
	return providers.ErrorLevel, nil
}

func loadWeightEngines() []providers.WeightEngine {
	var weightEngines []providers.WeightEngine

	weightEngines = append(weightEngines, pearktrue.WeightEngine{})
	weightEngines = append(weightEngines, aizhan.WeightEngine{})
	return weightEngines
}
