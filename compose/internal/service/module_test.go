// +build integration

package service

import (
	"context"
	"testing"

	"github.com/titpetric/factory"

	"github.com/crusttech/crust/compose/types"
	"github.com/crusttech/crust/internal/auth"
	"github.com/crusttech/crust/internal/test"
)

func TestModule(t *testing.T) {
	factory.Database.MustGet("compose").Profiler = newTestLogProfiler(t)

	ctx := context.WithValue(context.Background(), "testing", true)

	// Set Identity (required for permission checks).
	ctx = auth.SetIdentityToContext(ctx, auth.NewIdentity(1337))

	ns1, _ := createTestNamespaces(ctx, t)

	svc := Module().With(ctx)

	// the module object we're working with
	module := &types.Module{
		NamespaceID: ns1.ID,
		Name:        "Test",
		Fields: types.ModuleFieldSet{
			&types.ModuleField{
				Name: "name",
				Kind: "input",
			},
			&types.ModuleField{
				Name: "email",
				Kind: "email",
			},
			&types.ModuleField{
				Name: "options",
				Kind: "select_multi",
			},
			&types.ModuleField{
				Name: "description",
				Kind: "text",
			},
		},
	}

	prevModuleCount := 0

	{
		{
			m, err := svc.Update(module)
			test.Assert(t, m == nil, "Expected empty return for invalid update, got %#v", m)
			test.Assert(t, err != nil, "Expected error when updating invalid content")
		}

		// create module
		m, err := svc.Create(module)
		test.Assert(t, err == nil, "Error when creating module: %+v", err)
		test.Assert(t, m.ID > 0, "Expected auto generated ID")

		// fetch created module
		{
			ms, err := svc.FindByID(m.NamespaceID, m.ID)
			test.Assert(t, err == nil, "Error when retrieving module by id: %+v", err)
			test.Assert(t, ms.ID == m.ID, "Expected ID from database to match, %d != %d", m.ID, ms.ID)
			test.Assert(t, ms.Name == m.Name, "Expected Name from database to match, %s != %s", m.Name, ms.Name)
			test.Assert(t, len(ms.Fields) == 4, "Expected Fields count from database to match, 4 != %d", len(ms.Fields))
		}

		// update created module
		{
			m.Name = "Updated test"
			m.UpdatedAt = nil
			_, err := svc.Update(m)
			test.Assert(t, err == nil, "Error when updating module, %+v", err)
		}

		// fetch module fields
		{
			fl := m.Fields.Names()
			test.Assert(t, err == nil, "Error when retrieving module fields by module: %+v", err)
			test.Assert(t, len(fl) == 4, "Expected 4 fields, got %d", len(fl))
		}

		// re-fetch module
		{
			ms, err := svc.FindByID(m.NamespaceID, m.ID)
			test.Assert(t, err == nil, "Error when retrieving module by id: %+v", err)
			test.Assert(t, ms.ID == m.ID, "Expected ID from database to match, %d != %d", m.ID, ms.ID)
			test.Assert(t, ms.Name == m.Name, "Expected Name from database to match, %s != %s", m.Name, ms.Name)
		}

		// fetch all modules
		{
			ms, _, err := svc.Find(types.ModuleFilter{})
			test.Assert(t, err == nil, "Error when retrieving modules: %+v", err)
			test.Assert(t, len(ms) >= 1, "Expected at least one module, got %d", len(ms))
			prevModuleCount = len(ms)
		}

		// re-fetch module
		{
			err := svc.DeleteByID(m.NamespaceID, m.ID)
			test.Assert(t, err == nil, "Error when deleting module by id: %+v", err)
		}

		// fetch all modules
		{
			ms, _, err := svc.Find(types.ModuleFilter{})
			test.Assert(t, err == nil, "Error when retrieving modules: %+v", err)
			test.Assert(t, len(ms) < prevModuleCount, "Expected modules count to decrease after deletion, %d < %d", len(ms), prevModuleCount)
		}
	}
}
