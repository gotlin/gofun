package pipeline

import (
	"gofun/fun/scrago/data"
	"fmt"
)

type ConsolePipeline struct{}

func (cp *ConsolePipeline) Pipe(resp *data.Response) {

	fmt.Println(resp.StrBody)

}

func New() *ConsolePipeline{
	return &ConsolePipeline{}
}