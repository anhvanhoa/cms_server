package mailtemplate

import "github.com/cbroglie/mustache"

type MailTemplate interface {
	RenderWithLayout(layout, subject, body string, data map[string]any) (*Result, error)
	Render(subject, body string, data map[string]any) (*Result, error)
	RenderLayoutFile(fileLayout, subject, body string, data map[string]any) (*Result, error)
}

type Result struct {
	Subject string
	Body    string
}

type mailTemplate struct{}

func NewMailTemplate() MailTemplate {
	return &mailTemplate{}
}

func (m *mailTemplate) RenderWithLayout(layout, subject, body string, data map[string]any) (*Result, error) {
	var err error
	var result Result
	if subject, err = mustache.Render(subject, data); err != nil {
		return &result, err
	} else if body, err = mustache.RenderInLayout(body, layout, data); err != nil {
		return &result, err
	}
	result.Subject = subject
	result.Body = body
	return &result, nil
}

func (m *mailTemplate) Render(subject, body string, data map[string]any) (*Result, error) {
	var err error
	var result Result
	if subject, err = mustache.Render(subject, data); err != nil {
		return &result, err
	} else if body, err = mustache.Render(body, data); err != nil {
		return &result, err
	}
	result.Subject = subject
	result.Body = body
	return &result, nil
}

func (m *mailTemplate) RenderLayoutFile(fileLayout, subject, body string, data map[string]any) (*Result, error) {
	var result Result
	if layout, err := mustache.ParseFile(fileLayout); err != nil {
		return &result, err
	} else if subject, err = mustache.Render(subject, data); err != nil {
		return &result, err
	} else if renderedBody, err := layout.Render(data); err != nil {
		return &result, err
	} else {
		result.Subject = subject
		result.Body = renderedBody
	}
	return &result, nil
}
