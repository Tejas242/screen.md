package tui

import (
	"github.com/atotto/clipboard"
)

type ClipboardManager struct {
	content string
}

func NewClipboardManager() *ClipboardManager {
	return &ClipboardManager{}
}

func (c *ClipboardManager) Copy(text string) error {
	return clipboard.WriteAll(text)
}

func (c *ClipboardManager) Paste() (string, error) {
	return clipboard.ReadAll()
}

func (c *ClipboardManager) Cut(text string) error {
	err := c.Copy(text)
	if err != nil {
		return err
	}
	c.content = text
	return nil
}
