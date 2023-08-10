# GetGameSearch200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the hit. The query could still have matched with the description only. | [optional] 
**AnkamaId** | Pointer to **int32** | Ankama ID for retrieving more details in the type specific endpoint. | [optional] 
**Type** | Pointer to **string** | Type classification | [optional] 

## Methods

### NewGetGameSearch200ResponseInner

`func NewGetGameSearch200ResponseInner() *GetGameSearch200ResponseInner`

NewGetGameSearch200ResponseInner instantiates a new GetGameSearch200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGameSearch200ResponseInnerWithDefaults

`func NewGetGameSearch200ResponseInnerWithDefaults() *GetGameSearch200ResponseInner`

NewGetGameSearch200ResponseInnerWithDefaults instantiates a new GetGameSearch200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetGameSearch200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetGameSearch200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetGameSearch200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetGameSearch200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAnkamaId

`func (o *GetGameSearch200ResponseInner) GetAnkamaId() int32`

GetAnkamaId returns the AnkamaId field if non-nil, zero value otherwise.

### GetAnkamaIdOk

`func (o *GetGameSearch200ResponseInner) GetAnkamaIdOk() (*int32, bool)`

GetAnkamaIdOk returns a tuple with the AnkamaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnkamaId

`func (o *GetGameSearch200ResponseInner) SetAnkamaId(v int32)`

SetAnkamaId sets AnkamaId field to given value.

### HasAnkamaId

`func (o *GetGameSearch200ResponseInner) HasAnkamaId() bool`

HasAnkamaId returns a boolean if a field has been set.

### GetType

`func (o *GetGameSearch200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetGameSearch200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetGameSearch200ResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetGameSearch200ResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


