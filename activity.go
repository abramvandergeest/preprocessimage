package preprocessimage

import (
	"github.com/project-flogo/core/activity"
)

func init() {
	activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	ctx.Logger().Info("Input: Image here.")

	src := input.Image
	batchsize := 1

	bounds := src.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y

	// //Converting Image to array
	var img [][][][]uint8
	for j := 0; j < batchsize; j++ {
		var batch [][][]uint8
		for x := 0; x < w; x++ {
			var row [][]uint8
			for y := 0; y < h; y++ {
				var col []uint8
				for i := 0; i < 3; i++ {
					col = append(col, 0)
				}

				row = append(row, col)
			}
			batch = append(batch, row)
		}
		img = append(img, batch)
	}

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			imageColor := src.At(x, y)
			rr, bb, gg, _ := imageColor.RGBA()
			color := []uint8{uint8(rr / 256), uint8(bb / 255), uint8(gg / 256)}
			for i := 0; i < 3; i++ {
				img[0][x][y][i] = color[i]
			}

		}
	}

	output := &Output{Output: img}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}
