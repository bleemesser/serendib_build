package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("v4owycc4f7h4gkl")
		if err != nil {
			return err
		}

		// add
		new_referredby := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "robfijwg",
			"name": "referredby",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "v4owycc4f7h4gkl",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": 1,
				"displayFields": null
			}
		}`), new_referredby); err != nil {
			return err
		}
		collection.Schema.AddField(new_referredby)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("v4owycc4f7h4gkl")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("robfijwg")

		return dao.SaveCollection(collection)
	})
}
