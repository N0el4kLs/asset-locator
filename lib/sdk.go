package lib

import (
	"fmt"

	"github.com/N0el4kLs/asset-locator/runner"
)

// AssetLocatorEngine is the main struct for the asset locator engine
type AssetLocatorEngine struct {
	core *runner.Runner
}

// NewAssetLocatorEngine creates a new asset locator engine
func NewAssetLocatorEngine(resultCallback func(result runner.Result)) *AssetLocatorEngine {
	core := &runner.Runner{
		Options:  DefaultOptions(),
		Callback: resultCallback,
	}

	return &AssetLocatorEngine{
		core: core,
	}
}

// UnSetWeight sets the weight option to false
func (a *AssetLocatorEngine) UnSetWeight() *AssetLocatorEngine {
	a.core.Options.Weight = false
	return a
}

// LoadTargets loads the targets into the asset locator engine
func (a *AssetLocatorEngine) LoadTargets(targets []string) *AssetLocatorEngine {
	a.core.Target = targets
	return a
}

func (a *AssetLocatorEngine) RunSearch() {
	if err := a.core.Run(); err != nil {
		panic(err)
	}
}

// DefaultResultCallback is the default result callback
func DefaultResultCallback(result runner.Result) {
	fmt.Printf("%s [%s] [%s] \n", result.Target, result.ICP, result.Weight)
}

// DefaultOptions returns the default options for the asset locator engine
func DefaultOptions() *runner.Options {
	return &runner.Options{
		Weight: true,
	}
}
