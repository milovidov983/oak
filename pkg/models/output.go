package models

import "fmt"

type OutputType int

const (
	Console = iota
	Clipboard
)

var typeName = [...]string{"Console", "Clipboard"}

func (o OutputType) String() string {
	return typeName[o]
}

type Output struct {
	outputType OutputType
}

func (o *Output) Send(payload string) error {
	if o.outputType == Console {
		fmt.Println(payload)
	}
	if o.outputType == Clipboard {
		return fmt.Errorf("clipboard not implemented")
		//clipboard.WriteAll(payload)
	}
	return nil
}
