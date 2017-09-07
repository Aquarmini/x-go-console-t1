package main

import (
	"fmt"
	"os"
	"github.com/facebookgo/inject"
	"app/container"
	"app/providers"
	"app/controller"
)

func main() {

	var g inject.Graph

	// We provide our graph two "seed" objects, one our empty
	// HomePlanetRenderApp instance which we're hoping to get filled out,
	// and second our DefaultTransport to satisfy our HTTPTransport
	// dependency. We have to provide the DefaultTransport because the
	// dependency is defined in terms of the http.RoundTripper interface,
	// and since it is an interface the library cannot create an instance
	// for it. Instead it will use the given DefaultTransport to satisfy
	// the dependency since it implements the interface:
	var di container.Container
	err := g.Provide(
		&inject.Object{Value: &di, Name:"di"},
		&inject.Object{Value:providers.NewConfig(), Name:"config"},
		&inject.Object{Value:providers.NewLogger(), Name:"logger"},
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Here the Populate call is creating instances of NameAPI &
	// PlanetAPI, and setting the HTTPTransport on both to the
	// http.DefaultTransport provided above:
	if err := g.Populate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// There is a shorthand API for the simple case which combines the
	// three calls above is available as inject.Populate:
	//
	//   inject.Populate(&a, http.DefaultTransport)
	//
	// The above API shows the underlying API which also allows the use of
	// named instances for more complex scenarios.

	//fmt.Println(a.Config.GetKey("database", "adapter"))
	//fmt.Println(a.Config)

	var params []string
	if len(os.Args) <= 1 {
		fmt.Println("请输入方法名")
		return
	}

	TaskMap := make(map[string]controller.ControllerInterface)

	TaskMap["Test"] = &controller.TestController{DI:di}

	for i := 1; i < len(os.Args); i++ {
		params = append(params, os.Args[i])
	}

	if TaskMap[params[0]].Handle() != nil {
		fmt.Println(err.Error())
	}
}