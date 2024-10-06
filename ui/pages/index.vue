<template>
  <div class="todo-main">
    <h1>TODOリスト</h1>
    <div v-if="statusMessage" class="status-message">{{ statusMessage }}</div>
    <!-- Input section for adding a new task -->
    <div class="input-group">
      <input
        v-model="newTask"
        placeholder="新しいタスクを入力"
        @keyup.enter="addTodo"
      >
      <select v-model="newPriority">
        <option value="low">低い</option>
        <option value="medium">中くらい</option>
        <option value="high">たかい</option>
      </select>
      <button @click="addTodo">+ 追加</button>
    </div>
    <div class="input-group">
      <input
        v-model="searchTask"
        placeholder="🔍タスクの検索"
        @keyup.enter="searchTodos"
      >
      <button @click="toggleSortOrder">
        ↕️ 優先度順に並べ替える ({{ sortOrder === 'asc' ? '昇順' : '降順' }})
      </button>
    </div>
    <!-- Active Tasks Section -->
    <div class="tasks-section">
      <h2>Active Tasks</h2>
      <div v-if="activeTasks.length > 0">
        <div v-for="todo in activeTasks" :key="todo.id" class="todo-item">
          <input
            v-if="todo.isEditing"
            v-model="todo.task"
            class="edit-input"
            @blur="editTodo(todo)"
            @keyup.enter="editTodo(todo)"
          >
          <span
            v-else
            :class="{ 'done-task': todo.status === 'done' }"
            @click="enableEdit(todo)"
          >
            {{ todo.task }}
          </span>
          <div class="buttons">
            <!-- Priority display next to the task -->
            <span :class="priorityClass(todo.priority)" class="priority-label">
              {{ todo.priority }}
            </span>
            <button
              :class="{ 'done': todo.status === 'done' }"
              @click="updateStatus(todo)"
            >
              ✔️
            </button>
            <button class="delete-button" @click="deleteTodo(todo.id)">
              🗑️
            </button>
          </div>
        </div>
      </div>
      <div v-else>
        <p>タスクがありません。</p>
      </div>
    </div>
    <!-- Completed Tasks Section -->
    <div class="tasks-section">
      <h2>Completed Tasks</h2>
      <div v-if="completedTasks.length > 0">
        <div v-for="todo in completedTasks" :key="todo.id" class="todo-item">
          <input
            v-if="todo.isEditing"
            v-model="todo.task"
            class="edit-input"
            @blur="editTodo(todo)"
            @keyup.enter="editTodo(todo)"
          >
          <span
            v-else
            :class="{ 'done-task': todo.status === 'done' }"
            @click="enableEdit(todo)"
          >
            {{ todo.task }}
          </span>
          <div class="buttons">
            <!-- Priority display next to the task -->
            <span :class="priorityClass(todo.priority)" class="priority-label">
              {{ todo.priority }}
            </span>
            <button
              :class="{ 'done': todo.status === 'done' }"
              @click="updateStatus(todo)"
            >
              ✔️
            </button>
            <button class="delete-button" @click="deleteTodo(todo.id)">
              🗑️
            </button>
          </div>
        </div>
      </div>
      <div v-else>
        <p>完了したタスクはありません。</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
	data() {
		return {
			newTask: '',
			newPriority: 'medium',
			todos: [],
			searchTask: '',
			statusMessage: '',
			sortOrder: 'desc'
		};
	},
	computed: {
		// Active tasks filter (where status is not 'done')
		activeTasks() {
			return this.todos.filter(todo => todo.status !== 'done');
		},
		// Completed tasks filter (where status is 'done')
		completedTasks() {
			return this.todos.filter(todo => todo.status === 'done');
		}
	},
	mounted() {
		this.fetchTodos();
	},
	methods: {
		priorityClass(priority) {
			if (priority === 'high') return 'high-priority';
			if (priority === 'medium') return 'medium-priority';
			return 'low-priority';
		},
		async fetchTodos() {
			try {
				const response = await fetch(`/api/v1/todos`, {});
				if (!response.ok) throw new Error(`Failed to get todo list. statusCode: ${response.status}`);
				const data = await response.json();
				this.todos = data.data;
			} catch (error) {
				console.error(error);
				this.statusMessage = 'タスクの取得に失敗しました';
			}
		},
		async searchTodos() {
			try {
				const response = await fetch(`/api/v1/todos?task=${this.searchTask}`, {});
				if (!response.ok) throw new Error(`Failed to search todo list. statusCode: ${response.status}`);
				response.json().then(data => {
					this.todos = data.data;
				});
			} catch (error) {
				console.error(error);
				this.statusMessage = 'タスクの取得に失敗しました';
			}
		},
		sortTodos() {
			const priorityOrder = {
				low: 1,
				medium: 2,
				high: 3
			};

			this.todos.sort((a, b) => {
				if (this.sortOrder === 'asc') {
					// Ascending order (low to high)
					return priorityOrder[a.priority] - priorityOrder[b.priority];
				} else {
					// Descending order (high to low)
					return priorityOrder[b.priority] - priorityOrder[a.priority];
				}
			});

			this.statusMessage = `タスクが優先度順に ${this.sortOrder === 'asc' ? '昇順' : '降順'} で並べ替えられました`;
		},
		toggleSortOrder() {
			// Toggle between 'asc' and 'desc'
			this.sortOrder = this.sortOrder === 'asc' ? 'desc' : 'asc';
			this.sortTodos(); // Apply sorting after toggling
		},
		async addTodo() {
			if (this.newTask.trim() === '') return;
			if (this.newPriority.trim() === '') return;

			try {
				const response = await fetch('/api/v1/todos', {
					method: 'POST',
					headers: {
						'Content-Type': 'application/json',
					},
					body: JSON.stringify({
						task: this.newTask,
						priority: this.newPriority,
						Status: 'created' // Default status for new tasks
					})
				});

				if (!response.ok) throw new Error(`Failed to create todo. statusCode: ${response.status}`);

				const data = await response.json();
				this.todos.push(data.data);
				this.newTask = '';
				this.newPriority = '';
				this.statusMessage = 'タスクが追加されました';
			} catch (error) {
				console.error('Error creating todo:', error);
				this.statusMessage = 'タスクの作成に失敗しました';
			}
		},
		enableEdit(todo) {
			todo.isEditing = true;
		},
		async editTodo(todo) {
			todo.isEditing = false;

			try {
				const response = await fetch(`/api/v1/todos/${todo.id}`, {
					method: 'PUT',
					headers: {
						'Content-Type': 'application/json',
					},
					body: JSON.stringify({
						task: todo.task
					})
				});

				if (!response.ok) throw new Error(`Failed to edit todo. statusCode: ${response.status}`);

				this.statusMessage = 'タスクが編集されました';
			} catch (error) {
				console.error('Error editing todo:', error);
				this.statusMessage = 'タスクの編集に失敗しました';
			}
		},
		async updateStatus(todo) {
			try {
				const response = await fetch(`/api/v1/todos/${todo.id}`, {
					method: 'PUT',
					headers: {
						'Content-Type': 'application/json',
					},
					body: JSON.stringify({
						Status: todo.status === 'done' ? 'created' : 'done'
					})
				});

				if (!response.ok) throw new Error(`Failed to update todo Status. statusCode: ${response.status}`);

				todo.status = todo.status === 'done' ? 'created' : 'done';
				this.statusMessage = 'タスクのステータスが変更されました';
			} catch (error) {
				console.error('Error updating todo Status:', error);
				this.statusMessage = 'ステータスの更新に失敗しました';
			}
		},
		async deleteTodo(id) {
			try {
				const response = await fetch(`/api/v1/todos/${id}`, {
					method: 'DELETE',
					headers: {
						'Content-Type': 'application/json',
					},
				});

				if (!response.ok) throw new Error(`Failed to update todo Status. statusCode: ${response.status}`);

				this.todos = this.todos.filter(todo => todo.id !== id);
				this.statusMessage = 'タスクが削除されました';
			} catch (error) {
				console.error('Error deleting todo:', error);
				this.statusMessage = 'タスクの削除に失敗しました';
			}
		}
	}
};
</script>

<style scoped>
.todo-main {
  max-width: 400px;
  margin: 20px auto;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  background-color: #fff;
}

.input-group {
  display: flex;
  margin-bottom: 20px;
}

input {
  flex: 1;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
  margin-right: 10px;
}

button {
  padding: 10px;
  border: none;
  background-color: #333;
  color: #fff;
  border-radius: 4px;
  cursor: pointer;
}

.todo-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.buttons button {
  background-color: #f0f0f0;
  color: #333;
  margin-left: 5px;
  border-radius: 4px;
  padding: 5px 10px;
  transition: background-color 0.3s, color 0.3s;
}

.buttons button.done {
  background-color: #333;
  color: #fff;
}

.buttons button.done::before {
  color: #fff;
}

.buttons button.delete-button {
  color: white;
}

.status-message {
  margin-bottom: 20px;
  padding: 10px;
  background-color: #f0f0f0;
  border-radius: 4px;
}

.done-task {
  text-decoration: line-through;
}

.edit-input {
  flex: 1;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.select-container {
  width: 100%;
  position: relative;
}

select {
  padding: 8px 12px;
  border: 1px solid #ccc;
  margin-right: 5px;
  border-radius: 4px;
  background-color: #fff;
  color: #333;
  font-size: 14px;
  appearance: none;
}

select:focus {
  outline: none;
  border-color: #666;
}

.priority-label {
  margin-left: 10px;
  padding: 3px 8px;
  border-radius: 5px;
  font-size: 12px;
  font-weight: bold;
}

/* High priority (red) */
.high-priority {
  background-color: rgba(255, 0, 0, 0.2);
  color: red;
}

/* Medium priority (yellow) */
.medium-priority {
  background-color: rgba(255, 255, 0, 0.2);
  color: #ffa500; /* Orange for better contrast */
}

/* Low priority (dim gray) */
.low-priority {
  background-color: rgba(128, 128, 128, 0.2);
  color: dimgray;
}
</style>
