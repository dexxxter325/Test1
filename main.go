package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	privet = "вечер добрый, босс. я могу переводить валюты (usd-rub,btc-usd)Для этого просто напишите 'btc/usd' ИЛИ 'usd/rub'"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6033448258:AAGJt_Tk4lan7MaciD-7nd8bOQDMwkUDSXs")
	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true //доп инфа

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		if update.Message.Text == "/help" {

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, privet)
			bot.Send(msg)
		}
		if update.Message.Text == "btc/usd" {
			resp, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=BTCUSDT")

			if err != nil {
				log.Println(err)
				continue
			}

			var priceResp struct {
				Price float64 `json:"price,string"`
			}
			/*json.NewDecoder принимает в качестве аргумента поток данных, из которого будут читаться JSON-данные
			resp.Body - поток данных, содержащих ответ от сервера
			Decode(&priceResp) декодирует JSON-данные из потока и сохраняет их в переменную, на которую указывает переданный указатель*/
			json.NewDecoder(resp.Body).Decode(&priceResp)       //парсим данные
			stringPrice := fmt.Sprintf("%.2f", priceResp.Price) //число в строку,.2-2 знака после запятой
			if err != nil {
				log.Println(err)
				continue
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, stringPrice)
			bot.Send(msg)

		}

		if update.Message.Text == "usd/rub" { // If we got a message

			resp, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=USDTRUB")

			if err != nil {
				log.Println(err)
				continue
			}

			var priceResp struct {
				Price float64 `json:"price,string"`
			}
			/*json.NewDecoder принимает в качестве аргумента поток данных, из которого будут читаться JSON-данные
			resp.Body - поток данных, содержащих ответ от сервера
			Decode(&priceResp) декодирует JSON-данные из потока и сохраняет их в переменную, на которую указывает переданный указатель*/
			json.NewDecoder(resp.Body).Decode(&priceResp)       //парсим данные
			stringPrice := fmt.Sprintf("%.2f", priceResp.Price) //число в строку,.2-2 знака после запятой
			if err != nil {
				log.Println(err)
				continue
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, stringPrice)
			bot.Send(msg)

		}
	}
}
