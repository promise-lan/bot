package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/promise-lan/bot/events"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var token = ""

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Failed loading .env file, attempting to fallback to host environmental variables...")
	}
	token = os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("Env variable TOKEN not found!")
	}
	tmp := os.Getenv("PREFIX")
	if tmp == "" {
		log.Fatal("Env variable PREFIX not found!")
	}
}

func main() {
	client, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Error creating discord session: %v", err)
	}
	client.AddHandler(events.Ready)
	client.AddHandler(events.Message)
	client.AddHandler(events.ReactionAdd)
	client.AddHandler(events.ReactionRemove)
	err = client.Open()
	if err != nil {
		log.Fatalf("Error opening discord session: %v", err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-c
	_ = client.Close()
}
