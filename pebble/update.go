package pebble

type Option struct {
	Key       string `json:"cache"`
	ID        string `json:"key"`
	Min       int16  `json:"min"`
	Prefix    string `json:"prefix"`
	StartTime string `json:"startTime"`
}

type Conf = map[string]interface{}

func mapMerge(list Conf, data Conf) Conf {
	for key, value := range data {
		list[key] = value
	}
	return list
}

type FN func(interface{}) (interface{}, error)

func Put(opts Config, data interface{}, fn FN) error {

	return nil
}
