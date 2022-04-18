/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TODO: Add comments for each struct

type Schedule struct {
	Cron     string `json:"cron"`
	Duration string `json:"duration"`
}

// SpannerAutoscaleScheduleSpec defines the desired state of SpannerAutoscaleSchedule
type SpannerAutoscaleScheduleSpec struct {
	TargetResource            string   `json:"targetResource"`
	AdditionalProcessingUnits int      `json:"additionalProcessingUnits"`
	Schedule                  Schedule `json:"schedule"`
}

// SpannerAutoscaleScheduleStatus defines the observed state of SpannerAutoscaleSchedule
type SpannerAutoscaleScheduleStatus struct{}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Cron",type="string",JSONPath=".spec.schedule.cron"
//+kubebuilder:printcolumn:name="Duration",type="string",JSONPath=".spec.schedule.duration"
//+kubebuilder:printcolumn:name="Additional PU",type="integer",JSONPath=".spec.additionalProcessingUnits"
// SpannerAutoscaleSchedule is the Schema for the spannerautoscaleschedules API
type SpannerAutoscaleSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpannerAutoscaleScheduleSpec   `json:"spec,omitempty"`
	Status SpannerAutoscaleScheduleStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SpannerAutoscaleScheduleList contains a list of SpannerAutoscaleSchedule
type SpannerAutoscaleScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpannerAutoscaleSchedule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SpannerAutoscaleSchedule{}, &SpannerAutoscaleScheduleList{})
}
