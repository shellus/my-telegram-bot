package router

import (
	"gopkg.in/telegram-bot-api.v4"
	"regexp"
	"github.com/pkg/errors"
	"fmt"
	"github.com/astaxie/beego/logs"
)

type Route struct {
	Pattern string
	Comment string
	Handle  func(update tgbotapi.Update)
}
type router struct {
	def    *Route
	routes []*Route
}

func NewRouter() (*router) {
	return &router{}
}

func (r *Route) SetComment(str string) (ro *Route) {
	r.Comment = str
	ro = r
	return
}
func (r *router) Add(expr string, handle func(update tgbotapi.Update)) (ro *Route) {
	ro = &Route{Pattern:expr, Handle:handle}
	r.routes = append(r.routes, ro)
	return
}
func (r *router) Default(handle func(update tgbotapi.Update)) {
	r.def = &Route{Pattern:"", Handle:handle}
}

func (r *router) Routes() (routes []*Route) {
	routes = r.routes
	return
}

func (r *router) Dispatch(uri string)(ro *Route, err error) {

	logs.Debug("router Dispatch start uri: %s", uri)
	// 正则匹配路由
	for _, i := range r.routes {
		var ip *regexp.Regexp
		ip, err = regexp.Compile(i.Pattern)
		if err != nil {
			return
		}
		if  ip.MatchString(uri) {
			logs.Debug("router Dispatch hit pattern: %s", i.Pattern)
			ro = i
			return
		}
	}

	// 404
	if r.def != nil {
		ro = r.def
		return
	}else {
		return nil, errors.New(fmt.Sprintf("no def handle uri: %v", uri))
	}
}
