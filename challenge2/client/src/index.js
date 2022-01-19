import React from "react";
import ReactDOM from "react-dom";
import App from "./App";
import { Provider } from "react-redux";
import accountReducer from "./redux/account-slice.js";
import adminReducer from "./redux/admin-slice.js";
import { configureStore } from "@reduxjs/toolkit";
import { BrowserRouter } from "react-router-dom";

const store = configureStore({
  reducer: {
    accountSlice: accountReducer,
    adminSlice: adminReducer,
  },
});

ReactDOM.render(
  <React.StrictMode>
    <Provider store={store}>
      <BrowserRouter>
        <App />
      </BrowserRouter>
    </Provider>
  </React.StrictMode>,
  document.getElementById("root")
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
