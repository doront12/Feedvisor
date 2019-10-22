package DomainRequest

type (
	DomainRequest struct {
		Domain string `json:"domain"`
		Path   string `json:"path"`
	}
)

