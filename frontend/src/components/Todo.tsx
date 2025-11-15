interface TodoProps {
  todo: {
    id: number;
    text: string;
    completed: boolean;
  };
  toggleTodo: (id: number) => void;
  deleteTodo: (id: number) => void;
}

const Todo = ({ todo, toggleTodo, deleteTodo }: TodoProps) => {
  return (
    <div className="todo">
      <input
        type="checkbox"
        checked={todo.completed}
        onChange={() => toggleTodo(todo.id)}
      />
      <span style={{ textDecoration: todo.completed ? 'line-through' : 'none' }}>
        {todo.text}
      </span>
      <button onClick={() => deleteTodo(todo.id)}>Delete</button>
    </div>
  );
};

export default Todo;
