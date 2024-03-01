package icp

import (
	"github.com/N0el4kLs/asset-locator/pkg/sources/icp/providers"
	"github.com/N0el4kLs/asset-locator/pkg/sources/icp/providers/chinaz"
	"github.com/N0el4kLs/asset-locator/pkg/sources/icp/providers/icplishi"
	"github.com/N0el4kLs/asset-locator/pkg/sources/icp/providers/pearktrue"

	"github.com/projectdiscovery/gologger"
)

func SearchICP(domain string) (string, error) {
	engines := loadIcpEngines()
	for _, engine := range engines {
		icp, err := engine.SearchICP(domain)
		if err != nil {
			gologger.Debug().
				Label("Icp").
				Msgf("Engine %s get icp error", engine.Name(), err.Error())
			continue
		}
		return icp, nil
	}
	return providers.None, nil
}

func loadIcpEngines() []providers.ICPEngine {
	var icpEngines []providers.ICPEngine

	icpEngines = append(icpEngines, pearktrue.ICPEngine{})
	icpEngines = append(icpEngines, icplishi.ICPEngine{})
	icpEngines = append(icpEngines, chinaz.ICPEngine{})
	return icpEngines
}
