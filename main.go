package main

import (
	"context"
	"fmt"
	"reflect"
)

func log(ctx context.Context, t string) {
	email, err := getEmail(ctx)
	if err != nil {
		email = fmt.Sprintf("%v", err)
	}

	timeString := ""
	time, err := getTime(ctx)
	if err != nil {
		timeString = fmt.Sprintf("%v", err)
	} else {
		timeString = time.String()
	}

	id, err := getID(ctx)
	if err != nil {
		id = fmt.Sprintf("%v", err)
	}

	fmt.Printf("%s\nid: %s\nemail: %s\ntime: %s\n\n", t, id, email, timeString)
}

func main() {
	ctx := context.Background()

	signedInUser := &SignedInUser{
		id:    "12345",
		email: "example@grafana.com",
	}
	anonymousUser := &AnonymousUser{}
	pluginInstance := &PluginInstance{
		id: "abcdef",
	}

	ctx = setIdentity(ctx, signedInUser)
	log(ctx, reflect.TypeOf(signedInUser).String())

	ctx = setIdentity(ctx, anonymousUser)
	log(ctx, reflect.TypeOf(anonymousUser).String())

	ctx = setIdentity(ctx, pluginInstance)
	log(ctx, reflect.TypeOf(pluginInstance).String())
}
