package controller

import (
	"encoding/json"
)

const EquipmentBucket = "equipments"

type Equipment struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Outlet string `json:"outlet"`
	On     bool   `json:"on"`
	Value  int    `json:"value"`
}

func (c *Controller) GetEquipment(id string) (Equipment, error) {
	var eq Equipment
	return eq, c.store.Get(EquipmentBucket, id, &eq)
}

func (c Controller) ListEquipments() (*[]interface{}, error) {
	fn := func(v []byte) (interface{}, error) {
		var e Equipment
		return &e, json.Unmarshal(v, &e)
	}
	return c.store.List(EquipmentBucket, fn)
}

func (c *Controller) CreateEquipment(eq Equipment) error {
	fn := func(id string) interface{} {
		eq.ID = id
		return eq
	}
	if err := c.store.Create(EquipmentBucket, fn); err != nil {
		return err
	}
	return c.synOutlet(eq)
}

func (c *Controller) synOutlet(eq Equipment) error {
	return c.ConfigureOutlet(eq.Outlet, eq.On, eq.Value)
}

func (c *Controller) UpdateEquipment(id string, eq Equipment) error {
	if err := c.store.Update(EquipmentBucket, id, eq); err != nil {
		return err
	}
	eq.ID = id
	return c.synOutlet(eq)
}

func (c *Controller) DeleteEquipment(id string) error {
	return c.store.Delete(EquipmentBucket, id)
}