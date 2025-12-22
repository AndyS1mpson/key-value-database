package container

import (
	"runtime"
	"strings"
)

// MustOrGetNew retrieves a service instance from the container or creates a new one using the factory.
// Panics if the factory function returns an error.
func MustOrGetNew[T any](c *Container, factory func() T) T {
	instance, err := GetOrNew(c, func() (T, error) {
		return factory(), nil
	})
	if err != nil {
		panic(err)
	}

	return instance
}

// GetOrNew get service instance, if failed return error
func GetOrNew[T any](c *Container, factory func() (T, error)) (T, error) {
	name := serviceName()

	if srv, ok := c.services.Load(name); ok {
		return srv.(T), nil
	}

	srv, err := factory()
	if err != nil {
		var def T
		return def, err
	}

	c.services.Store(name, srv)

	return srv, nil
}

// serviceName determines the service name by inspecting the call stack to find the caller function.
func serviceName() ServiceName {
	pcs := make([]uintptr, 10000)
	n := runtime.Callers(1, pcs)

	if n == 0 {
		return ""
	}

	callerFrames := runtime.CallersFrames(pcs)

	more := true
	callerFrame := runtime.Frame{}

	for more {
		callerFrame, more = callerFrames.Next()
		name := callerFrame.Function

		if strings.Contains(name, "MustOrGetNew") ||
			strings.Contains(name, "GetOrNew") ||
			strings.Contains(name, "serviceName") ||
			name == "" {
			continue
		}

		parts := strings.Split(name, ".")

		return ServiceName(parts[len(parts)-1])
	}

	panic("can not detect correct name of service")
}
