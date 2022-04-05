package response

type Configurations struct {
	Configs []ConfigurationItem `json:"configs"`
}

type ConfigurationItem struct {
	Filename string `json:"filename"`
	Body     string `json:"body"`
}
