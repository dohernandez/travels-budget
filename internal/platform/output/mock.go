package output

// NewCallbackRendererMock creates a callback mock for tests
// nolint:unused
func NewCallbackRendererMock(
	rendererOutFunc func(output interface{}) error,
	rendererErrFunc func(err error) error,
) Renderer {
	return &rendererMock{
		rendererOutFunc: rendererOutFunc,
		rendererErrFunc: rendererErrFunc,
	}
}

// nolint:unused
type rendererMock struct {
	rendererOutFunc func(output interface{}) error
	rendererErrFunc func(err error) error
}

func (m *rendererMock) Render(output interface{}) error {
	return m.rendererOutFunc(output)
}

func (m *rendererMock) Error(err error) error {
	return m.rendererErrFunc(err)
}
