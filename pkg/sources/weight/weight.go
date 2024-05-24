package weight

import (
	"github.com/N0el4kLs/asset-locator/pkg/sources/weight/providers"
	"github.com/N0el4kLs/asset-locator/pkg/sources/weight/providers/aizhan"
	"github.com/N0el4kLs/asset-locator/pkg/sources/weight/providers/pearktrue"
	"github.com/N0el4kLs/asset-locator/pkg/util"

	"github.com/projectdiscovery/gologger"
)

// SearchWeight search weight via aizhan, only domain can search weight
func SearchWeight(domain string) (providers.WeightLevel, error) {
	engines := loadWeightEngines()
	for _, engine := range engines {
		weight, err := engine.SearchWeight(util.ExtractDomain(domain))
		if err != nil {
			gologger.Debug().
				Label("Weight").
				Msgf("Engine %s get weight error: %s\n", engine.Name(), err.Error())
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
