package dtos

import "errors"

type ShortnerRequest struct {
	Url          string `json:"url"`
	AutoRedirect bool   `json:"auto_redirect"`
}

func (s *ShortnerRequest) Validate() error {
	if s.Url == "" {
		return errors.New("url is required")
	}

	return nil
}
