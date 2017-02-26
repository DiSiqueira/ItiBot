package itibot

import (
	"github.com/mssola/user_agent"
	"net/http"
)

// BotMatcher Store the schema to match with the request Path
type BotMatcher struct {
	isBot bool
}

// New is the constructor to ItiBot
func New(isBot bool) *BotMatcher {
	return &BotMatcher{
		isBot: isBot,
	}
}

// Match returns if the request can be handled by this Route.
func (b *BotMatcher) Match(req *http.Request) bool {
	ua := user_agent.New(req.UserAgent())
	if ua.Bot() != b.isBot {
		return false
	}
	return true
}
