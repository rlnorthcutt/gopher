package core

import (
"os"

"github.com/pocketbase/pocketbase"
"github.com/pocketbase/pocketbase/models"
"gopkg.in/yaml.v3"

)

// RuleSet defines CRUD access rules for a PocketBase collection.
// Use nil to skip updating a rule. Use "" to explicitly clear a rule.
type RuleSet struct {
List   *string
View   *string
Create *string
Update *string
Delete *string
}

// PermissionsManager allows modules to define and sync collection rules.
type PermissionsManager struct {
pb *pocketbase.PocketBase
}

// NewPermissionsManager creates a new permissions manager.
func NewPermissionsManager(pb *pocketbase.PocketBase) *PermissionsManager {
return &PermissionsManager{pb: pb}
}

// SetRules applies rule overrides to a collection.
func (pm *PermissionsManager) SetRules(collectionName string, rules RuleSet) error {
collection, err := pm.pb.Dao().FindCollectionByNameOrId(collectionName)
if err != nil {
return err
}

if rules.List != nil {
	collection.ListRule = rules.List
}
if rules.View != nil {
	collection.ViewRule = rules.View
}
if rules.Create != nil {
	collection.CreateRule = rules.Create
}
if rules.Update != nil {
	collection.UpdateRule = rules.Update
}
if rules.Delete != nil {
	collection.DeleteRule = rules.Delete
}

return pm.pb.Dao().SaveCollection(collection)

}

// LoadFromYAML reads a permissions YAML file and applies it.
func (pm *PermissionsManager) LoadFromYAML(path string) error {
content, err := os.ReadFile(path)
if err != nil {
return err
}

var raw struct {
	Collection string            `yaml:"collection"`
	Rules      map[string]string `yaml:"rules"`
}

if err := yaml.Unmarshal(content, &raw); err != nil {
	return err
}

ruleSet := RuleSet{}

if v, ok := raw.Rules["list"]; ok {
	ruleSet.List = &v
}
if v, ok := raw.Rules["view"]; ok {
	ruleSet.View = &v
}
if v, ok := raw.Rules["create"]; ok {
	ruleSet.Create = &v
}
if v, ok := raw.Rules["update"]; ok {
	ruleSet.Update = &v
}
if v, ok := raw.Rules["delete"]; ok {
	ruleSet.Delete = &v
}

return pm.SetRules(raw.Collection, ruleSet)

}

