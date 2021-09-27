package serialization

type Person struct {
	Name           string          `yaml:"name"`
	Lastname       string          `yaml:"lastname"`
	Age            int             `yaml:"age"`
	Hobbies        []Hobby         `yaml:"hobbies"`
	SocialNetworks []SocialNetwork `yaml:"social-networks"`
}

type Hobby struct {
	Name string `yaml:"name"`
	Year string `yaml:"year"`
}

type SocialNetwork struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
}
