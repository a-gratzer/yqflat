package yqflat

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type YqFlat struct {
	separator string
}

func New(separator string) *YqFlat {
	return &YqFlat{
		separator: separator,
	}
}

func (y *YqFlat) Parse(file string) (map[string]string, error) {

	var structData map[string]interface{}
	var err error
	if structData, err = y.readStruct(file); err != nil {
		return nil, err
	}

	flat := make(map[string]string)
	y.flatten("", structData, flat)

	return flat, nil
}

// readStruct reads a YAML file from the given file path and unmarshals its contents
// into a map[string]interface{}. Returns the resulting map and any error encountered.
func (y *YqFlat) readStruct(filePath string) (map[string]interface{}, error) {
	// Read YAML file

	var bytes []byte
	var err error

	if bytes, err = os.ReadFile(filePath); err != nil {
		return nil, err
	}

	var data map[string]interface{}
	if err := yaml.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	return data, nil
}

// flatten recursively traverses a nested structure (maps and slices), building flat key-value pairs
// with keys composed using the given prefix and separator, and stores the results in the provided map.
func (y *YqFlat) flatten(prefix string, value interface{}, out map[string]string) {
	switch v := value.(type) {
	case map[string]interface{}:
		for key, val := range v {
			fullKey := key
			if prefix != "" {
				fullKey = prefix + y.separator + key
			}
			y.flatten(fullKey, val, out)
		}
	case []interface{}:
		for i, val := range v {
			fullKey := fmt.Sprintf("%s[%d]", prefix, i)
			y.flatten(fullKey, val, out)
		}
	default:
		out[prefix] = fmt.Sprintf("%v", v)
	}
}
