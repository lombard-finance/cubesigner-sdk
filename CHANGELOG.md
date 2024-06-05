<!--
Guiding Principles:

Changelogs are for humans, not machines.
There should be an entry for every single version.
The same types of changes should be grouped.
Versions and sections should be linkable.
The latest version comes first.
The release date of each version is displayed.
Mention whether you follow Semantic Versioning.

Usage:

Change log entries are to be added to the Unreleased section under the
appropriate stanza (see below). Each entry should ideally include a tag and
the Github issue reference in the following format:

* (<tag>) \#<issue-number> message

The issue numbers will later be link-ified during the release process so you do
not have to worry about including a link manually, but you can if you wish.

Types of changes (Stanzas):

"Features" for new features.
"Improvements" for changes in existing functionality.
"Deprecated" for soon-to-be removed features.
"Bug Fixes" for any bug fixes.
"Client Breaking" for breaking CLI commands and REST routes used by end-users.
"API Breaking" for breaking exported APIs used by developers building on SDK.
"State Machine Breaking" for any changes that result in a different AppState given same genesisState and txList.

Ref: https://keepachangelog.com/en/1.0.0/
-->

# Changelog

## [v0.0.1] - 2024-04-19

### Features

* (session) Implement session management
* (client) Implemented API methods:
  * `/v0/org/:org_id/user/me`
  * `/v0/about_me`
  * `/v1/org/:org_id/token/refresh`
  * `/v0/org/:org_id/keys`
  * `/v0/org/:org_id/keys/:key`

## [v0.0.2] - 2024-04-22

### Features

* (session) Implement base64 session decoding

## [v0.0.3] - 2024-04-22

### Features

* (api) Add common `IKeyInfo` interface
* (session) Implement `AutoRefresh` and `StopAutoRefresh`

## [v0.0.4] - 2024-05-29

### Bug Fixes

* (client) Escape key id

### Features

* (client) Ability to override headers
* (roles) Implemented API methods:
    * `PUT /v0/org/:org_id/roles/:role_id/add_keys`
    * `POST /v0/org/:org_id/roles/:role_id/tokens`

## [v0.0.5] - 2024-06-05

### Bug Fixes

* (roles) No need to use session token for `/add_keys` and `/tokens`
* (roles) Sanitize role id
* (roles) Add readable response for `/add_keys`

### Features

* (pagination) Implement pagination
* (roles) Implemented API methods:
    * `GET /v0/org/:org_id/roles/:role_id/keys`
* (btc) Implemented API methods: 
    * `POST /v0/org/:org_id/btc/taproot/sign/:pubkey` 