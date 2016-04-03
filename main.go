package main

import (
	"runtime"
	"github.com/sKudryashov/conc_strategies/callbacks"
	"github.com/sKudryashov/conc_strategies/promises"
	"github.com/sKudryashov/conc_strategies/filters"
	"github.com/sKudryashov/conc_strategies/etl"
	"github.com/sKudryashov/conc_strategies/mutex"
	"github.com/sKudryashov/conc_strategies/race_conditions"
	"github.com/sKudryashov/conc_strategies/events"
)

func main() {
    runtime.GOMAXPROCS(2)
	mutex.MutexPackageSync()
	race_conditions.GenerateRaceCondition()
	events.InitEventsFactory()
	callbacks.InitPurchaseCallback()
	promises.InitOrders()
	filters.InitFilters()
	etl.InitETL()
}



