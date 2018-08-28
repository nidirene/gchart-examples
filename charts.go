package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	chart "github.com/wcharczuk/go-chart"
)

func main() {
	sbc := chart.BarChart{
		Title:      "Test Bar Chart",
		TitleStyle: chart.StyleShow(),
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   512,
		BarWidth: 60,
		XAxis: chart.Style{
			Show: true,
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		Bars: []chart.Value{
			{Value: 5.25, Label: "Blue"},
			{Value: 4.88, Label: "Green"},
			{Value: 4.74, Label: "Gray"},
			{Value: 3.22, Label: "Orange"},
			{Value: 3, Label: "Test"},
			{Value: 2.27, Label: "??"},
			{Value: 1, Label: "!!"},
		},
	}
	buffer := bytes.NewBuffer([]byte{})
	err := sbc.Render(chart.PNG, buffer)
	if err != nil {
		log.Fatalf("error:%v", 2)
		panic(err)
	}

	file, err := os.Create("./chart.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// copy from reader data into writer file
	if _, err := io.Copy(file, buffer); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("file created")

}
