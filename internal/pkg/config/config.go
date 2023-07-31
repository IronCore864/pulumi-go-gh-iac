package config

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Organization struct {
	Org   string `yaml:"org"`
	Teams []Team `yaml:"teams"`
}

type Team struct {
	Name         string `yaml:"name"`
	Slug         string `yaml:"slug"`
	Description  string `yaml:"description"`
	ParentTeamId int
	Teams        []Team   `yaml:"teams"`
	Members      []Member `yaml:"members"`
}

type Member struct {
	UserName string `yaml:"username"`
	Role     string `yaml:"role"`
}

func LoadOrganizationConfig() (*Organization, error) {
	//import accurate team info from yaml
	orgFilePath, err := filepath.Abs("config.yaml")
	if err != nil {
		return nil, err
	}

	yamlFile, err := ioutil.ReadFile(orgFilePath)
	if err != nil {
		return nil, err
	}

	var org Organization
	err = yaml.Unmarshal(yamlFile, &org)
	if err != nil {
		return nil, err
	}

	return &org, nil
}
