package handlers

import (
	"github.com/bwmarrin/discordgo"
)

func CommuteHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == ";;출근" {
		goToWorkProcess()
		s.ChannelMessageSend(m.ChannelID, "")
	} else if m.Content == ";;퇴근" {
		getOffWorkProcess()
		s.ChannelMessageSend(m.ChannelID, "")
	}
}

func goToWorkProcess() {
}

func getOffWorkProcess() {
}
