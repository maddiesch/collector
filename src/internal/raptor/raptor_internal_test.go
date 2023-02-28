package raptor

/**

func Test_txConn(t *testing.T) {
	conn, err := New("file:internal-transaction?mode=memory&cache=shared")
	require.NoError(t, err)

	log := &test.CollectQueryLogger{}
	conn.SetLogger(log)

	ctx := context.Background()

	t.Run("multi-begin calls", func(t *testing.T) {
		defer log.Reset()

		conn.Transact(ctx, func(d DB) error {
			tx := d.(*txConn)

			err := tx.begin(ctx)
			assert.ErrorIs(t, err, ErrTransactionAlreadyStarted)

			return nil
		})
	})

	t.Run("multi-commit calls", func(t *testing.T) {
		defer log.Reset()

		conn.Transact(ctx, func(d DB) error {
			tx := d.(*txConn)

			err := tx.commit(ctx)
			assert.NoError(t, err)

			err = tx.commit(ctx)
			assert.NoError(t, err)

			return nil
		})

		assert.Len(t, log.Queries, 2)
	})

	t.Run("multi-rollback calls", func(t *testing.T) {
		defer log.Reset()

		conn.Transact(ctx, func(d DB) error {
			tx := d.(*txConn)

			err := tx.rollback(ctx)
			assert.NoError(t, err)

			err = tx.rollback(ctx)
			assert.NoError(t, err)

			return nil
		})

		assert.Len(t, log.Queries, 2)
	})
}

*/
