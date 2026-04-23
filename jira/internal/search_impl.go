package internal

import (
	"fmt"
	model "github.com/pandorian-ai/go-atlassian-fork/v2/pkg/infra/models"
	"github.com/pandorian-ai/go-atlassian-fork/v2/service"
)

// NewSearchService creates a new instance of SearchADFService and SearchRichTextService.
func NewSearchService(client service.Connector, version string) (*SearchADFService, *SearchRichTextService, error) {

	if version == "" {
		return nil, nil, fmt.Errorf("jira: %w", model.ErrNoVersionProvided)
	}

	rtService := &SearchRichTextService{
		internalClient: &internalSearchRichTextImpl{
			c:       client,
			version: version,
		},
	}

	adfService := &SearchADFService{
		internalClient: &internalSearchADFImpl{
			c:       client,
			version: version,
		},
	}

	return adfService, rtService, nil
}
