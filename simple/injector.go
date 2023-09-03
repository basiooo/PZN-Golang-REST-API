//go:build wireinject
// +build wireinject

package simple

import (
	"io"
	"os"

	"github.com/google/wire"
)

func InitializeService(isError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository, NewSimpleService,
	)
	return nil, nil
}

func InitializeDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabaseMongoDB,
		NewDatabaseRepository,
		NewDatabasePostgreSQL,
	)
	return nil
}

func InitializeFooBarService() *FooBarService {
	wire.Build(
		NewFooRepository,
		NewFooService,
		NewBarRepository,
		NewBarService,
		NewFooBarService,
	)
	return nil
}

// output sama seperti InitializeFooBarService() tapi yang ini digrouping
// NOTE: menggunakan grouping(newSet) atau tidak menggunakan grouping tidak mempengaruhi hasil akhir
var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializeFooBarServiceGroupProvider() *FooBarService {
	wire.Build(
		fooSet,
		barSet,
		NewFooBarService,
	)
	return nil
}

// binding interface
// func InitializeHelloService() *HelloService { //tanpa binding (error)
// 	wire.Build(
// 		NewHelloService,
// 		NewSayHelloImpl,
// 	)
// 	return nil
// }

var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)), //binding SayHello menjadi SayHelloImpl
)

func InizializeHelloService() *HelloService { //menggunakan binding (berhasil)
	wire.Build(helloSet, NewHelloService)
	return nil
}

// struct Provider
func InitializeFooBar() *FooBar {
	wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "Foo", "Bar"))
	return nil
}

// struct Provider Binding value
var fooValue = &Foo{}
var barValue = &Bar{}

func InitializeFooBarUsingValue() *FooBar {
	wire.Build(wire.Value(fooValue), wire.Value(barValue), wire.Struct(new(FooBar), "Foo", "Bar"))
	return nil
}

func InitializeReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

func InintializeConfiguration() *Configuration {
	// application := NewApplication()
	// configuration := application.Configuration
	// return configuration
	// result seperti diatas
	wire.Build(
		NewApplication,
		wire.FieldsOf(new(*Application), "Configuration"),
	)
	return nil
}

// cleanup Function

func InitializeConnection(name string) (*Connection, func()) {
	wire.Build(NewConnection, NewFile)
	return nil, nil
}
