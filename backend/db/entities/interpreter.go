package entities

type Interpreter struct {
	Interpreter_id *int64
	Name           *string
}

func NewInterpreter(id *int64, name *string) *Interpreter {
	interpreter := new(Interpreter)
	interpreter.Interpreter_id = id
	interpreter.Name = name
	return interpreter
}
