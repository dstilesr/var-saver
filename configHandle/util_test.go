package confighandle

func GetTestConfigObj() *SavedVariables {
	return &SavedVariables{
		Meta: GetMetadata(),
		Projects: map[string]*Project{
			"proj1": &Project{
				Name: "proj1",
				Variables: []*Variable{
					&Variable{
						Name:        "var1",
						Value:       "val1",
						Environment: "env1",
					},
					&Variable{
						Name:        "var2",
						Value:       "val2",
						Environment: "env1",
					},
					&Variable{
						Name:        "var1",
						Value:       "val3",
						Environment: "env2",
					},
				},
			},
			"proj2": &Project{
				Name: "proj1",
				Variables: []*Variable{
					&Variable{
						Name:        "var4",
						Value:       "val4",
						Environment: "env1",
					},
				},
			},
		},
	}
}
