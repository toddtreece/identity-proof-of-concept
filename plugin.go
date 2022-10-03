package main

type PluginInstance struct {
	id string
}

func (i *PluginInstance) GetID() string {
	return i.id
}

func (i *PluginInstance) HasAccess(scope string) bool {
	return true
}
