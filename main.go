package main

import (
	"fmt"
	"github.com/aid95/zaksim-discord-bot/bot/handlers"
	"github.com/aid95/zaksim-discord-bot/db"
	"github.com/aid95/zaksim-discord-bot/utils/gen/message"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Database 연결.
	if err := db.Open(); err != nil {
		log.Fatal(message.Error("Failed to connect to database. "), err)
	}
	defer db.Close()

	// DiscordGo를 위한 준비 및 실행
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		log.Fatal(message.Error("Failed creating Discord session. "), err)
	}
	defer dg.Close()

	//-- 봇 핸들러 등록
	dg.AddHandler(handlers.CommuteHandler)

	//-- 봇 연결
	if err = dg.Open(); err != nil {
		log.Fatal(message.Error("Failed opening connection. "), err)
	}

	fmt.Println(message.Info("🚀 Zaksim bot is now running, Press CTRL-C to exit."))
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
