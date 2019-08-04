import React from "react";
import "../App.css";

function Todo({ todo, completeTodo, removeTodo }) {
  return (
    <div
      className="todo"
      style={{ textDecoration: todo.is_completed ? "line-through" : "" }}
    >
      {todo.content}
      <div>
        <button onClick={() => completeTodo(todo.id)}>Complete</button>
        <button onClick={() => removeTodo(todo.id)}>x</button>
      </div>
    </div>
  );
}

export default Todo;
