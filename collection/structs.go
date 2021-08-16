package collection

type Collection struct {
	Name        string   `json:"name,omitempty"`
	AccountIDs  []string `json:"accountIDs,omitempty"`
	AppIDs      []string `json:"appIDs,omitempty"`
	Clusters    []string `json:"clusters,omitempty"`
	CodeRepos   []string `json:"codeRepos,omitempty"`
	Color       string   `json:"color,omitempty"`
	Containers  []string `json:"containers,omitempty"`
	Description string   `json:"description,omitempty"`
	Functions   []string `json:"functions,omitempty"`
	Hosts       []string `json:"hosts,omitempty"`
	Images      []string `json:"images,omitempty"`
	Labels      []string `json:"labels,omitempty"`
	Modified    string   `json:"modified,omitempty"`
	Namespaces  []string `json:"namespaces,omitempty"`
	Owner       string   `json:"owner,omitempty"`
	Prisma      bool     `json:"prisma,omitempty"`
	System      bool     `json:"system,omitempty"`
}
