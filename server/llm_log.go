package main

import (
	"fmt"

	"github.com/mattermost/mattermost-plugin-ai/server/ai"
	pluginapi "github.com/mattermost/mattermost-plugin-api"
)

type LanguageModelLogWrapper struct {
	log     pluginapi.LogService
	wrapped ai.LanguageModel
}

func NewLanguageModelLogWrapper(log pluginapi.LogService, wrapped ai.LanguageModel) *LanguageModelLogWrapper {
	return &LanguageModelLogWrapper{
		log:     log,
		wrapped: wrapped,
	}
}

func (w *LanguageModelLogWrapper) logInput(conversation ai.BotConversation, opts ...ai.LanguageModelOption) {
	prompt := fmt.Sprintf("\n%v", conversation)
	w.log.Info("LLM Call", "prompt", prompt)
}

func (w *LanguageModelLogWrapper) ChatCompletion(conversation ai.BotConversation, opts ...ai.LanguageModelOption) (*ai.TextStreamResult, error) {
	w.logInput(conversation, opts...)
	return w.wrapped.ChatCompletion(conversation, opts...)
}

func (w *LanguageModelLogWrapper) ChatCompletionNoStream(conversation ai.BotConversation, opts ...ai.LanguageModelOption) (string, error) {
	w.logInput(conversation, opts...)
	return w.wrapped.ChatCompletionNoStream(conversation, opts...)
}
