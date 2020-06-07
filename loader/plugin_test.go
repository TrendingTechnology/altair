package loader_test

import (
	"testing"

	"github.com/codefluence-x/altair/loader"
	"github.com/stretchr/testify/assert"
)

func TestPlugin(t *testing.T) {

	t.Run("Compile", func(t *testing.T) {
		t.Run("Given plugin path", func(t *testing.T) {
			t.Run("Normal scenario", func(t *testing.T) {
				t.Run("Return map string of entity.Plugin", func(t *testing.T) {
					pluginPath := "./normal_scenario_plugin_path/"

					generateTempTestFiles(pluginPath, PluginConfigNormal1, "oauth.yaml", 0666)
					generateTempTestFiles(pluginPath, PluginConfigNormal2, "cache.yaml", 0666)

					pluginList, err := loader.Plugin().Compile(pluginPath)

					assert.Nil(t, err)
					assert.Equal(t, 2, len(pluginList))
					assert.Contains(t, pluginList, "oauth")
					assert.Contains(t, pluginList, "cache")

					removeTempTestFiles(pluginPath)
				})
			})

			t.Run("Plugin already defined", func(t *testing.T) {
				t.Run("Return error", func(t *testing.T) {
					pluginPath := "./plugin_already_defined/"

					generateTempTestFiles(pluginPath, PluginConfigNormal1, "oauth.yaml", 0666)
					generateTempTestFiles(pluginPath, PluginConfigNormal1, "oauth_2.yaml", 0666)

					pluginList, err := loader.Plugin().Compile(pluginPath)

					assert.NotNil(t, err)
					assert.Equal(t, 1, len(pluginList))

					removeTempTestFiles(pluginPath)
				})
			})

			t.Run("Yaml unmarshal error", func(t *testing.T) {
				t.Run("Return error", func(t *testing.T) {
					pluginPath := "./plugin_config_yaml_unmarshal_error/"

					generateTempTestFiles(pluginPath, PluginConfigYamlUnmarshalError, "oauth.yaml", 0666)

					pluginList, err := loader.Plugin().Compile(pluginPath)

					assert.NotNil(t, err)
					assert.Equal(t, 0, len(pluginList))

					removeTempTestFiles(pluginPath)
				})
			})

			t.Run("Template parsing error error", func(t *testing.T) {
				t.Run("Return error", func(t *testing.T) {
					pluginPath := "./plugin_config_template_parsing_error/"

					generateTempTestFiles(pluginPath, PluginConfigTemplateParsingError, "oauth.yaml", 0666)

					pluginList, err := loader.Plugin().Compile(pluginPath)

					assert.NotNil(t, err)
					assert.Equal(t, 0, len(pluginList))

					removeTempTestFiles(pluginPath)
				})
			})

			t.Run("Dir is not exists", func(t *testing.T) {
				t.Run("Return error", func(t *testing.T) {
					pluginPath := "./plugin_config_dir_not_exists/"

					pluginList, err := loader.Plugin().Compile(pluginPath)

					assert.NotNil(t, err)
					assert.Equal(t, 0, len(pluginList))
				})
			})
		})
	})
}
