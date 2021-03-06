package handlers

import (
	"fmt"
	"github.com/aid95/zaksim-discord-bot/db"
	"github.com/bwmarrin/discordgo"
	"time"
)

func CommuteHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == ";;출근" {
		goToWorkProcess(s, m)
	} else if m.Content == ";;퇴근" {
		getOffWorkProcess(s, m)
	}
}

func startCommute(channel string, author string) error {
	conn := db.Instance()
	if _, err := conn.Client.Commute.
		Create().
		SetChannelID(channel).
		SetAuthorID(author).
		Save(conn.Ctx);
		err != nil {
		return err
	}
	return nil
}

func endCommute(channel string, author string) error {
	conn := db.Instance()
	if _, err := conn.Client.Commute.
		Update().
		Where().
		SetGetOffAt(time.Now()).
		Save(conn.Ctx);
		err != nil {
		return err
	}
	return nil
}

func goToWorkProcess(s *discordgo.Session, m *discordgo.MessageCreate) {
	startCommute(m.ChannelID, m.Author.ID)
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s님 출근! 오늘도 화이팅입니다! 👋", m.Author.Username))
}

func getOffWorkProcess(s *discordgo.Session, m *discordgo.MessageCreate) {
	endCommute(m.ChannelID, m.Author.ID)
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s님 퇴근! 오늘 하루도 고생하셨습니다. 🚌", m.Author.Username))
}
