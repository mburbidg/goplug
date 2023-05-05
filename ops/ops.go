package ops

type OpFunction func() string
type OpList map[string]OpFunction
