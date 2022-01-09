package yamlParser

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"

	"github.com/emirpasic/gods/maps/linkedhashmap"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

func parse(dirName string) error {
	yamlData, err := loadYamlData(dirName)
	if err != nil {
		fmt.Printf("%v", err)
		return err
	}
	process("service.yaml", yamlData)
	return nil
}

func process(targetFile string, yamlData map[string]*yamlData) (map[string]interface{}, error) {
	r := findIncludeFile(targetFile, yamlData)
	return merge(r), nil
}

// merge all depends yaml file with override value
func merge(data *linkedhashmap.Map) map[string]interface{} {
	yamlDataList := make([]interface{}, 0)
	it := data.Iterator()
	for it.Next() {
		yamlDataList = append(yamlDataList, it.Value().(*yamlData).data)
	}
	return merge0(yamlDataList, 0)
}

func merge0(values []interface{}, pos int) map[string]interface{} {
	if pos >= len(values) {
		return make(map[string]interface{})
	}
	to := merge0(values, pos+1)
	from := values[pos]
	for k, v := range from.(map[string]interface{}) {
		to[k] = v
	}
	return to
}

func findIncludeFile(targetFileName string, fileToData map[string]*yamlData) *linkedhashmap.Map {
	var result = linkedhashmap.New()
	targetFile := fileToData[targetFileName]
	findInclude0(targetFile, result, 0, fileToData)
	return result
}

func loadYamlData(dirName string) (map[string]*yamlData, error) {
	result := make(map[string]*yamlData)
	return result, filepath.Walk(dirName, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		fileName := info.Name()
		marshalData := make(map[string]interface{})
		file, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(file, marshalData)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("error marshal file %s", file))
		}
		result[fileName] = createYamlData(fileName, marshalData)
		return nil
	})
}

// findInclude0 find all included file recursively
func findInclude0(current *yamlData, result *linkedhashmap.Map, position int, fileToData map[string]*yamlData) {
	if _, ok := result.Get(current.filename); !ok {
		result.Put(current.filename, fileToData[current.filename])
	}
	if len(current.meta.include) != 0 {
		for _, v := range current.meta.include {
			if _, ok := result.Get(v); ok {
				continue
			}
			nextPos := position + 1
			findInclude0(fileToData[v], result, nextPos, fileToData)
		}
	}
	return
}
