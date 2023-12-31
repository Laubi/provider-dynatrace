/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ConditionObservation struct {

	// Possible Values: `CONTAINS_CASE_INSENSITIVE`, `CONTAINS_CASE_SENSITIVE`, `DOES_NOT_CONTAIN_CASE_INSENSITIVE`, `DOES_NOT_CONTAIN_CASE_SENSITIVE`, `DOES_NOT_EQUAL`, `DOES_NOT_START_WITH`, `EQUALS`, `STARTS_WITH`
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Possible Values: `CUSTOM_DEVICE_GROUP_NAME`, `ENTITY_ID`, `HOST_GROUP_NAME`, `HOST_NAME`, `MANAGEMENT_ZONE`, `NAME`, `PROCESS_GROUP_ID`, `PROCESS_GROUP_NAME`, `TAG`
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// no documentation available
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ConditionParameters struct {

	// Possible Values: `CONTAINS_CASE_INSENSITIVE`, `CONTAINS_CASE_SENSITIVE`, `DOES_NOT_CONTAIN_CASE_INSENSITIVE`, `DOES_NOT_CONTAIN_CASE_SENSITIVE`, `DOES_NOT_EQUAL`, `DOES_NOT_START_WITH`, `EQUALS`, `STARTS_WITH`
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// Possible Values: `CUSTOM_DEVICE_GROUP_NAME`, `ENTITY_ID`, `HOST_GROUP_NAME`, `HOST_NAME`, `MANAGEMENT_ZONE`, `NAME`, `PROCESS_GROUP_ID`, `PROCESS_GROUP_NAME`, `TAG`
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// no documentation available
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type ConditionsObservation struct {
	Condition []ConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`
}

type ConditionsParameters struct {

	// +kubebuilder:validation:Required
	Condition []ConditionParameters `json:"condition" tf:"condition,omitempty"`
}

type DimensionFilterObservation struct {
	Filter []FilterObservation `json:"filter,omitempty" tf:"filter,omitempty"`
}

type DimensionFilterParameters struct {

	// +kubebuilder:validation:Required
	Filter []FilterParameters `json:"filter" tf:"filter,omitempty"`
}

type EntityFilterObservation struct {

	// no documentation available
	Conditions []ConditionsObservation `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// Dimension key of entity type to filter
	DimensionKey *string `json:"dimensionKey,omitempty" tf:"dimension_key,omitempty"`
}

type EntityFilterParameters struct {

	// no documentation available
	// +kubebuilder:validation:Optional
	Conditions []ConditionsParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// Dimension key of entity type to filter
	// +kubebuilder:validation:Optional
	DimensionKey *string `json:"dimensionKey,omitempty" tf:"dimension_key,omitempty"`
}

type EventTemplateObservation struct {

	// Davis® AI will try to merge this event into existing problems, otherwise a new problem will always be created.
	DavisMerge *bool `json:"davisMerge,omitempty" tf:"davis_merge,omitempty"`

	// The description of the event to trigger.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Possible Values: `AVAILABILITY`, `CUSTOM_ALERT`, `CUSTOM_ANNOTATION`, `CUSTOM_CONFIGURATION`, `CUSTOM_DEPLOYMENT`, `ERROR`, `INFO`, `MARKED_FOR_TERMINATION`, `RESOURCE`, `SLOWDOWN`
	EventType *string `json:"eventType,omitempty" tf:"event_type,omitempty"`

	// Set of additional key-value properties to be attached to the triggered event.
	Metadata []MetadataObservation `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The title of the event to trigger.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type EventTemplateParameters struct {

	// Davis® AI will try to merge this event into existing problems, otherwise a new problem will always be created.
	// +kubebuilder:validation:Optional
	DavisMerge *bool `json:"davisMerge,omitempty" tf:"davis_merge,omitempty"`

	// The description of the event to trigger.
	// +kubebuilder:validation:Required
	Description *string `json:"description" tf:"description,omitempty"`

	// Possible Values: `AVAILABILITY`, `CUSTOM_ALERT`, `CUSTOM_ANNOTATION`, `CUSTOM_CONFIGURATION`, `CUSTOM_DEPLOYMENT`, `ERROR`, `INFO`, `MARKED_FOR_TERMINATION`, `RESOURCE`, `SLOWDOWN`
	// +kubebuilder:validation:Required
	EventType *string `json:"eventType" tf:"event_type,omitempty"`

	// Set of additional key-value properties to be attached to the triggered event.
	// +kubebuilder:validation:Optional
	Metadata []MetadataParameters `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The title of the event to trigger.
	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type FilterObservation struct {

	// Dimension key
	DimensionKey *string `json:"dimensionKey,omitempty" tf:"dimension_key,omitempty"`

	// Dimension value
	DimensionValue *string `json:"dimensionValue,omitempty" tf:"dimension_value,omitempty"`

	// Possible Values: `CONTAINS_CASE_SENSITIVE`, `DOES_NOT_CONTAIN_CASE_SENSITIVE`, `DOES_NOT_EQUAL`, `DOES_NOT_START_WITH`, `EQUALS`, `STARTS_WITH`
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`
}

type FilterParameters struct {

	// Dimension key
	// +kubebuilder:validation:Required
	DimensionKey *string `json:"dimensionKey" tf:"dimension_key,omitempty"`

	// Dimension value
	// +kubebuilder:validation:Required
	DimensionValue *string `json:"dimensionValue" tf:"dimension_value,omitempty"`

	// Possible Values: `CONTAINS_CASE_SENSITIVE`, `DOES_NOT_CONTAIN_CASE_SENSITIVE`, `DOES_NOT_EQUAL`, `DOES_NOT_START_WITH`, `EQUALS`, `STARTS_WITH`
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`
}

type MetadataObservation struct {

	// Type 'dt.' for key hints.
	MetadataKey *string `json:"metadataKey,omitempty" tf:"metadata_key,omitempty"`

	// no documentation available
	MetadataValue *string `json:"metadataValue,omitempty" tf:"metadata_value,omitempty"`
}

type MetadataParameters struct {

	// Type 'dt.' for key hints.
	// +kubebuilder:validation:Required
	MetadataKey *string `json:"metadataKey" tf:"metadata_key,omitempty"`

	// no documentation available
	// +kubebuilder:validation:Required
	MetadataValue *string `json:"metadataValue" tf:"metadata_value,omitempty"`
}

type MetricEventObservation struct {

	// This setting is enabled (`true`) or disabled (`false`)
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Controls the preferred entity type used for triggered events.
	EventEntityDimensionKey *string `json:"eventEntityDimensionKey,omitempty" tf:"event_entity_dimension_key,omitempty"`

	// Event template
	EventTemplate []EventTemplateObservation `json:"eventTemplate,omitempty" tf:"event_template,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Config id
	LegacyID *string `json:"legacyId,omitempty" tf:"legacy_id,omitempty"`

	// Monitoring strategy
	ModelProperties []ModelPropertiesObservation `json:"modelProperties,omitempty" tf:"model_properties,omitempty"`

	// Query definition
	QueryDefinition []QueryDefinitionObservation `json:"queryDefinition,omitempty" tf:"query_definition,omitempty"`

	// The textual summary of the metric event entry
	Summary *string `json:"summary,omitempty" tf:"summary,omitempty"`
}

type MetricEventParameters struct {

	// This setting is enabled (`true`) or disabled (`false`)
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Controls the preferred entity type used for triggered events.
	// +kubebuilder:validation:Optional
	EventEntityDimensionKey *string `json:"eventEntityDimensionKey,omitempty" tf:"event_entity_dimension_key,omitempty"`

	// Event template
	// +kubebuilder:validation:Optional
	EventTemplate []EventTemplateParameters `json:"eventTemplate,omitempty" tf:"event_template,omitempty"`

	// Config id
	// +kubebuilder:validation:Optional
	LegacyID *string `json:"legacyId,omitempty" tf:"legacy_id,omitempty"`

	// Monitoring strategy
	// +kubebuilder:validation:Optional
	ModelProperties []ModelPropertiesParameters `json:"modelProperties,omitempty" tf:"model_properties,omitempty"`

	// Query definition
	// +kubebuilder:validation:Optional
	QueryDefinition []QueryDefinitionParameters `json:"queryDefinition,omitempty" tf:"query_definition,omitempty"`

	// The textual summary of the metric event entry
	// +kubebuilder:validation:Optional
	Summary *string `json:"summary,omitempty" tf:"summary,omitempty"`
}

type ModelPropertiesObservation struct {

	// Possible Values: `ABOVE`, `BELOW`, `OUTSIDE`
	AlertCondition *string `json:"alertCondition,omitempty" tf:"alert_condition,omitempty"`

	// The ability to set an alert on missing data in a metric. When enabled, missing data samples will contribute as violating samples defined in advanced model properties. We recommend to not alert on missing data for sparse timeseries as this leads to alert spam.
	AlertOnNoData *bool `json:"alertOnNoData,omitempty" tf:"alert_on_no_data,omitempty"`

	// The number of one-minute samples within the evaluation window that must go back to normal to close the event.
	DealertingSamples *float64 `json:"dealertingSamples,omitempty" tf:"dealerting_samples,omitempty"`

	// The number of one-minute samples that form the sliding evaluation window.
	Samples *float64 `json:"samples,omitempty" tf:"samples,omitempty"`

	// Controls how many times the signal fluctuation is added to the baseline to produce the actual threshold for alerting
	SignalFluctuation *float64 `json:"signalFluctuation,omitempty" tf:"signal_fluctuation,omitempty"`

	// Raise an event if this value is violated
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`

	// Controls the width of the confidence band and larger values lead to a less sensitive model
	Tolerance *float64 `json:"tolerance,omitempty" tf:"tolerance,omitempty"`

	// Possible Values: `AUTO_ADAPTIVE_THRESHOLD`, `SEASONAL_BASELINE`, `STATIC_THRESHOLD`
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The number of one-minute samples within the evaluation window that must violate to trigger an event.
	ViolatingSamples *float64 `json:"violatingSamples,omitempty" tf:"violating_samples,omitempty"`
}

type ModelPropertiesParameters struct {

	// Possible Values: `ABOVE`, `BELOW`, `OUTSIDE`
	// +kubebuilder:validation:Required
	AlertCondition *string `json:"alertCondition" tf:"alert_condition,omitempty"`

	// The ability to set an alert on missing data in a metric. When enabled, missing data samples will contribute as violating samples defined in advanced model properties. We recommend to not alert on missing data for sparse timeseries as this leads to alert spam.
	// +kubebuilder:validation:Required
	AlertOnNoData *bool `json:"alertOnNoData" tf:"alert_on_no_data,omitempty"`

	// The number of one-minute samples within the evaluation window that must go back to normal to close the event.
	// +kubebuilder:validation:Required
	DealertingSamples *float64 `json:"dealertingSamples" tf:"dealerting_samples,omitempty"`

	// The number of one-minute samples that form the sliding evaluation window.
	// +kubebuilder:validation:Required
	Samples *float64 `json:"samples" tf:"samples,omitempty"`

	// Controls how many times the signal fluctuation is added to the baseline to produce the actual threshold for alerting
	// +kubebuilder:validation:Optional
	SignalFluctuation *float64 `json:"signalFluctuation,omitempty" tf:"signal_fluctuation,omitempty"`

	// Raise an event if this value is violated
	// +kubebuilder:validation:Optional
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`

	// Controls the width of the confidence band and larger values lead to a less sensitive model
	// +kubebuilder:validation:Optional
	Tolerance *float64 `json:"tolerance,omitempty" tf:"tolerance,omitempty"`

	// Possible Values: `AUTO_ADAPTIVE_THRESHOLD`, `SEASONAL_BASELINE`, `STATIC_THRESHOLD`
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// The number of one-minute samples within the evaluation window that must violate to trigger an event.
	// +kubebuilder:validation:Required
	ViolatingSamples *float64 `json:"violatingSamples" tf:"violating_samples,omitempty"`
}

type QueryDefinitionObservation struct {

	// Possible Values: `AVG`, `COUNT`, `MAX`, `MEDIAN`, `MIN`, `PERCENTILE90`, `SUM`, `VALUE`
	Aggregation *string `json:"aggregation,omitempty" tf:"aggregation,omitempty"`

	// Dimension filter
	DimensionFilter []DimensionFilterObservation `json:"dimensionFilter,omitempty" tf:"dimension_filter,omitempty"`

	// Use rule-based filters to define the scope this event monitors.
	EntityFilter []EntityFilterObservation `json:"entityFilter,omitempty" tf:"entity_filter,omitempty"`

	// Management zone
	ManagementZone *string `json:"managementZone,omitempty" tf:"management_zone,omitempty"`

	// Metric key
	MetricKey *string `json:"metricKey,omitempty" tf:"metric_key,omitempty"`

	// To learn more, visit [Metric Selector](https://dt-url.net/metselad)
	MetricSelector *string `json:"metricSelector,omitempty" tf:"metric_selector,omitempty"`

	// Minute offset of sliding evaluation window for metrics with latency
	QueryOffset *float64 `json:"queryOffset,omitempty" tf:"query_offset,omitempty"`

	// Possible Values: `METRIC_KEY`, `METRIC_SELECTOR`
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type QueryDefinitionParameters struct {

	// Possible Values: `AVG`, `COUNT`, `MAX`, `MEDIAN`, `MIN`, `PERCENTILE90`, `SUM`, `VALUE`
	// +kubebuilder:validation:Optional
	Aggregation *string `json:"aggregation,omitempty" tf:"aggregation,omitempty"`

	// Dimension filter
	// +kubebuilder:validation:Optional
	DimensionFilter []DimensionFilterParameters `json:"dimensionFilter,omitempty" tf:"dimension_filter,omitempty"`

	// Use rule-based filters to define the scope this event monitors.
	// +kubebuilder:validation:Optional
	EntityFilter []EntityFilterParameters `json:"entityFilter,omitempty" tf:"entity_filter,omitempty"`

	// Management zone
	// +kubebuilder:validation:Optional
	ManagementZone *string `json:"managementZone,omitempty" tf:"management_zone,omitempty"`

	// Metric key
	// +kubebuilder:validation:Optional
	MetricKey *string `json:"metricKey,omitempty" tf:"metric_key,omitempty"`

	// To learn more, visit [Metric Selector](https://dt-url.net/metselad)
	// +kubebuilder:validation:Optional
	MetricSelector *string `json:"metricSelector,omitempty" tf:"metric_selector,omitempty"`

	// Minute offset of sliding evaluation window for metrics with latency
	// +kubebuilder:validation:Optional
	QueryOffset *float64 `json:"queryOffset,omitempty" tf:"query_offset,omitempty"`

	// Possible Values: `METRIC_KEY`, `METRIC_SELECTOR`
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// MetricEventSpec defines the desired state of MetricEvent
type MetricEventSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MetricEventParameters `json:"forProvider"`
}

// MetricEventStatus defines the observed state of MetricEvent.
type MetricEventStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MetricEventObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MetricEvent is the Schema for the MetricEvents API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type MetricEvent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.enabled)",message="enabled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.eventTemplate)",message="eventTemplate is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.modelProperties)",message="modelProperties is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.queryDefinition)",message="queryDefinition is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.summary)",message="summary is a required parameter"
	Spec   MetricEventSpec   `json:"spec"`
	Status MetricEventStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MetricEventList contains a list of MetricEvents
type MetricEventList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MetricEvent `json:"items"`
}

// Repository type metadata.
var (
	MetricEvent_Kind             = "MetricEvent"
	MetricEvent_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MetricEvent_Kind}.String()
	MetricEvent_KindAPIVersion   = MetricEvent_Kind + "." + CRDGroupVersion.String()
	MetricEvent_GroupVersionKind = CRDGroupVersion.WithKind(MetricEvent_Kind)
)

func init() {
	SchemeBuilder.Register(&MetricEvent{}, &MetricEventList{})
}
