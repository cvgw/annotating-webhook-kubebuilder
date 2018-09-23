package annotator

import (
	"context"

	"net/http"

	"log"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/runtime/inject"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission/types"
)

type podAnnotator struct {
	client  client.Client
	decoder types.Decoder
}

// podAnnotator implements admission.Handler.
var _ admission.Handler = &podAnnotator{}

// Automatically generate RBAC rules to allow the Controller to read and write Deployments
// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=mutatingwebhookconfigurations,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=,resources=services,verbs=get;list;watch;create;update;patch;delete
func (a *podAnnotator) Handle(ctx context.Context, req types.Request) types.Response {
	pod := &corev1.Pod{}

	err := a.decoder.Decode(req, pod)
	if err != nil {
		return admission.ErrorResponse(http.StatusBadRequest, err)
	}
	copy := pod.DeepCopy()

	err = a.mutatePodsFn(ctx, copy)
	if err != nil {
		return admission.ErrorResponse(http.StatusInternalServerError, err)
	}

	// admission.PatchResponse generates a Response containing patches.
	return admission.PatchResponse(pod, copy)
}

// podAnnotator implements inject.Client.
var _ inject.Client = &podAnnotator{}

// InjectClient injects the client into the podAnnotator
func (a *podAnnotator) InjectClient(c client.Client) error {
	a.client = c
	return nil
}

// podAnnotator implements inject.Decoder.
var _ inject.Decoder = &podAnnotator{}

// InjectDecoder injects the decoder into the podAnnotator
func (a *podAnnotator) InjectDecoder(d types.Decoder) error {
	a.decoder = d
	return nil
}

func (a *podAnnotator) mutatePodsFn(ctx context.Context, copy *corev1.Pod) error {
	log.Println("I got a request")

	return nil
}
