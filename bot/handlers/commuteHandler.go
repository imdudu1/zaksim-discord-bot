package handlers

import (
	"fmt"
	"github.com/aid95/zaksim-discord-bot/db"
	"github.com/aid95/zaksim-discord-bot/ent/commute"
	"github.com/aid95/zaksim-discord-bot/utils/gen/message"
	"github.com/aid95/zaksim-discord-bot/utils/tiktok"
	"github.com/bwmarrin/discordgo"
)

func CommuteHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == ";;출근" {
		goToWorkProcess(s, m)
	} else if m.Content == ";;퇴근" {
		getOffWorkProcess(s, m)
	}
}

func isExistTodayCommuteData(channel string, author string) bool {
	conn := db.Instance()
	start, end, _ := tiktok.Today("Asia/Seoul")
	exist, err := conn.Client.Commute.Query().
		Where(
			commute.And(
				commute.ChannelIDEQ(channel),
				commute.AuthorIDEQ(author),
				commute.And(commute.GoToWorkAtGTE(start), commute.GoToWorkAtLT(end)))).
		Exist(conn.Ctx)
	if err != nil {
		return true
	} else if exist {
		return true
	}
	return false
}

func startCommute(channel string, author string) (bool, error) {
	// 출근 요청 시
	conn := db.Instance()
	//-- 1. 이미 생성된 금일 출근 기록이 있는지 검사한다.
	if exist := isExistTodayCommuteData(channel, author); exist {
		return false, nil
	}
	//-- 2. 금일 출근 기록이 없으면 생성
	if _, err := conn.Client.Commute.
		Create().
		SetChannelID(channel).
		SetAuthorID(author).
		Save(conn.Ctx);
		err != nil {
		return false, err
	}
	return true, nil
}

func endCommute(channel string, author string) (bool, error) {
	// 퇴근 요청 시
	conn := db.Instance()
	//-- 1. 이미 생성된 금일 출근 기록이 있는지 검사한다.
	if exist := isExistTodayCommuteData(channel, author); !exist {
		return false, nil
	}
	//-- 2. 오늘 출근 기록에서 퇴근 시간을 갱신한다.
	start, end, _ := tiktok.Today("Asia/Seoul")
	if _, err := conn.Client.Commute.
		Update().
		Where(
			commute.And(
				commute.ChannelIDEQ(channel),
				commute.AuthorIDEQ(author),
				commute.And(commute.GoToWorkAtGTE(start), commute.GoToWorkAtLT(end)))).
		SetGetOffAt(tiktok.LocNow("Asia/Seoul")).
		Save(conn.Ctx);
		err != nil {
		return false, err
	}
	return true, nil
}

func goToWorkProcess(s *discordgo.Session, m *discordgo.MessageCreate) {
	if complete, err := startCommute(m.ChannelID, m.Author.ID); err != nil {
		fmt.Println(message.Error(""), err)
	} else {
		if complete {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s님 출근! 오늘도 화이팅입니다! 👋", m.Author.Username))
		} else {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s님은 이미 출근 중 입니다. 😅", m.Author.Username))
		}
	}
}

func getOffWorkProcess(s *discordgo.Session, m *discordgo.MessageCreate) {
	if complate, err := endCommute(m.ChannelID, m.Author.ID); err != nil {
		fmt.Println(message.Error(""), err)
	} else {
		if complate {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s님 퇴근! 오늘 하루도 고생하셨습니다. 🚌", m.Author.Username))
		} else {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s님의 금일 출근 기록이 없습니다. 😅", m.Author.Username))
		}
	}
}
