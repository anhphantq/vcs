import { createSlice } from "@reduxjs/toolkit";
import { createAsyncThunk } from "@reduxjs/toolkit";

const getUsers = createAsyncThunk("/admin/getusers", async (_, thunkAPI) => {
  const jwt = "Bearer " + thunkAPI.getState().accountSlice.jwt;

  let res = await fetch("http://localhost:8080/user-management/user/all", {
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
  list: null,
};

const adminSlice = createSlice({
  name: "users",
  initialState,
  reducers: {
    deleteList(state) {
      state.list = null;
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(getUsers.rejected, (_, action) => {
        alert("Error went get users list" + action.error.message);
      })
      .addCase(getUsers.fulfilled, (state, action) => {
        if (!state.list) state.list = action.payload;
      });
  },
});

export const { deleteList } = adminSlice.actions;
export default adminSlice.reducer;
export { getUsers };
