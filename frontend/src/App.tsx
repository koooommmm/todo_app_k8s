interface Todo {
  id: number;
  text: string;
  completed: boolean;
}

import { useState, useEffect } from 'react';
import './App.css';
import AddTodoForm from './components/AddTodoForm';
import TodoList from './components/TodoList';

function App() {
  const [todos, setTodos] = useState<Todo[]>([]);

  useEffect(() => {
    fetch('/api/todos')
      .then(response => response.json())
      .then((data: Todo[]) => setTodos(data))
      .catch(error => console.error('Error fetching todos:', error));
  }, []);

  const addTodo = (text: string) => {
    const newTodo = {
      text,
      completed: false,
    };

    fetch('/api/todos', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(newTodo),
    })
      .then(response => response.json())
      .then((data: Todo) => {
        setTodos(prevTodos => [...prevTodos, data]);
      })
      .catch(error => console.error('Error adding todo:', error));
  };

  const toggleTodo = (id: number) => {
    setTodos(
      todos.map(todo =>
        todo.id === id ? { ...todo, completed: !todo.completed } : todo
      )
    );
  };

  const deleteTodo = (id: number) => {
    setTodos(todos.filter(todo => todo.id !== id));
  };

  return (
    <div className="App">
      <h1>Todo App</h1>
      <AddTodoForm addTodo={addTodo} />
      <TodoList todos={todos} toggleTodo={toggleTodo} deleteTodo={deleteTodo} />
    </div>
  );
}

export default App;
