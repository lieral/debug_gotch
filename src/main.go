/**
 * @file main.go
 * @brief
 * @author pom
 * @version 1.0
 * @date 2022-04-21
 */

package main

import (
	"fmt"

	"github.com/sugarme/gotch/ts"
)

var model *ts.CModule

func ModelManager() {
	var err error
	if model, err = ts.ModuleLoad("../model/epoch_2500.pt"); err != nil {
		fmt.Println(err)
	}
	model.SetEval()
	fmt.Println(model)
}

func ModelPredict() {
	obsVec_ := []float64{0.18, 0.32}
	inputTensor, _ := ts.NewTensorFromData(obsVec_, []int64{1, 2})
	inputIVal := ts.NewIValue(*inputTensor)
	if m, err := model.ForwardIs([]ts.IValue{*inputIVal}); err == nil {
		for _, outTensor := range m.Value().([]ts.Tensor) {
			fmt.Println(outTensor.Vals())
		}
	}
}

func main() {
	ModelManager()
	for i := 0; i < 100; i++ {
		ModelPredict()
	}
}
