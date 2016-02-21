package main

import (
	"runtime"
	"./mutex"
	"./race_conditions"
)

func main() {
	runtime.GOMAXPROCS(2)
	mutex.MutexNative()
	mutex.MutexPackageSync()
	race_conditions.GenerateRaceCondition()
}



