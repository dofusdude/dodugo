# Go API client for dodugo

# A project for you - the developer.
The free, always-up-to-date, low-latency, insert-buzzword-here Ankama API for your next cool project!

## Client SDKs
Don't write types or functions yourself - I already (kinda) did! 😉
- [Javascript](https://github.com/dofusdude/dofusdude-js) npm i dofusdude-js --save
- [Typescript](https://github.com/dofusdude/dofusdude-ts) npm i dofusdude-ts --save
- [Go](https://github.com/dofusdude/dodugo) go get -u github.com/dofusdude/dodugo
- [Python](https://github.com/dofusdude/dofusdude-py) pip install dofusdude
- [PHP](https://github.com/dofusdude/dofusdude-php)

Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.

Your favorite language is missing? Please let me know!

# Main Features
- 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 2 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.

- ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.

- 📨 **Discord Integration** Ankama related Twitter, RSS and Almanax feeds to post to Discord servers with advanced features like filters or mentions. Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.

- 🩸 **Dofus 2 Beta** from stable to bleeding edge by replacing /dofus2 with /dofus2beta.

- 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_ including the dropped languages from the Dofus website _de_ and _it_.

- 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing by the powerful [Meilisearch](https://www.meilisearch.com) written in Rust.

- 🕵️ **Complete** actual data from the game including items invisible to the encyclopedia like quest items.

- 🖼️ **HD Images** rendering vector graphics into PNGs up to 800x800 px in the background.


## Current state
- Weapons ✅
- Equipment ✅
- Sets ✅
- Resources ✅
- Consumables ✅
- Pets ✅
- Mounts ✅
- Cosmetics/Ceremonial Items ✅
- Harnesses ✅
- Quest Items ✅
- Almanax ✅
- Monsters ❌
- Spells ❌

... and much more on the Roadmap on my Discord. 

## Deploy now. Use forever.
Everything you see here on this site, you can use now and forever. Updates could introduce new fields, new paths or parameter but never break backwards compatibility, so no field or parameter will be deleted.

There is one exception! **The API will _always_ choose being up-to-date over everything else**. So if Ankama decides to drop languages from the game like they did with their website, the API will loose support for them, too.

## Only the beginning... 🤯
I want this project to be useful and not just add plain GET-categories no one needs.

There is a long list of features I want to add (see the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h)). But they are all focussed on you, the developers. So please let me know what you need. I will change the list based on demand.

# Get started! 🥳
Scroll down and try it for yourself!

Or see how these other awesome projects use it:
- [KaellyBot](https://github.com/Kaysoro/KaellyBot) by Kaysoro
- [Dofus Craftlist](https://dofuscraftlist-dev.netlify.app) by Lystina
- [AlmanaxApp](https://almanaxapp.netlify.app) by Lystina

I highly recommend using the SDKs for quick results. I use them myself for microservices for the API.

## Thank you!
I highly welcome everyone on my [Discord](https://discord.gg/3EtHskZD8h) to just talk about projects and use cases or give feedback of any kind.

The servers have a fixed monthly cost to provide very fast responses. If you want to help me keeping them running or simply  donate, consider becoming a [GitHub Sponsor](https://github.com/sponsors/dofusdude).


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.7.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://discord.gg/3EtHskZD8h](https://discord.gg/3EtHskZD8h)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import dodugo "github.com/dofusdude/dodugo"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), dodugo.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), dodugo.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
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
*AllItemsApi* | [**GetItemsAllSearch**](docs/AllItemsApi.md#getitemsallsearch) | **Get** /{game}/{language}/items/search | Search All Items
*AlmanaxApi* | [**GetAlmanaxDate**](docs/AlmanaxApi.md#getalmanaxdate) | **Get** /dofus2/{language}/almanax/{date} | Single Almanax Date
*AlmanaxApi* | [**GetAlmanaxRange**](docs/AlmanaxApi.md#getalmanaxrange) | **Get** /dofus2/{language}/almanax | Almanax Range
*ConsumablesApi* | [**GetAllItemsConsumablesList**](docs/ConsumablesApi.md#getallitemsconsumableslist) | **Get** /{game}/{language}/items/consumables/all | List All Consumables
*ConsumablesApi* | [**GetItemsConsumablesList**](docs/ConsumablesApi.md#getitemsconsumableslist) | **Get** /{game}/{language}/items/consumables | List Consumables
*ConsumablesApi* | [**GetItemsConsumablesSearch**](docs/ConsumablesApi.md#getitemsconsumablessearch) | **Get** /{game}/{language}/items/consumables/search | Search Consumables
*ConsumablesApi* | [**GetItemsConsumablesSingle**](docs/ConsumablesApi.md#getitemsconsumablessingle) | **Get** /{game}/{language}/items/consumables/{ankama_id} | Single Consumables
*CosmeticsApi* | [**GetAllCosmeticsList**](docs/CosmeticsApi.md#getallcosmeticslist) | **Get** /{game}/{language}/items/cosmetics/all | List All Cosmetics
*CosmeticsApi* | [**GetCosmeticsList**](docs/CosmeticsApi.md#getcosmeticslist) | **Get** /{game}/{language}/items/cosmetics | List Cosmetics
*CosmeticsApi* | [**GetCosmeticsSearch**](docs/CosmeticsApi.md#getcosmeticssearch) | **Get** /{game}/{language}/items/cosmetics/search | Search Cosmetics
*CosmeticsApi* | [**GetCosmeticsSingle**](docs/CosmeticsApi.md#getcosmeticssingle) | **Get** /{game}/{language}/items/cosmetics/{ankama_id} | Single Cosmetics
*EquipmentApi* | [**GetAllItemsEquipmentList**](docs/EquipmentApi.md#getallitemsequipmentlist) | **Get** /{game}/{language}/items/equipment/all | List All Equipment
*EquipmentApi* | [**GetItemsEquipmentList**](docs/EquipmentApi.md#getitemsequipmentlist) | **Get** /{game}/{language}/items/equipment | List Equipment
*EquipmentApi* | [**GetItemsEquipmentSearch**](docs/EquipmentApi.md#getitemsequipmentsearch) | **Get** /{game}/{language}/items/equipment/search | Search Equipment
*EquipmentApi* | [**GetItemsEquipmentSingle**](docs/EquipmentApi.md#getitemsequipmentsingle) | **Get** /{game}/{language}/items/equipment/{ankama_id} | Single Equipment
*MetaApi* | [**GetMetaAlmanaxBonuses**](docs/MetaApi.md#getmetaalmanaxbonuses) | **Get** /dofus2/meta/{language}/almanax/bonuses | Available Almanax Bonuses
*MetaApi* | [**GetMetaElements**](docs/MetaApi.md#getmetaelements) | **Get** /dofus2/meta/elements | Effects and Condition Elements
*MountsApi* | [**GetAllMountsList**](docs/MountsApi.md#getallmountslist) | **Get** /{game}/{language}/mounts/all | List All Mounts
*MountsApi* | [**GetMountsList**](docs/MountsApi.md#getmountslist) | **Get** /{game}/{language}/mounts | List Mounts
*MountsApi* | [**GetMountsSearch**](docs/MountsApi.md#getmountssearch) | **Get** /{game}/{language}/mounts/search | Search Mounts
*MountsApi* | [**GetMountsSingle**](docs/MountsApi.md#getmountssingle) | **Get** /{game}/{language}/mounts/{ankama_id} | Single Mounts
*QuestItemsApi* | [**GetAllItemsQuestList**](docs/QuestItemsApi.md#getallitemsquestlist) | **Get** /{game}/{language}/items/quest/all | List All Quest Items
*QuestItemsApi* | [**GetItemQuestSingle**](docs/QuestItemsApi.md#getitemquestsingle) | **Get** /{game}/{language}/items/quest/{ankama_id} | Single Quest Items
*QuestItemsApi* | [**GetItemsQuestList**](docs/QuestItemsApi.md#getitemsquestlist) | **Get** /{game}/{language}/items/quest | List Quest Items
*QuestItemsApi* | [**GetItemsQuestSearch**](docs/QuestItemsApi.md#getitemsquestsearch) | **Get** /{game}/{language}/items/quest/search | Search Quest Items
*ResourcesApi* | [**GetAllItemsResourcesList**](docs/ResourcesApi.md#getallitemsresourceslist) | **Get** /{game}/{language}/items/resources/all | List All Resources
*ResourcesApi* | [**GetItemsResourceSearch**](docs/ResourcesApi.md#getitemsresourcesearch) | **Get** /{game}/{language}/items/resources/search | Search Resources
*ResourcesApi* | [**GetItemsResourcesList**](docs/ResourcesApi.md#getitemsresourceslist) | **Get** /{game}/{language}/items/resources | List Resources
*ResourcesApi* | [**GetItemsResourcesSingle**](docs/ResourcesApi.md#getitemsresourcessingle) | **Get** /{game}/{language}/items/resources/{ankama_id} | Single Resources
*SetsApi* | [**GetAllSetsList**](docs/SetsApi.md#getallsetslist) | **Get** /{game}/{language}/sets/all | List All Sets
*SetsApi* | [**GetSetsList**](docs/SetsApi.md#getsetslist) | **Get** /{game}/{language}/sets | List Sets
*SetsApi* | [**GetSetsSearch**](docs/SetsApi.md#getsetssearch) | **Get** /{game}/{language}/sets/search | Search Sets
*SetsApi* | [**GetSetsSingle**](docs/SetsApi.md#getsetssingle) | **Get** /{game}/{language}/sets/{ankama_id} | Single Sets
*WebhooksApi* | [**DeleteWebhooksAlmanaxId**](docs/WebhooksApi.md#deletewebhooksalmanaxid) | **Delete** /webhooks/almanax/{id} | Unregister Almanax Hook
*WebhooksApi* | [**DeleteWebhooksRssId**](docs/WebhooksApi.md#deletewebhooksrssid) | **Delete** /webhooks/rss/{id} | Unregister RSS Hook
*WebhooksApi* | [**DeleteWebhooksTwitterId**](docs/WebhooksApi.md#deletewebhookstwitterid) | **Delete** /webhooks/twitter/{id} | Unregister Twitter Hook
*WebhooksApi* | [**GetMetaWebhooksAlmanax**](docs/WebhooksApi.md#getmetawebhooksalmanax) | **Get** /meta/webhooks/almanax | Get Almanax Hook Metainfo
*WebhooksApi* | [**GetMetaWebhooksRss**](docs/WebhooksApi.md#getmetawebhooksrss) | **Get** /meta/webhooks/rss | Get RSS Hook Metainfo
*WebhooksApi* | [**GetMetaWebhooksTwitter**](docs/WebhooksApi.md#getmetawebhookstwitter) | **Get** /meta/webhooks/twitter | Get Twitter Hook Metainfo
*WebhooksApi* | [**GetWebhooksAlmanaxId**](docs/WebhooksApi.md#getwebhooksalmanaxid) | **Get** /webhooks/almanax/{id} | Get Almanax Hook
*WebhooksApi* | [**GetWebhooksRssId**](docs/WebhooksApi.md#getwebhooksrssid) | **Get** /webhooks/rss/{id} | Get RSS Hook
*WebhooksApi* | [**GetWebhooksTwitterId**](docs/WebhooksApi.md#getwebhookstwitterid) | **Get** /webhooks/twitter/{id} | Get Twitter Hook
*WebhooksApi* | [**PostWebhooksAlmanax**](docs/WebhooksApi.md#postwebhooksalmanax) | **Post** /webhooks/almanax | Register Almanax Hook
*WebhooksApi* | [**PostWebhooksRss**](docs/WebhooksApi.md#postwebhooksrss) | **Post** /webhooks/rss | Register RSS Hook
*WebhooksApi* | [**PostWebhooksTwitter**](docs/WebhooksApi.md#postwebhookstwitter) | **Post** /webhooks/twitter | Register Twitter Hook
*WebhooksApi* | [**PutWebhooksAlmanaxId**](docs/WebhooksApi.md#putwebhooksalmanaxid) | **Put** /webhooks/almanax/{id} | Update Almanax Hook
*WebhooksApi* | [**PutWebhooksRssId**](docs/WebhooksApi.md#putwebhooksrssid) | **Put** /webhooks/rss/{id} | Update RSS Hook
*WebhooksApi* | [**PutWebhooksTwitterId**](docs/WebhooksApi.md#putwebhookstwitterid) | **Put** /webhooks/twitter/{id} | Update Twitter Hook


## Documentation For Models

 - [AlmanaxEntry](docs/AlmanaxEntry.md)
 - [AlmanaxEntryBonus](docs/AlmanaxEntryBonus.md)
 - [AlmanaxEntryTribute](docs/AlmanaxEntryTribute.md)
 - [AlmanaxEntryTributeItem](docs/AlmanaxEntryTributeItem.md)
 - [AlmanaxWebhook](docs/AlmanaxWebhook.md)
 - [AlmanaxWebhookDailySettings](docs/AlmanaxWebhookDailySettings.md)
 - [ConditionEntry](docs/ConditionEntry.md)
 - [ConditionEntryElement](docs/ConditionEntryElement.md)
 - [Cosmetic](docs/Cosmetic.md)
 - [CreateAlmanaxWebhook](docs/CreateAlmanaxWebhook.md)
 - [CreateAlmanaxWebhookDailySettings](docs/CreateAlmanaxWebhookDailySettings.md)
 - [CreateAlmanaxWebhookMentionsValueInner](docs/CreateAlmanaxWebhookMentionsValueInner.md)
 - [CreateRSSWebhook](docs/CreateRSSWebhook.md)
 - [CreateTwitterWebhook](docs/CreateTwitterWebhook.md)
 - [EffectsEntry](docs/EffectsEntry.md)
 - [EffectsEntryType](docs/EffectsEntryType.md)
 - [Equipment](docs/Equipment.md)
 - [EquipmentParentSet](docs/EquipmentParentSet.md)
 - [EquipmentSet](docs/EquipmentSet.md)
 - [GetMetaAlmanaxBonuses200ResponseInner](docs/GetMetaAlmanaxBonuses200ResponseInner.md)
 - [GetMetaWebhooksTwitter200Response](docs/GetMetaWebhooksTwitter200Response.md)
 - [ImageUrls](docs/ImageUrls.md)
 - [ItemListEntry](docs/ItemListEntry.md)
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
 - [RssWebhook](docs/RssWebhook.md)
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

