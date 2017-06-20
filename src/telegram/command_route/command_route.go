package command_route

import "gopkg.in/telegram-bot-api.v4"

type Route struct {
	pattern string
	handle func(update tgbotapi.Update)
}

var routes []Route
var texts []Route

func Command(pattern string, handle func(update tgbotapi.Update)){
	routes = append(routes, Route{pattern:pattern,handle:handle})
}
func Text(pattern string, handle func(update tgbotapi.Update)){
	texts = append(texts, Route{pattern:pattern,handle:handle})
}
func Handle(update tgbotapi.Update){
	if update.Message.IsCommand() {
		hit := false
		for _, r := range routes {
			if r.pattern == update.Message.Command() {
				r.handle(update)
				hit = true
				continue
			}
		}
		if hit == false {
			for _, r := range routes {
				if r.pattern == "default" {
					r.handle(update)
					hit = true
					continue
				}
			}
		}
	}else {
		hit := false
		for _, r := range texts {
			if r.pattern == update.Message.Text {
				r.handle(update)
				hit = true
				continue
			}
		}
		if hit == false {
			for _, r := range texts {
				if r.pattern == "default" {
					r.handle(update)
					hit = true
					continue
				}
			}
		}
	}
}

