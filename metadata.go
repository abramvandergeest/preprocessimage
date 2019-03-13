package preprocessImage

import (
	"image"
)

// type Settings struct {
// 	ASetting string `md:"aSetting,required"`
// }

type Input struct {
	Image image.Image `md:"image"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	// strVal, _ := coerce.ToString(values["anInput"])
	r.Image = values["image"].(image.Image)
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"image": r.Image,
	}
}

type Output struct {
	Output interface{} `md:"output"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	// strVal, _ := coerce.ToString()
	o.Output = values["output"]
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"output": o.Output,
	}
}
