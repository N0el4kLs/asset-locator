package main

import (
	"fmt"
	"sync"

	"github.com/N0el4kLs/asset-locator/lib"
	"github.com/N0el4kLs/asset-locator/runner"
)

type MyEngine struct {
	resultChan chan runner.Result
	engine     *lib.AssetLocatorEngine
}

func (m *MyEngine) DefalutCallback(rst runner.Result) {
	fmt.Printf("Put target %s into result channel\n", rst.Target)
	m.resultChan <- rst
}

func NewMyEngine() *MyEngine {
	myEngine := &MyEngine{
		resultChan: make(chan runner.Result),
	}
	engine := lib.NewAssetLocatorEngine(myEngine.DefalutCallback)

	myEngine.engine = engine
	return myEngine
}

func main() {
	target := []string{"https://www.baidu.com", "jd.com"}
	myEngine := NewMyEngine()
	myEngine.engine.LoadTargets(target)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		myEngine.engine.RunSearch()
	}()

	go func() {
		wg.Wait()
		close(myEngine.resultChan)
	}()

	for rst := range myEngine.resultChan {
		// do something with the result
		fmt.Println(rst.Target, rst.ICP, rst.Weight.ToString())
	}
}
