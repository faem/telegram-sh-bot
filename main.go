package main

import (
	"flag"
	"fmt"
	"image/png"
	"log"
	"os"
	"strings"
	"time"

	"github.com/codeskyblue/go-sh"
	"github.com/kbinani/screenshot"
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	admins     []string
	token      string
	pollTime   int64
	debug      bool
	filePrefix = "screenshots"
)

func init() {
	var tmp string
	flag.StringVar(&tmp, "admin", "fahim_abrar,Cauef", "Comma separated usernames of the admins")
	flag.StringVar(&token, "token", os.Getenv("TELEGRAM_BOT_TOKEN"), "Token of your bot")
	flag.Int64Var(&pollTime, "poll_time", 100, "Response time of bot")
	flag.BoolVar(&debug, "debug", false, "Print error info to debug")

	flag.Parse()
	for _, v := range strings.Split(tmp, ",") {
		admins = append(admins, v)
	}
	// fmt.Printf("%+v\n", admins)
	if len(os.Args) < 2 {
		flag.Usage()
	}
}

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: time.Duration(pollTime) * time.Millisecond},
		Reporter: func(err error) {
			if debug {
				log.Println(err)
			}
		},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(m *tb.Message) {
		log.Println(m.Sender.Username + ": " + m.Text)

		_, err := b.Send(m.Chat, `***Available Commands***:

- ***/hello***: simply greets the user (anyone can run it)
- ***/getss***: takes screenshots of the pc and sends it to the user (only admin can run it)
- ***/sh <valid-shell-command>***: runs the command where the bot is running and sends any output returned by the command or error (only admin can run it)`,
			&tb.SendOptions{
				ParseMode: "Markdown",
			})
		if err != nil {
			log.Fatal(err)
		}
	})

	b.Handle("/hello", func(m *tb.Message) {
		log.Println("Hi, " + m.Sender.FirstName + " " + m.Sender.LastName + "!")
		_, err := b.Send(m.Chat, "Hi, "+m.Sender.FirstName+" "+m.Sender.LastName+"!")
		if err != nil {
			log.Fatal(err)
		}
	})

	b.Handle("/sh", func(m *tb.Message) {
		log.Println(m.Sender.Username + ": " + m.Text)

		if !isAdmin(m.Sender.Username) {
			_, err := b.Send(m.Chat, "Only `"+fmt.Sprintf("%v", admins)+"` are authorized to run /sh command! You can run /hello", &tb.SendOptions{ParseMode: "Markdown"})
			if err != nil {
				log.Println(err)
			}
			return
		}

		stringArray := strings.Split(m.Text, " ")
		if len(stringArray) == 1 {
			_, err := b.Send(m.Chat, "Please Provide Command to run!")
			if err != nil {
				log.Fatal(err)
			}
			return
		}

		cmd := stringArray[1]
		args := make([]string, 0)

		if len(stringArray) > 2 {
			args = append(stringArray[2:])
		}

		resp, err := sh.Command(cmd, args).Output()
		if err != nil {
			log.Println(string(resp))
			_, err := b.Send(m.Chat, ""+err.Error())

			if err != nil {
				log.Fatal(err)
			}
			return
		}
		_, err = b.Send(m.Chat, "```\n"+string(resp)+"```", &tb.SendOptions{ParseMode: "Markdown"})
		if err != nil {
			log.Fatal(err)
		}
	})

	b.Handle("/getss", func(m *tb.Message) {
		log.Println(m.Sender.Username + ": " + m.Text)
		if !isAdmin(m.Sender.Username) {
			_, err := b.Send(m.Chat, "Only `"+fmt.Sprintf("%v", admins)+"` are authorized to run /getss command! You can run /hello 	\xF0\x9F\x98\x82", &tb.SendOptions{ParseMode: "Markdown"})
			if err != nil {
				log.Println(err)
			}
			return
		}

		_, err := b.Send(m.Chat, "Uploading ScreenShots Please Wait, "+m.Sender.FirstName)
		if err != nil {
			log.Println(err)
		}

		fileNames, err := GetScreenShots()
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < len(fileNames); i++ {
			ss := &tb.Photo{File: tb.FromDisk(fileNames[i])}

			_, err = b.SendAlbum(m.Sender, tb.Album{ss})
			if err != nil {
				log.Fatal(err)
			}
		}
	})

	log.Println("Bot is running. . . . .")
	b.Start()
}

func GetScreenShots() ([]string, error) {
	n := screenshot.NumActiveDisplays()
	var fileNames = make([]string, 0)

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			return []string{}, err
		}

		fileNames = append(fileNames, fmt.Sprintf("%s_%d.png", filePrefix, i))
		file, _ := os.Create(fileNames[i])
		defer file.Close()
		err = png.Encode(file, img)
		if err != nil {
			return []string{}, err
		}
	}

	return fileNames, nil
}

func isAdmin(username string) bool {
	for _, v := range admins {
		if username == v {
			return true
		}
	}

	return false
}
