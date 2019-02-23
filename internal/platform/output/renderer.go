package output

// Renderer renders the output
type Renderer interface {
	Render(output interface{}) error
	Error(err error) error
}
