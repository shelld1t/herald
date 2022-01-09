package yamlParser

type yamlData struct {
	filename string
	meta     *meta
	data     map[string]interface{}
}

type meta struct {
	include []string
}

func createYamlData(fileName string, data map[string]interface{}) *yamlData {
	con := &yamlData{}
	con.filename = fileName
	con.meta = extractMeta(data)
	con.data = data
	return con
}

func extractMeta(data map[string]interface{}) *meta {
	result := &meta{}
	if meta, ok := data["meta"]; ok {
		if include, ok := meta.(map[interface{}]interface{})["include"]; ok {
			// todo validate include array
			result.include = make([]string, 0)
			for _, v := range include.([]interface{}) {
				result.include = append(result.include, v.(string))
			}
		}
		// erase meta from yaml data
		delete(data, "meta")
	}
	return result
}
