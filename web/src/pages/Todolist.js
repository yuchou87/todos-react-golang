import React, { useState, useEffect } from "react";
import "../App.css";
import TodoForm from "../components/TodoForm";
import Todo from "../components/Todo";
import { getTodos, createTodo, updateTodo, deleteTodo } from "../api";

function TodoList() {
  const [todos, setTodos] = useState([]);

  useEffect(() => {
    getTodos().then(data => {
      setTodos(data.data);
    });
  }, []);

  const addTodo = text => {
    createTodo(text);
  };

  const completeTodo = id => {
    updateTodo(id);
  };

  const removeTodo = id => {
    deleteTodo(id);
  };

  return (
    <div className="todo-list">
      {todos.map((todo, index) => (
        <Todo
          key={index}
          todo={todo}
          completeTodo={completeTodo}
          removeTodo={removeTodo}
        />
      ))}
      <TodoForm addTodo={addTodo} />
    </div>
  );
}

export default TodoList;
