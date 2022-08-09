package main

import "flag"

func initParams(cfg *config) {
	flag.StringVar(&cfg.RunAddress, "a", cfg.RunAddress, "RunAddress")
	flag.StringVar(&cfg.DataBaseURI, "d", cfg.DataBaseURI, "DataBaseURI")
	flag.Parse()
}
