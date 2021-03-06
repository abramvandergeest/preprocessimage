package preprocessimage

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"os"
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestEval(t *testing.T) {
	fmt.Println("GETS HERE")
	f, err := os.Open("/Users/avanderg@tibco.com/working/image_recog_hands_on/IMG_20190214_152236108.jpg")
	if err != nil {
		log.Fatal("trouble loading test file")
	}
	pic, _, err := image.Decode(f)
	src := pic.(image.Image)

	// iCtx := test.NewActivityInitContext(nil, nil)
	// act, err := New(iCtx)
	act := &Activity{}
	assert.Nil(t, err)
	input := &Input{Image: src}
	tc := test.NewActivityContext(act.Metadata())

	tc.SetInputObject(input)

	done, err := act.Eval(tc)
	assert.True(t, done)
	assert.Nil(t, err)

	output := &Output{}
	tc.GetOutputObject(output)
	// fmt.Println(output.ResizedImage)
}
