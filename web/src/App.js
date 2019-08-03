import React from "react";
import logo from "./logo.svg";
import "./App.css";
import TodoList from "./pages/Todolist";

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
      </header>
      <div className="Todo-app">
        <TodoList />
      </div>
    </div>
  );
}

export default App;
