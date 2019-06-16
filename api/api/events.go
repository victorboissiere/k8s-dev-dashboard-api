package api

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

type Event struct {
	Name          string
	Reason        string
	Message       string
	Count         int32
	LastTimestamp metav1.Time
}

type EventList []Event

func GetEvents(namespace string, deploymentName string) EventList {
	result := &v1.EventList{}
	options := &metav1.ListOptions{
		FieldSelector: fields.OneTermEqualSelector("involvedObject.name", deploymentName).String(),
	}

	err := client.CoreV1().RESTClient().Get().
		Namespace(namespace).
		Resource("events").
		VersionedParams(options, scheme.ParameterCodec).
		Do().
		Into(result)

	checkError(err)

	var eventList = EventList{}
	for i := 0; i < len(result.Items); i++ {
		event := result.Items[i]
		eventList = append(eventList, Event{
			Name: event.ObjectMeta.Name,
			Reason: event.Reason,
			Message: event.Message,
			LastTimestamp: event.LastTimestamp,
			Count: event.Count,
		})
	}

	return eventList
}




