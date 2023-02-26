package conditional_test

/**

func TestLike(t *testing.T) {
	str, args, err := conditional.Like("Name", "Mad", false).Generate()

	require.NoError(t, err)

	assert.Equal(t, `"Name" LIKE ?`, str)
	assert.Equal(t, []any{"Mad%"}, args)
}

func TestStringContains(t *testing.T) {
	str, args, err := conditional.StringContains("Name", "Mad").Generate()

	require.NoError(t, err)

	assert.Equal(t, `("Name" LIKE ?) OR ("Name" LIKE ?)`, str)
	assert.Equal(t, []any{"%Mad", "Mad%"}, args)
}

*/
