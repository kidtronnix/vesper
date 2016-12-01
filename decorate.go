package vesper

// We define what a decorator to our client will look like
type Decorator func(Client) Client

// Decorate will decorate a client with a slice of passed decorators
func Decorate(c Client, ds ...Decorator) Client {
	decorated := c
	for _, decorate := range ds {
		decorated = decorate(decorated)
	}
	return decorated
}
