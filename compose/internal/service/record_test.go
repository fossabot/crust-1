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

func TestRecord(t *testing.T) {
	factory.Database.MustGet("compose").Profiler = newTestLogProfiler(t)

	ctx := context.WithValue(context.Background(), "testing", true)

	ctx = auth.SetIdentityToContext(ctx, auth.NewIdentity(1337))

	var err error
	ns1, ns2 := createTestNamespaces(ctx, t)

	moduleSvc := Module().With(ctx)
	svc := Record().With(ctx)

	module1 := &types.Module{
		NamespaceID: ns1.ID,
		Name:        "Test",
		Fields: types.ModuleFieldSet{
			&types.ModuleField{
				Name: "name",
			},
			&types.ModuleField{
				Name: "email",
			},
			&types.ModuleField{
				Name:  "options",
				Multi: true,
			},
			&types.ModuleField{
				Name: "description",
			},
			&types.ModuleField{
				Name: "another_record",
				Kind: "Record",
			},
		},
	}

	// set up a module1
	module1, err = moduleSvc.Create(module1)
	test.Assert(t, err == nil, "Error when creating module1: %+v", err)
	test.Assert(t, module1.ID > 0, "Expected auto generated ID")

	module2 := &types.Module{
		NamespaceID: ns2.ID,
		Name:        "Test Dummy",
		Fields: types.ModuleFieldSet{
			&types.ModuleField{
				Name: "name",
			},
			&types.ModuleField{
				Name: "email",
			},
			&types.ModuleField{
				Name:  "options",
				Multi: true,
			},
			&types.ModuleField{
				Name: "description",
			},
			&types.ModuleField{
				Name: "another_record",
				Kind: "Record",
			},
		},
	}
	module2, err = moduleSvc.Create(module2)
	test.Assert(t, err == nil, "Error when creating module1 in another namespace: %+v", err)

	record1 := &types.Record{
		NamespaceID: ns1.ID,
		ModuleID:    module1.ID,
	}

	record2 := &types.Record{
		NamespaceID: ns1.ID,
		ModuleID:    module1.ID,
		Values: types.RecordValueSet{
			&types.RecordValue{
				Name:  "name",
				Value: "John Doe",
			},
			&types.RecordValue{
				Name:  "email",
				Value: "john.doe@example.com",
			},
			&types.RecordValue{
				Name:  "options",
				Value: "1",
			},
			&types.RecordValue{
				Name:  "options",
				Value: "2",
			},
			&types.RecordValue{
				Name:  "options",
				Value: "3",
			},
			&types.RecordValue{
				Name:  "description",
				Value: "just an example",
			},
			&types.RecordValue{
				Name:  "another_record",
				Value: "918273645",
			},
		},
	}

	{
		// Let's put something to another namespace...
		_, err := svc.Create(&types.Record{
			NamespaceID: ns2.ID,
			ModuleID:    module2.ID,
		})
		test.Assert(t, err == nil, "Error when creating record: %+v", err)
	}

	// now work with records
	{
		{
			record1.UpdatedAt = nil
			m, err := svc.Update(record1)
			test.Assert(t, m == nil, "Expected empty return for invalid update, got %#v", m)
			test.Assert(t, err != nil, "Expected error when updating invalid record")
		}

		// create record
		m1, err := svc.Create(record1)
		test.Assert(t, err == nil, "Error when creating record: %+v", err)
		test.Assert(t, m1.ID > 0, "Expected auto generated ID")

		// create record
		m2, err := svc.Create(record2)
		test.Assert(t, err == nil, "Error when creating record: %+v", err)
		test.Assert(t, m2.ID > 0, "Expected auto generated ID")

		// fetch created record
		{
			ms, err := svc.FindByID(m1.NamespaceID, m1.ID)
			test.Assert(t, err == nil, "Error when retrieving record by id: %+v", err)
			test.Assert(t, ms.ID == m1.ID, "Expected ID from database to match, %d != %d", m1.ID, ms.ID)
			test.Assert(t, ms.ModuleID == m1.ModuleID, "Expected Module ID from database to match, %d != %d", m1.ModuleID, ms.ModuleID)
		}

		// update created record
		{
			m1.UpdatedAt = nil
			_, err := svc.Update(m1)
			test.Assert(t, err == nil, "Error when updating record, %+v", err)
		}

		// re-fetch record
		{
			ms, err := svc.FindByID(m1.NamespaceID, m1.ID)
			test.Assert(t, err == nil, "Error when retrieving record by id: %+v", err)
			test.Assert(t, ms.ID == m1.ID, "Expected ID from database to match, %d != %d", m1.ID, ms.ID)
			test.Assert(t, ms.ModuleID == m1.ModuleID, "Expected ID from database to match, %d != %d", m1.ModuleID, ms.ModuleID)
		}

		// fetch all records
		{
			mr, f, err := svc.Find(types.RecordFilter{ModuleID: module1.ID, Sort: "id desc"})
			test.Assert(t, err == nil, "Error when retrieving records: %+v", err)
			test.Assert(t, len(mr) == 2, "Expected two record, got %d", len(mr))
			test.Assert(t, f.Count == 2, "Expected Meta.Count == 2, got %d", f.Count)
			test.Assert(t, f.Sort == "id desc", "Expected Meta.Sort == id desc, got '%s'", f.Sort)
			test.Assert(t, mr[0].ModuleID == m1.ModuleID, "Expected record module1 to match, %d != %d", m1.ModuleID, mr[0].ModuleID)
			test.Assert(t, mr[0].ID > mr[1].ID, "Expected order to be descending")
		}

		// fetch all records
		{

			mr, f, err := svc.Find(types.RecordFilter{ModuleID: module1.ID, Sort: "name asc, email desc"})
			test.Assert(t, err == nil, "Error when retrieving records: %+v", err)
			test.Assert(t, len(mr) == 2, "Expected two record, got %d", len(mr))
			test.Assert(t, f.Count == 2, "Expected Meta.Count == 2, got %d", f.Count)
			test.Assert(t, f.Sort == "name asc, email desc", "Expected Meta.Sort == 'name asc, email desc' '%s'", f.Sort)
			test.Assert(t, mr[0].ModuleID == m1.ModuleID, "Expected record module1 to match, %d != %d", m1.ModuleID, mr[0].ModuleID)

			// @todo sort is not stable
			// test.Assert(t, mr[0].ID > mr[1].ID, "Expected order to be ascending")
		}

		// fetch all records
		{
			mr, f, err := svc.Find(types.RecordFilter{ModuleID: module1.ID, Sort: "created_at desc"})
			test.Assert(t, err == nil, "Error when retrieving records: %+v", err)
			test.Assert(t, len(mr) == 2, "Expected two record, got %d", len(mr))
			test.Assert(t, f.Count == 2, "Expected Meta.Count == 2, got %d", f.Count)
			test.Assert(t, f.Sort == "created_at desc", "Expected Meta.Sort == created_at desc, got '%s'", f.Sort)
			test.Assert(t, mr[0].ModuleID == m1.ModuleID, "Expected record module1 to match, %d != %d", m1.ModuleID, mr[0].ModuleID)

			// @todo sort is not stable
			// test.Assert(t, mr[0].ID > mr[1].ID, "Expected order to be ascending")
		}

		// fetch all records by query
		{
			filter := "name='John Doe' AND email='john.doe@example.com'"
			sort := "id desc"

			mr, f, err := svc.Find(types.RecordFilter{ModuleID: module1.ID, Sort: sort, Filter: filter})
			test.Assert(t, err == nil, "Error when retrieving records: %+v", err)
			test.Assert(t, len(mr) == 1, "Expected one record, got %d", len(mr))
			test.Assert(t, f.Count == 1, "Expected Meta.Count == 1, got %d", f.Count)
			test.Assert(t, f.Page == 0, "Expected Meta.Page == 0, got %d", f.Page)
			test.Assert(t, f.PerPage == 50, "Expected Meta.PerPage == 50, got %d", f.PerPage)
			test.Assert(t, f.Filter == filter, "Expected Meta.Filter == %q, got %q", filter, f.Filter)
			test.Assert(t, f.Sort == sort, "Expected Meta.Sort == %q, got %q", sort, f.Sort)
		}

		// fetch all records by query
		{
			mr, _, err := svc.Find(types.RecordFilter{ModuleID: module1.ID, Sort: "id asc", Filter: "name='niall'"})
			test.Assert(t, err == nil, "Error when retrieving records: %+v", err)
			test.Assert(t, len(mr) == 0, "Expected no records, got %d", len(mr))
		}

		// delete record
		{
			err := svc.DeleteByID(m1.NamespaceID, m1.ID)
			test.Assert(t, err == nil, "Error when retrieving record by id: %+v", err)

			err = svc.DeleteByID(m2.NamespaceID, m2.ID)
			test.Assert(t, err == nil, "Error when retrieving record by id: %+v", err)
		}

		// fetch all records
		{
			mr, _, err := svc.Find(types.RecordFilter{ModuleID: module1.ID})
			test.Assert(t, err == nil, "Error when retrieving records: %+v", err)
			test.Assert(t, len(mr) == 0, "Expected no record, got %d", len(mr))
		}
	}
}

func TestValueSanitizer(t *testing.T) {
	var (
		svc    = record{}
		module = &types.Module{
			Fields: types.ModuleFieldSet{
				&types.ModuleField{Name: "single1"},
				&types.ModuleField{Name: "multi1", Multi: true},
				&types.ModuleField{Name: "ref1", Kind: "Record"},
				&types.ModuleField{Name: "multiRef1", Kind: "Record", Multi: true},
			},
		}
		rvs types.RecordValueSet
	)

	rvs = types.RecordValueSet{{Name: "single1", Value: "single"}}
	test.NoError(t, svc.sanitizeValues(module, rvs), "unexpected error for sanitizeValues() call: %v")
	test.Assert(t, len(rvs) == 1, "expecting 1 record value after sanitization, got %d", len(rvs))

	rvs = types.RecordValueSet{{Name: "unknown", Value: "single"}}
	test.Assert(t, svc.sanitizeValues(module, rvs) != nil, "expecting sanitizeValues() to return an error, got nil")

	rvs = types.RecordValueSet{{Name: "single1", Value: "single"}, {Name: "single1", Value: "single2"}}
	test.Assert(t, svc.sanitizeValues(module, rvs) != nil, "expecting sanitizeValues() to return an error, got nil")

	rvs = types.RecordValueSet{{Name: "multi1", Value: "multi1"}, {Name: "multi1", Value: "multi1"}}
	test.NoError(t, svc.sanitizeValues(module, rvs), "unexpected error for sanitizeValues() call: %v")
	test.Assert(t, len(rvs) == 2, "expecting 2 record values after sanitization, got %d", len(rvs))
	test.Assert(t, rvs[0].Place == 0, "expecting first value to have place value 0, got %d", rvs[0].Place)
	test.Assert(t, rvs[1].Place == 1, "expecting second value to have place value 1, got %d", rvs[1].Place)

	rvs = types.RecordValueSet{{Name: "ref1", Value: "multi1"}}
	test.Assert(t, svc.sanitizeValues(module, rvs) != nil, "expecting sanitizeValues() to return an error, got nil")

	rvs = types.RecordValueSet{{Name: "ref1", Value: "12345"}}
	test.NoError(t, svc.sanitizeValues(module, rvs), "unexpected error for sanitizeValues() call: %v")
	test.Assert(t, len(rvs) == 1, "expecting 1 record values after sanitization, got %d", len(rvs))
	test.Assert(t, rvs[0].Ref == 12345, "expecting parsed ref value to match, got %d", rvs[0].Ref)

	rvs = types.RecordValueSet{{Name: "multiRef1", Value: "12345"}, {Name: "multiRef1", Value: "67890"}}
	test.NoError(t, svc.sanitizeValues(module, rvs), "unexpected error for sanitizeValues() call: %v")
	test.Assert(t, len(rvs) == 2, "expecting 2 record values after sanitization, got %d", len(rvs))
	test.Assert(t, rvs[0].Ref == 12345, "expecting parsed ref value to match, got %d", rvs[0].Ref)
	test.Assert(t, rvs[1].Ref == 67890, "expecting parsed ref value to match, got %d", rvs[1].Ref)
}
