package config

type Output struct {
	Name string `name`
	Type string `type`
}

type Outputs struct {
	OutputsArray []Output `outputs`
}
