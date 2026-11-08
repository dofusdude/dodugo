# TargetSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | Human-readable target category or relationship. | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Negated** | Pointer to **bool** |  | [optional] 
**Reference** | Pointer to [**CombatReference**](CombatReference.md) |  | [optional] 
**Conditions** | Pointer to [**[]CombatCondition**](CombatCondition.md) |  | [optional] 

## Methods

### NewTargetSelector

`func NewTargetSelector() *TargetSelector`

NewTargetSelector instantiates a new TargetSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetSelectorWithDefaults

`func NewTargetSelectorWithDefaults() *TargetSelector`

NewTargetSelectorWithDefaults instantiates a new TargetSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *TargetSelector) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TargetSelector) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TargetSelector) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TargetSelector) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetCode

`func (o *TargetSelector) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TargetSelector) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TargetSelector) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TargetSelector) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *TargetSelector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetSelector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetSelector) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TargetSelector) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *TargetSelector) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetSelector) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetSelector) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetSelector) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSubject

`func (o *TargetSelector) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TargetSelector) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TargetSelector) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *TargetSelector) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetNegated

`func (o *TargetSelector) GetNegated() bool`

GetNegated returns the Negated field if non-nil, zero value otherwise.

### GetNegatedOk

`func (o *TargetSelector) GetNegatedOk() (*bool, bool)`

GetNegatedOk returns a tuple with the Negated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegated

`func (o *TargetSelector) SetNegated(v bool)`

SetNegated sets Negated field to given value.

### HasNegated

`func (o *TargetSelector) HasNegated() bool`

HasNegated returns a boolean if a field has been set.

### GetReference

`func (o *TargetSelector) GetReference() CombatReference`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *TargetSelector) GetReferenceOk() (*CombatReference, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *TargetSelector) SetReference(v CombatReference)`

SetReference sets Reference field to given value.

### HasReference

`func (o *TargetSelector) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetConditions

`func (o *TargetSelector) GetConditions() []CombatCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *TargetSelector) GetConditionsOk() (*[]CombatCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *TargetSelector) SetConditions(v []CombatCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *TargetSelector) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


