import { createSlice } from "@reduxjs/toolkit";
import { createAsyncThunk } from "@reduxjs/toolkit";

const loginThunk = createAsyncThunk(
  "/account/login",
  async (account, thunkAPI) => {
    let res = await fetch("http://localhost:8080/user-management/user/signin", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(account),
    });

    if (!res.ok) {
      res = await res.json();
      throw new Error(`${JSON.stringify(res)}`);
    }

    res = await res.json();

    return res;
  }
);

const getInfo = createAsyncThunk("/account/getinfo", async (_, thunkAPI) => {
  const jwt = "Bearer " + thunkAPI.getState().accountSlice.jwt;

  let res = await fetch("http://localhost:8080/user-management/user", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Authorization: jwt,
    },
  });

  if (!res.ok) {
    res = await res.json();
    throw new Error(`${JSON.stringify(res)}`);
  }

  res = await res.json();

  return res;
});

const initialState = {
  isLoggedIn: false,
  jwt: null,
  info: null,
};

const accountSlice = createSlice({
  name: "account",
  initialState,
  reducers: {
    login(state) {
      state.isLoggedIn = true;
    },
    logout(state) {
      state.isLoggedIn = false;
      state.jwt = null;
      state.info = null;
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(loginThunk.fulfilled, (state, action) => {
        state.isLoggedIn = true;
        state.jwt = action.payload["jwt"];
      })
      .addCase(loginThunk.rejected, (_, action) => {
        alert(action.error.message);
      })
      .addCase(getInfo.rejected, (_, action) => {
        alert(action.error.message);
      })
      .addCase(getInfo.fulfilled, (state, action) => {
        if (!state.info) state.info = action.payload;
      });
  },
});

export const { login, logout } = accountSlice.actions;
export default accountSlice.reducer;
export { loginThunk, getInfo };
