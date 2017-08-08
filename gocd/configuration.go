package gocd

import (
	"context"
)

// ConfigurationService describes the HAL _link resource for the api response object for a pipelineconfig
type ConfigurationService service

type ConfigXML struct {
	Repositories       []ConfigMaterialRepository `xml:"repositories>repository"`
	Server             ConfigServer               `xml:"server"`
	SCMS               []ConfigSCM                `xml:"scms>scm"`
	ConfigRepositories []ConfigRepository         `xml:"config-repos>config-repo"`
	Pipelines          []ConfigPipeline           `xml:"pipelines>pipeline"`
}

type ConfigPipeline struct {
	Name                 string                      `xml:"name,attr"`
	LabelTemplate        string                      `xml:"labeltemplate,attr"`
	Params               []ConfigParam               `xml:"params>param"`
	GitMaterials         []GitRepositoryMaterial     `xml:"materials>git,omitempty"`
	PipelineMaterials    []PipelineMaterial          `xml:"materials>pipeline,omitempty"`
	Timer                string                      `xml:"timer"`
	EnvironmentVariables []ConfigEnvironmentVariable `xml:"environmentvariables>variable"`
	Stages               []ConfigStage               `xml:"stage"`
}

type ConfigTasks struct {
	Exec []ConfigExecTask `xml:"exec"`
}

type ConfigExecTask struct {
	Command    string `xml:"command,attr"`
	WorkingDir string `xml:"workingdir"`
}

type ConfigStage struct {
	Name     string         `xml:"name,attr"`
	Approval ConfigApproval `xml:"approval"`
	Jobs     []ConfigJob    `xml:"jobs>job"`
}

type ConfigJob struct {
	Name                 string                      `xml:"name,attr"`
	EnvironmentVariables []ConfigEnvironmentVariable `xml:"environmentvariables>variable"`
	Tasks                ConfigTasks                 `xml:"tasks"`
}

type ConfigApproval struct {
	Type string `xml:"type,attr"`
}

type ConfigEnvironmentVariable struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value"`
}

type PipelineMaterial struct {
	Name         string `xml:"pipelineName,attr"`
	StageName    string `xml:"stageName,attr"`
	MaterialName string `xml:"materialName,attr"`
}

type GitRepositoryMaterial struct {
	URL    string         `xml:"url,attr"`
	Filter []ConfigFilter `xml:"filter>ignore"`
}

type ConfigFilter struct {
	Ignore string `xml:"pattern,attr"`
}

type ConfigParam struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

type ConfigRepository struct {
	Plugin string              `xml:"plugin,attr"`
	ID     string              `xml:"id,attr"`
	Git    ConfigRepositoryGit `xml:"git"`
}

type ConfigRepositoryGit struct {
	URL string `xml:"url,attr"`
}

type ConfigSCM struct {
	ID                  string                    `xml:"id,attr"`
	Name                string                    `xml:"name,attr"`
	PluginConfiguration ConfigPluginConfiguration `xml:"pluginConfiguration"`
	Configuration       []ConfigProperty          `xml:"configuration>property"`
}

type ConfigMaterialRepository struct {
	ID                  string                    `xml:"id,attr"`
	Name                string                    `xml:"name,attr"`
	PluginConfiguration ConfigPluginConfiguration `xml:"pluginConfiguration"`
	Configuration       []ConfigProperty          `xml:"configuration>property"`
	Packages            []ConfigPackage           `xml:"packages>package"`
}

type ConfigPackage struct {
	ID            string           `xml:"id,attr"`
	Name          string           `xml:"name,attr"`
	Configuration []ConfigProperty `xml:"configuration>property"`
}

type ConfigPluginConfiguration struct {
	ID      string `xml:"id,attr"`
	Version string `xml:"version,attr"`
}

type ConfigServer struct {
	MailHost                  MailHost       `xml:"mailhost"`
	Security                  ConfigSecurity `xml:"security"`
	Elastic                   ConfigElastic  `xml:"elastic"`
	ArtifactsDir              string         `xml:"artifactsdir,attr"`
	SiteUrl                   string         `xml:"siteUrl,attr"`
	SecureSiteUrl             string         `xml:"secureSiteUrl,attr"`
	PurgeStart                string         `xml:"purgeStart,attr"`
	PurgeUpTo                 string         `xml:"purgeUpto,attr"`
	JobTimeout                int            `xml:"jobTimeout,attr"`
	AgentAutoRegisterKey      string         `xml:"agentAutoRegisterKey,attr"`
	WebhookSecret             string         `xml:"webhookSecret,attr"`
	CommandRepositoryLocation string         `xml:"commandRepositoryLocation,attr"`
	ServerId                  string         `xml:"serverId,attr"`
}

type MailHost struct {
	Hostname string `xml:"hostname,attr"`
	Port     int    `xml:"port,attr"`
	TLS      bool   `xml:"tls,attr"`
	From     string `xml:"from,attr"`
	Admin    string `xml:"admin,attr"`
}

type ConfigSecurity struct {
	AuthConfigs []ConfigAuthConfig `xml:"authConfigs>authConfig"`
	Roles       []ConfigRole       `xml:"roles>role"`
	Admins      []string           `xml:"admins>user"`
}

type ConfigRole struct {
	Name  string   `xml:"name,attr"`
	Users []string `xml:"users>user"`
}

type ConfigAuthConfig struct {
	ID         string           `xml:"id,attr"`
	PluginId   string           `xml:"pluginId,attr"`
	Properties []ConfigProperty `xml:"property"`
}

type ConfigElastic struct {
	Profiles []ConfigElasticProfile `xml:"profiles>profile"`
}

type ConfigElasticProfile struct {
	ID         string           `xml:"id,attr"`
	PluginID   string           `xml:"pluginId,attr"`
	Properties []ConfigProperty `xml:"property"`
}

type ConfigProperty struct {
	Key   string `xml:"key"`
	Value string `xml:"value"`
}

// Get will retrieve all agents, their status, and metadata from the GoCD Server.
// Get returns a list of pipeline instanves describing the pipeline history.
func (cd *ConfigurationService) Get(ctx context.Context) (*ConfigXML, *APIResponse, error) {
	req, err := cd.client.NewRequest("GET", "admin/config.xml", nil, "")
	if err != nil {
		return nil, nil, err
	}

	cx := ConfigXML{}
	resp, err := cd.client.Do(ctx, req, &cx, responseTypeXML)
	if err != nil {
		return nil, resp, err
	}

	return &cx, resp, nil
}
