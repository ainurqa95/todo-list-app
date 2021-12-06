package repositories

import (
	"errors"
	"github.com/ainurqa95/todo-list-app"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"testing"
)

func TestTodoItem_Create(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type Args struct {
		listId int
		item   todo.TodoItem
	}

	type mockBehavior func(args Args, id int)

	testsTable := []struct {
		name         string
		mockBehavior mockBehavior
		args         Args
		expectedId   int
		wantErr      bool
	}{
		{
			name: "Ok",
			args: Args{
				listId: 1,
				item: todo.TodoItem{
					Title:       "Test Title",
					Description: "Test description",
				},
			},
			expectedId: 2,
			wantErr:    false,
			mockBehavior: func(args Args, id int) {
				mock.ExpectBegin()
				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO todo_items").
					WithArgs(args.item.Title, args.item.Description).WillReturnRows(rows)

				mock.ExpectExec("INSERT INTO lists_items").
					WithArgs(args.listId, id).WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
		},
		{
			name: "Empty fields",
			args: Args{
				listId: 1,
				item: todo.TodoItem{
					Title:       "",
					Description: "Test description",
				},
			},
			expectedId: 2,
			wantErr:    true,
			mockBehavior: func(args Args, id int) {
				mock.ExpectBegin()
				rows := sqlmock.NewRows([]string{"id"}).RowError(0, errors.New("insert error"))
				mock.ExpectQuery("INSERT INTO todo_items").
					WithArgs(args.item.Title, args.item.Description).WillReturnRows(rows)

				mock.ExpectRollback()
			},
		},
		{
			name: "2nd insert fails",
			args: Args{
				listId: 1,
				item: todo.TodoItem{
					Title:       "Test ",
					Description: "Test description",
				},
			},
			expectedId: 2,
			wantErr:    true,
			mockBehavior: func(args Args, id int) {
				mock.ExpectBegin()
				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO todo_items").
					WithArgs(args.item.Title, args.item.Description).WillReturnRows(rows)

				mock.ExpectExec("INSERT INTO lists_items").
					WithArgs(args.listId, id).WillReturnError(errors.New("insert error"))

				mock.ExpectCommit()
			},

		},
	}
	repository := NewDbListItemRepository(db)

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args, testCase.expectedId)
			id, err := repository.Create(testCase.args.item, testCase.args.listId)
			if testCase.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, testCase.expectedId, id)
		})
	}

}
