package transloadit

var verdictMap map[string]map[string]interface{} = map[string]map[string]interface{}{
	"resize": map[string]interface{}{
		"robot":           "/image/resize",
		"width":           75,
		"height":          75,
		"resize_strategy": "pad",
		"background":      "#000000",
	},
}
