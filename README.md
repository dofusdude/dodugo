# Go API client for dodugo

# A project for you - the developer.
The all-in-one toolbelt for your next Ankama related project.

## Client SDKs
- [Javascript](https://github.com/dofusdude/dofusdude-js) `npm i dofusdude-js --save`
- [Typescript](https://github.com/dofusdude/dofusdude-ts) `npm i dofusdude-ts --save`
- [Go](https://github.com/dofusdude/dodugo) `go get -u github.com/dofusdude/dodugo`
- [Python](https://github.com/dofusdude/dofusdude-py) `pip install dofusdude`
- [PHP](https://github.com/dofusdude/dofusdude-php)
- [Java](https://github.com/dofusdude/dofusdude-java) Maven with GitHub packages setup

Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.

Your favorite language is missing? Please let me know!

# Main Features
- 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 10 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.

- ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.

- 📨 **Discord Integration** Ankama related RSS and Almanax feeds to post to Discord servers with advanced features like filters or mentions. Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.

- 🩸 **Dofus 2 Beta** from stable to bleeding edge by replacing /dofus2 with /dofus2beta.

- 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_ including the dropped languages from the Dofus website _de_ and _it_.

- 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing.

- 🕵️ **Complete** actual data from the game including items invisible to the encyclopedia like quest items.

- 🖼️ **HD Images** rendering game assets to high-res images with up to 800x800 px.

... and much more on the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h).


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.9.4
- Package version: 1.0.0
- Generator version: 7.10.0-SNAPSHOT
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://discord.gg/3EtHskZD8h](https://discord.gg/3EtHskZD8h)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import dodugo "github.com/dofusdude/dodugo"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `dodugo.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), dodugo.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `dodugo.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), dodugo.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `dodugo.ContextOperationServerIndices` and `dodugo.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), dodugo.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), dodugo.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.dofusdu.de*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AlmanaxAPI* | [**GetAlmanaxDate**](docs/AlmanaxAPI.md#getalmanaxdate) | **Get** /dofus2/{language}/almanax/{date} | Single Almanax Date
*AlmanaxAPI* | [**GetAlmanaxRange**](docs/AlmanaxAPI.md#getalmanaxrange) | **Get** /dofus2/{language}/almanax | Almanax Range
*ConsumablesAPI* | [**GetAllItemsConsumablesList**](docs/ConsumablesAPI.md#getallitemsconsumableslist) | **Get** /{game}/{language}/items/consumables/all | List All Consumables
*ConsumablesAPI* | [**GetItemsConsumablesList**](docs/ConsumablesAPI.md#getitemsconsumableslist) | **Get** /{game}/{language}/items/consumables | List Consumables
*ConsumablesAPI* | [**GetItemsConsumablesSearch**](docs/ConsumablesAPI.md#getitemsconsumablessearch) | **Get** /{game}/{language}/items/consumables/search | Search Consumables
*ConsumablesAPI* | [**GetItemsConsumablesSingle**](docs/ConsumablesAPI.md#getitemsconsumablessingle) | **Get** /{game}/{language}/items/consumables/{ankama_id} | Single Consumables
*CosmeticsAPI* | [**GetAllCosmeticsList**](docs/CosmeticsAPI.md#getallcosmeticslist) | **Get** /{game}/{language}/items/cosmetics/all | List All Cosmetics
*CosmeticsAPI* | [**GetCosmeticsList**](docs/CosmeticsAPI.md#getcosmeticslist) | **Get** /{game}/{language}/items/cosmetics | List Cosmetics
*CosmeticsAPI* | [**GetCosmeticsSearch**](docs/CosmeticsAPI.md#getcosmeticssearch) | **Get** /{game}/{language}/items/cosmetics/search | Search Cosmetics
*CosmeticsAPI* | [**GetCosmeticsSingle**](docs/CosmeticsAPI.md#getcosmeticssingle) | **Get** /{game}/{language}/items/cosmetics/{ankama_id} | Single Cosmetics
*EquipmentAPI* | [**GetAllItemsEquipmentList**](docs/EquipmentAPI.md#getallitemsequipmentlist) | **Get** /{game}/{language}/items/equipment/all | List All Equipment
*EquipmentAPI* | [**GetItemsEquipmentList**](docs/EquipmentAPI.md#getitemsequipmentlist) | **Get** /{game}/{language}/items/equipment | List Equipment
*EquipmentAPI* | [**GetItemsEquipmentSearch**](docs/EquipmentAPI.md#getitemsequipmentsearch) | **Get** /{game}/{language}/items/equipment/search | Search Equipment
*EquipmentAPI* | [**GetItemsEquipmentSingle**](docs/EquipmentAPI.md#getitemsequipmentsingle) | **Get** /{game}/{language}/items/equipment/{ankama_id} | Single Equipment
*GameAPI* | [**GetGameSearch**](docs/GameAPI.md#getgamesearch) | **Get** /{game}/{language}/search | Game Search
*GameAPI* | [**GetItemsAllSearch**](docs/GameAPI.md#getitemsallsearch) | **Get** /{game}/{language}/items/search | Search All Items
*MetaAPI* | [**GetGameSearchTypes**](docs/MetaAPI.md#getgamesearchtypes) | **Get** /dofus2/meta/search/types | Available Game Search Types
*MetaAPI* | [**GetItemTypes**](docs/MetaAPI.md#getitemtypes) | **Get** /dofus2/meta/items/types | Available Item Types
*MetaAPI* | [**GetMetaAlmanaxBonuses**](docs/MetaAPI.md#getmetaalmanaxbonuses) | **Get** /dofus2/meta/{language}/almanax/bonuses | Available Almanax Bonuses
*MetaAPI* | [**GetMetaAlmanaxBonusesSearch**](docs/MetaAPI.md#getmetaalmanaxbonusessearch) | **Get** /dofus2/meta/{language}/almanax/bonuses/search | Search Available Almanax Bonuses
*MetaAPI* | [**GetMetaElements**](docs/MetaAPI.md#getmetaelements) | **Get** /dofus2/meta/elements | Effects and Condition Elements
*MetaAPI* | [**GetMetaVersion**](docs/MetaAPI.md#getmetaversion) | **Get** /dofus2/meta/version | Game Version
*MountsAPI* | [**GetAllMountsList**](docs/MountsAPI.md#getallmountslist) | **Get** /{game}/{language}/mounts/all | List All Mounts
*MountsAPI* | [**GetMountsList**](docs/MountsAPI.md#getmountslist) | **Get** /{game}/{language}/mounts | List Mounts
*MountsAPI* | [**GetMountsSearch**](docs/MountsAPI.md#getmountssearch) | **Get** /{game}/{language}/mounts/search | Search Mounts
*MountsAPI* | [**GetMountsSingle**](docs/MountsAPI.md#getmountssingle) | **Get** /{game}/{language}/mounts/{ankama_id} | Single Mounts
*QuestItemsAPI* | [**GetAllItemsQuestList**](docs/QuestItemsAPI.md#getallitemsquestlist) | **Get** /{game}/{language}/items/quest/all | List All Quest Items
*QuestItemsAPI* | [**GetItemQuestSingle**](docs/QuestItemsAPI.md#getitemquestsingle) | **Get** /{game}/{language}/items/quest/{ankama_id} | Single Quest Items
*QuestItemsAPI* | [**GetItemsQuestList**](docs/QuestItemsAPI.md#getitemsquestlist) | **Get** /{game}/{language}/items/quest | List Quest Items
*QuestItemsAPI* | [**GetItemsQuestSearch**](docs/QuestItemsAPI.md#getitemsquestsearch) | **Get** /{game}/{language}/items/quest/search | Search Quest Items
*ResourcesAPI* | [**GetAllItemsResourcesList**](docs/ResourcesAPI.md#getallitemsresourceslist) | **Get** /{game}/{language}/items/resources/all | List All Resources
*ResourcesAPI* | [**GetItemsResourceSearch**](docs/ResourcesAPI.md#getitemsresourcesearch) | **Get** /{game}/{language}/items/resources/search | Search Resources
*ResourcesAPI* | [**GetItemsResourcesList**](docs/ResourcesAPI.md#getitemsresourceslist) | **Get** /{game}/{language}/items/resources | List Resources
*ResourcesAPI* | [**GetItemsResourcesSingle**](docs/ResourcesAPI.md#getitemsresourcessingle) | **Get** /{game}/{language}/items/resources/{ankama_id} | Single Resources
*SetsAPI* | [**GetAllSetsList**](docs/SetsAPI.md#getallsetslist) | **Get** /{game}/{language}/sets/all | List All Sets
*SetsAPI* | [**GetSetsList**](docs/SetsAPI.md#getsetslist) | **Get** /{game}/{language}/sets | List Sets
*SetsAPI* | [**GetSetsSearch**](docs/SetsAPI.md#getsetssearch) | **Get** /{game}/{language}/sets/search | Search Sets
*SetsAPI* | [**GetSetsSingle**](docs/SetsAPI.md#getsetssingle) | **Get** /{game}/{language}/sets/{ankama_id} | Single Sets
*WebhooksAPI* | [**DeleteWebhooksAlmanaxId**](docs/WebhooksAPI.md#deletewebhooksalmanaxid) | **Delete** /webhooks/almanax/{id} | Unregister Almanax Hook
*WebhooksAPI* | [**DeleteWebhooksRssId**](docs/WebhooksAPI.md#deletewebhooksrssid) | **Delete** /webhooks/rss/{id} | Unregister RSS Hook
*WebhooksAPI* | [**DeleteWebhooksTwitterId**](docs/WebhooksAPI.md#deletewebhookstwitterid) | **Delete** /webhooks/twitter/{id} | Unregister Twitter Hook
*WebhooksAPI* | [**GetMetaWebhooksAlmanax**](docs/WebhooksAPI.md#getmetawebhooksalmanax) | **Get** /meta/webhooks/almanax | Get Almanax Hook Metainfo
*WebhooksAPI* | [**GetMetaWebhooksRss**](docs/WebhooksAPI.md#getmetawebhooksrss) | **Get** /meta/webhooks/rss | Get RSS Hook Metainfo
*WebhooksAPI* | [**GetMetaWebhooksTwitter**](docs/WebhooksAPI.md#getmetawebhookstwitter) | **Get** /meta/webhooks/twitter | Get Twitter Hook Metainfo
*WebhooksAPI* | [**GetWebhooksAlmanaxId**](docs/WebhooksAPI.md#getwebhooksalmanaxid) | **Get** /webhooks/almanax/{id} | Get Almanax Hook
*WebhooksAPI* | [**GetWebhooksRssId**](docs/WebhooksAPI.md#getwebhooksrssid) | **Get** /webhooks/rss/{id} | Get RSS Hook
*WebhooksAPI* | [**GetWebhooksTwitterId**](docs/WebhooksAPI.md#getwebhookstwitterid) | **Get** /webhooks/twitter/{id} | Get Twitter Hook
*WebhooksAPI* | [**PostWebhooksAlmanax**](docs/WebhooksAPI.md#postwebhooksalmanax) | **Post** /webhooks/almanax | Register Almanax Hook
*WebhooksAPI* | [**PostWebhooksRss**](docs/WebhooksAPI.md#postwebhooksrss) | **Post** /webhooks/rss | Register RSS Hook
*WebhooksAPI* | [**PostWebhooksTwitter**](docs/WebhooksAPI.md#postwebhookstwitter) | **Post** /webhooks/twitter | Register Twitter Hook
*WebhooksAPI* | [**PutWebhooksAlmanaxId**](docs/WebhooksAPI.md#putwebhooksalmanaxid) | **Put** /webhooks/almanax/{id} | Update Almanax Hook
*WebhooksAPI* | [**PutWebhooksRssId**](docs/WebhooksAPI.md#putwebhooksrssid) | **Put** /webhooks/rss/{id} | Update RSS Hook
*WebhooksAPI* | [**PutWebhooksTwitterId**](docs/WebhooksAPI.md#putwebhookstwitterid) | **Put** /webhooks/twitter/{id} | Update Twitter Hook


## Documentation For Models

 - [AlmanaxEntry](docs/AlmanaxEntry.md)
 - [AlmanaxEntryBonus](docs/AlmanaxEntryBonus.md)
 - [AlmanaxEntryTribute](docs/AlmanaxEntryTribute.md)
 - [AlmanaxEntryTributeItem](docs/AlmanaxEntryTributeItem.md)
 - [AlmanaxWebhook](docs/AlmanaxWebhook.md)
 - [AlmanaxWebhookDailySettings](docs/AlmanaxWebhookDailySettings.md)
 - [ConditionEntry](docs/ConditionEntry.md)
 - [ConditionTreeLeaf](docs/ConditionTreeLeaf.md)
 - [ConditionTreeNode](docs/ConditionTreeNode.md)
 - [ConditionTreeRelation](docs/ConditionTreeRelation.md)
 - [CreateAlmanaxWebhook](docs/CreateAlmanaxWebhook.md)
 - [CreateAlmanaxWebhookDailySettings](docs/CreateAlmanaxWebhookDailySettings.md)
 - [CreateAlmanaxWebhookMentionsValueInner](docs/CreateAlmanaxWebhookMentionsValueInner.md)
 - [CreateRSSWebhook](docs/CreateRSSWebhook.md)
 - [CreateTwitterWebhook](docs/CreateTwitterWebhook.md)
 - [EffectsEntry](docs/EffectsEntry.md)
 - [Equipment](docs/Equipment.md)
 - [EquipmentSet](docs/EquipmentSet.md)
 - [GetGameSearch200ResponseInner](docs/GetGameSearch200ResponseInner.md)
 - [GetMetaAlmanaxBonuses200ResponseInner](docs/GetMetaAlmanaxBonuses200ResponseInner.md)
 - [GetMetaVersion200Response](docs/GetMetaVersion200Response.md)
 - [GetMetaWebhooksTwitter200Response](docs/GetMetaWebhooksTwitter200Response.md)
 - [ImageUrls](docs/ImageUrls.md)
 - [ItemListEntry](docs/ItemListEntry.md)
 - [ItemListEntryParentSet](docs/ItemListEntryParentSet.md)
 - [ItemListEntryRange](docs/ItemListEntryRange.md)
 - [ItemsListEntryTyped](docs/ItemsListEntryTyped.md)
 - [ItemsListEntryTypedType](docs/ItemsListEntryTypedType.md)
 - [ItemsListPaged](docs/ItemsListPaged.md)
 - [LinksPaged](docs/LinksPaged.md)
 - [Mount](docs/Mount.md)
 - [MountListEntry](docs/MountListEntry.md)
 - [MountsListPaged](docs/MountsListPaged.md)
 - [PutAlmanaxWebhook](docs/PutAlmanaxWebhook.md)
 - [PutRSSWebhook](docs/PutRSSWebhook.md)
 - [PutTwitterWebhook](docs/PutTwitterWebhook.md)
 - [RecipeEntry](docs/RecipeEntry.md)
 - [Resource](docs/Resource.md)
 - [ResourceType](docs/ResourceType.md)
 - [RssWebhook](docs/RssWebhook.md)
 - [SetEffectsEntry](docs/SetEffectsEntry.md)
 - [SetEffectsEntryType](docs/SetEffectsEntryType.md)
 - [SetListEntry](docs/SetListEntry.md)
 - [SetsListPaged](docs/SetsListPaged.md)
 - [TwitterWebhook](docs/TwitterWebhook.md)
 - [Weapon](docs/Weapon.md)
 - [WeaponRange](docs/WeaponRange.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

stelzo@steado.de

