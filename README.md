# Go API client for dodugo

# Open Ankama Developer Community
The all-in-one toolbelt for your next Ankama related project.

## Versions
- [Dofus 2](https://docs.dofusdu.de/dofus2/)
- [Dofus 3](https://docs.dofusdu.de/dofus3/)
  - v1 [latest] (you are here) 

## Client SDKs
- [Javascript](https://github.com/dofusdude/dofusdude-js) `npm i dofusdude-js --save`
- [Typescript](https://github.com/dofusdude/dofusdude-ts) `npm i dofusdude-ts --save`
- [Go](https://github.com/dofusdude/dodugo) `go get -u github.com/dofusdude/dodugo`
- [Python](https://github.com/dofusdude/dofusdude-py) `pip install dofusdude`
- [Java](https://github.com/dofusdude/dofusdude-java) Maven with GitHub packages setup

Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.

Your favorite language is missing? Please let me know!

# Main Features
- 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 10 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.

- ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.

- 📨 **Almanax Discord Integration** Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.

- 🩸 **Dofus 3 Beta** from stable to bleeding edge by replacing /dofus3 with /dofus3beta.

- 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_, _de_.

- 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing.

- 🕵️ **Official Sources** generated from actual data from the game.

... and much more on the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h).


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0-rc.3
- Package version: 1.0.0
- Generator version: 7.11.0-SNAPSHOT
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
*ConsumablesAPI* | [**GetAllItemsConsumablesList**](docs/ConsumablesAPI.md#getallitemsconsumableslist) | **Get** /{game}/v1/{language}/items/consumables/all | List All Consumables
*ConsumablesAPI* | [**GetItemsConsumablesList**](docs/ConsumablesAPI.md#getitemsconsumableslist) | **Get** /{game}/v1/{language}/items/consumables | List Consumables
*ConsumablesAPI* | [**GetItemsConsumablesSearch**](docs/ConsumablesAPI.md#getitemsconsumablessearch) | **Get** /{game}/v1/{language}/items/consumables/search | Search Consumables
*ConsumablesAPI* | [**GetItemsConsumablesSingle**](docs/ConsumablesAPI.md#getitemsconsumablessingle) | **Get** /{game}/v1/{language}/items/consumables/{ankama_id} | Single Consumables
*CosmeticsAPI* | [**GetAllCosmeticsList**](docs/CosmeticsAPI.md#getallcosmeticslist) | **Get** /{game}/v1/{language}/items/cosmetics/all | List All Cosmetics
*CosmeticsAPI* | [**GetCosmeticsList**](docs/CosmeticsAPI.md#getcosmeticslist) | **Get** /{game}/v1/{language}/items/cosmetics | List Cosmetics
*CosmeticsAPI* | [**GetCosmeticsSearch**](docs/CosmeticsAPI.md#getcosmeticssearch) | **Get** /{game}/v1/{language}/items/cosmetics/search | Search Cosmetics
*CosmeticsAPI* | [**GetCosmeticsSingle**](docs/CosmeticsAPI.md#getcosmeticssingle) | **Get** /{game}/v1/{language}/items/cosmetics/{ankama_id} | Single Cosmetics
*EquipmentAPI* | [**GetAllItemsEquipmentList**](docs/EquipmentAPI.md#getallitemsequipmentlist) | **Get** /{game}/v1/{language}/items/equipment/all | List All Equipment
*EquipmentAPI* | [**GetItemsEquipmentList**](docs/EquipmentAPI.md#getitemsequipmentlist) | **Get** /{game}/v1/{language}/items/equipment | List Equipment
*EquipmentAPI* | [**GetItemsEquipmentSearch**](docs/EquipmentAPI.md#getitemsequipmentsearch) | **Get** /{game}/v1/{language}/items/equipment/search | Search Equipment
*EquipmentAPI* | [**GetItemsEquipmentSingle**](docs/EquipmentAPI.md#getitemsequipmentsingle) | **Get** /{game}/v1/{language}/items/equipment/{ankama_id} | Single Equipment
*GameAPI* | [**GetGameSearch**](docs/GameAPI.md#getgamesearch) | **Get** /{game}/v1/{language}/search | Game Search
*GameAPI* | [**GetItemsAllSearch**](docs/GameAPI.md#getitemsallsearch) | **Get** /{game}/v1/{language}/items/search | Search All Items
*MetaAPI* | [**GetGameSearchTypes**](docs/MetaAPI.md#getgamesearchtypes) | **Get** /dofus3beta/v1/meta/search/types | Available Game Search Types
*MetaAPI* | [**GetItemTypes**](docs/MetaAPI.md#getitemtypes) | **Get** /dofus3beta/v1/meta/items/types | Available Item Types
*MetaAPI* | [**GetMetaAlmanaxBonuses**](docs/MetaAPI.md#getmetaalmanaxbonuses) | **Get** /dofus2/meta/{language}/almanax/bonuses | Available Almanax Bonuses
*MetaAPI* | [**GetMetaAlmanaxBonusesSearch**](docs/MetaAPI.md#getmetaalmanaxbonusessearch) | **Get** /dofus2/meta/{language}/almanax/bonuses/search | Search Available Almanax Bonuses
*MetaAPI* | [**GetMetaElements**](docs/MetaAPI.md#getmetaelements) | **Get** /dofus3beta/v1/meta/elements | Effects and Condition Elements
*MetaAPI* | [**GetMetaVersion**](docs/MetaAPI.md#getmetaversion) | **Get** /{game}/v1/meta/version | Game Version
*MountsAPI* | [**GetAllMountsList**](docs/MountsAPI.md#getallmountslist) | **Get** /{game}/v1/{language}/mounts/all | List All Mounts
*MountsAPI* | [**GetMountsList**](docs/MountsAPI.md#getmountslist) | **Get** /{game}/v1/{language}/mounts | List Mounts
*MountsAPI* | [**GetMountsSearch**](docs/MountsAPI.md#getmountssearch) | **Get** /{game}/v1/{language}/mounts/search | Search Mounts
*MountsAPI* | [**GetMountsSingle**](docs/MountsAPI.md#getmountssingle) | **Get** /{game}/v1/{language}/mounts/{ankama_id} | Single Mounts
*QuestItemsAPI* | [**GetAllItemsQuestList**](docs/QuestItemsAPI.md#getallitemsquestlist) | **Get** /{game}/v1/{language}/items/quest/all | List All Quest Items
*QuestItemsAPI* | [**GetItemQuestSingle**](docs/QuestItemsAPI.md#getitemquestsingle) | **Get** /{game}/v1/{language}/items/quest/{ankama_id} | Single Quest Items
*QuestItemsAPI* | [**GetItemsQuestList**](docs/QuestItemsAPI.md#getitemsquestlist) | **Get** /{game}/v1/{language}/items/quest | List Quest Items
*QuestItemsAPI* | [**GetItemsQuestSearch**](docs/QuestItemsAPI.md#getitemsquestsearch) | **Get** /{game}/v1/{language}/items/quest/search | Search Quest Items
*ResourcesAPI* | [**GetAllItemsResourcesList**](docs/ResourcesAPI.md#getallitemsresourceslist) | **Get** /{game}/v1/{language}/items/resources/all | List All Resources
*ResourcesAPI* | [**GetItemsResourceSearch**](docs/ResourcesAPI.md#getitemsresourcesearch) | **Get** /{game}/v1/{language}/items/resources/search | Search Resources
*ResourcesAPI* | [**GetItemsResourcesList**](docs/ResourcesAPI.md#getitemsresourceslist) | **Get** /{game}/v1/{language}/items/resources | List Resources
*ResourcesAPI* | [**GetItemsResourcesSingle**](docs/ResourcesAPI.md#getitemsresourcessingle) | **Get** /{game}/v1/{language}/items/resources/{ankama_id} | Single Resources
*SetsAPI* | [**GetAllSetsList**](docs/SetsAPI.md#getallsetslist) | **Get** /{game}/v1/{language}/sets/all | List All Sets
*SetsAPI* | [**GetSetsList**](docs/SetsAPI.md#getsetslist) | **Get** /{game}/v1/{language}/sets | List Sets
*SetsAPI* | [**GetSetsSearch**](docs/SetsAPI.md#getsetssearch) | **Get** /{game}/v1/{language}/sets/search | Search Sets
*SetsAPI* | [**GetSetsSingle**](docs/SetsAPI.md#getsetssingle) | **Get** /{game}/v1/{language}/sets/{ankama_id} | Single Sets
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

 - [Almanax](docs/Almanax.md)
 - [AlmanaxBonus](docs/AlmanaxBonus.md)
 - [AlmanaxTribute](docs/AlmanaxTribute.md)
 - [AlmanaxTributeItem](docs/AlmanaxTributeItem.md)
 - [AlmanaxWebhook](docs/AlmanaxWebhook.md)
 - [AlmanaxWebhookDailySettings](docs/AlmanaxWebhookDailySettings.md)
 - [Condition](docs/Condition.md)
 - [ConditionLeaf](docs/ConditionLeaf.md)
 - [ConditionNode](docs/ConditionNode.md)
 - [ConditionRelation](docs/ConditionRelation.md)
 - [CreateAlmanaxWebhook](docs/CreateAlmanaxWebhook.md)
 - [CreateAlmanaxWebhookDailySettings](docs/CreateAlmanaxWebhookDailySettings.md)
 - [CreateAlmanaxWebhookMentionsValueInner](docs/CreateAlmanaxWebhookMentionsValueInner.md)
 - [CreateRSSWebhook](docs/CreateRSSWebhook.md)
 - [CreateTwitterWebhook](docs/CreateTwitterWebhook.md)
 - [Effect](docs/Effect.md)
 - [EffectType](docs/EffectType.md)
 - [Equipment](docs/Equipment.md)
 - [EquipmentSet](docs/EquipmentSet.md)
 - [Error](docs/Error.md)
 - [GameSearch](docs/GameSearch.md)
 - [GameSearchItem](docs/GameSearchItem.md)
 - [GameSearchType](docs/GameSearchType.md)
 - [GetMetaAlmanaxBonuses200ResponseInner](docs/GetMetaAlmanaxBonuses200ResponseInner.md)
 - [GetMetaWebhooksTwitter200Response](docs/GetMetaWebhooksTwitter200Response.md)
 - [Images](docs/Images.md)
 - [ItemSubtype](docs/ItemSubtype.md)
 - [ListEquipmentSet](docs/ListEquipmentSet.md)
 - [ListEquipmentSets](docs/ListEquipmentSets.md)
 - [ListItem](docs/ListItem.md)
 - [ListItemGeneral](docs/ListItemGeneral.md)
 - [ListItems](docs/ListItems.md)
 - [ListMounts](docs/ListMounts.md)
 - [Mount](docs/Mount.md)
 - [MountFamily](docs/MountFamily.md)
 - [PagedLinks](docs/PagedLinks.md)
 - [PutAlmanaxWebhook](docs/PutAlmanaxWebhook.md)
 - [PutRSSWebhook](docs/PutRSSWebhook.md)
 - [PutTwitterWebhook](docs/PutTwitterWebhook.md)
 - [Range](docs/Range.md)
 - [Recipe](docs/Recipe.md)
 - [Resource](docs/Resource.md)
 - [RssWebhook](docs/RssWebhook.md)
 - [TranslatedId](docs/TranslatedId.md)
 - [TwitterWebhook](docs/TwitterWebhook.md)
 - [Version](docs/Version.md)
 - [Weapon](docs/Weapon.md)


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

