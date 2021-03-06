package main

import (
	"flag"
	"fmt"
	"github.com/aid95/zaksim-discord-bot/bot/handlers"
	"github.com/aid95/zaksim-discord-bot/db"
	"github.com/aid95/zaksim-discord-bot/utils/gen/message"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Discord bot token")
}

func main() {
	// Database 연결.
	if err := db.Connect(); err != nil {
		fmt.Println(message.Error("💀 Failed to connect to database."))
		return
	}
	defer db.Close()

	// DiscordGo를 위한 준비 및 실행
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println(message.Error("💀 Failed creating Discord session."))
		return
	}
	defer dg.Close()

	//-- 봇 핸들러 등록
	dg.AddHandler(handlers.CommuteHandler)

	//-- 봇 연결
	if err = dg.Open(); err != nil {
		fmt.Println(message.Error("💀 Failed opening connection."))
		return
	}

	fmt.Println(message.Info("🚀 Zaksim bot is now running, Press CTRL-C to exit."))
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
