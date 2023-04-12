package ui

import (
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

const exec_color = "#ffa657"
const config_color = "#ffffff"
const chat_color = "#66b3ff"
const help_color = "#aaaaaa"
const error_color = "#cc3333"
const warning_color = "#ffcc00"
const success_color = "#46b946"

type Renderer struct {
	contentRenderer *glamour.TermRenderer
	successRenderer lipgloss.Style
	warningRenderer lipgloss.Style
	errorRenderer   lipgloss.Style
	helpRenderer    lipgloss.Style
}

func NewRenderer(options ...glamour.TermRendererOption) *Renderer {

	contentRenderer, err := glamour.NewTermRenderer(options...)
	if err != nil {
		return nil
	}

	successRenderer := lipgloss.NewStyle().Foreground(lipgloss.Color(success_color))
	warningRenderer := lipgloss.NewStyle().Foreground(lipgloss.Color(warning_color))
	errorRenderer := lipgloss.NewStyle().Foreground(lipgloss.Color(error_color))
	helpRenderer := lipgloss.NewStyle().Foreground(lipgloss.Color(help_color)).Italic(true)

	return &Renderer{
		contentRenderer: contentRenderer,
		successRenderer: successRenderer,
		warningRenderer: warningRenderer,
		errorRenderer:   errorRenderer,
		helpRenderer:    helpRenderer,
	}
}

func (r *Renderer) RenderContent(in string) string {
	out, _ := r.contentRenderer.Render(in)

	return out
}

func (r *Renderer) RenderSuccess(in string) string {
	return r.successRenderer.Render(in)
}

func (r *Renderer) RenderWarning(in string) string {
	return r.warningRenderer.Render(in)
}

func (r *Renderer) RenderError(in string) string {
	return r.errorRenderer.Render(in)
}

func (r *Renderer) RenderHelp(in string) string {
	return r.helpRenderer.Render(in)
}

func (r *Renderer) RenderConfigMessage() string {
	welcome := "**Yo**, welcome! 👋  \n\n"
	welcome += "I cannot find a configuration file, please enter an **OpenAI API key** "
	welcome += "from https://platform.openai.com/account/api-keys so I can generate it for you."

	return welcome
}