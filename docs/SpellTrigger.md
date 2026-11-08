# SpellTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | Resolved event that activates the effect. | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Reference** | Pointer to [**CombatReference**](CombatReference.md) |  | [optional] 

## Methods

### NewSpellTrigger

`func NewSpellTrigger() *SpellTrigger`

NewSpellTrigger instantiates a new SpellTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpellTriggerWithDefaults

`func NewSpellTriggerWithDefaults() *SpellTrigger`

NewSpellTriggerWithDefaults instantiates a new SpellTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *SpellTrigger) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SpellTrigger) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SpellTrigger) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SpellTrigger) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetCode

`func (o *SpellTrigger) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SpellTrigger) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SpellTrigger) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *SpellTrigger) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *SpellTrigger) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpellTrigger) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpellTrigger) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SpellTrigger) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SpellTrigger) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SpellTrigger) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SpellTrigger) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SpellTrigger) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSubject

`func (o *SpellTrigger) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SpellTrigger) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SpellTrigger) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *SpellTrigger) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetReference

`func (o *SpellTrigger) GetReference() CombatReference`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *SpellTrigger) GetReferenceOk() (*CombatReference, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *SpellTrigger) SetReference(v CombatReference)`

SetReference sets Reference field to given value.

### HasReference

`func (o *SpellTrigger) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


