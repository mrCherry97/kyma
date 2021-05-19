// Code generated by injection-gen. DO NOT EDIT.

package fake

import (
	"context"

	fake "github.com/kyma-project/kyma/components/event-sources/client/generated/clientset/internalclientset/fake"
	client "github.com/kyma-project/kyma/components/event-sources/client/generated/injection/client"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	injection "knative.dev/pkg/injection"
	logging "knative.dev/pkg/logging"
)

func init() {
	injection.Fake.RegisterClient(withClient)
}

func withClient(ctx context.Context, cfg *rest.Config) context.Context {
	ctx, _ = With(ctx)
	return ctx
}

func With(ctx context.Context, objects ...runtime.Object) (context.Context, *fake.Clientset) {
	cs := fake.NewSimpleClientset(objects...)
	return context.WithValue(ctx, client.Key{}, cs), cs
}

// Get extracts the Kubernetes client from the context.
func Get(ctx context.Context) *fake.Clientset {
	untyped := ctx.Value(client.Key{})
	if untyped == nil {
		logging.FromContext(ctx).Panic(
			"Unable to fetch github.com/kyma-project/kyma/components/event-sources/client/generated/clientset/internalclientset/fake.Clientset from context.")
	}
	return untyped.(*fake.Clientset)
}