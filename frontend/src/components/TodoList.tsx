import Todo from './Todo';

interface TodoListProps {
  todos: {
    id: number;
    text: string;
    completed: boolean;
  }[];
  toggleTodo: (id: number) => void;
  deleteTodo: (id: number) => void;
}

const TodoList = ({ todos, toggleTodo, deleteTodo }: TodoListProps) => {
  return (
    <div className="todo-list">
      {todos.map(todo => (
        <Todo key={todo.id} todo={todo} toggleTodo={toggleTodo} deleteTodo={deleteTodo} />
      ))}
    </div>
  );
};

export default TodoList;
