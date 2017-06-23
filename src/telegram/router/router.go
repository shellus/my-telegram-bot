package router

import "gopkg.in/telegram-bot-api.v4"

type route struct {
	Pattern string
	Comment string
	Handle  func(update tgbotapi.Update)
}
type router struct {
	def    *route
	routes []*route
}

func NewRouter() (r *router) {
	return &router{}
}

func (r *route) SetComment(str string) (ro *route) {
	r.Comment = str
	ro = r
	return
}
func (r *router) Add(pattern string, handle func(update tgbotapi.Update)) (ro *route) {
	ro = &route{Pattern:pattern, Handle:handle}
	r.routes = append(r.routes, ro)
	return
}
func (r *router) Default(handle func(update tgbotapi.Update)) {
	r.def = &route{Pattern:"", Handle:handle}
}

func (r *router) Routes() (routes []*route) {
	routes = r.routes
	return
}

func (r *router) Handle(update tgbotapi.Update) {
	for _, i := range r.routes {
		if i.Pattern == update.Message.Text {
			i.Handle(update)
			return
		}
	}
	r.def.Handle(update)
}
