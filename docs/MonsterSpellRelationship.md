# MonsterSpellRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnkamaId** | **int32** |  | 
**Name** | **string** |  | 
**Grade** | **int32** | The spell grade assigned to this monster grade. | 
**Url** | **string** | Grade-specific spell URL; already includes the &#x60;grade&#x60; query parameter. | 
**ImageUrls** | Pointer to [**Images**](Images.md) |  | [optional] 

## Methods

### NewMonsterSpellRelationship

`func NewMonsterSpellRelationship(ankamaId int32, name string, grade int32, url string, ) *MonsterSpellRelationship`

NewMonsterSpellRelationship instantiates a new MonsterSpellRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonsterSpellRelationshipWithDefaults

`func NewMonsterSpellRelationshipWithDefaults() *MonsterSpellRelationship`

NewMonsterSpellRelationshipWithDefaults instantiates a new MonsterSpellRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnkamaId

`func (o *MonsterSpellRelationship) GetAnkamaId() int32`

GetAnkamaId returns the AnkamaId field if non-nil, zero value otherwise.

### GetAnkamaIdOk

`func (o *MonsterSpellRelationship) GetAnkamaIdOk() (*int32, bool)`

GetAnkamaIdOk returns a tuple with the AnkamaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnkamaId

`func (o *MonsterSpellRelationship) SetAnkamaId(v int32)`

SetAnkamaId sets AnkamaId field to given value.


### GetName

`func (o *MonsterSpellRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonsterSpellRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonsterSpellRelationship) SetName(v string)`

SetName sets Name field to given value.


### GetGrade

`func (o *MonsterSpellRelationship) GetGrade() int32`

GetGrade returns the Grade field if non-nil, zero value otherwise.

### GetGradeOk

`func (o *MonsterSpellRelationship) GetGradeOk() (*int32, bool)`

GetGradeOk returns a tuple with the Grade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrade

`func (o *MonsterSpellRelationship) SetGrade(v int32)`

SetGrade sets Grade field to given value.


### GetUrl

`func (o *MonsterSpellRelationship) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MonsterSpellRelationship) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MonsterSpellRelationship) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetImageUrls

`func (o *MonsterSpellRelationship) GetImageUrls() Images`

GetImageUrls returns the ImageUrls field if non-nil, zero value otherwise.

### GetImageUrlsOk

`func (o *MonsterSpellRelationship) GetImageUrlsOk() (*Images, bool)`

GetImageUrlsOk returns a tuple with the ImageUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrls

`func (o *MonsterSpellRelationship) SetImageUrls(v Images)`

SetImageUrls sets ImageUrls field to given value.

### HasImageUrls

`func (o *MonsterSpellRelationship) HasImageUrls() bool`

HasImageUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


