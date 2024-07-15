package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `[
			{
				"id": "7x6bj5bnbjafmog",
				"created": "2024-07-08 00:36:29.595Z",
				"updated": "2024-07-08 00:36:29.595Z",
				"name": "places",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "cva1x71e",
						"name": "geocoded",
						"type": "json",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSize": 2048
						}
					},
					{
						"system": false,
						"id": "r147rb8z",
						"name": "raw_input",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					}
				],
				"indexes": [],
				"listRule": "",
				"viewRule": "",
				"createRule": "",
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "v4owycc4f7h4gkl",
				"created": "2024-07-08 00:36:29.595Z",
				"updated": "2024-07-15 04:25:54.288Z",
				"name": "user",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "ga7n2ool",
						"name": "firstname",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "dikmx2tv",
						"name": "lastname",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "pg9yqosq",
						"name": "email",
						"type": "email",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"exceptDomains": null,
							"onlyDomains": null
						}
					},
					{
						"system": false,
						"id": "iriaar5k",
						"name": "residence",
						"type": "relation",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "7x6bj5bnbjafmog",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": 1,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "pajwhlz1",
						"name": "to_visit",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "7x6bj5bnbjafmog",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": null,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "wgmsuugp",
						"name": "to_visit_draw",
						"type": "text",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"pattern": ""
						}
					},
					{
						"system": false,
						"id": "lz1itdb7",
						"name": "trips_per_year",
						"type": "number",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"min": null,
							"max": null,
							"noDecimal": false
						}
					},
					{
						"system": false,
						"id": "ozlnm2qo",
						"name": "known_well",
						"type": "relation",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
							"collectionId": "7x6bj5bnbjafmog",
							"cascadeDelete": false,
							"minSelect": null,
							"maxSelect": null,
							"displayFields": null
						}
					},
					{
						"system": false,
						"id": "9ziz61fp",
						"name": "referrals",
						"type": "json",
						"required": false,
						"presentable": false,
						"unique": false,
						"options": {
							"maxSize": 2048
						}
					}
				],
				"indexes": [],
				"listRule": null,
				"viewRule": "",
				"createRule": "",
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		return nil
	})
}
