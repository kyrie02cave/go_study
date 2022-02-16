package time

import (
	"testing"
)

func TestAll(t *testing.T){
	t.Run("timeDemo", func(t *testing.T){
		timeDemo()
	})

	t.Run("timestampDemo", func(t *testing.T){
		timestampDemo()
	})

	t.Run("timestampDemo2", func(t *testing.T){
		timestampDemo2(1636511048)
	})

	t.Run("timeAdd", func(t *testing.T){
		timeAdd()
	})

	t.Run("tickDemo", func(t *testing.T){
		tickDemo()
	})

	t.Run("timeParse", func(t *testing.T){
		timeParse()
	})
}
