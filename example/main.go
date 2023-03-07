package main

import (
	"context"
	"log"
	"net/http"

	"github.com/caarlos0/env/v7"
	vg "github.com/lll-lll-lll-lll/valorant-go"
)

func main() {
	summonerName := ""
	client := &http.Client{}
	cfg, err := vg.NewValConfig(vg.JP1)
	if err != nil {
		log.Fatal(err)
	}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("config env error : %v", err)
	}
	valorant := vg.New(client, cfg)
	res, err := valorant.DoToSummoner(context.Background(), summonerName)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}
