package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/michaelfraser99/randomgeneratordiscordbot/structs"
	"github.com/michaelfraser99/randomgeneratordiscordbot/util"
)

var (
	err      error
	BotToken string
	Config   structs.Config
)

var session *discordgo.Session

func init() {
	flag.StringVar(&BotToken, "t", "", "Bot Token")
	flag.Parse()
	configJson := util.ReadJson("./data/config.json")

	json.Unmarshal(configJson, &Config)
}

//go run main.go -t ODU1NzcxMjIzNjY0NzU0NzM4.YM3VDw.5Q32nSycfX7FA1aByLv3Yh_CMqk

func main() {
	fmt.Printf("Starting bot with token: %v \n", BotToken)

	session, err = discordgo.New("Bot " + BotToken)

	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
		return
	}

	session.AddHandler(messageCreate)
	session.AddHandler(voiceStateUpdate)

	err = session.Open()

	if err != nil {
		log.Fatalf("Failed to open session: %v", err)
		return
	}

	fmt.Println("Bot is running, press CTRL+C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	session.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	var err error

	if m.Author.ID == s.State.User.ID || len(m.Content) < len(Config.CliKey) || m.Content[:len(Config.CliKey)] != Config.CliKey {
		return
	} else {
		message := util.ProcessArgument(m, Config)

		_, err = s.ChannelMessageSendComplex(m.ChannelID, message)

		if err != nil {
			log.Fatalf("Failed to send message: %v", err)
		}
	}
}

func voiceStateUpdate(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {

	if v.ChannelID != "" {

		//Exit if update orginated from same channel (i.e. mute & unmute)
		if v.BeforeUpdate != nil {
			if v.BeforeUpdate.ChannelID == v.ChannelID {
				return
			}
		}

		var message discordgo.MessageSend

		//Get username
		user, err := s.User(v.UserID)

		//If user is part of users to ignore from config
		for _, element := range Config.Greetings.AccountsIgnored {
			if element == user.Username {
				return
			}
		}

		if err != nil {
			log.Fatalf("Failed to get user")
		}

		//Get channel joined
		channel, err := s.Channel(v.ChannelID)

		if err != nil {
			log.Fatalf("Failed to get channel name")
		}

		message.Content = user.Username + " joined " + channel.Name + " by the server"
		message.TTS = true

		//Send to general text channel
		channels, _ := s.GuildChannels(v.GuildID)

		for _, c := range channels {
			if c.Type == discordgo.ChannelTypeGuildText {
				log.Println("Sending tts...")
				_, err = s.ChannelMessageSendComplex(c.ID, &message)

				if err != nil {
					log.Println(err)
					log.Fatalf("Failed to send user connected tts")
				}
				return
			}
		}
	}
}
