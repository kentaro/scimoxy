package scimoxy

var handlers = map[string]ServiceHandler{
	"slack": NewSlackHandler(),
}

type ServiceHandlerContainer struct {
}

func NewServiceHandlerContainer() *ServiceHandlerContainer {
	return &ServiceHandlerContainer{}
}

func (c *ServiceHandlerContainer) Count() int {
	return len(handlers)
}

func (c *ServiceHandlerContainer) Add(key string, handler ServiceHandler) {
	handlers[key] = handler
}

func (c *ServiceHandlerContainer) Get(path string) ServiceHandler {
	return handlers[path]
}
