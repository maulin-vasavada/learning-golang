package main

import (
	"gobyexamples/misc"
)

func main() {
	/*
		myIntList := generics.List[int]{}
		myIntList.Push(10)
		myIntList.Push(20)
		myIntList.Push(15)
		myIntList.Push(8)
		intElements := myIntList.GetAll()
		fmt.Println(intElements)

		myStringList := generics.List[string]{}
		myStringList.Push("jerry")
		myStringList.Push("tom")
		myStringList.Push("harry")

		stringElements := myStringList.GetAll()
		fmt.Println(stringElements)
	*/

	/*
		fmt.Println(generics.Max[int](45, 67))
		var m = map[int]string{1: "2", 2: "4", 4: "8"}
		fmt.Println("keys:", generics.MapKeys(m))
	*/

	//channels.PingPong()
	//channels.RangeOverChannel()
	//selecttest.Select()
	//workerpool.Test()
	//workerpool.WaitGroupTest()
	//workerpool.WaitGroupTest2()
	//selecttest.SelectWithTimeout()
	//timers.TimerTest()
	//timers.TickerTest()

	/*
		ratelimiter.SimpleLimiter()
		ratelimiter.BurstyLimiter()
	*/

	/*
		concurrency.MutexTest()
		concurrency.AtomCounterTest()
		concurrency.StatefulGoRoutineTest()
	*/

	/*
		sorting.SortBuiltInTest()
		sorting.SortByFuncTest()
	*/

	/*
		misc.PrintEnvVars()
		misc.CmdLineArgs()
	*/
	misc.LogTest()

	//web.StartHttpServer()
}
