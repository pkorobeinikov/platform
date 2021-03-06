package deployment

import (
	"bytes"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"

	"github.com/pkorobeinikov/platform/platform-lib/service/env"
)

func (d *dockerComposeGeneratorV2) Generate(request SpecGenerationRequest) (SpecGenerationResponse, error) {
	var (
		response SpecGenerationResponse
		spec     dockerComposeSpecV2

		b bytes.Buffer
	)

	spec.Services = make(map[string]dockerComposeServiceV2)

	for _, platformComponent := range request.PlatformComponentList {
		dcsList, err := platformComponent.dockerComposeServiceSpecList(request)
		if err != nil {
			return response, err
		}

		for _, dcs := range dcsList {
			spec.Services[dcs.ContainerName] = dcs
		}
	}

	for _, serviceComponent := range request.ServiceComponentList {
		dcsList, err := serviceComponent.dockerComposeServiceSpecList(request)
		if err != nil {
			return response, err
		}

		for _, dcs := range dcsList {
			spec.Services[dcs.ContainerName] = dcs
		}
	}

	ye := yaml.NewEncoder(&b)
	ye.SetIndent(2)
	if err := ye.Encode(&spec); err != nil {
		return response, err
	}

	env.Registry().
		RegisterMany(request.Environment).
		Register("SERVICE", request.ServiceName)

	envFileContent, err := godotenv.Marshal(env.Registry().All())
	if err != nil {
		return response, err
	}

	response.FileList = map[string]string{
		DockerComposeFile: b.String(),
		env.File:          envFileContent,
	}

	return response, nil
}

func NewDockerComposeGeneratorV2() *dockerComposeGeneratorV2 {
	return &dockerComposeGeneratorV2{}
}

type (
	dockerComposeGeneratorV2 struct {
	}

	dockerComposeSpecV2 struct {
		Services map[string]dockerComposeServiceV2 `yaml:"services"`
	}

	dockerComposeServiceV2 struct {
		ContainerName string            `yaml:"container_name"`
		Image         string            `yaml:"image"`
		Restart       string            `yaml:"restart,omitempty"`
		DependsOn     []string          `yaml:"depends_on,omitempty"`
		Ports         []string          `yaml:"ports,omitempty"`
		Environment   map[string]string `yaml:"environment,omitempty"`
		CapAdd        []string          `yaml:"cap_add,omitempty"`
		Command       string            `yaml:"command,omitempty"`
	}
)
