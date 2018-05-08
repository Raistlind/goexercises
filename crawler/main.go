package main

import (
	"GolandProjects/goexercises/crawler/engine"
	"GolandProjects/goexercises/crawler/scheduler"
	"GolandProjects/goexercises/crawler/zhenai/parser"
	"GolandProjects/goexercises/crawler/persist"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 1,
		ItemChan:    persist.ItemSaver(),
	}

	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})
}
