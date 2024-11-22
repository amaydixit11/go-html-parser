package dom

type AbortController struct {
	signal *AbortSignal
}

type AbortControllerInterface interface {
	Abort(reason any)
}
