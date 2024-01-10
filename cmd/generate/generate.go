package generate

import (
	"errors"
	"fmt"
	"os"
	"path"
	"protocomp/cmd/config"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var GenerateCmd = &cobra.Command{
	Use: "generate",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(config.Logo)
		return generateCompFile()
	},
}

func generateCompFile() error {
	config.InfoLog.Printf("Generating comp file, name=%s.yaml, path=%s, dependencies=%v", config.CompFileName, config.CompFilePath, config.DependencyProtos)
	deps, err := parseDependencies()
	if err != nil {
		return err
	}
	return generateYamlFile(deps)
}

func parseDependencies() (map[string]interface{}, error) {
	res := map[string]interface{}{}
	for _, command := range config.DependencyProtos {
		var deps []string
		values := strings.Split(command, ":")
		parentPath := values[0]
		if len(values) > 1 {
			if err := checkIfPathExist(parentPath); err != nil {
				return nil, err
			}
			depPaths := strings.Split(values[1], ",")
			for _, elPath := range depPaths {
				if err := checkIfPathExist(elPath); err != nil {
					return nil, err
				}
				name := path.Base(elPath)
				deps = append(deps, name[:len(name)-6])
			}
		}
		parentName := path.Base(parentPath)
		data := map[string]interface{}{
			"path":         parentPath,
			"dependencies": deps,
		}
		res[parentName[:len(parentName)-6]] = data
	}
	return res, nil
}

func checkIfPathExist(path string) error {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		config.ErrorLog.Printf("File doesn't exist!, file_path=%s", path)
		return err
	}
	if err := validatePath(path); err != nil {
		return err
	}
	return nil
}

func validatePath(elPath string) error {
	if path.Ext(elPath) != ".proto" {
		config.ErrorLog.Printf("Wrong file extension! Extendion should be .go, provided_path=%s", elPath)
		return errors.New("wrong file extension")
	}
	return nil
}

func generateYamlFile(deps map[string]interface{}) error {
	if config.CompFilePath[len(config.CompFilePath)-1] != '/' {
		config.CompFilePath = fmt.Sprintf("%s/", config.CompFilePath)
	}
	filePath := fmt.Sprintf("%s%s.yaml", config.CompFilePath, config.CompFileName)
	fmt.Println(filePath)
	f, err := os.Create(filePath)
	defer f.Close()
	if err != nil {
		config.ErrorLog.Printf("Couldn't open comp file! Details: %v", err)
		return err
	}
	yaml.NewEncoder(f).Encode(deps)
	return nil
}

func init() {
	GenerateCmd.PersistentFlags().StringVarP(&config.CompFileName, "name", "n", "comp", "name of output comp file")
	GenerateCmd.PersistentFlags().StringVarP(&config.CompFilePath, "path", "p", ".", "path where output comp file will be generated")
	GenerateCmd.PersistentFlags().StringArrayVar(&config.DependencyProtos, "dep", []string{}, "All the proto dependecies."+
		" Input format: Path_to_parent: path_to_child,path_to_child Path_to_parent: path_to_child,path_to_child")
	GenerateCmd.MarkPersistentFlagRequired("dep")
}
