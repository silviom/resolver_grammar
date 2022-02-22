package main

import "testing"

func TestParse(t *testing.T) {
	type useCase struct {
		name  string
		input string
		ast   expr
	}

	useCases := []useCase{
		{
			name:  "text",
			input: "some text",
			ast: &stmtsAST{
				list: []*stmtAST{
					{before: "some text"},
				},
			},
		},
		{
			name:  "resolver",
			input: "{{ foo() }}",
			ast: &stmtsAST{
				list: []*stmtAST{
					{
						resolver: &resolverAST{
							name:      "foo",
							arguments: &argsAST{},
						},
					},
				},
			},
		},
		{
			name:  "resolver space between name",
			input: "{{ foo () }}",
			ast: &stmtsAST{
				list: []*stmtAST{
					{
						resolver: &resolverAST{
							name:      "foo",
							arguments: &argsAST{},
						},
					},
				},
			},
		}, {
			name:  "resolver with an argument",
			input: "{{ foo(bar) }}",
			ast: &stmtsAST{
				list: []*stmtAST{
					{
						resolver: &resolverAST{
							name: "foo",
							arguments: &argsAST{
								list: []*argAST{
									{before: "bar"},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "resolver with two arguments",
			input: "{{ foo(bar,baz) }}",
			ast: &stmtsAST{
				list: []*stmtAST{
					{
						resolver: &resolverAST{
							name: "foo",
							arguments: &argsAST{
								list: []*argAST{
									{before: "bar"},
									{before: "baz"},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "resolver with spaces in arguments",
			input: "{{ foo(bar, baz, and more spaces ) }}",
			ast: &stmtsAST{
				list: []*stmtAST{
					{
						resolver: &resolverAST{
							name: "foo",
							arguments: &argsAST{
								list: []*argAST{
									{before: "bar"},
									{before: " baz"},
									{before: " and more spaces "},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "text resolver",
			input: "my text {{ foo() }}",
			ast: &stmtsAST{
				list: []*stmtAST{
					{
						before: "my text ",
						resolver: &resolverAST{
							name:      "foo",
							arguments: &argsAST{},
						},
					},
				},
			},
		},
		{
			name:  "resolver text",
			input: "{{ foo() }} some text",
			ast: &stmtsAST{
				list: []*stmtAST{
					{
						resolver: &resolverAST{
							name:      "foo",
							arguments: &argsAST{},
						},
						after: " some text",
					},
				},
			},
		},
		{
			name:  "text resolver text",
			input: "my text {{ foo() }} some text",
			ast: &stmtsAST{
				list: []*stmtAST{
					{
						before: "my text ",
						resolver: &resolverAST{
							name:      "foo",
							arguments: &argsAST{},
						},
						after: " some text",
					},
				},
			},
		},
		{
			name:  "resolver resolver",
			input: "{{foo()}}{{bar()}}",
			ast: &stmtsAST{
				list: []*stmtAST{
					{
						resolver: &resolverAST{
							name:      "foo",
							arguments: &argsAST{},
						},
					},
					{
						resolver: &resolverAST{
							name:      "bar",
							arguments: &argsAST{},
						},
					},
				},
			},
		},
		{
			name:  "text resolver resolver",
			input: "my text {{foo()}}{{bar()}}",
			ast: &stmtsAST{
				list: []*stmtAST{
					{
						before: "my text ",
						resolver: &resolverAST{
							name:      "foo",
							arguments: &argsAST{},
						},
					},
					{
						resolver: &resolverAST{
							name:      "bar",
							arguments: &argsAST{},
						},
					},
				},
			},
		},
		{
			name:  "text resolver text resolver",
			input: "my text {{foo()}} something between {{bar()}}",
			ast: &stmtsAST{
				list: []*stmtAST{
					{
						before: "my text ",
						resolver: &resolverAST{
							name:      "foo",
							arguments: &argsAST{},
						},
						after: " something between ",
					},
					{
						resolver: &resolverAST{
							name:      "bar",
							arguments: &argsAST{},
						},
					},
				},
			},
		},
		{
			name:  "text resolver text resolver resolver",
			input: "my text {{foo()}} something between {{bar()}} and the final words",
			ast: &stmtsAST{
				list: []*stmtAST{
					{
						before: "my text ",
						resolver: &resolverAST{
							name:      "foo",
							arguments: &argsAST{},
						},
						after: " something between ",
					},
					{
						resolver: &resolverAST{
							name:      "bar",
							arguments: &argsAST{},
						},
						after: " and the final words",
					},
				},
			},
		},
		{
			name:  "resolver inner_resolver",
			input: "{{ foo({{ bar() }}) }}",
			ast: &stmtsAST{
				list: []*stmtAST{
					{
						resolver: &resolverAST{
							name: "foo",
							arguments: &argsAST{
								list: []*argAST{
									{
										resolver: &resolverAST{
											name:      "bar",
											arguments: &argsAST{},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "resolver ( text inner_resolver )",
			input: "{{ foo(hello {{ bar() }}) }}",
			ast: &stmtsAST{
				list: []*stmtAST{
					{
						resolver: &resolverAST{
							name: "foo",
							arguments: &argsAST{
								list: []*argAST{
									{
										before: "hello ",
										resolver: &resolverAST{
											name:      "bar",
											arguments: &argsAST{},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "resolver (text inner_resolver text)",
			input: "{{ foo(hello {{ bar() }} world) }}",
			ast: &stmtsAST{
				list: []*stmtAST{
					{
						resolver: &resolverAST{
							name: "foo",
							arguments: &argsAST{
								list: []*argAST{
									{
										before: "hello ",
										resolver: &resolverAST{
											name:      "bar",
											arguments: &argsAST{},
										},
										after: " world",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "resolver (inner_resolver text)",
			input: "{{ foo({{ bar() }} some text) }}",
			ast: &stmtsAST{
				list: []*stmtAST{
					{
						resolver: &resolverAST{
							name: "foo",
							arguments: &argsAST{
								list: []*argAST{
									{
										resolver: &resolverAST{
											name:      "bar",
											arguments: &argsAST{},
										},
										after: " some text",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "resolver(inner_resolver(inner_inner_resolver()))",
			input: "{{ foo({{ bar({{ baz() }}) }}) }}",
			ast: &stmtsAST{
				list: []*stmtAST{
					{
						resolver: &resolverAST{
							name: "foo",
							arguments: &argsAST{
								list: []*argAST{
									{
										resolver: &resolverAST{
											name: "bar",
											arguments: &argsAST{
												list: []*argAST{
													{
														resolver: &resolverAST{
															name:      "baz",
															arguments: &argsAST{},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name:  "resolver (text, text inner_resolver text, text)",
			input: "{{ foo(arg1, hello {{ bar() }} world, another argument) }}",
			ast: &stmtsAST{
				list: []*stmtAST{
					{
						resolver: &resolverAST{
							name: "foo",
							arguments: &argsAST{
								list: []*argAST{
									{
										before: "arg1",
									},
									{
										before: " hello ",
										resolver: &resolverAST{
											name:      "bar",
											arguments: &argsAST{},
										},
										after: " world",
									},
									{
										before: " another argument",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, uc := range useCases {
		t.Run(uc.name, func(t *testing.T) {
			interpreter := interpreter{}
			ast := interpreter.Parse(uc.input)
			result := interpreter.eval(ast)
			expected := interpreter.eval(uc.ast)
			if result != expected {
				t.Errorf("\nresult:\t%s\nexpected:\t%s\n", result, expected)
			}
		})
	}
}
