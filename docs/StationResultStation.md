# StationResultStation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**SystemName** | Pointer to **string** |  | [optional] 
**ShortName** | Pointer to **string** |  | [optional] 
**MediumName** | Pointer to **string** |  | [optional] 
**Website** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Keywords** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Sms** | Pointer to **string** |  | [optional] 
**Telephone** | Pointer to **string** |  | [optional] 
**GenreId** | Pointer to **int32** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**LogoRectangle** | Pointer to **string** |  | [optional] 
**Logo128x128** | Pointer to **string** |  | [optional] 
**Logo320x320** | Pointer to **string** |  | [optional] 
**Logo600x600** | Pointer to **string** |  | [optional] 
**PayOff** | Pointer to **string** |  | [optional] 
**PtyCode** | Pointer to **int32** |  | [optional] 
**PtyType** | Pointer to **string** |  | [optional] 
**StationKey** | Pointer to **string** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**MetadataradioOrganisation** | Pointer to **string** |  | [optional] 
**MetadataradioStationId** | Pointer to **string** |  | [optional] 
**StartDays** | Pointer to [**StationResultStationStartDays**](StationResultStationStartDays.md) |  | [optional] 

## Methods

### NewStationResultStation

`func NewStationResultStation() *StationResultStation`

NewStationResultStation instantiates a new StationResultStation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStationResultStationWithDefaults

`func NewStationResultStationWithDefaults() *StationResultStation`

NewStationResultStationWithDefaults instantiates a new StationResultStation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StationResultStation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StationResultStation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StationResultStation) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *StationResultStation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *StationResultStation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StationResultStation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StationResultStation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StationResultStation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StationResultStation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StationResultStation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StationResultStation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StationResultStation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *StationResultStation) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StationResultStation) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StationResultStation) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *StationResultStation) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetSystemName

`func (o *StationResultStation) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *StationResultStation) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *StationResultStation) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *StationResultStation) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetShortName

`func (o *StationResultStation) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *StationResultStation) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *StationResultStation) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *StationResultStation) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetMediumName

`func (o *StationResultStation) GetMediumName() string`

GetMediumName returns the MediumName field if non-nil, zero value otherwise.

### GetMediumNameOk

`func (o *StationResultStation) GetMediumNameOk() (*string, bool)`

GetMediumNameOk returns a tuple with the MediumName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediumName

`func (o *StationResultStation) SetMediumName(v string)`

SetMediumName sets MediumName field to given value.

### HasMediumName

`func (o *StationResultStation) HasMediumName() bool`

HasMediumName returns a boolean if a field has been set.

### GetWebsite

`func (o *StationResultStation) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *StationResultStation) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *StationResultStation) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *StationResultStation) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetEmail

`func (o *StationResultStation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *StationResultStation) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *StationResultStation) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *StationResultStation) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetKeywords

`func (o *StationResultStation) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *StationResultStation) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *StationResultStation) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *StationResultStation) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetDescription

`func (o *StationResultStation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StationResultStation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StationResultStation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StationResultStation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSms

`func (o *StationResultStation) GetSms() string`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *StationResultStation) GetSmsOk() (*string, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *StationResultStation) SetSms(v string)`

SetSms sets Sms field to given value.

### HasSms

`func (o *StationResultStation) HasSms() bool`

HasSms returns a boolean if a field has been set.

### GetTelephone

`func (o *StationResultStation) GetTelephone() string`

GetTelephone returns the Telephone field if non-nil, zero value otherwise.

### GetTelephoneOk

`func (o *StationResultStation) GetTelephoneOk() (*string, bool)`

GetTelephoneOk returns a tuple with the Telephone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephone

`func (o *StationResultStation) SetTelephone(v string)`

SetTelephone sets Telephone field to given value.

### HasTelephone

`func (o *StationResultStation) HasTelephone() bool`

HasTelephone returns a boolean if a field has been set.

### GetGenreId

`func (o *StationResultStation) GetGenreId() int32`

GetGenreId returns the GenreId field if non-nil, zero value otherwise.

### GetGenreIdOk

`func (o *StationResultStation) GetGenreIdOk() (*int32, bool)`

GetGenreIdOk returns a tuple with the GenreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenreId

`func (o *StationResultStation) SetGenreId(v int32)`

SetGenreId sets GenreId field to given value.

### HasGenreId

`func (o *StationResultStation) HasGenreId() bool`

HasGenreId returns a boolean if a field has been set.

### GetLanguage

`func (o *StationResultStation) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *StationResultStation) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *StationResultStation) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *StationResultStation) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetActive

`func (o *StationResultStation) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *StationResultStation) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *StationResultStation) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *StationResultStation) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetLogoRectangle

`func (o *StationResultStation) GetLogoRectangle() string`

GetLogoRectangle returns the LogoRectangle field if non-nil, zero value otherwise.

### GetLogoRectangleOk

`func (o *StationResultStation) GetLogoRectangleOk() (*string, bool)`

GetLogoRectangleOk returns a tuple with the LogoRectangle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoRectangle

`func (o *StationResultStation) SetLogoRectangle(v string)`

SetLogoRectangle sets LogoRectangle field to given value.

### HasLogoRectangle

`func (o *StationResultStation) HasLogoRectangle() bool`

HasLogoRectangle returns a boolean if a field has been set.

### GetLogo128x128

`func (o *StationResultStation) GetLogo128x128() string`

GetLogo128x128 returns the Logo128x128 field if non-nil, zero value otherwise.

### GetLogo128x128Ok

`func (o *StationResultStation) GetLogo128x128Ok() (*string, bool)`

GetLogo128x128Ok returns a tuple with the Logo128x128 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo128x128

`func (o *StationResultStation) SetLogo128x128(v string)`

SetLogo128x128 sets Logo128x128 field to given value.

### HasLogo128x128

`func (o *StationResultStation) HasLogo128x128() bool`

HasLogo128x128 returns a boolean if a field has been set.

### GetLogo320x320

`func (o *StationResultStation) GetLogo320x320() string`

GetLogo320x320 returns the Logo320x320 field if non-nil, zero value otherwise.

### GetLogo320x320Ok

`func (o *StationResultStation) GetLogo320x320Ok() (*string, bool)`

GetLogo320x320Ok returns a tuple with the Logo320x320 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo320x320

`func (o *StationResultStation) SetLogo320x320(v string)`

SetLogo320x320 sets Logo320x320 field to given value.

### HasLogo320x320

`func (o *StationResultStation) HasLogo320x320() bool`

HasLogo320x320 returns a boolean if a field has been set.

### GetLogo600x600

`func (o *StationResultStation) GetLogo600x600() string`

GetLogo600x600 returns the Logo600x600 field if non-nil, zero value otherwise.

### GetLogo600x600Ok

`func (o *StationResultStation) GetLogo600x600Ok() (*string, bool)`

GetLogo600x600Ok returns a tuple with the Logo600x600 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo600x600

`func (o *StationResultStation) SetLogo600x600(v string)`

SetLogo600x600 sets Logo600x600 field to given value.

### HasLogo600x600

`func (o *StationResultStation) HasLogo600x600() bool`

HasLogo600x600 returns a boolean if a field has been set.

### GetPayOff

`func (o *StationResultStation) GetPayOff() string`

GetPayOff returns the PayOff field if non-nil, zero value otherwise.

### GetPayOffOk

`func (o *StationResultStation) GetPayOffOk() (*string, bool)`

GetPayOffOk returns a tuple with the PayOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayOff

`func (o *StationResultStation) SetPayOff(v string)`

SetPayOff sets PayOff field to given value.

### HasPayOff

`func (o *StationResultStation) HasPayOff() bool`

HasPayOff returns a boolean if a field has been set.

### GetPtyCode

`func (o *StationResultStation) GetPtyCode() int32`

GetPtyCode returns the PtyCode field if non-nil, zero value otherwise.

### GetPtyCodeOk

`func (o *StationResultStation) GetPtyCodeOk() (*int32, bool)`

GetPtyCodeOk returns a tuple with the PtyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtyCode

`func (o *StationResultStation) SetPtyCode(v int32)`

SetPtyCode sets PtyCode field to given value.

### HasPtyCode

`func (o *StationResultStation) HasPtyCode() bool`

HasPtyCode returns a boolean if a field has been set.

### GetPtyType

`func (o *StationResultStation) GetPtyType() string`

GetPtyType returns the PtyType field if non-nil, zero value otherwise.

### GetPtyTypeOk

`func (o *StationResultStation) GetPtyTypeOk() (*string, bool)`

GetPtyTypeOk returns a tuple with the PtyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtyType

`func (o *StationResultStation) SetPtyType(v string)`

SetPtyType sets PtyType field to given value.

### HasPtyType

`func (o *StationResultStation) HasPtyType() bool`

HasPtyType returns a boolean if a field has been set.

### GetStationKey

`func (o *StationResultStation) GetStationKey() string`

GetStationKey returns the StationKey field if non-nil, zero value otherwise.

### GetStationKeyOk

`func (o *StationResultStation) GetStationKeyOk() (*string, bool)`

GetStationKeyOk returns a tuple with the StationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationKey

`func (o *StationResultStation) SetStationKey(v string)`

SetStationKey sets StationKey field to given value.

### HasStationKey

`func (o *StationResultStation) HasStationKey() bool`

HasStationKey returns a boolean if a field has been set.

### GetTimezone

`func (o *StationResultStation) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *StationResultStation) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *StationResultStation) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *StationResultStation) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetMetadataradioOrganisation

`func (o *StationResultStation) GetMetadataradioOrganisation() string`

GetMetadataradioOrganisation returns the MetadataradioOrganisation field if non-nil, zero value otherwise.

### GetMetadataradioOrganisationOk

`func (o *StationResultStation) GetMetadataradioOrganisationOk() (*string, bool)`

GetMetadataradioOrganisationOk returns a tuple with the MetadataradioOrganisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataradioOrganisation

`func (o *StationResultStation) SetMetadataradioOrganisation(v string)`

SetMetadataradioOrganisation sets MetadataradioOrganisation field to given value.

### HasMetadataradioOrganisation

`func (o *StationResultStation) HasMetadataradioOrganisation() bool`

HasMetadataradioOrganisation returns a boolean if a field has been set.

### GetMetadataradioStationId

`func (o *StationResultStation) GetMetadataradioStationId() string`

GetMetadataradioStationId returns the MetadataradioStationId field if non-nil, zero value otherwise.

### GetMetadataradioStationIdOk

`func (o *StationResultStation) GetMetadataradioStationIdOk() (*string, bool)`

GetMetadataradioStationIdOk returns a tuple with the MetadataradioStationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataradioStationId

`func (o *StationResultStation) SetMetadataradioStationId(v string)`

SetMetadataradioStationId sets MetadataradioStationId field to given value.

### HasMetadataradioStationId

`func (o *StationResultStation) HasMetadataradioStationId() bool`

HasMetadataradioStationId returns a boolean if a field has been set.

### GetStartDays

`func (o *StationResultStation) GetStartDays() StationResultStationStartDays`

GetStartDays returns the StartDays field if non-nil, zero value otherwise.

### GetStartDaysOk

`func (o *StationResultStation) GetStartDaysOk() (*StationResultStationStartDays, bool)`

GetStartDaysOk returns a tuple with the StartDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDays

`func (o *StationResultStation) SetStartDays(v StationResultStationStartDays)`

SetStartDays sets StartDays field to given value.

### HasStartDays

`func (o *StationResultStation) HasStartDays() bool`

HasStartDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


